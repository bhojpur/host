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

// HostUIClient is the client API for HostUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostUIClient interface {
	// ListInstanceSpecs returns a list of Host Instance(s) that can be started through the UI.
	ListInstanceSpecs(ctx context.Context, in *ListInstanceSpecsRequest, opts ...grpc.CallOption) (HostUI_ListInstanceSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type hostUIClient struct {
	cc grpc.ClientConnInterface
}

func NewHostUIClient(cc grpc.ClientConnInterface) HostUIClient {
	return &hostUIClient{cc}
}

func (c *hostUIClient) ListInstanceSpecs(ctx context.Context, in *ListInstanceSpecsRequest, opts ...grpc.CallOption) (HostUI_ListInstanceSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &HostUI_ServiceDesc.Streams[0], "/v1.HostUI/ListInstanceSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &hostUIListInstanceSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HostUI_ListInstanceSpecsClient interface {
	Recv() (*ListInstanceSpecsResponse, error)
	grpc.ClientStream
}

type hostUIListInstanceSpecsClient struct {
	grpc.ClientStream
}

func (x *hostUIListInstanceSpecsClient) Recv() (*ListInstanceSpecsResponse, error) {
	m := new(ListInstanceSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hostUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.HostUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostUIServer is the server API for HostUI service.
// All implementations must embed UnimplementedHostUIServer
// for forward compatibility
type HostUIServer interface {
	// ListInstanceSpecs returns a list of Host Instance(s) that can be started through the UI.
	ListInstanceSpecs(*ListInstanceSpecsRequest, HostUI_ListInstanceSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedHostUIServer()
}

// UnimplementedHostUIServer must be embedded to have forward compatible implementations.
type UnimplementedHostUIServer struct {
}

func (UnimplementedHostUIServer) ListInstanceSpecs(*ListInstanceSpecsRequest, HostUI_ListInstanceSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListInstanceSpecs not implemented")
}
func (UnimplementedHostUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedHostUIServer) mustEmbedUnimplementedHostUIServer() {}

// UnsafeHostUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostUIServer will
// result in compilation errors.
type UnsafeHostUIServer interface {
	mustEmbedUnimplementedHostUIServer()
}

func RegisterHostUIServer(s grpc.ServiceRegistrar, srv HostUIServer) {
	s.RegisterService(&HostUI_ServiceDesc, srv)
}

func _HostUI_ListInstanceSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListInstanceSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HostUIServer).ListInstanceSpecs(m, &hostUIListInstanceSpecsServer{stream})
}

type HostUI_ListInstanceSpecsServer interface {
	Send(*ListInstanceSpecsResponse) error
	grpc.ServerStream
}

type hostUIListInstanceSpecsServer struct {
	grpc.ServerStream
}

func (x *hostUIListInstanceSpecsServer) Send(m *ListInstanceSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HostUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HostUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostUI_ServiceDesc is the grpc.ServiceDesc for HostUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.HostUI",
	HandlerType: (*HostUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _HostUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListInstanceSpecs",
			Handler:       _HostUI_ListInstanceSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "host-ui.proto",
}
