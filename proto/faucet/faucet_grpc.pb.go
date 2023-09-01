// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: faucet/faucet.proto

package faucet

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
	Faucet_Request_FullMethodName = "/faucet.faucet/Request"
)

// FaucetClient is the client API for Faucet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaucetClient interface {
	Request(ctx context.Context, in *FaucetRequest, opts ...grpc.CallOption) (*FaucetResponse, error)
}

type faucetClient struct {
	cc grpc.ClientConnInterface
}

func NewFaucetClient(cc grpc.ClientConnInterface) FaucetClient {
	return &faucetClient{cc}
}

func (c *faucetClient) Request(ctx context.Context, in *FaucetRequest, opts ...grpc.CallOption) (*FaucetResponse, error) {
	out := new(FaucetResponse)
	err := c.cc.Invoke(ctx, Faucet_Request_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaucetServer is the server API for Faucet service.
// All implementations must embed UnimplementedFaucetServer
// for forward compatibility
type FaucetServer interface {
	Request(context.Context, *FaucetRequest) (*FaucetResponse, error)
	mustEmbedUnimplementedFaucetServer()
}

// UnimplementedFaucetServer must be embedded to have forward compatible implementations.
type UnimplementedFaucetServer struct {
}

func (UnimplementedFaucetServer) Request(context.Context, *FaucetRequest) (*FaucetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedFaucetServer) mustEmbedUnimplementedFaucetServer() {}

// UnsafeFaucetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaucetServer will
// result in compilation errors.
type UnsafeFaucetServer interface {
	mustEmbedUnimplementedFaucetServer()
}

func RegisterFaucetServer(s grpc.ServiceRegistrar, srv FaucetServer) {
	s.RegisterService(&Faucet_ServiceDesc, srv)
}

func _Faucet_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FaucetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Faucet_Request_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServer).Request(ctx, req.(*FaucetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Faucet_ServiceDesc is the grpc.ServiceDesc for Faucet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Faucet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "faucet.faucet",
	HandlerType: (*FaucetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Faucet_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faucet/faucet.proto",
}