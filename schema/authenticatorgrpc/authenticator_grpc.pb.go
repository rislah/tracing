// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authenticatorgrpc

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

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticatorClient interface {
	AuthenticatePassword(ctx context.Context, in *AuthenticatePasswordReq, opts ...grpc.CallOption) (*AuthenticatePasswordResponse, error)
}

type authenticatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticatorClient(cc grpc.ClientConnInterface) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) AuthenticatePassword(ctx context.Context, in *AuthenticatePasswordReq, opts ...grpc.CallOption) (*AuthenticatePasswordResponse, error) {
	out := new(AuthenticatePasswordResponse)
	err := c.cc.Invoke(ctx, "/toy.Authenticator/AuthenticatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
// All implementations must embed UnimplementedAuthenticatorServer
// for forward compatibility
type AuthenticatorServer interface {
	AuthenticatePassword(context.Context, *AuthenticatePasswordReq) (*AuthenticatePasswordResponse, error)
	mustEmbedUnimplementedAuthenticatorServer()
}

// UnimplementedAuthenticatorServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorServer struct {
}

func (UnimplementedAuthenticatorServer) AuthenticatePassword(context.Context, *AuthenticatePasswordReq) (*AuthenticatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticatePassword not implemented")
}
func (UnimplementedAuthenticatorServer) mustEmbedUnimplementedAuthenticatorServer() {}

// UnsafeAuthenticatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticatorServer will
// result in compilation errors.
type UnsafeAuthenticatorServer interface {
	mustEmbedUnimplementedAuthenticatorServer()
}

func RegisterAuthenticatorServer(s grpc.ServiceRegistrar, srv AuthenticatorServer) {
	s.RegisterService(&Authenticator_ServiceDesc, srv)
}

func _Authenticator_AuthenticatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).AuthenticatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toy.Authenticator/AuthenticatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).AuthenticatePassword(ctx, req.(*AuthenticatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Authenticator_ServiceDesc is the grpc.ServiceDesc for Authenticator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authenticator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "toy.Authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticatePassword",
			Handler:    _Authenticator_AuthenticatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/authenticator.proto",
}
