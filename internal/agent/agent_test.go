package agent

import (
	"context"
	api "github.com/sirkrypt0/pyro/api/agent/v1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"net"
	"testing"
)

type ServerTestSuite struct {
	suite.Suite
	client         api.AgentServiceClient
	teardownServer func()
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (s *ServerTestSuite) SetupTest() {
	client, teardown := newTestServer(s.T())
	s.client = client
	s.teardownServer = teardown
}

func (s *ServerTestSuite) TeardownTest() {
	s.teardownServer()
}

func (s *ServerTestSuite) TestExecuteCommand() {
	ctx := context.Background()
	req := &api.ExecuteCommandRequest{
		Command:     "echo 'hello world'",
		Environment: map[string]string{"PATH": "/bin"},
	}
	expectedResponse := &api.ExecuteCommandResponse{
		Stdout:   &api.ExecuteIO{Close: true, Data: []byte("stdout")},
		Stderr:   &api.ExecuteIO{Close: true, Data: []byte("stderr")},
		ExitCode: 0,
	}
	resp, err := s.client.ExecuteCommand(ctx, req)
	s.NoError(err)
	s.Require().NotNil(resp)
	s.Require().NotNil(resp.Stdout)
	s.Equal(expectedResponse.Stdout.Close, resp.Stdout.Close)
	s.Equal(expectedResponse.Stdout.Data, resp.Stdout.Data)
	s.Require().NotNil(resp.Stderr)
	s.Equal(expectedResponse.Stderr.Close, resp.Stderr.Close)
	s.Equal(expectedResponse.Stderr.Data, resp.Stderr.Data)
}

func (s *ServerTestSuite) TestExecuteCommandStream() {
	ctx := context.Background()
	req := &api.ExecuteCommandStreamRequest{
		Prepare: &api.ExecuteCommandStreamRequest_Prepare{
			Command:     "echo 'hello world'",
			Environment: map[string]string{"PATH": "/bin"},
		},
		Stdin: nil,
	}
	expectedResponse := &api.ExecuteCommandStreamResponse{
		Stdout: &api.ExecuteIO{Close: true, Data: []byte("stdout")},
		Stderr: &api.ExecuteIO{Close: true, Data: []byte("stderr")},
		Result: &api.ExecuteResult{
			Exited:   true,
			ExitCode: 0,
		},
	}

	stream, err := s.client.ExecuteCommandStream(ctx)
	s.Require().NoError(err)

	err = stream.Send(req)
	s.Require().NoError(err)

	resp, err := stream.Recv()
	s.Require().NoError(err)

	s.Require().NotNil(resp)
	s.Require().NotNil(resp.Stdout)
	s.Equal(expectedResponse.Stdout.Close, resp.Stdout.Close)
	s.Equal(expectedResponse.Stdout.Data, resp.Stdout.Data)
	s.Require().NotNil(resp.Stderr)
	s.Equal(expectedResponse.Stderr.Close, resp.Stderr.Close)
	s.Equal(expectedResponse.Stderr.Data, resp.Stderr.Data)
	s.Require().NotNil(resp.Result)
	s.Equal(expectedResponse.Result.Exited, resp.Result.Exited)
	s.Equal(expectedResponse.Result.ExitCode, resp.Result.ExitCode)
}

func newTestServer(t *testing.T) (client api.AgentServiceClient, teardown func()) {
	t.Helper()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	clientOptions := []grpc.DialOption{grpc.WithInsecure()}
	cc, err := grpc.Dial(l.Addr().String(), clientOptions...)
	require.NoError(t, err)

	server, err := NewGRPCServer()
	require.NoError(t, err)

	go func() {
		require.NoError(t, server.Serve(l))
	}()

	client = api.NewAgentServiceClient(cc)

	teardown = func() {
		server.Stop()
		require.NoError(t, cc.Close())
		require.NoError(t, l.Close())
	}

	return client, teardown
}
