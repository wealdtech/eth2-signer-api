// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accountmanager.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UnlockAccountRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Passphrase           []byte   `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnlockAccountRequest) Reset()         { *m = UnlockAccountRequest{} }
func (m *UnlockAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UnlockAccountRequest) ProtoMessage()    {}
func (*UnlockAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba6db06a1aaf83a, []int{0}
}

func (m *UnlockAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockAccountRequest.Unmarshal(m, b)
}
func (m *UnlockAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockAccountRequest.Marshal(b, m, deterministic)
}
func (m *UnlockAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockAccountRequest.Merge(m, src)
}
func (m *UnlockAccountRequest) XXX_Size() int {
	return xxx_messageInfo_UnlockAccountRequest.Size(m)
}
func (m *UnlockAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockAccountRequest proto.InternalMessageInfo

func (m *UnlockAccountRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UnlockAccountRequest) GetPassphrase() []byte {
	if m != nil {
		return m.Passphrase
	}
	return nil
}

type LockAccountRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockAccountRequest) Reset()         { *m = LockAccountRequest{} }
func (m *LockAccountRequest) String() string { return proto.CompactTextString(m) }
func (*LockAccountRequest) ProtoMessage()    {}
func (*LockAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba6db06a1aaf83a, []int{1}
}

func (m *LockAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockAccountRequest.Unmarshal(m, b)
}
func (m *LockAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockAccountRequest.Marshal(b, m, deterministic)
}
func (m *LockAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockAccountRequest.Merge(m, src)
}
func (m *LockAccountRequest) XXX_Size() int {
	return xxx_messageInfo_LockAccountRequest.Size(m)
}
func (m *LockAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockAccountRequest proto.InternalMessageInfo

func (m *LockAccountRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func init() {
	proto.RegisterType((*UnlockAccountRequest)(nil), "v1.UnlockAccountRequest")
	proto.RegisterType((*LockAccountRequest)(nil), "v1.LockAccountRequest")
}

func init() { proto.RegisterFile("accountmanager.proto", fileDescriptor_aba6db06a1aaf83a) }

var fileDescriptor_aba6db06a1aaf83a = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xc9, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24,
	0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x14, 0x28, 0x05, 0x70, 0x89, 0x84, 0xe6,
	0xe5, 0xe4, 0x27, 0x67, 0x3b, 0x42, 0xb4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49,
	0x70, 0xb1, 0x43, 0x0d, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xe4, 0xb8,
	0xb8, 0x0a, 0x12, 0x8b, 0x8b, 0x0b, 0x32, 0x8a, 0x12, 0x8b, 0x53, 0x25, 0x98, 0x80, 0x92, 0x3c,
	0x41, 0x48, 0x22, 0x4a, 0x7a, 0x5c, 0x42, 0x3e, 0x24, 0x98, 0x67, 0xd4, 0xc1, 0xc8, 0xc5, 0x07,
	0x55, 0xec, 0x0b, 0x71, 0xbb, 0x90, 0x0d, 0x17, 0x1b, 0xc4, 0x51, 0x42, 0x12, 0x7a, 0x65, 0x86,
	0x7a, 0xd8, 0x1c, 0x28, 0x25, 0xa6, 0x07, 0xf1, 0x96, 0x1e, 0xcc, 0x5b, 0x7a, 0xae, 0x20, 0x6f,
	0x29, 0x31, 0x08, 0x59, 0x70, 0xb1, 0x80, 0x1c, 0x20, 0x24, 0x06, 0xd2, 0xeb, 0x43, 0x82, 0xce,
	0x24, 0x36, 0xb0, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x47, 0xfa, 0x56, 0xa9, 0x4c, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountManagerClient is the client API for AccountManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountManagerClient interface {
	Unlock(ctx context.Context, in *UnlockAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Lock(ctx context.Context, in *LockAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type accountManagerClient struct {
	cc *grpc.ClientConn
}

func NewAccountManagerClient(cc *grpc.ClientConn) AccountManagerClient {
	return &accountManagerClient{cc}
}

func (c *accountManagerClient) Unlock(ctx context.Context, in *UnlockAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.AccountManager/Unlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountManagerClient) Lock(ctx context.Context, in *LockAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.AccountManager/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountManagerServer is the server API for AccountManager service.
type AccountManagerServer interface {
	Unlock(context.Context, *UnlockAccountRequest) (*empty.Empty, error)
	Lock(context.Context, *LockAccountRequest) (*empty.Empty, error)
}

// UnimplementedAccountManagerServer can be embedded to have forward compatible implementations.
type UnimplementedAccountManagerServer struct {
}

func (*UnimplementedAccountManagerServer) Unlock(ctx context.Context, req *UnlockAccountRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (*UnimplementedAccountManagerServer) Lock(ctx context.Context, req *LockAccountRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}

func RegisterAccountManagerServer(s *grpc.Server, srv AccountManagerServer) {
	s.RegisterService(&_AccountManager_serviceDesc, srv)
}

func _AccountManager_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountManager/Unlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).Unlock(ctx, req.(*UnlockAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountManager_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountManagerServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountManager/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountManagerServer).Lock(ctx, req.(*LockAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountManager",
	HandlerType: (*AccountManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unlock",
			Handler:    _AccountManager_Unlock_Handler,
		},
		{
			MethodName: "Lock",
			Handler:    _AccountManager_Lock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accountmanager.proto",
}