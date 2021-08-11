package agent

import (
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
