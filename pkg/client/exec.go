package client

import (
	"context"
	"errors"
	"fmt"
	agentv1 "github.com/sirkrypt0/pyro/api/agent/v1"
	"io"
)

const (
	streamingInputBufferSize = 2048
	exitCodeError            = -2
)

func (c *Client) Execute(command []string, ctx context.Context, outw, errw io.WriteCloser) (exitCode int, err error) {
	req := &agentv1.ExecuteCommandRequest{
		Command:     command,
		Environment: nil,
	}
	resp, err := c.agent.ExecuteCommand(ctx, req)
	if err != nil {
		return exitCodeError, fmt.Errorf("error executing command: %w", err)
	}
	if resp.Stdout != nil {
		if _, err := outw.Write(resp.Stdout.Data); err != nil {
			return exitCodeError, fmt.Errorf("error writing to out: %w", err)
		}
	}
	if resp.Stderr != nil {
		_, err := errw.Write(resp.Stderr.Data)
		if err != nil {
			return exitCodeError, fmt.Errorf("error writing to err: %w", err)
		}
	}
	if err := outw.Close(); err != nil {
		return exitCodeError, fmt.Errorf("error closing out writer: %w", err)
	}
	if err := errw.Close(); err != nil {
		return exitCodeError, fmt.Errorf("error closing err writer: %w", err)
	}
	return int(resp.ExitCode), nil
}

func (c *Client) ExecuteInteractively(
	command []string, ctx context.Context, inr io.Reader, outw, errw io.WriteCloser,
) (exitCode int, err error) {
	prep := &agentv1.ExecuteCommandStreamRequest{
		Prepare: &agentv1.ExecuteCommandStreamRequest_Prepare{
			Command:     command,
			Environment: nil,
		},
	}
	stream, err := c.agent.ExecuteCommandStream(ctx)
	if err != nil {
		return exitCodeError, fmt.Errorf("error executing command stream: %w", err)
	}

	errCh := make(chan error, 1)
	exitCh := make(chan int)

	// start listeners first
	go streamInput(ctx, inr, errCh, stream)
	go streamOutput(ctx, outw, errw, errCh, exitCh, stream)

	// kick off execution next
	if err := stream.Send(prep); err != nil {
		return exitCodeError, fmt.Errorf("error sending command: %w", err)
	}

	// finally wait for exit or any failure
	select {
	case err := <-errCh:
		return exitCodeError, err
	case <-ctx.Done():
		return exitCodeError, fmt.Errorf("context is done: %w", ctx.Err())
	case exit := <-exitCh:
		return exit, nil
	}
}

//nolint:gocognit // Currently, the function is quite readable and straight forward.
func streamInput(
	ctx context.Context, inr io.Reader, errCh chan error,
	stream agentv1.AgentService_ExecuteCommandStreamClient,
) {
	//nolint:makezero // We can't initialize with zero here as otherwise the reader won't block
	bytes := make([]byte, streamingInputBufferSize)
	for {
		if ctx.Err() != nil {
			return
		}
		input := agentv1.ExecuteCommandStreamRequest{Stdin: &agentv1.ExecuteIO{}}
		n, err := inr.Read(bytes)

		if n != 0 {
			input.Stdin.Data = bytes[:n]
			if err := stream.Send(&input); err != nil {
				errCh <- err
				return
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) && n == 0 {
				input.Stdin.Close = true
				if err := stream.Send(&input); err != nil {
					errCh <- err
					return
				}
			} else {
				errCh <- err
				return
			}
		}
	}
}

func streamOutput(
	ctx context.Context, outw, errw io.WriteCloser, errCh chan error, exitCh chan int,
	stream agentv1.AgentService_ExecuteCommandStreamClient,
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
		if err := writeReceivedData(recv.Stdout, outw); err != nil {
			errCh <- err
			return
		}
		if err := writeReceivedData(recv.Stderr, errw); err != nil {
			errCh <- err
			return
		}
		if recv.Result != nil && recv.Result.Exited {
			exitCh <- int(recv.Result.ExitCode)
			return
		}
	}
}

func writeReceivedData(recv *agentv1.ExecuteIO, out io.WriteCloser) error {
	if recv != nil {
		if len(recv.Data) != 0 {
			if _, err := out.Write(recv.Data); err != nil {
				return fmt.Errorf("error writing received data: %w", err)
			}
		}
		if recv.Close {
			if err := out.Close(); err != nil {
				return fmt.Errorf("error closing out writer: %w", err)
			}
		}
	}
	return nil
}
