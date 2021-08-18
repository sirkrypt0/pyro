package streamio

import (
	"bytes"
	"errors"
	"fmt"
	api "github.com/sirkrypt0/pyro/api/agent/v1"
	"io"
)

var (
	ErrAlreadyReceiving = errors.New("a receiving goroutine is already running")
	ErrAlreadyClosed    = errors.New("already closed")
)

type IoToExecuteCommandStreamServerReader struct {
	stream api.AgentService_ExecuteCommandStreamServer
	inBuff    bytes.Buffer
	receiving bool
	err       error
	closed    bool
}

type IoToExecuteCommandStreamServerWriter struct {
	stream api.AgentService_ExecuteCommandStreamServer
	stderr bool
	closed bool
}

func NewIoToExecuteCommandStreamServerReader(
	stream api.AgentService_ExecuteCommandStreamServer,
) (*IoToExecuteCommandStreamServerReader, error) {
	r := &IoToExecuteCommandStreamServerReader{stream: stream}
	if err := r.BeginReceiving(); err != nil {
		return nil, fmt.Errorf("error starting to receive stdin: %w", err)
	}
	return r, nil
}

func NewIoToExecuteCommandStreamServerWriter(
	stream api.AgentService_ExecuteCommandStreamServer, stderr bool,
) (*IoToExecuteCommandStreamServerWriter, error) {
	w := &IoToExecuteCommandStreamServerWriter{
		stream: stream,
		stderr: stderr,
	}
	return w, nil
}

func (r *IoToExecuteCommandStreamServerReader) BeginReceiving() error {
	if r.receiving {
		return ErrAlreadyReceiving
	}
	go func() {
		r.receiving = true
		var recv *api.ExecuteCommandStreamRequest
		for {
			recv, r.err = r.stream.Recv()
			if r.err != nil {
				return
			} else if recv.Stdin != nil {
				r.inBuff.Write(recv.Stdin.Data)
				if recv.Stdin.Close {
					r.closed = true
					return
				}
			}
		}
	}()
	return nil
}

func (r *IoToExecuteCommandStreamServerReader) Read(p []byte) (n int, err error) {
	// we need to send EOF when the process exited for cmd.Wait() to finish
	if r.closed {
		return 0, io.EOF
	}
	if r.err != nil {
		return 0, r.err
	}
	if r.inBuff.Len() == 0 {
		// nothing to read right now, ignore. Otherwise the next read returns io.EOF
		return 0, nil
	}
	n, err = r.inBuff.Read(p)
	if err != nil {
		return n, fmt.Errorf("error reading from in buffer: %w", err)
	}
	return n, nil
}

func (r *IoToExecuteCommandStreamServerReader) Close() error {
	if r.closed {
		return ErrAlreadyClosed
	}
	r.closed = true
	return nil
}

func (w *IoToExecuteCommandStreamServerWriter) Write(p []byte) (n int, err error) {
	output := api.ExecuteCommandStreamResponse{}
	executeIO := &api.ExecuteIO{Data: p}
	if w.stderr {
		output.Stderr = executeIO
	} else {
		output.Stdout = executeIO
	}
	if err := w.stream.Send(&output); err != nil {
		return 0, fmt.Errorf("error sending output to stream: %w", err)
	}
	return len(p), nil
}

func (w *IoToExecuteCommandStreamServerWriter) Close() error {
	if w.closed {
		return ErrAlreadyClosed
	}
	w.closed = true

	closeMsg := api.ExecuteCommandStreamResponse{}
	executeIO := &api.ExecuteIO{Close: true}
	if w.stderr {
		closeMsg.Stderr = executeIO
	} else {
		closeMsg.Stdout = executeIO
	}

	err := w.stream.Send(&closeMsg)
	if err != nil {
		return fmt.Errorf("error sending close message to client: %w", err)
	}

	return nil
}
