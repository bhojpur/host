// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostServiceClient interface {
	// StartLocalInstance starts an Instance on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the host/config.yaml
	//   3. all bytes constituting the Instance YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalInstance(ctx context.Context, opts ...grpc.CallOption) (HostService_StartLocalInstanceClient, error)
	// StartFromPreviousInstance starts a new Instance based on a previous one.
	// If the previous Instance does not have the can-replay condition set this call will result in an error.
	StartFromPreviousInstance(ctx context.Context, in *StartFromPreviousInstanceRequest, opts ...grpc.CallOption) (*StartInstanceResponse, error)
	// StartInstanceRequest starts a new Instance based on its specification.
	StartInstance(ctx context.Context, in *StartInstanceRequest, opts ...grpc.CallOption) (*StartInstanceResponse, error)
	// Searches for Instance(s) known to this instance
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Subscribe listens to new Instance(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (HostService_SubscribeClient, error)
	// GetInstance retrieves details of a single Instance
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error)
	// Listen listens to Instance updates and log output of a running Instance
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (HostService_ListenClient, error)
	// StopInstance stops a currently running Instance
	StopInstance(ctx context.Context, in *StopInstanceRequest, opts ...grpc.CallOption) (*StopInstanceResponse, error)
}

type hostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostServiceClient(cc grpc.ClientConnInterface) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) StartLocalInstance(ctx context.Context, opts ...grpc.CallOption) (HostService_StartLocalInstanceClient, error) {
	stream, err := c.cc.NewStream(ctx, &HostService_ServiceDesc.Streams[0], "/v1.HostService/StartLocalInstance", opts...)
	if err != nil {
		return nil, err
	}
	x := &hostServiceStartLocalInstanceClient{stream}
	return x, nil
}

type HostService_StartLocalInstanceClient interface {
	Send(*StartLocalInstanceRequest) error
	CloseAndRecv() (*StartInstanceResponse, error)
	grpc.ClientStream
}

type hostServiceStartLocalInstanceClient struct {
	grpc.ClientStream
}

func (x *hostServiceStartLocalInstanceClient) Send(m *StartLocalInstanceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hostServiceStartLocalInstanceClient) CloseAndRecv() (*StartInstanceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartInstanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hostServiceClient) StartFromPreviousInstance(ctx context.Context, in *StartFromPreviousInstanceRequest, opts ...grpc.CallOption) (*StartInstanceResponse, error) {
	out := new(StartInstanceResponse)
	err := c.cc.Invoke(ctx, "/v1.HostService/StartFromPreviousInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) StartInstance(ctx context.Context, in *StartInstanceRequest, opts ...grpc.CallOption) (*StartInstanceResponse, error) {
	out := new(StartInstanceResponse)
	err := c.cc.Invoke(ctx, "/v1.HostService/StartInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/v1.HostService/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (HostService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &HostService_ServiceDesc.Streams[1], "/v1.HostService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &hostServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HostService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type hostServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *hostServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hostServiceClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error) {
	out := new(GetInstanceResponse)
	err := c.cc.Invoke(ctx, "/v1.HostService/GetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (HostService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &HostService_ServiceDesc.Streams[2], "/v1.HostService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &hostServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HostService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type hostServiceListenClient struct {
	grpc.ClientStream
}

func (x *hostServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hostServiceClient) StopInstance(ctx context.Context, in *StopInstanceRequest, opts ...grpc.CallOption) (*StopInstanceResponse, error) {
	out := new(StopInstanceResponse)
	err := c.cc.Invoke(ctx, "/v1.HostService/StopInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
// All implementations must embed UnimplementedHostServiceServer
// for forward compatibility
type HostServiceServer interface {
	// StartLocalInstance starts an Instance on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the host/config.yaml
	//   3. all bytes constituting the Instance YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalInstance(HostService_StartLocalInstanceServer) error
	// StartFromPreviousInstance starts a new Instance based on a previous one.
	// If the previous Instance does not have the can-replay condition set this call will result in an error.
	StartFromPreviousInstance(context.Context, *StartFromPreviousInstanceRequest) (*StartInstanceResponse, error)
	// StartInstanceRequest starts a new Instance based on its specification.
	StartInstance(context.Context, *StartInstanceRequest) (*StartInstanceResponse, error)
	// Searches for Instance(s) known to this instance
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Subscribe listens to new Instance(s) updates
	Subscribe(*SubscribeRequest, HostService_SubscribeServer) error
	// GetInstance retrieves details of a single Instance
	GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error)
	// Listen listens to Instance updates and log output of a running Instance
	Listen(*ListenRequest, HostService_ListenServer) error
	// StopInstance stops a currently running Instance
	StopInstance(context.Context, *StopInstanceRequest) (*StopInstanceResponse, error)
	mustEmbedUnimplementedHostServiceServer()
}

// UnimplementedHostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHostServiceServer struct {
}

func (UnimplementedHostServiceServer) StartLocalInstance(HostService_StartLocalInstanceServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalInstance not implemented")
}
func (UnimplementedHostServiceServer) StartFromPreviousInstance(context.Context, *StartFromPreviousInstanceRequest) (*StartInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousInstance not implemented")
}
func (UnimplementedHostServiceServer) StartInstance(context.Context, *StartInstanceRequest) (*StartInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartInstance not implemented")
}
func (UnimplementedHostServiceServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedHostServiceServer) Subscribe(*SubscribeRequest, HostService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedHostServiceServer) GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedHostServiceServer) Listen(*ListenRequest, HostService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedHostServiceServer) StopInstance(context.Context, *StopInstanceRequest) (*StopInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopInstance not implemented")
}
func (UnimplementedHostServiceServer) mustEmbedUnimplementedHostServiceServer() {}

// UnsafeHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostServiceServer will
// result in compilation errors.
type UnsafeHostServiceServer interface {
	mustEmbedUnimplementedHostServiceServer()
}

func RegisterHostServiceServer(s grpc.ServiceRegistrar, srv HostServiceServer) {
	s.RegisterService(&HostService_ServiceDesc, srv)
}

func _HostService_StartLocalInstance_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HostServiceServer).StartLocalInstance(&hostServiceStartLocalInstanceServer{stream})
}

type HostService_StartLocalInstanceServer interface {
	SendAndClose(*StartInstanceResponse) error
	Recv() (*StartLocalInstanceRequest, error)
	grpc.ServerStream
}

type hostServiceStartLocalInstanceServer struct {
	grpc.ServerStream
}

func (x *hostServiceStartLocalInstanceServer) SendAndClose(m *StartInstanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hostServiceStartLocalInstanceServer) Recv() (*StartLocalInstanceRequest, error) {
	m := new(StartLocalInstanceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HostService_StartFromPreviousInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).StartFromPreviousInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HostService/StartFromPreviousInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).StartFromPreviousInstance(ctx, req.(*StartFromPreviousInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_StartInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).StartInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HostService/StartInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).StartInstance(ctx, req.(*StartInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HostService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HostServiceServer).Subscribe(m, &hostServiceSubscribeServer{stream})
}

type HostService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type hostServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *hostServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HostService_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HostService/GetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HostServiceServer).Listen(m, &hostServiceListenServer{stream})
}

type HostService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type hostServiceListenServer struct {
	grpc.ServerStream
}

func (x *hostServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HostService_StopInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).StopInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HostService/StopInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).StopInstance(ctx, req.(*StopInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostService_ServiceDesc is the grpc.ServiceDesc for HostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousInstance",
			Handler:    _HostService_StartFromPreviousInstance_Handler,
		},
		{
			MethodName: "StartInstance",
			Handler:    _HostService_StartInstance_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _HostService_ListInstances_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _HostService_GetInstance_Handler,
		},
		{
			MethodName: "StopInstance",
			Handler:    _HostService_StopInstance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalInstance",
			Handler:       _HostService_StartLocalInstance_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _HostService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _HostService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "host.proto",
}
