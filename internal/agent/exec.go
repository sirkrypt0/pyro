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

	stdinReader := streamio.IoToExecuteCommandStreamServerReader{Stream: stream, ProcState: &cmd.ProcessState}
	if err := stdinReader.BeginReceiving(); err != nil {
		return fmt.Errorf("error starting to receive stdin: %w", err)
	}

	// TODO: Currently, we don't inform the client about closed streams. Is this fine?
	// We could use cmd.StdinPipe() with io.Copy() and our readers instead (also see func (c *Cmd) stdin())
	cmd.Stdin = &stdinReader
	cmd.Stdout = &streamio.IoToExecuteCommandStreamServerWriter{Stream: stream}
	cmd.Stderr = &streamio.IoToExecuteCommandStreamServerWriter{Stream: stream, Stderr: true}

	var exitCode = 0
	if err := cmd.Run(); err != nil {
		switch err := err.(type) {
		case *exec.Error:
			err2 := stream.Send(&api.ExecuteCommandStreamResponse{
				Stderr: &api.ExecuteIO{Close: true, Data: []byte(err.Error())},
				Result: &api.ExecuteResult{
					Exited:   true,
					ExitCode: -2,
				},
			})
			if err2 != nil {
				return fmt.Errorf("error sending exec error: %w", err2)
			}
			return nil
		case *exec.ExitError:
			exitCode = err.ExitCode()
		default:
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
