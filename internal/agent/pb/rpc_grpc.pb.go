// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: rpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AgentService_ListContainers_FullMethodName         = "/protobuf.AgentService/ListContainers"
	AgentService_FindContainer_FullMethodName          = "/protobuf.AgentService/FindContainer"
	AgentService_StreamLogs_FullMethodName             = "/protobuf.AgentService/StreamLogs"
	AgentService_LogsBetweenDates_FullMethodName       = "/protobuf.AgentService/LogsBetweenDates"
	AgentService_StreamRawBytes_FullMethodName         = "/protobuf.AgentService/StreamRawBytes"
	AgentService_StreamEvents_FullMethodName           = "/protobuf.AgentService/StreamEvents"
	AgentService_StreamStats_FullMethodName            = "/protobuf.AgentService/StreamStats"
	AgentService_StreamContainerStarted_FullMethodName = "/protobuf.AgentService/StreamContainerStarted"
	AgentService_HostInfo_FullMethodName               = "/protobuf.AgentService/HostInfo"
	AgentService_ContainerAction_FullMethodName        = "/protobuf.AgentService/ContainerAction"
	AgentService_ContainerExec_FullMethodName          = "/protobuf.AgentService/ContainerExec"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	ListContainers(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error)
	FindContainer(ctx context.Context, in *FindContainerRequest, opts ...grpc.CallOption) (*FindContainerResponse, error)
	StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamLogsResponse], error)
	LogsBetweenDates(ctx context.Context, in *LogsBetweenDatesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamLogsResponse], error)
	StreamRawBytes(ctx context.Context, in *StreamRawBytesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamRawBytesResponse], error)
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamEventsResponse], error)
	StreamStats(ctx context.Context, in *StreamStatsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamStatsResponse], error)
	StreamContainerStarted(ctx context.Context, in *StreamContainerStartedRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamContainerStartedResponse], error)
	HostInfo(ctx context.Context, in *HostInfoRequest, opts ...grpc.CallOption) (*HostInfoResponse, error)
	ContainerAction(ctx context.Context, in *ContainerActionRequest, opts ...grpc.CallOption) (*ContainerActionResponse, error)
	ContainerExec(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ContainerExecRequest, ContainerExecResponse], error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) ListContainers(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListContainersResponse)
	err := c.cc.Invoke(ctx, AgentService_ListContainers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) FindContainer(ctx context.Context, in *FindContainerRequest, opts ...grpc.CallOption) (*FindContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindContainerResponse)
	err := c.cc.Invoke(ctx, AgentService_FindContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamLogsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], AgentService_StreamLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamLogsRequest, StreamLogsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamLogsClient = grpc.ServerStreamingClient[StreamLogsResponse]

func (c *agentServiceClient) LogsBetweenDates(ctx context.Context, in *LogsBetweenDatesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamLogsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[1], AgentService_LogsBetweenDates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LogsBetweenDatesRequest, StreamLogsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_LogsBetweenDatesClient = grpc.ServerStreamingClient[StreamLogsResponse]

func (c *agentServiceClient) StreamRawBytes(ctx context.Context, in *StreamRawBytesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamRawBytesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[2], AgentService_StreamRawBytes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRawBytesRequest, StreamRawBytesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamRawBytesClient = grpc.ServerStreamingClient[StreamRawBytesResponse]

func (c *agentServiceClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamEventsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[3], AgentService_StreamEvents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamEventsRequest, StreamEventsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamEventsClient = grpc.ServerStreamingClient[StreamEventsResponse]

func (c *agentServiceClient) StreamStats(ctx context.Context, in *StreamStatsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamStatsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[4], AgentService_StreamStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamStatsRequest, StreamStatsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamStatsClient = grpc.ServerStreamingClient[StreamStatsResponse]

func (c *agentServiceClient) StreamContainerStarted(ctx context.Context, in *StreamContainerStartedRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamContainerStartedResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[5], AgentService_StreamContainerStarted_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamContainerStartedRequest, StreamContainerStartedResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamContainerStartedClient = grpc.ServerStreamingClient[StreamContainerStartedResponse]

func (c *agentServiceClient) HostInfo(ctx context.Context, in *HostInfoRequest, opts ...grpc.CallOption) (*HostInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HostInfoResponse)
	err := c.cc.Invoke(ctx, AgentService_HostInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ContainerAction(ctx context.Context, in *ContainerActionRequest, opts ...grpc.CallOption) (*ContainerActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContainerActionResponse)
	err := c.cc.Invoke(ctx, AgentService_ContainerAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ContainerExec(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ContainerExecRequest, ContainerExecResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[6], AgentService_ContainerExec_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ContainerExecRequest, ContainerExecResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_ContainerExecClient = grpc.BidiStreamingClient[ContainerExecRequest, ContainerExecResponse]

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility.
type AgentServiceServer interface {
	ListContainers(context.Context, *ListContainersRequest) (*ListContainersResponse, error)
	FindContainer(context.Context, *FindContainerRequest) (*FindContainerResponse, error)
	StreamLogs(*StreamLogsRequest, grpc.ServerStreamingServer[StreamLogsResponse]) error
	LogsBetweenDates(*LogsBetweenDatesRequest, grpc.ServerStreamingServer[StreamLogsResponse]) error
	StreamRawBytes(*StreamRawBytesRequest, grpc.ServerStreamingServer[StreamRawBytesResponse]) error
	StreamEvents(*StreamEventsRequest, grpc.ServerStreamingServer[StreamEventsResponse]) error
	StreamStats(*StreamStatsRequest, grpc.ServerStreamingServer[StreamStatsResponse]) error
	StreamContainerStarted(*StreamContainerStartedRequest, grpc.ServerStreamingServer[StreamContainerStartedResponse]) error
	HostInfo(context.Context, *HostInfoRequest) (*HostInfoResponse, error)
	ContainerAction(context.Context, *ContainerActionRequest) (*ContainerActionResponse, error)
	ContainerExec(grpc.BidiStreamingServer[ContainerExecRequest, ContainerExecResponse]) error
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAgentServiceServer struct{}

func (UnimplementedAgentServiceServer) ListContainers(context.Context, *ListContainersRequest) (*ListContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContainers not implemented")
}
func (UnimplementedAgentServiceServer) FindContainer(context.Context, *FindContainerRequest) (*FindContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindContainer not implemented")
}
func (UnimplementedAgentServiceServer) StreamLogs(*StreamLogsRequest, grpc.ServerStreamingServer[StreamLogsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}
func (UnimplementedAgentServiceServer) LogsBetweenDates(*LogsBetweenDatesRequest, grpc.ServerStreamingServer[StreamLogsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method LogsBetweenDates not implemented")
}
func (UnimplementedAgentServiceServer) StreamRawBytes(*StreamRawBytesRequest, grpc.ServerStreamingServer[StreamRawBytesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamRawBytes not implemented")
}
func (UnimplementedAgentServiceServer) StreamEvents(*StreamEventsRequest, grpc.ServerStreamingServer[StreamEventsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedAgentServiceServer) StreamStats(*StreamStatsRequest, grpc.ServerStreamingServer[StreamStatsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamStats not implemented")
}
func (UnimplementedAgentServiceServer) StreamContainerStarted(*StreamContainerStartedRequest, grpc.ServerStreamingServer[StreamContainerStartedResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamContainerStarted not implemented")
}
func (UnimplementedAgentServiceServer) HostInfo(context.Context, *HostInfoRequest) (*HostInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostInfo not implemented")
}
func (UnimplementedAgentServiceServer) ContainerAction(context.Context, *ContainerActionRequest) (*ContainerActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerAction not implemented")
}
func (UnimplementedAgentServiceServer) ContainerExec(grpc.BidiStreamingServer[ContainerExecRequest, ContainerExecResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ContainerExec not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}
func (UnimplementedAgentServiceServer) testEmbeddedByValue()                      {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	// If the following call pancis, it indicates UnimplementedAgentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_ListContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ListContainers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListContainers(ctx, req.(*ListContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_FindContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).FindContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_FindContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).FindContainer(ctx, req.(*FindContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamLogs(m, &grpc.GenericServerStream[StreamLogsRequest, StreamLogsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamLogsServer = grpc.ServerStreamingServer[StreamLogsResponse]

func _AgentService_LogsBetweenDates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsBetweenDatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).LogsBetweenDates(m, &grpc.GenericServerStream[LogsBetweenDatesRequest, StreamLogsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_LogsBetweenDatesServer = grpc.ServerStreamingServer[StreamLogsResponse]

func _AgentService_StreamRawBytes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRawBytesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamRawBytes(m, &grpc.GenericServerStream[StreamRawBytesRequest, StreamRawBytesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamRawBytesServer = grpc.ServerStreamingServer[StreamRawBytesResponse]

func _AgentService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamEvents(m, &grpc.GenericServerStream[StreamEventsRequest, StreamEventsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamEventsServer = grpc.ServerStreamingServer[StreamEventsResponse]

func _AgentService_StreamStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamStats(m, &grpc.GenericServerStream[StreamStatsRequest, StreamStatsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamStatsServer = grpc.ServerStreamingServer[StreamStatsResponse]

func _AgentService_StreamContainerStarted_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamContainerStartedRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamContainerStarted(m, &grpc.GenericServerStream[StreamContainerStartedRequest, StreamContainerStartedResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_StreamContainerStartedServer = grpc.ServerStreamingServer[StreamContainerStartedResponse]

func _AgentService_HostInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).HostInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_HostInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).HostInfo(ctx, req.(*HostInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ContainerAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ContainerAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ContainerAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ContainerAction(ctx, req.(*ContainerActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ContainerExec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).ContainerExec(&grpc.GenericServerStream[ContainerExecRequest, ContainerExecResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentService_ContainerExecServer = grpc.BidiStreamingServer[ContainerExecRequest, ContainerExecResponse]

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContainers",
			Handler:    _AgentService_ListContainers_Handler,
		},
		{
			MethodName: "FindContainer",
			Handler:    _AgentService_FindContainer_Handler,
		},
		{
			MethodName: "HostInfo",
			Handler:    _AgentService_HostInfo_Handler,
		},
		{
			MethodName: "ContainerAction",
			Handler:    _AgentService_ContainerAction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLogs",
			Handler:       _AgentService_StreamLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogsBetweenDates",
			Handler:       _AgentService_LogsBetweenDates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamRawBytes",
			Handler:       _AgentService_StreamRawBytes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEvents",
			Handler:       _AgentService_StreamEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamStats",
			Handler:       _AgentService_StreamStats_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamContainerStarted",
			Handler:       _AgentService_StreamContainerStarted_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ContainerExec",
			Handler:       _AgentService_ContainerExec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}
