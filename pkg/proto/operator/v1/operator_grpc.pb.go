// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package operator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorClient interface {
	// Sends events to Dapr sidecars upon component changes.
	ComponentUpdate(ctx context.Context, in *ComponentUpdateRequest, opts ...grpc.CallOption) (Operator_ComponentUpdateClient, error)
	// Returns a list of available components
	ListComponents(ctx context.Context, in *ListComponentsRequest, opts ...grpc.CallOption) (*ListComponentResponse, error)
	// Returns a given configuration by name
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	// Returns a list of pub/sub subscriptions
	ListSubscriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error)
	// Returns a list of pub/sub subscriptions, ListSubscriptionsRequest to expose pod info
	ListSubscriptionsV2(ctx context.Context, in *ListSubscriptionsRequest, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error)
}

type operatorClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorClient(cc grpc.ClientConnInterface) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) ComponentUpdate(ctx context.Context, in *ComponentUpdateRequest, opts ...grpc.CallOption) (Operator_ComponentUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Operator_ServiceDesc.Streams[0], "/dapr.proto.operator.v1.Operator/ComponentUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &operatorComponentUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Operator_ComponentUpdateClient interface {
	Recv() (*ComponentUpdateEvent, error)
	grpc.ClientStream
}

type operatorComponentUpdateClient struct {
	grpc.ClientStream
}

func (x *operatorComponentUpdateClient) Recv() (*ComponentUpdateEvent, error) {
	m := new(ComponentUpdateEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *operatorClient) ListComponents(ctx context.Context, in *ListComponentsRequest, opts ...grpc.CallOption) (*ListComponentResponse, error) {
	out := new(ListComponentResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/ListComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ListSubscriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error) {
	out := new(ListSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/ListSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ListSubscriptionsV2(ctx context.Context, in *ListSubscriptionsRequest, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error) {
	out := new(ListSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/ListSubscriptionsV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
// All implementations should embed UnimplementedOperatorServer
// for forward compatibility
type OperatorServer interface {
	// Sends events to Dapr sidecars upon component changes.
	ComponentUpdate(*ComponentUpdateRequest, Operator_ComponentUpdateServer) error
	// Returns a list of available components
	ListComponents(context.Context, *ListComponentsRequest) (*ListComponentResponse, error)
	// Returns a given configuration by name
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	// Returns a list of pub/sub subscriptions
	ListSubscriptions(context.Context, *emptypb.Empty) (*ListSubscriptionsResponse, error)
	// Returns a list of pub/sub subscriptions, ListSubscriptionsRequest to expose pod info
	ListSubscriptionsV2(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)
}

// UnimplementedOperatorServer should be embedded to have forward compatible implementations.
type UnimplementedOperatorServer struct {
}

func (UnimplementedOperatorServer) ComponentUpdate(*ComponentUpdateRequest, Operator_ComponentUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method ComponentUpdate not implemented")
}
func (UnimplementedOperatorServer) ListComponents(context.Context, *ListComponentsRequest) (*ListComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComponents not implemented")
}
func (UnimplementedOperatorServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedOperatorServer) ListSubscriptions(context.Context, *emptypb.Empty) (*ListSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}
func (UnimplementedOperatorServer) ListSubscriptionsV2(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptionsV2 not implemented")
}

// UnsafeOperatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServer will
// result in compilation errors.
type UnsafeOperatorServer interface {
	mustEmbedUnimplementedOperatorServer()
}

func RegisterOperatorServer(s grpc.ServiceRegistrar, srv OperatorServer) {
	s.RegisterService(&Operator_ServiceDesc, srv)
}

func _Operator_ComponentUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ComponentUpdateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperatorServer).ComponentUpdate(m, &operatorComponentUpdateServer{stream})
}

type Operator_ComponentUpdateServer interface {
	Send(*ComponentUpdateEvent) error
	grpc.ServerStream
}

type operatorComponentUpdateServer struct {
	grpc.ServerStream
}

func (x *operatorComponentUpdateServer) Send(m *ComponentUpdateEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Operator_ListComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComponentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/ListComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListComponents(ctx, req.(*ListComponentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ListSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/ListSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListSubscriptions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ListSubscriptionsV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListSubscriptionsV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/ListSubscriptionsV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListSubscriptionsV2(ctx, req.(*ListSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operator_ServiceDesc is the grpc.ServiceDesc for Operator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.operator.v1.Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListComponents",
			Handler:    _Operator_ListComponents_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _Operator_GetConfiguration_Handler,
		},
		{
			MethodName: "ListSubscriptions",
			Handler:    _Operator_ListSubscriptions_Handler,
		},
		{
			MethodName: "ListSubscriptionsV2",
			Handler:    _Operator_ListSubscriptionsV2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComponentUpdate",
			Handler:       _Operator_ComponentUpdate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dapr/proto/operator/v1/operator.proto",
}
