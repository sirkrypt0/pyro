package agent

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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

	go io.Copy(&stdoutBuffer, stdout)
	go io.Copy(&stderrBuffer, stderr)

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

// TODO: think of using custom readers/writers instead of doing the copying stuff
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

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("error retrieving stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error retrieving stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error retrieving stderr pipe: %w", err)
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
				return fmt.Errorf("error sending exec error: %w", err)
			}
			return nil
		}
		return fmt.Errorf("error starting command: %w", err)
	}

	errCh := make(chan error, 3)
	go streamStdin(ctx, stdin, errCh, stream)
	go streamStdout(ctx, stdout, errCh, stream)
	go streamStderr(ctx, stderr, errCh, stream)

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

func streamStdin(
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

func streamStdout(
	ctx context.Context, stdout io.ReadCloser, errCh chan error,
	stream api.AgentService_ExecuteCommandStreamServer,
) {
	//nolint:makezero // We can't initialize with zero here as otherwise the reader won't block
	outputBuffer := make([]byte, streamingOutputBufferSize)
	for {
		if ctx.Err() != nil {
			return
		}
		output := api.ExecuteCommandStreamResponse{Stdout: &api.ExecuteIO{}}
		if err := readOutput(stdout, outputBuffer, output.Stdout); err != nil {
			errCh <- err
			return
		}

		if err := stream.Send(&output); err != nil {
			errCh <- err
			return
		}
	}
}

func streamStderr(
	ctx context.Context, stderr io.ReadCloser, errCh chan error,
	stream api.AgentService_ExecuteCommandStreamServer,
) {
	//nolint:makezero // We can't initialize with zero here as otherwise the reader won't block
	outputBuffer := make([]byte, streamingOutputBufferSize)
	for {
		if ctx.Err() != nil {
			return
		}
		output := api.ExecuteCommandStreamResponse{Stderr: &api.ExecuteIO{}}
		if err := readOutput(stderr, outputBuffer, output.Stderr); err != nil {
			errCh <- err
			return
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
