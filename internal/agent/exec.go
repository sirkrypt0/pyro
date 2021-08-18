package agent

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	api "github.com/sirkrypt0/pyro/api/agent/v1"
	"github.com/sirkrypt0/pyro/pkg/streamio"
	"io"
	"os/exec"
	"sync"
)

var (
	ErrNoCommandGiven      = errors.New("no command given")
	ErrExpectedPrepMessage = errors.New("expected first message to be a preparation message")
)

func (s *server) ExecuteCommand(
	ctx context.Context, req *api.ExecuteCommandRequest,
) (*api.ExecuteCommandResponse, error) {
	s.logger.WithField("executeCommandRequest", req).Debug("Got execute command request")

	if len(req.Command) == 0 {
		return nil, ErrNoCommandGiven
	}

	//nolint:gosec // G204: It is intended that arbitrary commands can be executed here
	cmd := exec.CommandContext(ctx, req.Command[0], req.Command[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("error retrieving stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("error retrieving stderr pipe: %w", err)
	}

	var stdoutBuffer, stderrBuffer bytes.Buffer

	if err := cmd.Start(); err != nil {
		var execErr *exec.Error
		if errors.As(err, &execErr) {
			return &api.ExecuteCommandResponse{
				Stderr:   &api.ExecuteIO{Close: true, Data: []byte(execErr.Error())},
				ExitCode: -2,
			}, nil
		}
		return nil, fmt.Errorf("error starting command: %w", err)
	}

	go func() {
		if _, err := io.Copy(&stdoutBuffer, stdout); err != nil {
			s.logger.WithError(err).Warn("error copying stdout")
		}
	}()
	go func() {
		if _, err := io.Copy(&stderrBuffer, stderr); err != nil {
			s.logger.WithError(err).Warn("error copying stderr")
		}
	}()

	var exitCode = 0
	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			exitCode = exitErr.ExitCode()
		} else {
			return nil, fmt.Errorf("error running command: %w", err)
		}
	}

	return &api.ExecuteCommandResponse{
		Stdout:   &api.ExecuteIO{Close: true, Data: stdoutBuffer.Bytes()},
		Stderr:   &api.ExecuteIO{Close: true, Data: stderrBuffer.Bytes()},
		ExitCode: int32(exitCode),
	}, nil
}

func (s *server) ExecuteCommandStream(stream api.AgentService_ExecuteCommandStreamServer) error {
	s.logger.WithField("executeCommandStreamServer", stream).Debug("Got execute command request")

	prepare, err := prepareExecution(stream)
	if err != nil {
		return err
	}

	ctx := stream.Context()

	//nolint:gosec // G204: It is intended that arbitrary commands can be executed here
	cmd := exec.CommandContext(ctx, prepare.Command[0], prepare.Command[1:]...)
	cmd.Env = append(cmd.Env, prepare.Environment...)

	var wg sync.WaitGroup
	if err := s.connectIO(stream, cmd, &wg); err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		var execErr *exec.Error
		if errors.As(err, &execErr) {
			err = stream.Send(&api.ExecuteCommandStreamResponse{
				Stderr: &api.ExecuteIO{Close: true, Data: []byte(execErr.Error())},
				Result: &api.ExecuteResult{
					Exited:   true,
					ExitCode: -2,
				},
			})
			if err != nil {
				return fmt.Errorf("error sending exec error: %w", err)
			}
			return nil
		}
		return fmt.Errorf("error starting command: %w", err)
	}

	// cmd.Wait should only be called after all reading is done
	wg.Wait()

	var exitCode = 0
	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			exitCode = exitErr.ExitCode()
		} else {
			return fmt.Errorf("error running command: %w", err)
		}
	}

	exitResponse := api.ExecuteCommandStreamResponse{Result: &api.ExecuteResult{
		Exited:   true,
		ExitCode: int32(exitCode),
	}}
	if err := stream.Send(&exitResponse); err != nil {
		return fmt.Errorf("error sending exit response: %w", err)
	}
	return nil
}

func prepareExecution(
	stream api.AgentService_ExecuteCommandStreamServer,
) (*api.ExecuteCommandStreamRequest_Prepare, error) {
	recv, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("receiving from stream failed: %w", err)
	}
	if recv.Prepare == nil {
		return nil, ErrExpectedPrepMessage
	}
	if len(recv.Prepare.Command) == 0 {
		return nil, ErrNoCommandGiven
	}
	return recv.Prepare, nil
}

func (s *server) connectIO(
	stream api.AgentService_ExecuteCommandStreamServer, cmd *exec.Cmd, wg *sync.WaitGroup,
) (err error) {
	var stdinWriter, stdoutWriter, stderrWriter io.WriteCloser
	var stdinReader, stdoutReader, stderrReader io.ReadCloser

	if stdinWriter, stdoutReader, stderrReader, err = createPipes(cmd); err != nil {
		return fmt.Errorf("error creating pipes: %w", err)
	}

	if stdinReader, stdoutWriter, stderrWriter, err = createIOToStreamConverters(stream, cmd); err != nil {
		return fmt.Errorf("error creating IO to stream converters: %w", err)
	}

	go func() {
		if _, err := io.Copy(stdinWriter, stdinReader); err != nil {
			s.logger.WithError(err).Error("error copying stdin")
		}
		s.logger.Debug("Stdin copying finished")
	}()

	const pipeReadingGoRoutines = 2
	wg.Add(pipeReadingGoRoutines)
	go func() {
		if _, err := io.Copy(stdoutWriter, stdoutReader); err != nil {
			s.logger.WithError(err).Errorf("error copying stdout")
		}
		if err := stdoutWriter.Close(); err != nil {
			s.logger.WithError(err).Errorf("error closing stdout writer")
		}
		wg.Done()
		s.logger.Debug("Stdout copying finished")
	}()
	go func() {
		if _, err := io.Copy(stderrWriter, stderrReader); err != nil {
			s.logger.WithError(err).Errorf("error copying stderr")
		}
		if err := stderrWriter.Close(); err != nil {
			s.logger.WithError(err).Errorf("error closing stderr writer")
		}
		wg.Done()
		s.logger.Debug("Stderr copying finished")
	}()

	return nil
}

func createPipes(cmd *exec.Cmd) (stdinWriter io.WriteCloser, stdoutReader, stderrReader io.ReadCloser, err error) {
	if stdinWriter, err = cmd.StdinPipe(); err != nil {
		return nil, nil, nil, fmt.Errorf("failed creating stdin pipe: %w", err)
	}
	if stdoutReader, err = cmd.StdoutPipe(); err != nil {
		return nil, nil, nil, fmt.Errorf("failed creating stdout pipe: %w", err)
	}
	if stderrReader, err = cmd.StderrPipe(); err != nil {
		return nil, nil, nil, fmt.Errorf("failed creating stdout pipe: %w", err)
	}
	return stdinWriter, stdoutReader, stderrReader, nil
}

func createIOToStreamConverters(
	stream api.AgentService_ExecuteCommandStreamServer, cmd *exec.Cmd,
) (stdinReader io.ReadCloser, stdoutWriter, stderrWriter io.WriteCloser, err error) {
	if stdinReader, err = streamio.NewIoToExecuteCommandStreamServerReader(stream); err != nil {
		return nil, nil, nil, fmt.Errorf("error creating new stdin reader: %w", err)
	}
	if stdoutWriter, err = streamio.NewIoToExecuteCommandStreamServerWriter(stream, false); err != nil {
		return nil, nil, nil, fmt.Errorf("error creating new stdout writer: %w", err)
	}
	if stderrWriter, err = streamio.NewIoToExecuteCommandStreamServerWriter(stream, true); err != nil {
		return nil, nil, nil, fmt.Errorf("error creating new stderr writer: %w", err)
	}
	return stdinReader, stdoutWriter, stderrWriter, nil
}
