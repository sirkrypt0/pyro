package streamio

import (
	"bytes"
	"errors"
	"fmt"
	api "github.com/sirkrypt0/pyro/api/agent/v1"
	"io"
	"os"
)

var ErrAlreadyReceiving = errors.New("a receiving goroutine is already running")

type IoToExecuteCommandStreamServerReader struct {
	Stream api.AgentService_ExecuteCommandStreamServer
	// procState represents the state of the underlying process that consumes the stdin.
	// It is a pointer to a pointer, as initially, when creating the Command, no process exists.
	ProcState **os.ProcessState

	inBuff    bytes.Buffer
	receiving bool
	err       error
	closed    bool
}

type IoToExecuteCommandStreamServerWriter struct {
	Stream api.AgentService_ExecuteCommandStreamServer
	Stderr bool
}

func (r *IoToExecuteCommandStreamServerReader) BeginReceiving() error {
	if r.receiving {
		return ErrAlreadyReceiving
	}
	go func() {
		r.receiving = true
		var recv *api.ExecuteCommandStreamRequest
		for {
			recv, r.err = r.Stream.Recv()
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
	if r.closed || (r.ProcState != nil && *r.ProcState != nil && (*r.ProcState).Exited()) {
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

func (w *IoToExecuteCommandStreamServerWriter) Write(p []byte) (n int, err error) {
	output := api.ExecuteCommandStreamResponse{}
	executeIO := &api.ExecuteIO{Data: p}
	if w.Stderr {
		output.Stderr = executeIO
	} else {
		output.Stdout = executeIO
	}
	if err := w.Stream.Send(&output); err != nil {
		return 0, fmt.Errorf("error sending output to stream: %w", err)
	}
	return len(p), nil
}
