// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: flow/executiondata/executiondata.proto

package access

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

// ExecutionDataAPIClient is the client API for ExecutionDataAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutionDataAPIClient interface {
	// GetExecutionDataByBlockID returns execution data for a specific block ID.
	//
	// Errors:
	// - InvalidArgument is returned if the request contains an invalid block ID.
	// - NotFound is returned if the start block or execution data are not
	// currently available on the
	//
	//	node. This may happen if the block was from a previous spork, or if the
	//	block has yet not been received.
	GetExecutionDataByBlockID(ctx context.Context, in *GetExecutionDataByBlockIDRequest, opts ...grpc.CallOption) (*GetExecutionDataByBlockIDResponse, error)
	// SubscribeExecutionData streams execution data for all blocks starting at
	// the requested start block, up until the latest available block. Once the
	// latest is reached, the stream will remain open and responses are sent for
	// each new execution data as it becomes available.
	//
	// Errors:
	// - InvalidArgument is returned if the request contains an invalid start
	// block.
	// - NotFound is returned if the start block is not currently available on the
	// node. This may
	//
	//	happen if the block was from a previous spork, or if the block has yet
	//	not been received.
	SubscribeExecutionData(ctx context.Context, in *SubscribeExecutionDataRequest, opts ...grpc.CallOption) (ExecutionDataAPI_SubscribeExecutionDataClient, error)
	// SubscribeEvents streams events for all blocks starting at the requested
	// start block, up until the latest available block. Once the latest is
	// reached, the stream will remain open and responses are sent for each new
	// block as it becomes available.
	//
	// Events within each block are filtered by the provided EventFilter, and only
	// those events that match the filter are returned. If no filter is provided,
	// all events are returned.
	//
	// Responses are returned for each block containing at least one event that
	// matches the filter. Additionally, heatbeat responses
	// (SubscribeEventsResponse with no events) are returned periodically to allow
	// clients to track which blocks were searched. Clients can use this
	// information to determine which block to start from when reconnecting.
	//
	// Errors:
	// - InvalidArgument is returned if the request contains an invalid
	// EventFilter or start block.
	// - NotFound is returned if the start block is not currently available on the
	// node. This may
	//
	//	happen if the block was from a previous spork, or if the block has yet
	//	not been received.
	SubscribeEvents(ctx context.Context, in *SubscribeEventsRequest, opts ...grpc.CallOption) (ExecutionDataAPI_SubscribeEventsClient, error)
	// GetRegisterValues gets the register values for the given Ids as of the given block height
	GetRegisterValues(ctx context.Context, in *GetRegisterValuesRequest, opts ...grpc.CallOption) (*GetRegisterValuesResponse, error)
}

type executionDataAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionDataAPIClient(cc grpc.ClientConnInterface) ExecutionDataAPIClient {
	return &executionDataAPIClient{cc}
}

func (c *executionDataAPIClient) GetExecutionDataByBlockID(ctx context.Context, in *GetExecutionDataByBlockIDRequest, opts ...grpc.CallOption) (*GetExecutionDataByBlockIDResponse, error) {
	out := new(GetExecutionDataByBlockIDResponse)
	err := c.cc.Invoke(ctx, "/flow.access.ExecutionDataAPI/GetExecutionDataByBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionDataAPIClient) SubscribeExecutionData(ctx context.Context, in *SubscribeExecutionDataRequest, opts ...grpc.CallOption) (ExecutionDataAPI_SubscribeExecutionDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExecutionDataAPI_ServiceDesc.Streams[0], "/flow.access.ExecutionDataAPI/SubscribeExecutionData", opts...)
	if err != nil {
		return nil, err
	}
	x := &executionDataAPISubscribeExecutionDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExecutionDataAPI_SubscribeExecutionDataClient interface {
	Recv() (*SubscribeExecutionDataResponse, error)
	grpc.ClientStream
}

type executionDataAPISubscribeExecutionDataClient struct {
	grpc.ClientStream
}

func (x *executionDataAPISubscribeExecutionDataClient) Recv() (*SubscribeExecutionDataResponse, error) {
	m := new(SubscribeExecutionDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executionDataAPIClient) SubscribeEvents(ctx context.Context, in *SubscribeEventsRequest, opts ...grpc.CallOption) (ExecutionDataAPI_SubscribeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExecutionDataAPI_ServiceDesc.Streams[1], "/flow.access.ExecutionDataAPI/SubscribeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &executionDataAPISubscribeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExecutionDataAPI_SubscribeEventsClient interface {
	Recv() (*SubscribeEventsResponse, error)
	grpc.ClientStream
}

type executionDataAPISubscribeEventsClient struct {
	grpc.ClientStream
}

func (x *executionDataAPISubscribeEventsClient) Recv() (*SubscribeEventsResponse, error) {
	m := new(SubscribeEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executionDataAPIClient) GetRegisterValues(ctx context.Context, in *GetRegisterValuesRequest, opts ...grpc.CallOption) (*GetRegisterValuesResponse, error) {
	out := new(GetRegisterValuesResponse)
	err := c.cc.Invoke(ctx, "/flow.access.ExecutionDataAPI/GetRegisterValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionDataAPIServer is the server API for ExecutionDataAPI service.
// All implementations should embed UnimplementedExecutionDataAPIServer
// for forward compatibility
type ExecutionDataAPIServer interface {
	// GetExecutionDataByBlockID returns execution data for a specific block ID.
	//
	// Errors:
	// - InvalidArgument is returned if the request contains an invalid block ID.
	// - NotFound is returned if the start block or execution data are not
	// currently available on the
	//
	//	node. This may happen if the block was from a previous spork, or if the
	//	block has yet not been received.
	GetExecutionDataByBlockID(context.Context, *GetExecutionDataByBlockIDRequest) (*GetExecutionDataByBlockIDResponse, error)
	// SubscribeExecutionData streams execution data for all blocks starting at
	// the requested start block, up until the latest available block. Once the
	// latest is reached, the stream will remain open and responses are sent for
	// each new execution data as it becomes available.
	//
	// Errors:
	// - InvalidArgument is returned if the request contains an invalid start
	// block.
	// - NotFound is returned if the start block is not currently available on the
	// node. This may
	//
	//	happen if the block was from a previous spork, or if the block has yet
	//	not been received.
	SubscribeExecutionData(*SubscribeExecutionDataRequest, ExecutionDataAPI_SubscribeExecutionDataServer) error
	// SubscribeEvents streams events for all blocks starting at the requested
	// start block, up until the latest available block. Once the latest is
	// reached, the stream will remain open and responses are sent for each new
	// block as it becomes available.
	//
	// Events within each block are filtered by the provided EventFilter, and only
	// those events that match the filter are returned. If no filter is provided,
	// all events are returned.
	//
	// Responses are returned for each block containing at least one event that
	// matches the filter. Additionally, heatbeat responses
	// (SubscribeEventsResponse with no events) are returned periodically to allow
	// clients to track which blocks were searched. Clients can use this
	// information to determine which block to start from when reconnecting.
	//
	// Errors:
	// - InvalidArgument is returned if the request contains an invalid
	// EventFilter or start block.
	// - NotFound is returned if the start block is not currently available on the
	// node. This may
	//
	//	happen if the block was from a previous spork, or if the block has yet
	//	not been received.
	SubscribeEvents(*SubscribeEventsRequest, ExecutionDataAPI_SubscribeEventsServer) error
	// GetRegisterValues gets the register values for the given Ids as of the given block height
	GetRegisterValues(context.Context, *GetRegisterValuesRequest) (*GetRegisterValuesResponse, error)
}

// UnimplementedExecutionDataAPIServer should be embedded to have forward compatible implementations.
type UnimplementedExecutionDataAPIServer struct {
}

func (UnimplementedExecutionDataAPIServer) GetExecutionDataByBlockID(context.Context, *GetExecutionDataByBlockIDRequest) (*GetExecutionDataByBlockIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecutionDataByBlockID not implemented")
}
func (UnimplementedExecutionDataAPIServer) SubscribeExecutionData(*SubscribeExecutionDataRequest, ExecutionDataAPI_SubscribeExecutionDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeExecutionData not implemented")
}
func (UnimplementedExecutionDataAPIServer) SubscribeEvents(*SubscribeEventsRequest, ExecutionDataAPI_SubscribeEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeEvents not implemented")
}
func (UnimplementedExecutionDataAPIServer) GetRegisterValues(context.Context, *GetRegisterValuesRequest) (*GetRegisterValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisterValues not implemented")
}

// UnsafeExecutionDataAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutionDataAPIServer will
// result in compilation errors.
type UnsafeExecutionDataAPIServer interface {
	mustEmbedUnimplementedExecutionDataAPIServer()
}

func RegisterExecutionDataAPIServer(s grpc.ServiceRegistrar, srv ExecutionDataAPIServer) {
	s.RegisterService(&ExecutionDataAPI_ServiceDesc, srv)
}

func _ExecutionDataAPI_GetExecutionDataByBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExecutionDataByBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionDataAPIServer).GetExecutionDataByBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.ExecutionDataAPI/GetExecutionDataByBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionDataAPIServer).GetExecutionDataByBlockID(ctx, req.(*GetExecutionDataByBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionDataAPI_SubscribeExecutionData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeExecutionDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutionDataAPIServer).SubscribeExecutionData(m, &executionDataAPISubscribeExecutionDataServer{stream})
}

type ExecutionDataAPI_SubscribeExecutionDataServer interface {
	Send(*SubscribeExecutionDataResponse) error
	grpc.ServerStream
}

type executionDataAPISubscribeExecutionDataServer struct {
	grpc.ServerStream
}

func (x *executionDataAPISubscribeExecutionDataServer) Send(m *SubscribeExecutionDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExecutionDataAPI_SubscribeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutionDataAPIServer).SubscribeEvents(m, &executionDataAPISubscribeEventsServer{stream})
}

type ExecutionDataAPI_SubscribeEventsServer interface {
	Send(*SubscribeEventsResponse) error
	grpc.ServerStream
}

type executionDataAPISubscribeEventsServer struct {
	grpc.ServerStream
}

func (x *executionDataAPISubscribeEventsServer) Send(m *SubscribeEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExecutionDataAPI_GetRegisterValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionDataAPIServer).GetRegisterValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.ExecutionDataAPI/GetRegisterValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionDataAPIServer).GetRegisterValues(ctx, req.(*GetRegisterValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecutionDataAPI_ServiceDesc is the grpc.ServiceDesc for ExecutionDataAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecutionDataAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.access.ExecutionDataAPI",
	HandlerType: (*ExecutionDataAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExecutionDataByBlockID",
			Handler:    _ExecutionDataAPI_GetExecutionDataByBlockID_Handler,
		},
		{
			MethodName: "GetRegisterValues",
			Handler:    _ExecutionDataAPI_GetRegisterValues_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeExecutionData",
			Handler:       _ExecutionDataAPI_SubscribeExecutionData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeEvents",
			Handler:       _ExecutionDataAPI_SubscribeEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flow/executiondata/executiondata.proto",
}
