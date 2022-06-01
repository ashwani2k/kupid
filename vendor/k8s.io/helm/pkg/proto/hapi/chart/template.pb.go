// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/chart/template.proto

package chart

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

// Template represents a template as a name/value pair.
//
// By convention, name is a relative path within the scope of the chart's
// base directory.
type Template struct {
	// Name is the path-like name of the template.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data is the template as byte data.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Template) Reset()         { *m = Template{} }
func (m *Template) String() string { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()    {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_926c98477d6df5e9, []int{0}
}
func (m *Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Template.Unmarshal(m, b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Template.Marshal(b, m, deterministic)
}
func (dst *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(dst, src)
}
func (m *Template) XXX_Size() int {
	return xxx_messageInfo_Template.Size(m)
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Template) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Template)(nil), "hapi.chart.Template")
}

func init() {
	proto.RegisterFile("hapi/chart/template.proto", fileDescriptor_template_926c98477d6df5e9)
}

var fileDescriptor_template_926c98477d6df5e9 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x48, 0x2c, 0xc8,
	0xd4, 0x4f, 0xce, 0x48, 0x2c, 0x2a, 0xd1, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x49, 0xe9, 0x81, 0xa5, 0x94, 0x8c, 0xb8, 0x38,
	0x42, 0xa0, 0xb2, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x60, 0x36, 0x48, 0x2c, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27,
	0x08, 0xcc, 0x76, 0x62, 0x8f, 0x62, 0x05, 0x6b, 0x4e, 0x62, 0x03, 0x9b, 0x67, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xee, 0x0e, 0x67, 0x6c, 0x00, 0x00, 0x00,
}
