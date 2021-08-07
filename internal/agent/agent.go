package agent

import (
	"context"
	"fmt"
	api "github.com/sirkrypt0/pyro/api/agent/v1"
	"github.com/sirkrypt0/pyro/pkg/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var _ api.AgentServiceServer = (*server)(nil)

type server struct {
	api.UnimplementedAgentServiceServer
	logger *logrus.Entry
}

func NewGRPCServer() (gsrv *grpc.Server, err error) {
	gsrv = grpc.NewServer()
	srv, err := NewServer()
	if err != nil {
		return nil, err
	}
	api.RegisterAgentServiceServer(gsrv, srv)
	return gsrv, nil
}

func NewServer() (srv *server, err error) {
	srv = &server{logger: logging.GetLogger("agent")}
	return srv, nil
}

func (s *server) ExecuteCommand(
	ctx context.Context, req *api.ExecuteCommandRequest,
) (*api.ExecuteCommandResponse, error) {
	s.logger.WithField("executeCommandRequest", req).Debug("Got execute command request")
	return &api.ExecuteCommandResponse{
		Stdout:   &api.ExecuteIO{Close: true, Data: []byte("stdout")},
		Stderr:   &api.ExecuteIO{Close: true, Data: []byte("stderr")},
		ExitCode: 0,
	}, nil
}

func (s *server) ExecuteCommandStream(stream api.AgentService_ExecuteCommandStreamServer) error {
	s.logger.WithField("executeCommandStreamServer", stream).Debug("Got execute command request")
	err := stream.Send(&api.ExecuteCommandStreamResponse{
		Stdout: &api.ExecuteIO{Close: true, Data: []byte("stdout")},
		Stderr: &api.ExecuteIO{Close: true, Data: []byte("stderr")},
		Result: &api.ExecuteResult{
			Exited:   true,
			ExitCode: 0,
		},
	})
	if err != nil {
		return fmt.Errorf("failed sending execute command stream response: %w", err)
	}
	return nil
}
