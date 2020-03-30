// Code generated by protoc-gen-go. DO NOT EDIT.
// source: responsestate.proto

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

type ResponseState int32

const (
	// UNKNOWN occurs when no information about the response is available.
	ResponseState_UNKNOWN ResponseState = 0
	// SUCCEEDED occurs when a request was successful.
	ResponseState_SUCCEEDED ResponseState = 1
	// DENIED occurs when a request was denied.
	ResponseState_DENIED ResponseState = 2
	// FAILED occurs when a request failed to complete.
	ResponseState_FAILED ResponseState = 3
)

var ResponseState_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCEEDED",
	2: "DENIED",
	3: "FAILED",
}

var ResponseState_value = map[string]int32{
	"UNKNOWN":   0,
	"SUCCEEDED": 1,
	"DENIED":    2,
	"FAILED":    3,
}

func (x ResponseState) String() string {
	return proto.EnumName(ResponseState_name, int32(x))
}

func (ResponseState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fd8bc0c69b30697, []int{0}
}

func init() {
	proto.RegisterEnum("v1.ResponseState", ResponseState_name, ResponseState_value)
}

func init() {
	proto.RegisterFile("responsestate.proto", fileDescriptor_2fd8bc0c69b30697)
}

var fileDescriptor_2fd8bc0c69b30697 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x2d, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x33, 0xd4, 0x72, 0xe6, 0xe2, 0x0d, 0x82, 0x4a, 0x05, 0x83, 0xa4, 0x84, 0xb8, 0xb9,
	0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0x78, 0xb9, 0x38, 0x83, 0x43,
	0x9d, 0x9d, 0x5d, 0x5d, 0x5d, 0x5c, 0x5d, 0x04, 0x18, 0x85, 0xb8, 0xb8, 0xd8, 0x5c, 0x5c, 0xfd,
	0x3c, 0x5d, 0x5d, 0x04, 0x98, 0x40, 0x6c, 0x37, 0x47, 0x4f, 0x1f, 0x57, 0x17, 0x01, 0x66, 0xa7,
	0x26, 0x46, 0x2e, 0xb9, 0xe4, 0xfc, 0x5c, 0xbd, 0xf2, 0xd4, 0xc4, 0x9c, 0x94, 0x92, 0xd4, 0xe4,
	0x0c, 0xbd, 0xd4, 0x92, 0x0c, 0xa3, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0xa2, 0xc4, 0x82, 0x4c, 0xbd,
	0x32, 0x43, 0x27, 0x21, 0x14, 0x5b, 0x02, 0x40, 0xf6, 0x07, 0x30, 0x46, 0x69, 0xa5, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0x35, 0xeb, 0x83, 0x34, 0xeb, 0x42, 0x74,
	0xeb, 0x26, 0x16, 0x64, 0xea, 0x17, 0x24, 0xe9, 0x97, 0x19, 0xae, 0x62, 0xe2, 0x75, 0x2d, 0xc9,
	0x30, 0x0a, 0x06, 0x0b, 0xeb, 0x95, 0x19, 0x9e, 0x42, 0xe6, 0xc7, 0x94, 0x19, 0x26, 0xb1, 0x81,
	0x3d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x85, 0x5c, 0x52, 0x07, 0xeb, 0x00, 0x00, 0x00,
}
