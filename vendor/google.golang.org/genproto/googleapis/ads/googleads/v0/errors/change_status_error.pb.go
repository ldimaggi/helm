// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/change_status_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible change status errors.
type ChangeStatusErrorEnum_ChangeStatusError int32

const (
	// Enum unspecified.
	ChangeStatusErrorEnum_UNSPECIFIED ChangeStatusErrorEnum_ChangeStatusError = 0
	// The received error code is not known in this version.
	ChangeStatusErrorEnum_UNKNOWN ChangeStatusErrorEnum_ChangeStatusError = 1
	// The requested start date is too old.
	ChangeStatusErrorEnum_START_DATE_TOO_OLD ChangeStatusErrorEnum_ChangeStatusError = 3
)

var ChangeStatusErrorEnum_ChangeStatusError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "START_DATE_TOO_OLD",
}
var ChangeStatusErrorEnum_ChangeStatusError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"START_DATE_TOO_OLD": 3,
}

func (x ChangeStatusErrorEnum_ChangeStatusError) String() string {
	return proto.EnumName(ChangeStatusErrorEnum_ChangeStatusError_name, int32(x))
}
func (ChangeStatusErrorEnum_ChangeStatusError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_change_status_error_3510b0503a390638, []int{0, 0}
}

// Container for enum describing possible change status errors.
type ChangeStatusErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusErrorEnum) Reset()         { *m = ChangeStatusErrorEnum{} }
func (m *ChangeStatusErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusErrorEnum) ProtoMessage()    {}
func (*ChangeStatusErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_change_status_error_3510b0503a390638, []int{0}
}
func (m *ChangeStatusErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusErrorEnum.Unmarshal(m, b)
}
func (m *ChangeStatusErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ChangeStatusErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusErrorEnum.Merge(dst, src)
}
func (m *ChangeStatusErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusErrorEnum.Size(m)
}
func (m *ChangeStatusErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChangeStatusErrorEnum)(nil), "google.ads.googleads.v0.errors.ChangeStatusErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.ChangeStatusErrorEnum_ChangeStatusError", ChangeStatusErrorEnum_ChangeStatusError_name, ChangeStatusErrorEnum_ChangeStatusError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/change_status_error.proto", fileDescriptor_change_status_error_3510b0503a390638)
}

var fileDescriptor_change_status_error_3510b0503a390638 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xe4, 0x8c, 0xc4, 0xbc, 0xf4, 0xd4, 0xf8, 0xe2, 0x92, 0xc4, 0x92,
	0xd2, 0xe2, 0x78, 0xb0, 0xa0, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44, 0xb9, 0x5e,
	0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xa7, 0x5e, 0x99, 0x81, 0x1e, 0x44, 0xa7, 0x52, 0x12, 0x97, 0xa8,
	0x33, 0x58, 0x73, 0x30, 0x58, 0xaf, 0x2b, 0x48, 0xd4, 0x35, 0xaf, 0x34, 0x57, 0xc9, 0x93, 0x4b,
	0x10, 0x43, 0x42, 0x88, 0x9f, 0x8b, 0x3b, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9, 0xd3, 0xcd, 0xd3,
	0xd5, 0x45, 0x80, 0x41, 0x88, 0x9b, 0x8b, 0x3d, 0xd4, 0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80,
	0x51, 0x48, 0x8c, 0x4b, 0x28, 0x38, 0xc4, 0x31, 0x28, 0x24, 0xde, 0xc5, 0x31, 0xc4, 0x35, 0x3e,
	0xc4, 0xdf, 0x3f, 0xde, 0xdf, 0xc7, 0x45, 0x80, 0xd9, 0xe9, 0x0c, 0x23, 0x97, 0x52, 0x72, 0x7e,
	0xae, 0x1e, 0x7e, 0xa7, 0x38, 0x89, 0x61, 0xd8, 0x17, 0x00, 0xf2, 0x42, 0x00, 0x63, 0x94, 0x0b,
	0x54, 0x67, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e,
	0xd8, 0x83, 0xb0, 0xe0, 0x28, 0xc8, 0x2c, 0xc6, 0x15, 0x3a, 0xd6, 0x10, 0x6a, 0x11, 0x13, 0xb3,
	0xbb, 0xa3, 0xe3, 0x2a, 0x26, 0x39, 0x77, 0x88, 0x61, 0x8e, 0x29, 0xc5, 0x7a, 0x10, 0x26, 0x88,
	0x15, 0x66, 0xa0, 0x07, 0xb6, 0xb2, 0xf8, 0x14, 0x4c, 0x41, 0x8c, 0x63, 0x4a, 0x71, 0x0c, 0x5c,
	0x41, 0x4c, 0x98, 0x41, 0x0c, 0x44, 0x41, 0x12, 0x1b, 0xd8, 0x62, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0x6d, 0xde, 0x35, 0x95, 0x01, 0x00, 0x00,
}
