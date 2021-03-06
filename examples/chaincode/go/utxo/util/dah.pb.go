// Code generated by protoc-gen-go.
// source: dah.proto
// DO NOT EDIT!

/*
Package util is a generated protocol buffer package.

It is generated from these files:
	dah.proto

It has these top-level messages:
	TX
*/
package util

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TX struct {
	Version  uint32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LockTime uint32      `protobuf:"varint,2,opt,name=lockTime" json:"lockTime,omitempty"`
	Txin     []*TX_TXIN  `protobuf:"bytes,3,rep,name=txin" json:"txin,omitempty"`
	Txout    []*TX_TXOUT `protobuf:"bytes,4,rep,name=txout" json:"txout,omitempty"`
	Blocks   [][]byte    `protobuf:"bytes,5,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Fee      uint64      `protobuf:"varint,6,opt,name=fee" json:"fee,omitempty"`
}

func (m *TX) Reset()                    { *m = TX{} }
func (m *TX) String() string            { return proto.CompactTextString(m) }
func (*TX) ProtoMessage()               {}
func (*TX) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TX) GetTxin() []*TX_TXIN {
	if m != nil {
		return m.Txin
	}
	return nil
}

func (m *TX) GetTxout() []*TX_TXOUT {
	if m != nil {
		return m.Txout
	}
	return nil
}

type TX_TXIN struct {
	Ix         uint32 `protobuf:"varint,1,opt,name=ix" json:"ix,omitempty"`
	SourceHash []byte `protobuf:"bytes,2,opt,name=sourceHash,proto3" json:"sourceHash,omitempty"`
	Script     []byte `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	Sequence   uint32 `protobuf:"varint,4,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *TX_TXIN) Reset()                    { *m = TX_TXIN{} }
func (m *TX_TXIN) String() string            { return proto.CompactTextString(m) }
func (*TX_TXIN) ProtoMessage()               {}
func (*TX_TXIN) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TX_TXOUT struct {
	Value    uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Script   []byte `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Color    []byte `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Quantity uint64 `protobuf:"varint,4,opt,name=quantity" json:"quantity,omitempty"`
}

func (m *TX_TXOUT) Reset()                    { *m = TX_TXOUT{} }
func (m *TX_TXOUT) String() string            { return proto.CompactTextString(m) }
func (*TX_TXOUT) ProtoMessage()               {}
func (*TX_TXOUT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func init() {
	proto.RegisterType((*TX)(nil), "util.TX")
	proto.RegisterType((*TX_TXIN)(nil), "util.TX.TXIN")
	proto.RegisterType((*TX_TXOUT)(nil), "util.TX.TXOUT")
}

func init() { proto.RegisterFile("dah.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4e, 0xf3, 0x30,
	0x10, 0xc4, 0x95, 0xd8, 0xc9, 0xf7, 0xb1, 0xb4, 0x15, 0x5a, 0x55, 0xc8, 0xca, 0x01, 0x05, 0xc4,
	0x21, 0xa7, 0x1c, 0xe0, 0x25, 0xe0, 0x02, 0x92, 0x15, 0xa4, 0x5c, 0xd3, 0x60, 0xa8, 0x21, 0xc4,
	0xad, 0xff, 0x54, 0xe1, 0xca, 0x93, 0x23, 0x3b, 0x69, 0x15, 0x6e, 0xfe, 0x79, 0x56, 0x33, 0xb3,
	0x5a, 0x38, 0x7b, 0x6d, 0xb6, 0xe5, 0x4e, 0x2b, 0xab, 0x90, 0x3a, 0x2b, 0xbb, 0x9b, 0x1f, 0x02,
	0x71, 0x55, 0x23, 0x83, 0x7f, 0x07, 0xa1, 0x8d, 0x54, 0x3d, 0x8b, 0xf2, 0xa8, 0x58, 0xf2, 0x23,
	0x62, 0x06, 0xff, 0x3b, 0xd5, 0x7e, 0x56, 0xf2, 0x4b, 0xb0, 0x38, 0x48, 0x27, 0xc6, 0x6b, 0xa0,
	0x76, 0x90, 0x3d, 0x23, 0x39, 0x29, 0xce, 0xef, 0x96, 0xa5, 0x77, 0x2c, 0xab, 0xba, 0xac, 0xea,
	0xc7, 0x27, 0x1e, 0x24, 0xbc, 0x85, 0xc4, 0x0e, 0xca, 0x59, 0x46, 0xc3, 0xcc, 0x6a, 0x36, 0xf3,
	0xfc, 0x52, 0xf1, 0x51, 0xc4, 0x4b, 0x48, 0x37, 0xde, 0xd5, 0xb0, 0x24, 0x27, 0xc5, 0x82, 0x4f,
	0x84, 0x17, 0x40, 0xde, 0x84, 0x60, 0x69, 0x1e, 0x15, 0x94, 0xfb, 0x67, 0xf6, 0x01, 0xd4, 0xbb,
	0xe3, 0x0a, 0x62, 0x39, 0x4c, 0x5d, 0x63, 0x39, 0xe0, 0x15, 0x80, 0x51, 0x4e, 0xb7, 0xe2, 0xa1,
	0x31, 0xdb, 0x50, 0x74, 0xc1, 0x67, 0x3f, 0x3e, 0xc1, 0xb4, 0x5a, 0xee, 0x2c, 0x23, 0x41, 0x9b,
	0xc8, 0xaf, 0x67, 0xc4, 0xde, 0x89, 0xbe, 0x15, 0x8c, 0x8e, 0xeb, 0x1d, 0x39, 0x7b, 0x87, 0x24,
	0xb4, 0xc4, 0x35, 0x24, 0x87, 0xa6, 0x73, 0x22, 0xe4, 0x51, 0x3e, 0xc2, 0xcc, 0x32, 0xfe, 0x63,
	0xb9, 0x86, 0xa4, 0x55, 0x9d, 0xd2, 0x53, 0xd2, 0x08, 0x3e, 0x68, 0xef, 0x9a, 0xde, 0x4a, 0xfb,
	0x1d, 0x82, 0x28, 0x3f, 0xf1, 0x26, 0x0d, 0x17, 0xb9, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x91, 0x0f, 0xdd, 0x9e, 0x01, 0x00, 0x00,
}
