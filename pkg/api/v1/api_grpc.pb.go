// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api.proto

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

// AppleServiceClient is the client API for AppleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppleServiceClient interface {
	AppleCreate(ctx context.Context, in *AppleCreateRequest, opts ...grpc.CallOption) (*AppleCreateResponse, error)
	AppleGet(ctx context.Context, in *AppleGetRequest, opts ...grpc.CallOption) (*AppleGetResponse, error)
	AppleList(ctx context.Context, in *AppleListRequest, opts ...grpc.CallOption) (AppleService_AppleListClient, error)
	AppleUpdate(ctx context.Context, in *AppleUpdateRequest, opts ...grpc.CallOption) (*AppleUpdateResponse, error)
	AppleDelete(ctx context.Context, in *AppleDeleteRequest, opts ...grpc.CallOption) (*AppleDeleteResponse, error)
}

type appleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppleServiceClient(cc grpc.ClientConnInterface) AppleServiceClient {
	return &appleServiceClient{cc}
}

func (c *appleServiceClient) AppleCreate(ctx context.Context, in *AppleCreateRequest, opts ...grpc.CallOption) (*AppleCreateResponse, error) {
	out := new(AppleCreateResponse)
	err := c.cc.Invoke(ctx, "/api.v1.AppleService/AppleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appleServiceClient) AppleGet(ctx context.Context, in *AppleGetRequest, opts ...grpc.CallOption) (*AppleGetResponse, error) {
	out := new(AppleGetResponse)
	err := c.cc.Invoke(ctx, "/api.v1.AppleService/AppleGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appleServiceClient) AppleList(ctx context.Context, in *AppleListRequest, opts ...grpc.CallOption) (AppleService_AppleListClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppleService_ServiceDesc.Streams[0], "/api.v1.AppleService/AppleList", opts...)
	if err != nil {
		return nil, err
	}
	x := &appleServiceAppleListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AppleService_AppleListClient interface {
	Recv() (*AppleGetResponse, error)
	grpc.ClientStream
}

type appleServiceAppleListClient struct {
	grpc.ClientStream
}

func (x *appleServiceAppleListClient) Recv() (*AppleGetResponse, error) {
	m := new(AppleGetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appleServiceClient) AppleUpdate(ctx context.Context, in *AppleUpdateRequest, opts ...grpc.CallOption) (*AppleUpdateResponse, error) {
	out := new(AppleUpdateResponse)
	err := c.cc.Invoke(ctx, "/api.v1.AppleService/AppleUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appleServiceClient) AppleDelete(ctx context.Context, in *AppleDeleteRequest, opts ...grpc.CallOption) (*AppleDeleteResponse, error) {
	out := new(AppleDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.v1.AppleService/AppleDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppleServiceServer is the server API for AppleService service.
// All implementations must embed UnimplementedAppleServiceServer
// for forward compatibility
type AppleServiceServer interface {
	AppleCreate(context.Context, *AppleCreateRequest) (*AppleCreateResponse, error)
	AppleGet(context.Context, *AppleGetRequest) (*AppleGetResponse, error)
	AppleList(*AppleListRequest, AppleService_AppleListServer) error
	AppleUpdate(context.Context, *AppleUpdateRequest) (*AppleUpdateResponse, error)
	AppleDelete(context.Context, *AppleDeleteRequest) (*AppleDeleteResponse, error)
	mustEmbedUnimplementedAppleServiceServer()
}

// UnimplementedAppleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppleServiceServer struct {
}

func (UnimplementedAppleServiceServer) AppleCreate(context.Context, *AppleCreateRequest) (*AppleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppleCreate not implemented")
}
func (UnimplementedAppleServiceServer) AppleGet(context.Context, *AppleGetRequest) (*AppleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppleGet not implemented")
}
func (UnimplementedAppleServiceServer) AppleList(*AppleListRequest, AppleService_AppleListServer) error {
	return status.Errorf(codes.Unimplemented, "method AppleList not implemented")
}
func (UnimplementedAppleServiceServer) AppleUpdate(context.Context, *AppleUpdateRequest) (*AppleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppleUpdate not implemented")
}
func (UnimplementedAppleServiceServer) AppleDelete(context.Context, *AppleDeleteRequest) (*AppleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppleDelete not implemented")
}
func (UnimplementedAppleServiceServer) mustEmbedUnimplementedAppleServiceServer() {}

// UnsafeAppleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppleServiceServer will
// result in compilation errors.
type UnsafeAppleServiceServer interface {
	mustEmbedUnimplementedAppleServiceServer()
}

func RegisterAppleServiceServer(s grpc.ServiceRegistrar, srv AppleServiceServer) {
	s.RegisterService(&AppleService_ServiceDesc, srv)
}

func _AppleService_AppleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppleServiceServer).AppleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.AppleService/AppleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppleServiceServer).AppleCreate(ctx, req.(*AppleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppleService_AppleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppleServiceServer).AppleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.AppleService/AppleGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppleServiceServer).AppleGet(ctx, req.(*AppleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppleService_AppleList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AppleListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppleServiceServer).AppleList(m, &appleServiceAppleListServer{stream})
}

type AppleService_AppleListServer interface {
	Send(*AppleGetResponse) error
	grpc.ServerStream
}

type appleServiceAppleListServer struct {
	grpc.ServerStream
}

func (x *appleServiceAppleListServer) Send(m *AppleGetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AppleService_AppleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppleServiceServer).AppleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.AppleService/AppleUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppleServiceServer).AppleUpdate(ctx, req.(*AppleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppleService_AppleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppleServiceServer).AppleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.AppleService/AppleDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppleServiceServer).AppleDelete(ctx, req.(*AppleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppleService_ServiceDesc is the grpc.ServiceDesc for AppleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.AppleService",
	HandlerType: (*AppleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppleCreate",
			Handler:    _AppleService_AppleCreate_Handler,
		},
		{
			MethodName: "AppleGet",
			Handler:    _AppleService_AppleGet_Handler,
		},
		{
			MethodName: "AppleUpdate",
			Handler:    _AppleService_AppleUpdate_Handler,
		},
		{
			MethodName: "AppleDelete",
			Handler:    _AppleService_AppleDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AppleList",
			Handler:       _AppleService_AppleList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
