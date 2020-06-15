// Code generated by protoc-gen-go. DO NOT EDIT.
// source: endpoint.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Endpoint struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_e90a83b3c7ac16f6, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "v1.Endpoint")
}

func init() {
	proto.RegisterFile("endpoint.proto", fileDescriptor_e90a83b3c7ac16f6)
}

var fileDescriptor_e90a83b3c7ac16f6 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x4b, 0x29,
	0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54, 0x72,
	0xe2, 0xe2, 0x70, 0x85, 0x8a, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x04, 0x01, 0x59, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x40, 0x11,
	0xce, 0x20, 0x30, 0x1b, 0x24, 0x56, 0x90, 0x5f, 0x54, 0x22, 0xc1, 0x0c, 0x14, 0xe3, 0x0d, 0x02,
	0xb3, 0x9d, 0x6a, 0xb9, 0xe4, 0x92, 0xf3, 0x73, 0xf5, 0xca, 0x53, 0x13, 0x73, 0x52, 0x4a, 0x52,
	0x93, 0x33, 0xf4, 0x52, 0x4b, 0x32, 0x8c, 0x8a, 0x33, 0xd3, 0xf3, 0x52, 0x8b, 0x12, 0x0b, 0x32,
	0xf5, 0xca, 0x0c, 0x9d, 0x78, 0x61, 0x76, 0x04, 0x80, 0x2c, 0x0e, 0x60, 0x8c, 0xd2, 0x4a, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x03, 0xea, 0xd3, 0x87, 0xeb, 0xd3, 0x07, 0xe9, 0xd3, 0x85, 0x68,
	0xd4, 0x05, 0xea, 0xd4, 0x2f, 0x48, 0xd2, 0x2f, 0x33, 0x5c, 0xc5, 0xc4, 0xeb, 0x0a, 0x14, 0x0f,
	0x06, 0x0b, 0x03, 0x0d, 0x3b, 0x85, 0xcc, 0x8f, 0x29, 0x33, 0x4c, 0x62, 0x03, 0xfb, 0xc6, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x97, 0xb2, 0x3b, 0xdf, 0x00, 0x00, 0x00,
}
