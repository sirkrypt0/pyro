// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agent_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error)
	ExecuteCommandStream(ctx context.Context, opts ...grpc.CallOption) (AgentService_ExecuteCommandStreamClient, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error) {
	out := new(ExecuteCommandResponse)
	err := c.cc.Invoke(ctx, "/api.agent.v1.AgentService/ExecuteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ExecuteCommandStream(ctx context.Context, opts ...grpc.CallOption) (AgentService_ExecuteCommandStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], "/api.agent.v1.AgentService/ExecuteCommandStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceExecuteCommandStreamClient{stream}
	return x, nil
}

type AgentService_ExecuteCommandStreamClient interface {
	Send(*ExecuteCommandStreamRequest) error
	Recv() (*ExecuteCommandStreamResponse, error)
	grpc.ClientStream
}

type agentServiceExecuteCommandStreamClient struct {
	grpc.ClientStream
}

func (x *agentServiceExecuteCommandStreamClient) Send(m *ExecuteCommandStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServiceExecuteCommandStreamClient) Recv() (*ExecuteCommandStreamResponse, error) {
	m := new(ExecuteCommandStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations should embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error)
	ExecuteCommandStream(AgentService_ExecuteCommandStreamServer) error
}

// UnimplementedAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedAgentServiceServer) ExecuteCommandStream(AgentService_ExecuteCommandStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteCommandStream not implemented")
}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.agent.v1.AgentService/ExecuteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ExecuteCommand(ctx, req.(*ExecuteCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ExecuteCommandStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).ExecuteCommandStream(&agentServiceExecuteCommandStreamServer{stream})
}

type AgentService_ExecuteCommandStreamServer interface {
	Send(*ExecuteCommandStreamResponse) error
	Recv() (*ExecuteCommandStreamRequest, error)
	grpc.ServerStream
}

type agentServiceExecuteCommandStreamServer struct {
	grpc.ServerStream
}

func (x *agentServiceExecuteCommandStreamServer) Send(m *ExecuteCommandStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServiceExecuteCommandStreamServer) Recv() (*ExecuteCommandStreamRequest, error) {
	m := new(ExecuteCommandStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.agent.v1.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommand",
			Handler:    _AgentService_ExecuteCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteCommandStream",
			Handler:       _AgentService_ExecuteCommandStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/agent/v1/agent.proto",
}
