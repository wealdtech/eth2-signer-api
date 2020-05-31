// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dkg.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type PrepareRequest struct {
	// account is the name of the account.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// threshold is the number of participants required to generate a valid signature.
	Threshold uint64 `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// participants contains the IDs of all participants.
	Participants         []uint64 `protobuf:"varint,3,rep,packed,name=participants,proto3" json:"participants,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareRequest) Reset()         { *m = PrepareRequest{} }
func (m *PrepareRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareRequest) ProtoMessage()    {}
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{0}
}

func (m *PrepareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareRequest.Unmarshal(m, b)
}
func (m *PrepareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareRequest.Marshal(b, m, deterministic)
}
func (m *PrepareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareRequest.Merge(m, src)
}
func (m *PrepareRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareRequest.Size(m)
}
func (m *PrepareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareRequest proto.InternalMessageInfo

func (m *PrepareRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *PrepareRequest) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PrepareRequest) GetParticipants() []uint64 {
	if m != nil {
		return m.Participants
	}
	return nil
}

type ExecuteRequest struct {
	// account is the name of the account.
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{1}
}

func (m *ExecuteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteRequest.Unmarshal(m, b)
}
func (m *ExecuteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteRequest.Merge(m, src)
}
func (m *ExecuteRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteRequest.Size(m)
}
func (m *ExecuteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteRequest proto.InternalMessageInfo

func (m *ExecuteRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type CommitRequest struct {
	// account is the name of the account.
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitRequest) Reset()         { *m = CommitRequest{} }
func (m *CommitRequest) String() string { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()    {}
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{2}
}

func (m *CommitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitRequest.Unmarshal(m, b)
}
func (m *CommitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitRequest.Marshal(b, m, deterministic)
}
func (m *CommitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitRequest.Merge(m, src)
}
func (m *CommitRequest) XXX_Size() int {
	return xxx_messageInfo_CommitRequest.Size(m)
}
func (m *CommitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitRequest proto.InternalMessageInfo

func (m *CommitRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type CommitResponse struct {
	// public_key is the key generated by the process.
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// confirmation_signature is the signature generated by the individual secret key.
	ConfirmationSignature []byte   `protobuf:"bytes,2,opt,name=confirmation_signature,json=confirmationSignature,proto3" json:"confirmation_signature,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *CommitResponse) Reset()         { *m = CommitResponse{} }
func (m *CommitResponse) String() string { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()    {}
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{3}
}

func (m *CommitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitResponse.Unmarshal(m, b)
}
func (m *CommitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitResponse.Marshal(b, m, deterministic)
}
func (m *CommitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitResponse.Merge(m, src)
}
func (m *CommitResponse) XXX_Size() int {
	return xxx_messageInfo_CommitResponse.Size(m)
}
func (m *CommitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitResponse proto.InternalMessageInfo

func (m *CommitResponse) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *CommitResponse) GetConfirmationSignature() []byte {
	if m != nil {
		return m.ConfirmationSignature
	}
	return nil
}

type AbortRequest struct {
	// account is the name of the account.
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AbortRequest) Reset()         { *m = AbortRequest{} }
func (m *AbortRequest) String() string { return proto.CompactTextString(m) }
func (*AbortRequest) ProtoMessage()    {}
func (*AbortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{4}
}

func (m *AbortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AbortRequest.Unmarshal(m, b)
}
func (m *AbortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AbortRequest.Marshal(b, m, deterministic)
}
func (m *AbortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AbortRequest.Merge(m, src)
}
func (m *AbortRequest) XXX_Size() int {
	return xxx_messageInfo_AbortRequest.Size(m)
}
func (m *AbortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AbortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AbortRequest proto.InternalMessageInfo

func (m *AbortRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// ContributeRequest is sent by each part to all other parties with a contribution.
type ContributeRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Secret               []byte   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	VerificationVector   [][]byte `protobuf:"bytes,3,rep,name=verification_vector,json=verificationVector,proto3" json:"verification_vector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContributeRequest) Reset()         { *m = ContributeRequest{} }
func (m *ContributeRequest) String() string { return proto.CompactTextString(m) }
func (*ContributeRequest) ProtoMessage()    {}
func (*ContributeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{5}
}

func (m *ContributeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContributeRequest.Unmarshal(m, b)
}
func (m *ContributeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContributeRequest.Marshal(b, m, deterministic)
}
func (m *ContributeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContributeRequest.Merge(m, src)
}
func (m *ContributeRequest) XXX_Size() int {
	return xxx_messageInfo_ContributeRequest.Size(m)
}
func (m *ContributeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContributeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContributeRequest proto.InternalMessageInfo

func (m *ContributeRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *ContributeRequest) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *ContributeRequest) GetVerificationVector() [][]byte {
	if m != nil {
		return m.VerificationVector
	}
	return nil
}

// ContributeResponse provides the contribution from a participant.
type ContributeResponse struct {
	Secret               []byte   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	VerificationVector   [][]byte `protobuf:"bytes,2,rep,name=verification_vector,json=verificationVector,proto3" json:"verification_vector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContributeResponse) Reset()         { *m = ContributeResponse{} }
func (m *ContributeResponse) String() string { return proto.CompactTextString(m) }
func (*ContributeResponse) ProtoMessage()    {}
func (*ContributeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_584d508ed20653c0, []int{6}
}

func (m *ContributeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContributeResponse.Unmarshal(m, b)
}
func (m *ContributeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContributeResponse.Marshal(b, m, deterministic)
}
func (m *ContributeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContributeResponse.Merge(m, src)
}
func (m *ContributeResponse) XXX_Size() int {
	return xxx_messageInfo_ContributeResponse.Size(m)
}
func (m *ContributeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContributeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContributeResponse proto.InternalMessageInfo

func (m *ContributeResponse) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *ContributeResponse) GetVerificationVector() [][]byte {
	if m != nil {
		return m.VerificationVector
	}
	return nil
}

func init() {
	proto.RegisterType((*PrepareRequest)(nil), "v1.PrepareRequest")
	proto.RegisterType((*ExecuteRequest)(nil), "v1.ExecuteRequest")
	proto.RegisterType((*CommitRequest)(nil), "v1.CommitRequest")
	proto.RegisterType((*CommitResponse)(nil), "v1.CommitResponse")
	proto.RegisterType((*AbortRequest)(nil), "v1.AbortRequest")
	proto.RegisterType((*ContributeRequest)(nil), "v1.ContributeRequest")
	proto.RegisterType((*ContributeResponse)(nil), "v1.ContributeResponse")
}

func init() {
	proto.RegisterFile("dkg.proto", fileDescriptor_584d508ed20653c0)
}

var fileDescriptor_584d508ed20653c0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x8f, 0x93, 0x40,
	0x14, 0xb6, 0x54, 0xbb, 0xe1, 0x05, 0x1b, 0xf7, 0x99, 0x6d, 0x48, 0xd5, 0x64, 0xc3, 0x69, 0xf5,
	0x40, 0x53, 0x37, 0xc6, 0x93, 0x07, 0xb3, 0x6e, 0x3c, 0xf4, 0x62, 0x30, 0xf1, 0x66, 0x1a, 0x98,
	0x3e, 0xe8, 0xa4, 0xc0, 0xe0, 0x30, 0x10, 0xfb, 0x67, 0xfc, 0xad, 0x02, 0x03, 0x29, 0x34, 0xd9,
	0xb6, 0xc7, 0xf9, 0xbe, 0xf7, 0x7d, 0x6f, 0xde, 0x7b, 0x1f, 0x98, 0x9b, 0x5d, 0xe4, 0x66, 0x52,
	0x28, 0x81, 0x46, 0xb9, 0x9c, 0xbf, 0x89, 0x84, 0x88, 0x62, 0x5a, 0x34, 0x48, 0x50, 0x84, 0x0b,
	0x4a, 0x32, 0xb5, 0xd7, 0x05, 0x4e, 0x0c, 0xd3, 0x1f, 0x92, 0x32, 0x5f, 0x92, 0x47, 0x7f, 0x0a,
	0xca, 0x15, 0xda, 0x70, 0xe5, 0x33, 0x26, 0x8a, 0x54, 0xd9, 0xa3, 0xdb, 0xd1, 0x9d, 0xe9, 0x75,
	0x4f, 0x7c, 0x0b, 0xa6, 0xda, 0x4a, 0xca, 0xb7, 0x22, 0xde, 0xd8, 0x46, 0xc5, 0x3d, 0xf7, 0x0e,
	0x00, 0x3a, 0x60, 0x55, 0x36, 0x8a, 0x33, 0x9e, 0xf9, 0xa9, 0xca, 0xed, 0xf1, 0xed, 0xb8, 0x2a,
	0x18, 0x60, 0xce, 0x07, 0x98, 0x3e, 0xfe, 0x25, 0x56, 0xa8, 0xf3, 0xdd, 0x9c, 0xf7, 0xf0, 0xf2,
	0x41, 0x24, 0x09, 0x57, 0xe7, 0x4b, 0x43, 0x98, 0x76, 0xa5, 0x79, 0x26, 0xd2, 0x9c, 0xf0, 0x1d,
	0x40, 0x56, 0x04, 0x31, 0x67, 0xeb, 0x1d, 0xed, 0x9b, 0x72, 0xcb, 0x33, 0x35, 0xb2, 0xa2, 0x3d,
	0x7e, 0x82, 0x19, 0x13, 0x69, 0xc8, 0x65, 0xe2, 0x2b, 0x2e, 0xd2, 0x75, 0xce, 0xa3, 0xd4, 0x57,
	0x85, 0xa4, 0x66, 0x2c, 0xcb, 0xbb, 0xe9, 0xb3, 0x3f, 0x3b, 0xd2, 0xb9, 0x03, 0xeb, 0x6b, 0x20,
	0xe4, 0x05, 0x3f, 0x2a, 0xe1, 0xfa, 0x41, 0xa4, 0x4a, 0xf2, 0xe0, 0x92, 0x59, 0x71, 0x06, 0x93,
	0x9c, 0x98, 0x24, 0xd5, 0xf6, 0x6f, 0x5f, 0xb8, 0x80, 0xd7, 0x25, 0x49, 0x1e, 0x72, 0xa6, 0xff,
	0x59, 0x12, 0x53, 0x42, 0x36, 0xab, 0xb5, 0x3c, 0xec, 0x53, 0xbf, 0x1a, 0xc6, 0xf9, 0x0d, 0xd8,
	0xef, 0xdb, 0x6e, 0xe3, 0x60, 0x3f, 0xba, 0xc4, 0xde, 0x78, 0xca, 0xfe, 0xe3, 0x3f, 0x03, 0xc6,
	0xdf, 0x56, 0xdf, 0xf1, 0x33, 0x5c, 0xb5, 0xa9, 0x41, 0x74, 0xcb, 0xa5, 0x3b, 0x8c, 0xd0, 0x7c,
	0xe6, 0xea, 0xc8, 0xb9, 0x5d, 0xe4, 0xdc, 0xc7, 0x3a, 0x72, 0xce, 0xb3, 0x5a, 0xd8, 0x06, 0x40,
	0x0b, 0x87, 0x69, 0x38, 0x21, 0x5c, 0xc2, 0x44, 0x9f, 0x18, 0xaf, 0x6b, 0xdd, 0x20, 0x19, 0x73,
	0xec, 0x43, 0x7a, 0xe6, 0x4a, 0x72, 0x0f, 0x2f, 0x9a, 0x6b, 0xe1, 0xab, 0x9a, 0xee, 0x1f, 0xee,
	0x44, 0x9f, 0x2f, 0x00, 0x87, 0x05, 0xe2, 0x8d, 0x36, 0x3e, 0x3a, 0x64, 0x25, 0x3f, 0x82, 0xbb,
	0x9e, 0xc1, 0xa4, 0x31, 0xbc, 0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x89, 0x77, 0xe9, 0xcd, 0x83,
	0x03, 0x00, 0x00,
}
