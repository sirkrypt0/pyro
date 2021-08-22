package agent

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/kr/pty"
	api "github.com/sirkrypt0/pyro/api/agent/v1"
	"io"
	"os/exec"
)

const streamingOutputBufferSize = 2048

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

	var stdoutBuffer, stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer

	var exitCode = 0
	if err := cmd.Run(); err != nil {
		switch err := err.(type) {
		case *exec.Error:
			return &api.ExecuteCommandResponse{
				Stderr:   &api.ExecuteIO{Close: true, Data: []byte(err.Error())},
				ExitCode: -2,
			}, nil
		case *exec.ExitError:
			exitCode = err.ExitCode()
		default:
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

	var exitCode int
	exitCode, err = runCommand(cmd, prepare.Tty, stream, ctx)
	if err != nil {
		return fmt.Errorf("error running command: %w", err)
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

func runCommand(
	cmd *exec.Cmd, tty bool, stream api.AgentService_ExecuteCommandStreamServer, ctx context.Context,
) (exitCode int, err error) {
	var stdin io.WriteCloser
	var stdout, stderr io.ReadCloser
	if tty {
		ptmx, err := pty.Start(cmd)
		if err != nil {
			return -2, fmt.Errorf("error creating pty: %w", err)
		}
		// Make sure to close the pty at the end.
		defer func() { _ = ptmx.Close() }()
	} else {
		stdin, stdout, stderr, err = connectCommandPipes(cmd)
		if err != nil {
			return -2, fmt.Errorf("error connecting pipes: %w", err)
		}
	}

	if err := cmd.Start(); err != nil {
		var execErr *exec.Error
		if errors.As(err, &execErr) {
			err := stream.Send(&api.ExecuteCommandStreamResponse{
				Stderr: &api.ExecuteIO{Close: true, Data: []byte(execErr.Error())},
				Result: &api.ExecuteResult{
					Exited:   true,
					ExitCode: -2,
				},
			})
			if err != nil {
				return -2, fmt.Errorf("error sending exec error %v: %w", execErr, err)
			}
			return -2, execErr
		}
		return -2, fmt.Errorf("error starting command: %w", err)
	}

	const concurrentStreams = 3
	errCh := make(chan error, concurrentStreams)
	go streamInput(ctx, stdin, errCh, stream)
	go streamOutput(ctx, stdout, errCh, stream, false)
	go streamOutput(ctx, stderr, errCh, stream, true)

	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			exitCode = exitErr.ExitCode()
		} else {
			return -2, fmt.Errorf("error running command: %w", err)
		}
	}
	return exitCode, nil
}

func connectCommandPipes(cmd *exec.Cmd) (stdin io.WriteCloser, stdout, stderr io.ReadCloser, err error) {
	pty.Start()
	stdin, err = cmd.StdinPipe()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error retrieving stdin pipe: %w", err)
	}
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error retrieving stdout pipe: %w", err)
	}
	stderr, err = cmd.StderrPipe()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error retrieving stderr pipe: %w", err)
	}
	return stdin, stdout, stderr, nil
}

//nolint:gocognit // It is not that hard to read.
func streamInput(
	ctx context.Context, stdin io.WriteCloser, errCh chan error,
	stream api.AgentService_ExecuteCommandStreamServer,
) {
	for {
		if ctx.Err() != nil {
			return
		}
		recv, err := stream.Recv()
		if err != nil {
			errCh <- err
			return
		}

		if recv.Stdin != nil {
			if len(recv.Stdin.Data) != 0 {
				if _, err := stdin.Write(recv.Stdin.Data); err != nil {
					errCh <- err
					return
				}
			}
			if recv.Stdin.Close {
				if err := stdin.Close(); err != nil {
					errCh <- err
					return
				}
			}
		}
	}
}

func streamOutput(
	ctx context.Context, outReader io.ReadCloser, errCh chan error,
	stream api.AgentService_ExecuteCommandStreamServer, stderr bool,
) {
	//nolint:makezero // We can't initialize with zero here as otherwise the reader won't block
	outputBuffer := make([]byte, streamingOutputBufferSize)
	for {
		if ctx.Err() != nil {
			return
		}
		output := api.ExecuteCommandStreamResponse{}
		executeIO := &api.ExecuteIO{}
		if err := readOutput(outReader, outputBuffer, executeIO); err != nil {
			errCh <- err
			return
		}

		if stderr {
			output.Stderr = executeIO
		} else {
			output.Stdout = executeIO
		}

		if err := stream.Send(&output); err != nil {
			errCh <- err
			return
		}
	}
}

func readOutput(outr io.ReadCloser, buffer []byte, output *api.ExecuteIO) error {
	n, err := outr.Read(buffer)

	if n != 0 {
		output.Data = buffer[:n]
	}

	if err != nil {
		if errors.Is(err, io.EOF) && n == 0 {
			output.Close = true
		} else {
			return fmt.Errorf("error reading output: %w", err)
		}
	}
	return nil
}
