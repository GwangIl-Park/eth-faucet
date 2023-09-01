// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: faucetToken/faucetToken.proto

package faucetToken

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

const (
	FaucetToken_RequestToken_FullMethodName = "/faucetToken.faucetToken/RequestToken"
)

// FaucetTokenClient is the client API for FaucetToken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaucetTokenClient interface {
	RequestToken(ctx context.Context, in *FaucetTokenRequest, opts ...grpc.CallOption) (*FaucetTokenResponse, error)
}

type faucetTokenClient struct {
	cc grpc.ClientConnInterface
}

func NewFaucetTokenClient(cc grpc.ClientConnInterface) FaucetTokenClient {
	return &faucetTokenClient{cc}
}

func (c *faucetTokenClient) RequestToken(ctx context.Context, in *FaucetTokenRequest, opts ...grpc.CallOption) (*FaucetTokenResponse, error) {
	out := new(FaucetTokenResponse)
	err := c.cc.Invoke(ctx, FaucetToken_RequestToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaucetTokenServer is the server API for FaucetToken service.
// All implementations must embed UnimplementedFaucetTokenServer
// for forward compatibility
type FaucetTokenServer interface {
	RequestToken(context.Context, *FaucetTokenRequest) (*FaucetTokenResponse, error)
	mustEmbedUnimplementedFaucetTokenServer()
}

// UnimplementedFaucetTokenServer must be embedded to have forward compatible implementations.
type UnimplementedFaucetTokenServer struct {
}

func (UnimplementedFaucetTokenServer) RequestToken(context.Context, *FaucetTokenRequest) (*FaucetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToken not implemented")
}
func (UnimplementedFaucetTokenServer) mustEmbedUnimplementedFaucetTokenServer() {}

// UnsafeFaucetTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaucetTokenServer will
// result in compilation errors.
type UnsafeFaucetTokenServer interface {
	mustEmbedUnimplementedFaucetTokenServer()
}

func RegisterFaucetTokenServer(s grpc.ServiceRegistrar, srv FaucetTokenServer) {
	s.RegisterService(&FaucetToken_ServiceDesc, srv)
}

func _FaucetToken_RequestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FaucetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetTokenServer).RequestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FaucetToken_RequestToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetTokenServer).RequestToken(ctx, req.(*FaucetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FaucetToken_ServiceDesc is the grpc.ServiceDesc for FaucetToken service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FaucetToken_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "faucetToken.faucetToken",
	HandlerType: (*FaucetTokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestToken",
			Handler:    _FaucetToken_RequestToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faucetToken/faucetToken.proto",
}
