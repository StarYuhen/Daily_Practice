// Code generated by protoc-gen-go. DO NOT EDIT.
// source: set.proto

//当前目录的包名

package protoc

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

//Grodens 请求结构
type Grodens struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Up                   *UploadStar `protobuf:"bytes,2,opt,name=up,proto3" json:"up,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Grodens) Reset()         { *m = Grodens{} }
func (m *Grodens) String() string { return proto.CompactTextString(m) }
func (*Grodens) ProtoMessage()    {}
func (*Grodens) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d650fd95c5da449, []int{0}
}

func (m *Grodens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grodens.Unmarshal(m, b)
}
func (m *Grodens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grodens.Marshal(b, m, deterministic)
}
func (m *Grodens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grodens.Merge(m, src)
}
func (m *Grodens) XXX_Size() int {
	return xxx_messageInfo_Grodens.Size(m)
}
func (m *Grodens) XXX_DiscardUnknown() {
	xxx_messageInfo_Grodens.DiscardUnknown(m)
}

var xxx_messageInfo_Grodens proto.InternalMessageInfo

func (m *Grodens) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Grodens) GetUp() *UploadStar {
	if m != nil {
		return m.Up
	}
	return nil
}

//Gou响应结构
type Gou struct {
	Mess                 string   `protobuf:"bytes,1,opt,name=mess,proto3" json:"mess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gou) Reset()         { *m = Gou{} }
func (m *Gou) String() string { return proto.CompactTextString(m) }
func (*Gou) ProtoMessage()    {}
func (*Gou) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d650fd95c5da449, []int{1}
}

func (m *Gou) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gou.Unmarshal(m, b)
}
func (m *Gou) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gou.Marshal(b, m, deterministic)
}
func (m *Gou) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gou.Merge(m, src)
}
func (m *Gou) XXX_Size() int {
	return xxx_messageInfo_Gou.Size(m)
}
func (m *Gou) XXX_DiscardUnknown() {
	xxx_messageInfo_Gou.DiscardUnknown(m)
}

var xxx_messageInfo_Gou proto.InternalMessageInfo

func (m *Gou) GetMess() string {
	if m != nil {
		return m.Mess
	}
	return ""
}

func init() {
	proto.RegisterType((*Grodens)(nil), "main.Grodens")
	proto.RegisterType((*Gou)(nil), "main.Gou")
}

func init() { proto.RegisterFile("set.proto", fileDescriptor_2d650fd95c5da449) }

var fileDescriptor_2d650fd95c5da449 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4e, 0x2d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93, 0xe2, 0x29, 0x2d, 0xc8,
	0xc9, 0x4f, 0x4c, 0x81, 0x88, 0x29, 0xd9, 0x73, 0xb1, 0xbb, 0x17, 0xe5, 0xa7, 0xa4, 0xe6, 0x15,
	0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x42, 0x0a, 0x5c, 0x4c, 0xa5, 0x05, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x02, 0x7a,
	0x20, 0xfd, 0x7a, 0xa1, 0x60, 0xed, 0xc1, 0x25, 0x89, 0x45, 0x41, 0x4c, 0xa5, 0x05, 0x4a, 0x92,
	0x5c, 0xcc, 0xee, 0xf9, 0xa5, 0x20, 0xcd, 0xb9, 0xa9, 0xc5, 0xc5, 0x30, 0xcd, 0x20, 0xb6, 0x91,
	0x36, 0x17, 0x5b, 0x00, 0xd8, 0x6c, 0x21, 0x45, 0x2e, 0xe6, 0xc4, 0x94, 0x14, 0x21, 0x5e, 0x88,
	0x09, 0x50, 0x0b, 0xa5, 0x38, 0xa1, 0xdc, 0xfc, 0x52, 0x25, 0x06, 0x27, 0xce, 0x28, 0x76, 0x7d,
	0xb0, 0x93, 0x92, 0x93, 0xd8, 0xc0, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x51, 0x6a,
	0xda, 0xbb, 0x00, 0x00, 0x00,
}