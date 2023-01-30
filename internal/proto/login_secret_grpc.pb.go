// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: login_secret.proto

package proto

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	CreateLoginSecret(ctx context.Context, in *CreateLoginSecretRequest, opts ...grpc.CallOption) (*SecretResponse, error)
	UpdateLoginSecret(ctx context.Context, in *UpdateLoginSecretRequest, opts ...grpc.CallOption) (*SecretResponse, error)
	DeleteLoginSecret(ctx context.Context, in *DeleteLoginSecretRequest, opts ...grpc.CallOption) (*SecretResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) CreateLoginSecret(ctx context.Context, in *CreateLoginSecretRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/proto.LoginService/CreateLoginSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) UpdateLoginSecret(ctx context.Context, in *UpdateLoginSecretRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/proto.LoginService/UpdateLoginSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) DeleteLoginSecret(ctx context.Context, in *DeleteLoginSecretRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/proto.LoginService/DeleteLoginSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	CreateLoginSecret(context.Context, *CreateLoginSecretRequest) (*SecretResponse, error)
	UpdateLoginSecret(context.Context, *UpdateLoginSecretRequest) (*SecretResponse, error)
	DeleteLoginSecret(context.Context, *DeleteLoginSecretRequest) (*SecretResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) CreateLoginSecret(context.Context, *CreateLoginSecretRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoginSecret not implemented")
}
func (UnimplementedLoginServiceServer) UpdateLoginSecret(context.Context, *UpdateLoginSecretRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLoginSecret not implemented")
}
func (UnimplementedLoginServiceServer) DeleteLoginSecret(context.Context, *DeleteLoginSecretRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLoginSecret not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_CreateLoginSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoginSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CreateLoginSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LoginService/CreateLoginSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CreateLoginSecret(ctx, req.(*CreateLoginSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_UpdateLoginSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLoginSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).UpdateLoginSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LoginService/UpdateLoginSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).UpdateLoginSecret(ctx, req.(*UpdateLoginSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_DeleteLoginSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLoginSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).DeleteLoginSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LoginService/DeleteLoginSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).DeleteLoginSecret(ctx, req.(*DeleteLoginSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLoginSecret",
			Handler:    _LoginService_CreateLoginSecret_Handler,
		},
		{
			MethodName: "UpdateLoginSecret",
			Handler:    _LoginService_UpdateLoginSecret_Handler,
		},
		{
			MethodName: "DeleteLoginSecret",
			Handler:    _LoginService_DeleteLoginSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login_secret.proto",
}
