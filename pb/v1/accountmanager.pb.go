// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accountmanager.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UnlockAccountResponse struct {
	State                ResponseState `protobuf:"varint,1,opt,name=state,proto3,enum=v1.ResponseState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UnlockAccountResponse) Reset()         { *m = UnlockAccountResponse{} }
func (m *UnlockAccountResponse) String() string { return proto.CompactTextString(m) }
func (*UnlockAccountResponse) ProtoMessage()    {}
func (*UnlockAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba6db06a1aaf83a, []int{2}
}

func (m *UnlockAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockAccountResponse.Unmarshal(m, b)
}
func (m *UnlockAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockAccountResponse.Marshal(b, m, deterministic)
}
func (m *UnlockAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockAccountResponse.Merge(m, src)
}
func (m *UnlockAccountResponse) XXX_Size() int {
	return xxx_messageInfo_UnlockAccountResponse.Size(m)
}
func (m *UnlockAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockAccountResponse proto.InternalMessageInfo

func (m *UnlockAccountResponse) GetState() ResponseState {
	if m != nil {
		return m.State
	}
	return ResponseState_UNKNOWN
}

type LockAccountResponse struct {
	State                ResponseState `protobuf:"varint,1,opt,name=state,proto3,enum=v1.ResponseState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LockAccountResponse) Reset()         { *m = LockAccountResponse{} }
func (m *LockAccountResponse) String() string { return proto.CompactTextString(m) }
func (*LockAccountResponse) ProtoMessage()    {}
func (*LockAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba6db06a1aaf83a, []int{3}
}

func (m *LockAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockAccountResponse.Unmarshal(m, b)
}
func (m *LockAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockAccountResponse.Marshal(b, m, deterministic)
}
func (m *LockAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockAccountResponse.Merge(m, src)
}
func (m *LockAccountResponse) XXX_Size() int {
	return xxx_messageInfo_LockAccountResponse.Size(m)
}
func (m *LockAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockAccountResponse proto.InternalMessageInfo

func (m *LockAccountResponse) GetState() ResponseState {
	if m != nil {
		return m.State
	}
	return ResponseState_UNKNOWN
}

func init() {
	proto.RegisterType((*UnlockAccountRequest)(nil), "v1.UnlockAccountRequest")
	proto.RegisterType((*LockAccountRequest)(nil), "v1.LockAccountRequest")
	proto.RegisterType((*UnlockAccountResponse)(nil), "v1.UnlockAccountResponse")
	proto.RegisterType((*LockAccountResponse)(nil), "v1.LockAccountResponse")
}

func init() {
	proto.RegisterFile("accountmanager.proto", fileDescriptor_aba6db06a1aaf83a)
}

var fileDescriptor_aba6db06a1aaf83a = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0xf8, 0xff, 0x8a, 0x83, 0x16, 0x9c, 0x56, 0x9b, 0x46, 0xa9, 0x35, 0x1b, 0x8b,
	0xd0, 0x19, 0x12, 0xf7, 0xa2, 0x05, 0x77, 0x15, 0x4a, 0x8a, 0x20, 0xe2, 0xc2, 0x69, 0x1c, 0x92,
	0x60, 0x3b, 0x13, 0x33, 0x93, 0x71, 0xaf, 0x8f, 0xe0, 0x5b, 0xf8, 0x28, 0xe2, 0xce, 0x57, 0xf0,
	0x41, 0x64, 0x32, 0xad, 0xb4, 0xb6, 0x1b, 0x71, 0x79, 0xcf, 0xc9, 0xf9, 0x72, 0x72, 0x73, 0x41,
	0x9d, 0x44, 0x11, 0x2f, 0x98, 0x9c, 0x10, 0x46, 0x62, 0x9a, 0xa3, 0x2c, 0xe7, 0x92, 0x43, 0x5b,
	0xf9, 0xee, 0x5e, 0xcc, 0x79, 0x3c, 0xa6, 0x98, 0x64, 0x29, 0x26, 0x8c, 0x71, 0x49, 0x64, 0xca,
	0x99, 0x30, 0x4f, 0xb8, 0xb5, 0x9c, 0x8a, 0x8c, 0x33, 0x41, 0x85, 0x24, 0x92, 0x1a, 0xd1, 0x1b,
	0x80, 0xfa, 0x25, 0x1b, 0xf3, 0xe8, 0xfe, 0xcc, 0x40, 0x43, 0xfa, 0x50, 0x50, 0x21, 0xa1, 0x03,
	0xd6, 0xa6, 0xaf, 0x71, 0xac, 0xb6, 0xd5, 0x59, 0x0f, 0x67, 0x23, 0x6c, 0x01, 0x90, 0x11, 0x21,
	0xb2, 0x24, 0x27, 0x82, 0x3a, 0x76, 0xdb, 0xea, 0x6c, 0x84, 0x73, 0x8a, 0x87, 0x00, 0xec, 0xff,
	0x82, 0xe7, 0x9d, 0x82, 0xed, 0x1f, 0x0d, 0x4c, 0x4b, 0x78, 0x08, 0xfe, 0x97, 0x4d, 0xcb, 0x40,
	0x35, 0xd8, 0x42, 0xca, 0x47, 0x33, 0x73, 0xa8, 0x8d, 0xd0, 0xf8, 0xde, 0x09, 0xa8, 0xf5, 0xff,
	0x90, 0x0f, 0xde, 0x2d, 0x50, 0x9d, 0x86, 0x2f, 0xcc, 0x4e, 0xe1, 0x2d, 0xa8, 0x98, 0x52, 0xd0,
	0xd1, 0xb1, 0x55, 0x2b, 0x72, 0x9b, 0x2b, 0x1c, 0x43, 0xf7, 0x0e, 0x9e, 0x3e, 0x3e, 0x5f, 0xec,
	0x5d, 0xd8, 0xc4, 0xca, 0xc7, 0x8b, 0xbf, 0x0b, 0x17, 0x86, 0x7b, 0x05, 0xfe, 0xe9, 0xd2, 0x70,
	0x47, 0x53, 0x96, 0x17, 0xe6, 0x36, 0x96, 0xf4, 0x29, 0x7b, 0xbf, 0x64, 0x37, 0x61, 0x63, 0x05,
	0x5b, 0x93, 0x7b, 0xcf, 0x16, 0x68, 0x45, 0x7c, 0x82, 0x1e, 0x29, 0x19, 0xdf, 0x49, 0x1a, 0x25,
	0x88, 0xca, 0x24, 0x10, 0x69, 0xcc, 0x68, 0x4e, 0xb2, 0x14, 0x29, 0xbf, 0x57, 0x5b, 0xfc, 0xdc,
	0x81, 0x3e, 0x85, 0x81, 0x75, 0x7d, 0x14, 0xa7, 0x32, 0x29, 0x46, 0x28, 0xe2, 0x13, 0xfc, 0x9d,
	0xc6, 0x3a, 0xdd, 0x35, 0xf1, 0xae, 0xbe, 0xac, 0x6c, 0x84, 0x95, 0xff, 0x6a, 0x6f, 0x9e, 0xcb,
	0x24, 0x18, 0x96, 0x32, 0x52, 0xfe, 0xdb, 0xfc, 0x7c, 0xa3, 0xfc, 0x51, 0xa5, 0xbc, 0xaf, 0xe3,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xa4, 0xcf, 0xf3, 0xae, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountManagerClient is the client API for AccountManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountManagerClient interface {
	Unlock(ctx context.Context, in *UnlockAccountRequest, opts ...grpc.CallOption) (*UnlockAccountResponse, error)
	Lock(ctx context.Context, in *LockAccountRequest, opts ...grpc.CallOption) (*LockAccountResponse, error)
}

type accountManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountManagerClient(cc grpc.ClientConnInterface) AccountManagerClient {
	return &accountManagerClient{cc}
}

func (c *accountManagerClient) Unlock(ctx context.Context, in *UnlockAccountRequest, opts ...grpc.CallOption) (*UnlockAccountResponse, error) {
	out := new(UnlockAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountManager/Unlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountManagerClient) Lock(ctx context.Context, in *LockAccountRequest, opts ...grpc.CallOption) (*LockAccountResponse, error) {
	out := new(LockAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountManager/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountManagerServer is the server API for AccountManager service.
type AccountManagerServer interface {
	Unlock(context.Context, *UnlockAccountRequest) (*UnlockAccountResponse, error)
	Lock(context.Context, *LockAccountRequest) (*LockAccountResponse, error)
}

// UnimplementedAccountManagerServer can be embedded to have forward compatible implementations.
type UnimplementedAccountManagerServer struct {
}

func (*UnimplementedAccountManagerServer) Unlock(ctx context.Context, req *UnlockAccountRequest) (*UnlockAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (*UnimplementedAccountManagerServer) Lock(ctx context.Context, req *LockAccountRequest) (*LockAccountResponse, error) {
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
