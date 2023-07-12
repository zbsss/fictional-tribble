// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device-manager.proto

package generated

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

type GetTokenRequest struct {
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{0}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type GetTokenReply struct {
	ExpiresAt            int64    `protobuf:"varint,1,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenReply) Reset()         { *m = GetTokenReply{} }
func (m *GetTokenReply) String() string { return proto.CompactTextString(m) }
func (*GetTokenReply) ProtoMessage()    {}
func (*GetTokenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{1}
}

func (m *GetTokenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenReply.Unmarshal(m, b)
}
func (m *GetTokenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenReply.Marshal(b, m, deterministic)
}
func (m *GetTokenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenReply.Merge(m, src)
}
func (m *GetTokenReply) XXX_Size() int {
	return xxx_messageInfo_GetTokenReply.Size(m)
}
func (m *GetTokenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenReply proto.InternalMessageInfo

func (m *GetTokenReply) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

type ReturnTokenRequest struct {
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnTokenRequest) Reset()         { *m = ReturnTokenRequest{} }
func (m *ReturnTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ReturnTokenRequest) ProtoMessage()    {}
func (*ReturnTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{2}
}

func (m *ReturnTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnTokenRequest.Unmarshal(m, b)
}
func (m *ReturnTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnTokenRequest.Marshal(b, m, deterministic)
}
func (m *ReturnTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnTokenRequest.Merge(m, src)
}
func (m *ReturnTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ReturnTokenRequest.Size(m)
}
func (m *ReturnTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnTokenRequest proto.InternalMessageInfo

func (m *ReturnTokenRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type ReturnTokenReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnTokenReply) Reset()         { *m = ReturnTokenReply{} }
func (m *ReturnTokenReply) String() string { return proto.CompactTextString(m) }
func (*ReturnTokenReply) ProtoMessage()    {}
func (*ReturnTokenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{3}
}

func (m *ReturnTokenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnTokenReply.Unmarshal(m, b)
}
func (m *ReturnTokenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnTokenReply.Marshal(b, m, deterministic)
}
func (m *ReturnTokenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnTokenReply.Merge(m, src)
}
func (m *ReturnTokenReply) XXX_Size() int {
	return xxx_messageInfo_ReturnTokenReply.Size(m)
}
func (m *ReturnTokenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnTokenReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnTokenReply proto.InternalMessageInfo

type GetMemoryQuotaRequest struct {
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Memory               uint64   `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemoryQuotaRequest) Reset()         { *m = GetMemoryQuotaRequest{} }
func (m *GetMemoryQuotaRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemoryQuotaRequest) ProtoMessage()    {}
func (*GetMemoryQuotaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{4}
}

func (m *GetMemoryQuotaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemoryQuotaRequest.Unmarshal(m, b)
}
func (m *GetMemoryQuotaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemoryQuotaRequest.Marshal(b, m, deterministic)
}
func (m *GetMemoryQuotaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemoryQuotaRequest.Merge(m, src)
}
func (m *GetMemoryQuotaRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemoryQuotaRequest.Size(m)
}
func (m *GetMemoryQuotaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemoryQuotaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemoryQuotaRequest proto.InternalMessageInfo

func (m *GetMemoryQuotaRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *GetMemoryQuotaRequest) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

type GetMemoryQuotaReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemoryQuotaReply) Reset()         { *m = GetMemoryQuotaReply{} }
func (m *GetMemoryQuotaReply) String() string { return proto.CompactTextString(m) }
func (*GetMemoryQuotaReply) ProtoMessage()    {}
func (*GetMemoryQuotaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{5}
}

func (m *GetMemoryQuotaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemoryQuotaReply.Unmarshal(m, b)
}
func (m *GetMemoryQuotaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemoryQuotaReply.Marshal(b, m, deterministic)
}
func (m *GetMemoryQuotaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemoryQuotaReply.Merge(m, src)
}
func (m *GetMemoryQuotaReply) XXX_Size() int {
	return xxx_messageInfo_GetMemoryQuotaReply.Size(m)
}
func (m *GetMemoryQuotaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemoryQuotaReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemoryQuotaReply proto.InternalMessageInfo

type ReturnMemoryQuotaRequest struct {
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Memory               uint64   `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnMemoryQuotaRequest) Reset()         { *m = ReturnMemoryQuotaRequest{} }
func (m *ReturnMemoryQuotaRequest) String() string { return proto.CompactTextString(m) }
func (*ReturnMemoryQuotaRequest) ProtoMessage()    {}
func (*ReturnMemoryQuotaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{6}
}

func (m *ReturnMemoryQuotaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnMemoryQuotaRequest.Unmarshal(m, b)
}
func (m *ReturnMemoryQuotaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnMemoryQuotaRequest.Marshal(b, m, deterministic)
}
func (m *ReturnMemoryQuotaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnMemoryQuotaRequest.Merge(m, src)
}
func (m *ReturnMemoryQuotaRequest) XXX_Size() int {
	return xxx_messageInfo_ReturnMemoryQuotaRequest.Size(m)
}
func (m *ReturnMemoryQuotaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnMemoryQuotaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnMemoryQuotaRequest proto.InternalMessageInfo

func (m *ReturnMemoryQuotaRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ReturnMemoryQuotaRequest) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

type ReturnMemoryQuotaReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnMemoryQuotaReply) Reset()         { *m = ReturnMemoryQuotaReply{} }
func (m *ReturnMemoryQuotaReply) String() string { return proto.CompactTextString(m) }
func (*ReturnMemoryQuotaReply) ProtoMessage()    {}
func (*ReturnMemoryQuotaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a1c4951d394250d, []int{7}
}

func (m *ReturnMemoryQuotaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnMemoryQuotaReply.Unmarshal(m, b)
}
func (m *ReturnMemoryQuotaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnMemoryQuotaReply.Marshal(b, m, deterministic)
}
func (m *ReturnMemoryQuotaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnMemoryQuotaReply.Merge(m, src)
}
func (m *ReturnMemoryQuotaReply) XXX_Size() int {
	return xxx_messageInfo_ReturnMemoryQuotaReply.Size(m)
}
func (m *ReturnMemoryQuotaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnMemoryQuotaReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnMemoryQuotaReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTokenRequest)(nil), "device_manager.GetTokenRequest")
	proto.RegisterType((*GetTokenReply)(nil), "device_manager.GetTokenReply")
	proto.RegisterType((*ReturnTokenRequest)(nil), "device_manager.ReturnTokenRequest")
	proto.RegisterType((*ReturnTokenReply)(nil), "device_manager.ReturnTokenReply")
	proto.RegisterType((*GetMemoryQuotaRequest)(nil), "device_manager.GetMemoryQuotaRequest")
	proto.RegisterType((*GetMemoryQuotaReply)(nil), "device_manager.GetMemoryQuotaReply")
	proto.RegisterType((*ReturnMemoryQuotaRequest)(nil), "device_manager.ReturnMemoryQuotaRequest")
	proto.RegisterType((*ReturnMemoryQuotaReply)(nil), "device_manager.ReturnMemoryQuotaReply")
}

func init() {
	proto.RegisterFile("device-manager.proto", fileDescriptor_1a1c4951d394250d)
}

var fileDescriptor_1a1c4951d394250d = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xf6, 0x93, 0x62, 0xaf, 0xb4, 0xea, 0x68, 0x4b, 0x28, 0x14, 0xc3, 0x88, 0x92,
	0xa2, 0x26, 0xa0, 0x4f, 0xa0, 0x08, 0x01, 0xb1, 0x0b, 0x83, 0x6e, 0x5c, 0x58, 0x92, 0xf6, 0x12,
	0x83, 0x49, 0x26, 0x4e, 0x26, 0x62, 0x7c, 0x77, 0x41, 0x92, 0x89, 0x68, 0xfe, 0x68, 0xbb, 0x70,
	0x39, 0x87, 0x5f, 0xce, 0xb9, 0x39, 0x97, 0x0b, 0xbb, 0x0b, 0x7c, 0xf1, 0xe6, 0x78, 0x12, 0xd8,
	0xa1, 0xed, 0x22, 0xd7, 0x23, 0xce, 0x04, 0x23, 0x7d, 0xa9, 0xce, 0x0a, 0x95, 0x4e, 0x60, 0xd3,
	0x44, 0x71, 0xcb, 0x9e, 0x30, 0xb4, 0xf0, 0x39, 0xc1, 0x58, 0x90, 0x21, 0x74, 0x24, 0xa4, 0xb4,
	0xd4, 0x96, 0xd6, 0xb5, 0x8a, 0x17, 0xd5, 0xa1, 0xf7, 0x85, 0x46, 0x7e, 0x4a, 0xc6, 0x00, 0xf8,
	0x1a, 0x79, 0x1c, 0xe3, 0x99, 0x2d, 0x72, 0xf8, 0xbf, 0xd5, 0x2d, 0x94, 0x73, 0x41, 0x8f, 0x81,
	0x58, 0x28, 0x12, 0x1e, 0xae, 0xe4, 0x4e, 0x60, 0xab, 0x44, 0x47, 0x7e, 0x4a, 0x4d, 0x18, 0x98,
	0x28, 0xa6, 0x18, 0x30, 0x9e, 0xde, 0x24, 0x4c, 0xd8, 0x4b, 0x4c, 0x32, 0x3d, 0xc8, 0x69, 0xa5,
	0xad, 0xb6, 0xb4, 0x35, 0xab, 0x78, 0xd1, 0x01, 0xec, 0x54, 0x8d, 0x32, 0xff, 0x2b, 0x50, 0x64,
	0xe6, 0x1f, 0x44, 0x28, 0x30, 0x6c, 0xf0, 0x8a, 0xfc, 0xf4, 0xf4, 0xbd, 0x0d, 0xbd, 0xcb, 0xfc,
	0xe3, 0xa9, 0x2c, 0x9d, 0x5c, 0xc3, 0xfa, 0x67, 0x93, 0x64, 0x4f, 0x2f, 0x6f, 0x44, 0xaf, 0xac,
	0x63, 0x34, 0xfe, 0x19, 0xc8, 0xfe, 0xe1, 0x1f, 0xb9, 0x83, 0x8d, 0x6f, 0xcd, 0x11, 0x5a, 0xe5,
	0xeb, 0x4b, 0x18, 0xa9, 0xbf, 0x32, 0xd2, 0xf6, 0x01, 0xfa, 0xe5, 0xce, 0xc8, 0x41, 0xc3, 0x24,
	0xf5, 0xe6, 0x46, 0xfb, 0xcb, 0x30, 0xe9, 0xef, 0xc2, 0x76, 0xad, 0x30, 0xa2, 0x35, 0x0f, 0xd6,
	0x90, 0x72, 0xb8, 0x02, 0x99, 0x07, 0x5d, 0x1c, 0xdd, 0x4f, 0x5c, 0x4f, 0x3c, 0x26, 0x8e, 0x3e,
	0x67, 0x81, 0xf1, 0xe6, 0xc4, 0x71, 0x6c, 0x94, 0x6f, 0xc3, 0x70, 0x31, 0x44, 0x6e, 0x0b, 0x5c,
	0x38, 0x9d, 0xfc, 0x4c, 0xce, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x46, 0xf7, 0x13, 0x7f, 0x3e,
	0x03, 0x00, 0x00,
}
