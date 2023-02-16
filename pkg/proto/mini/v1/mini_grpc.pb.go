// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: mini/v1/mini.proto

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

// MiniClient is the client API for Mini service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiniClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type miniClient struct {
	cc grpc.ClientConnInterface
}

func NewMiniClient(cc grpc.ClientConnInterface) MiniClient {
	return &miniClient{cc}
}

func (c *miniClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/v1.Mini/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiniServer is the server API for Mini service.
// All implementations must embed UnimplementedMiniServer
// for forward compatibility
type MiniServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedMiniServer()
}

// UnimplementedMiniServer must be embedded to have forward compatible implementations.
type UnimplementedMiniServer struct {
}

func (UnimplementedMiniServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedMiniServer) mustEmbedUnimplementedMiniServer() {}

// UnsafeMiniServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiniServer will
// result in compilation errors.
type UnsafeMiniServer interface {
	mustEmbedUnimplementedMiniServer()
}

func RegisterMiniServer(s grpc.ServiceRegistrar, srv MiniServer) {
	s.RegisterService(&Mini_ServiceDesc, srv)
}

func _Mini_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Mini/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mini_ServiceDesc is the grpc.ServiceDesc for Mini service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mini_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Mini",
	HandlerType: (*MiniServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _Mini_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mini/v1/mini.proto",
}