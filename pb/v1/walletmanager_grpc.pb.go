// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: walletmanager.proto

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

const (
	WalletManager_Unlock_FullMethodName = "/v1.WalletManager/Unlock"
	WalletManager_Lock_FullMethodName   = "/v1.WalletManager/Lock"
)

// WalletManagerClient is the client API for WalletManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletManagerClient interface {
	Unlock(ctx context.Context, in *UnlockWalletRequest, opts ...grpc.CallOption) (*UnlockWalletResponse, error)
	Lock(ctx context.Context, in *LockWalletRequest, opts ...grpc.CallOption) (*LockWalletResponse, error)
}

type walletManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletManagerClient(cc grpc.ClientConnInterface) WalletManagerClient {
	return &walletManagerClient{cc}
}

func (c *walletManagerClient) Unlock(ctx context.Context, in *UnlockWalletRequest, opts ...grpc.CallOption) (*UnlockWalletResponse, error) {
	out := new(UnlockWalletResponse)
	err := c.cc.Invoke(ctx, WalletManager_Unlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletManagerClient) Lock(ctx context.Context, in *LockWalletRequest, opts ...grpc.CallOption) (*LockWalletResponse, error) {
	out := new(LockWalletResponse)
	err := c.cc.Invoke(ctx, WalletManager_Lock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerServer is the server API for WalletManager service.
// All implementations must embed UnimplementedWalletManagerServer
// for forward compatibility
type WalletManagerServer interface {
	Unlock(context.Context, *UnlockWalletRequest) (*UnlockWalletResponse, error)
	Lock(context.Context, *LockWalletRequest) (*LockWalletResponse, error)
	mustEmbedUnimplementedWalletManagerServer()
}

// UnimplementedWalletManagerServer must be embedded to have forward compatible implementations.
type UnimplementedWalletManagerServer struct {
}

func (UnimplementedWalletManagerServer) Unlock(context.Context, *UnlockWalletRequest) (*UnlockWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (UnimplementedWalletManagerServer) Lock(context.Context, *LockWalletRequest) (*LockWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedWalletManagerServer) mustEmbedUnimplementedWalletManagerServer() {}

// UnsafeWalletManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletManagerServer will
// result in compilation errors.
type UnsafeWalletManagerServer interface {
	mustEmbedUnimplementedWalletManagerServer()
}

func RegisterWalletManagerServer(s grpc.ServiceRegistrar, srv WalletManagerServer) {
	s.RegisterService(&WalletManager_ServiceDesc, srv)
}

func _WalletManager_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagerServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletManager_Unlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagerServer).Unlock(ctx, req.(*UnlockWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletManager_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagerServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletManager_Lock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagerServer).Lock(ctx, req.(*LockWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletManager_ServiceDesc is the grpc.ServiceDesc for WalletManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WalletManager",
	HandlerType: (*WalletManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unlock",
			Handler:    _WalletManager_Unlock_Handler,
		},
		{
			MethodName: "Lock",
			Handler:    _WalletManager_Lock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "walletmanager.proto",
}
