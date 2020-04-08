// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth2.proto

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

// AttestationData is defined at https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#attestationdata
type AttestationData struct {
	Slot                 uint64      `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex       uint64      `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
	BeaconBlockRoot      []byte      `protobuf:"bytes,3,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty"`
	Source               *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target               *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c52217a1e941062, []int{0}
}

func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (m *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(m, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *AttestationData) GetCommitteeIndex() uint64 {
	if m != nil {
		return m.CommitteeIndex
	}
	return 0
}

func (m *AttestationData) GetBeaconBlockRoot() []byte {
	if m != nil {
		return m.BeaconBlockRoot
	}
	return nil
}

func (m *AttestationData) GetSource() *Checkpoint {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttestationData) GetTarget() *Checkpoint {
	if m != nil {
		return m.Target
	}
	return nil
}

// Checkpoint is defined at https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#checkpoint
type Checkpoint struct {
	Epoch                uint64   `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Root                 []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c52217a1e941062, []int{1}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Checkpoint) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

// BeaconBlockheader is defined at https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#beaconblockheader
type BeaconBlockHeader struct {
	Slot                 uint64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	ProposerIndex        uint64   `protobuf:"varint,2,opt,name=proposer_index,json=proposerIndex,proto3" json:"proposer_index,omitempty"`
	ParentRoot           []byte   `protobuf:"bytes,3,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty"`
	StateRoot            []byte   `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	BodyRoot             []byte   `protobuf:"bytes,5,opt,name=body_root,json=bodyRoot,proto3" json:"body_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconBlockHeader) Reset()         { *m = BeaconBlockHeader{} }
func (m *BeaconBlockHeader) String() string { return proto.CompactTextString(m) }
func (*BeaconBlockHeader) ProtoMessage()    {}
func (*BeaconBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c52217a1e941062, []int{2}
}

func (m *BeaconBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconBlockHeader.Unmarshal(m, b)
}
func (m *BeaconBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconBlockHeader.Marshal(b, m, deterministic)
}
func (m *BeaconBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconBlockHeader.Merge(m, src)
}
func (m *BeaconBlockHeader) XXX_Size() int {
	return xxx_messageInfo_BeaconBlockHeader.Size(m)
}
func (m *BeaconBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconBlockHeader proto.InternalMessageInfo

func (m *BeaconBlockHeader) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *BeaconBlockHeader) GetProposerIndex() uint64 {
	if m != nil {
		return m.ProposerIndex
	}
	return 0
}

func (m *BeaconBlockHeader) GetParentRoot() []byte {
	if m != nil {
		return m.ParentRoot
	}
	return nil
}

func (m *BeaconBlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BeaconBlockHeader) GetBodyRoot() []byte {
	if m != nil {
		return m.BodyRoot
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestationData)(nil), "v1.AttestationData")
	proto.RegisterType((*Checkpoint)(nil), "v1.Checkpoint")
	proto.RegisterType((*BeaconBlockHeader)(nil), "v1.BeaconBlockHeader")
}

func init() {
	proto.RegisterFile("eth2.proto", fileDescriptor_5c52217a1e941062)
}

var fileDescriptor_5c52217a1e941062 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0x49, 0xcc, 0x15, 0x6f, 0x6a, 0xef, 0x68, 0xf0, 0x21, 0x20, 0xea, 0x71, 0xa0, 0x1e,
	0x85, 0x6e, 0xc8, 0x09, 0xbe, 0x1b, 0x15, 0xf4, 0xad, 0xc4, 0x37, 0x11, 0x8e, 0xcd, 0x66, 0xb8,
	0x2c, 0xbd, 0x64, 0x96, 0xcd, 0x34, 0xda, 0xaf, 0xa4, 0xdf, 0xa4, 0x9f, 0x4a, 0x76, 0xb7, 0xf6,
	0xae, 0x70, 0x6f, 0xbb, 0xbf, 0xdf, 0x0e, 0xcc, 0xfc, 0x67, 0x01, 0x90, 0xdb, 0xb5, 0x30, 0x96,
	0x98, 0xd2, 0x78, 0x2c, 0x96, 0x77, 0x11, 0xcc, 0x3f, 0x32, 0xe3, 0xc0, 0x92, 0x35, 0xf5, 0x9f,
	0x25, 0xcb, 0x34, 0x85, 0x64, 0xd8, 0x11, 0x67, 0xd1, 0x22, 0x5a, 0x25, 0x95, 0x3f, 0xa7, 0xef,
	0x60, 0xae, 0xa8, 0xeb, 0x34, 0x33, 0xe2, 0x46, 0xf7, 0x0d, 0xfe, 0xce, 0x62, 0xaf, 0x67, 0x0f,
	0xf8, 0x9b, 0xa3, 0xe9, 0x05, 0x9c, 0xd7, 0x28, 0x15, 0xf5, 0x9b, 0x7a, 0x47, 0xea, 0x7a, 0x63,
	0x89, 0x38, 0x7b, 0xb2, 0x88, 0x56, 0xcf, 0xaa, 0x79, 0x10, 0xa5, 0xe3, 0x15, 0x11, 0xa7, 0x6f,
	0xe1, 0x64, 0xa0, 0x1b, 0xab, 0x30, 0x4b, 0x16, 0xd1, 0xea, 0x74, 0x3d, 0x13, 0x63, 0x21, 0x3e,
	0xb5, 0xa8, 0xae, 0x0d, 0xe9, 0x9e, 0xab, 0x7b, 0xeb, 0xde, 0xb1, 0xb4, 0x5b, 0xe4, 0x6c, 0x72,
	0xfc, 0x5d, 0xb0, 0xcb, 0x0f, 0x00, 0x7b, 0x9a, 0x3e, 0x87, 0x09, 0x1a, 0x52, 0xed, 0xfd, 0x1c,
	0xe1, 0xe2, 0x86, 0xf3, 0x2d, 0xc5, 0xbe, 0x25, 0x7f, 0x5e, 0xfe, 0x8d, 0xe0, 0xbc, 0xdc, 0xf7,
	0xf6, 0x15, 0x65, 0x83, 0xf6, 0x68, 0x0c, 0x6f, 0x60, 0x66, 0x2c, 0x19, 0x1a, 0xd0, 0x3e, 0x4a,
	0xe1, 0xec, 0x3f, 0x0d, 0x21, 0xbc, 0x86, 0x53, 0x23, 0x2d, 0xf6, 0x7c, 0x38, 0x3e, 0x04, 0xe4,
	0x27, 0x7f, 0x09, 0xe0, 0x12, 0xc7, 0xe0, 0x13, 0xef, 0xa7, 0x9e, 0x78, 0xfd, 0x02, 0xa6, 0x35,
	0x35, 0xb7, 0xc1, 0x4e, 0xbc, 0x7d, 0xea, 0x80, 0x93, 0xe5, 0x2d, 0xbc, 0x52, 0xd4, 0x89, 0x5f,
	0x28, 0x77, 0x0d, 0xa3, 0x6a, 0x85, 0x5b, 0xe9, 0xa0, 0xb7, 0x3d, 0x5a, 0x69, 0xb4, 0x18, 0x8b,
	0x72, 0xfa, 0x85, 0xdb, 0xf5, 0x95, 0xdb, 0xf1, 0x55, 0xf4, 0xe3, 0x62, 0xab, 0xb9, 0xbd, 0xa9,
	0x85, 0xa2, 0x2e, 0x7f, 0xa8, 0xc9, 0x5d, 0xcd, 0x65, 0x28, 0xba, 0x94, 0x46, 0xe7, 0xa6, 0xce,
	0xc7, 0xe2, 0x4f, 0x7c, 0xe6, 0x2a, 0xbf, 0x7b, 0x2c, 0xc6, 0xe2, 0xee, 0xf0, 0xfe, 0x73, 0x2c,
	0xea, 0x13, 0xff, 0x71, 0xde, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xda, 0xe9, 0x60, 0x54, 0x46,
	0x02, 0x00, 0x00,
}
