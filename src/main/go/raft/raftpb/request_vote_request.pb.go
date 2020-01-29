// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request_vote_request.proto

package raftpb

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

type RequestVoteRequest struct {
	Term                 *uint64  `protobuf:"varint,1,req,name=Term" json:"Term,omitempty"`
	LastLogIndex         *uint64  `protobuf:"varint,2,req,name=LastLogIndex" json:"LastLogIndex,omitempty"`
	LastLogTerm          *uint64  `protobuf:"varint,3,req,name=LastLogTerm" json:"LastLogTerm,omitempty"`
	CandidateName        *string  `protobuf:"bytes,4,req,name=CandidateName" json:"CandidateName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteRequest) Reset()         { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()    {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bda23e41b68e9a, []int{0}
}

func (m *RequestVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteRequest.Unmarshal(m, b)
}
func (m *RequestVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteRequest.Marshal(b, m, deterministic)
}
func (m *RequestVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteRequest.Merge(m, src)
}
func (m *RequestVoteRequest) XXX_Size() int {
	return xxx_messageInfo_RequestVoteRequest.Size(m)
}
func (m *RequestVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteRequest proto.InternalMessageInfo

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogIndex() uint64 {
	if m != nil && m.LastLogIndex != nil {
		return *m.LastLogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogTerm() uint64 {
	if m != nil && m.LastLogTerm != nil {
		return *m.LastLogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateName() string {
	if m != nil && m.CandidateName != nil {
		return *m.CandidateName
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestVoteRequest)(nil), "raftpb.RequestVoteRequest")
}

func init() { proto.RegisterFile("request_vote_request.proto", fileDescriptor_21bda23e41b68e9a) }

var fileDescriptor_21bda23e41b68e9a = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x89, 0x2f, 0xcb, 0x2f, 0x49, 0x8d, 0x87, 0x72, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0xd8, 0x8a, 0x12, 0xd3, 0x4a, 0x0a, 0x92, 0x94, 0x52, 0xb8, 0x84, 0x82, 0x20, 0x12,
	0x61, 0xf9, 0x25, 0xa9, 0x50, 0xa6, 0x10, 0x0f, 0x17, 0x4b, 0x48, 0x6a, 0x51, 0xae, 0x04, 0xa3,
	0x02, 0x93, 0x06, 0x8b, 0x90, 0x08, 0x17, 0x8f, 0x4f, 0x62, 0x71, 0x89, 0x4f, 0x7e, 0xba, 0x67,
	0x5e, 0x4a, 0x6a, 0x85, 0x04, 0x13, 0x58, 0x54, 0x98, 0x8b, 0x1b, 0x2a, 0x0a, 0x56, 0xca, 0x0c,
	0x16, 0x14, 0xe5, 0xe2, 0x75, 0x4e, 0xcc, 0x4b, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xf5, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x51, 0x60, 0xd2, 0xe0, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0xab, 0x5c, 0xe5,
	0xd8, 0x8a, 0x00, 0x00, 0x00,
}