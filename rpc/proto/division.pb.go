// Code generated by protoc-gen-go. DO NOT EDIT.
// source: division.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Status int32

const (
	Status_SUCCEED Status = 0
	Status_FAILED  Status = 1
)

var Status_name = map[int32]string{
	0: "SUCCEED",
	1: "FAILED",
}

var Status_value = map[string]int32{
	"SUCCEED": 0,
	"FAILED":  1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{0}
}

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

var Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}

var Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{1}
}

type User struct {
	Udid                 string   `protobuf:"bytes,1,opt,name=Udid,proto3" json:"Udid,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	Channel              string   `protobuf:"bytes,4,opt,name=Channel,proto3" json:"Channel,omitempty"`
	Gender               Gender   `protobuf:"varint,5,opt,name=Gender,proto3,enum=proto.Gender" json:"Gender,omitempty"`
	BusinessId           string   `protobuf:"bytes,6,opt,name=BusinessId,proto3" json:"BusinessId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUdid() string {
	if m != nil {
		return m.Udid
	}
	return ""
}

func (m *User) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *User) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *User) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *User) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_MALE
}

func (m *User) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

type BusinessInfo struct {
	BusinessId           string   `protobuf:"bytes,1,opt,name=BusinessId,proto3" json:"BusinessId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusinessInfo) Reset()         { *m = BusinessInfo{} }
func (m *BusinessInfo) String() string { return proto.CompactTextString(m) }
func (*BusinessInfo) ProtoMessage()    {}
func (*BusinessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{1}
}

func (m *BusinessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessInfo.Unmarshal(m, b)
}
func (m *BusinessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessInfo.Marshal(b, m, deterministic)
}
func (m *BusinessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessInfo.Merge(m, src)
}
func (m *BusinessInfo) XXX_Size() int {
	return xxx_messageInfo_BusinessInfo.Size(m)
}
func (m *BusinessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessInfo proto.InternalMessageInfo

func (m *BusinessInfo) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

type Passthrough struct {
	ExpId                string   `protobuf:"bytes,1,opt,name=ExpId,proto3" json:"ExpId,omitempty"`
	BucketId             string   `protobuf:"bytes,2,opt,name=BucketId,proto3" json:"BucketId,omitempty"`
	BucketName           string   `protobuf:"bytes,3,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
	Uid                  int64    `protobuf:"varint,4,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Udid                 string   `protobuf:"bytes,5,opt,name=Udid,proto3" json:"Udid,omitempty"`
	ServeTime            int64    `protobuf:"varint,6,opt,name=ServeTime,proto3" json:"ServeTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Passthrough) Reset()         { *m = Passthrough{} }
func (m *Passthrough) String() string { return proto.CompactTextString(m) }
func (*Passthrough) ProtoMessage()    {}
func (*Passthrough) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{2}
}

func (m *Passthrough) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Passthrough.Unmarshal(m, b)
}
func (m *Passthrough) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Passthrough.Marshal(b, m, deterministic)
}
func (m *Passthrough) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Passthrough.Merge(m, src)
}
func (m *Passthrough) XXX_Size() int {
	return xxx_messageInfo_Passthrough.Size(m)
}
func (m *Passthrough) XXX_DiscardUnknown() {
	xxx_messageInfo_Passthrough.DiscardUnknown(m)
}

var xxx_messageInfo_Passthrough proto.InternalMessageInfo

func (m *Passthrough) GetExpId() string {
	if m != nil {
		return m.ExpId
	}
	return ""
}

func (m *Passthrough) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

func (m *Passthrough) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *Passthrough) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Passthrough) GetUdid() string {
	if m != nil {
		return m.Udid
	}
	return ""
}

func (m *Passthrough) GetServeTime() int64 {
	if m != nil {
		return m.ServeTime
	}
	return 0
}

type ABDivision struct {
	Assignment           string       `protobuf:"bytes,1,opt,name=Assignment,proto3" json:"Assignment,omitempty"`
	Passthrough          *Passthrough `protobuf:"bytes,2,opt,name=Passthrough,proto3" json:"Passthrough,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ABDivision) Reset()         { *m = ABDivision{} }
func (m *ABDivision) String() string { return proto.CompactTextString(m) }
func (*ABDivision) ProtoMessage()    {}
func (*ABDivision) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{3}
}

func (m *ABDivision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABDivision.Unmarshal(m, b)
}
func (m *ABDivision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABDivision.Marshal(b, m, deterministic)
}
func (m *ABDivision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABDivision.Merge(m, src)
}
func (m *ABDivision) XXX_Size() int {
	return xxx_messageInfo_ABDivision.Size(m)
}
func (m *ABDivision) XXX_DiscardUnknown() {
	xxx_messageInfo_ABDivision.DiscardUnknown(m)
}

var xxx_messageInfo_ABDivision proto.InternalMessageInfo

func (m *ABDivision) GetAssignment() string {
	if m != nil {
		return m.Assignment
	}
	return ""
}

func (m *ABDivision) GetPassthrough() *Passthrough {
	if m != nil {
		return m.Passthrough
	}
	return nil
}

type ABResponse struct {
	ResponseStatus       Status        `protobuf:"varint,1,opt,name=ResponseStatus,proto3,enum=proto.Status" json:"ResponseStatus,omitempty"`
	Massage              string        `protobuf:"bytes,2,opt,name=Massage,proto3" json:"Massage,omitempty"`
	Data                 []*ABDivision `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ABResponse) Reset()         { *m = ABResponse{} }
func (m *ABResponse) String() string { return proto.CompactTextString(m) }
func (*ABResponse) ProtoMessage()    {}
func (*ABResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{4}
}

func (m *ABResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABResponse.Unmarshal(m, b)
}
func (m *ABResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABResponse.Marshal(b, m, deterministic)
}
func (m *ABResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABResponse.Merge(m, src)
}
func (m *ABResponse) XXX_Size() int {
	return xxx_messageInfo_ABResponse.Size(m)
}
func (m *ABResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ABResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ABResponse proto.InternalMessageInfo

func (m *ABResponse) GetResponseStatus() Status {
	if m != nil {
		return m.ResponseStatus
	}
	return Status_SUCCEED
}

func (m *ABResponse) GetMassage() string {
	if m != nil {
		return m.Massage
	}
	return ""
}

func (m *ABResponse) GetData() []*ABDivision {
	if m != nil {
		return m.Data
	}
	return nil
}

type ExperimentInfo struct {
	ExpId                string   `protobuf:"bytes,1,opt,name=ExpId,proto3" json:"ExpId,omitempty"`
	ExpName              string   `protobuf:"bytes,2,opt,name=ExpName,proto3" json:"ExpName,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	NewUserOnly          bool     `protobuf:"varint,4,opt,name=NewUserOnly,proto3" json:"NewUserOnly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExperimentInfo) Reset()         { *m = ExperimentInfo{} }
func (m *ExperimentInfo) String() string { return proto.CompactTextString(m) }
func (*ExperimentInfo) ProtoMessage()    {}
func (*ExperimentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{5}
}

func (m *ExperimentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentInfo.Unmarshal(m, b)
}
func (m *ExperimentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentInfo.Marshal(b, m, deterministic)
}
func (m *ExperimentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentInfo.Merge(m, src)
}
func (m *ExperimentInfo) XXX_Size() int {
	return xxx_messageInfo_ExperimentInfo.Size(m)
}
func (m *ExperimentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentInfo proto.InternalMessageInfo

func (m *ExperimentInfo) GetExpId() string {
	if m != nil {
		return m.ExpId
	}
	return ""
}

func (m *ExperimentInfo) GetExpName() string {
	if m != nil {
		return m.ExpName
	}
	return ""
}

func (m *ExperimentInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ExperimentInfo) GetNewUserOnly() bool {
	if m != nil {
		return m.NewUserOnly
	}
	return false
}

type ExperimentInfoResponse struct {
	ResponseStatus       Status            `protobuf:"varint,1,opt,name=ResponseStatus,proto3,enum=proto.Status" json:"ResponseStatus,omitempty"`
	Massage              string            `protobuf:"bytes,2,opt,name=Massage,proto3" json:"Massage,omitempty"`
	Data                 []*ExperimentInfo `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExperimentInfoResponse) Reset()         { *m = ExperimentInfoResponse{} }
func (m *ExperimentInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ExperimentInfoResponse) ProtoMessage()    {}
func (*ExperimentInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f675e9d1cd0eb12, []int{6}
}

func (m *ExperimentInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentInfoResponse.Unmarshal(m, b)
}
func (m *ExperimentInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentInfoResponse.Marshal(b, m, deterministic)
}
func (m *ExperimentInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentInfoResponse.Merge(m, src)
}
func (m *ExperimentInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ExperimentInfoResponse.Size(m)
}
func (m *ExperimentInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentInfoResponse proto.InternalMessageInfo

func (m *ExperimentInfoResponse) GetResponseStatus() Status {
	if m != nil {
		return m.ResponseStatus
	}
	return Status_SUCCEED
}

func (m *ExperimentInfoResponse) GetMassage() string {
	if m != nil {
		return m.Massage
	}
	return ""
}

func (m *ExperimentInfoResponse) GetData() []*ExperimentInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.Status", Status_name, Status_value)
	proto.RegisterEnum("proto.Gender", Gender_name, Gender_value)
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*BusinessInfo)(nil), "proto.BusinessInfo")
	proto.RegisterType((*Passthrough)(nil), "proto.Passthrough")
	proto.RegisterType((*ABDivision)(nil), "proto.ABDivision")
	proto.RegisterType((*ABResponse)(nil), "proto.ABResponse")
	proto.RegisterType((*ExperimentInfo)(nil), "proto.ExperimentInfo")
	proto.RegisterType((*ExperimentInfoResponse)(nil), "proto.ExperimentInfoResponse")
}

func init() { proto.RegisterFile("division.proto", fileDescriptor_3f675e9d1cd0eb12) }

var fileDescriptor_3f675e9d1cd0eb12 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xed, 0x60, 0xc7, 0x49, 0xae, 0x69, 0x14, 0x2e, 0x0f, 0x59, 0x11, 0x54, 0xc6, 0x52, 0xa5,
	0xd0, 0x45, 0x24, 0x02, 0x7c, 0x40, 0x1e, 0xa6, 0x8a, 0xd4, 0x96, 0x6a, 0x42, 0xd8, 0xbb, 0x78,
	0x48, 0x46, 0x34, 0x63, 0xcb, 0xe3, 0x94, 0xb0, 0x47, 0xe2, 0x07, 0xf8, 0x03, 0x16, 0xfc, 0x26,
	0x9a, 0xf1, 0x38, 0x76, 0x42, 0xb7, 0xac, 0x7c, 0x5f, 0xbe, 0xe7, 0x9e, 0x73, 0x6c, 0xe8, 0xc4,
	0xfc, 0x8e, 0x4b, 0x9e, 0x88, 0x41, 0x9a, 0x25, 0x79, 0x82, 0x0d, 0xfd, 0x08, 0xfe, 0x10, 0xb0,
	0x17, 0x92, 0x65, 0x88, 0x60, 0x2f, 0x62, 0x1e, 0x7b, 0xc4, 0x27, 0xfd, 0x36, 0xd5, 0x31, 0x76,
	0xc1, 0x5a, 0xf0, 0xd8, 0x7b, 0xe0, 0x93, 0xbe, 0x45, 0x55, 0x88, 0x1e, 0x34, 0x3f, 0xb1, 0x4c,
	0xad, 0xf1, 0x2c, 0x3d, 0x58, 0xa6, 0xaa, 0x33, 0x59, 0x45, 0x42, 0xb0, 0x5b, 0xcf, 0x2e, 0x3a,
	0x26, 0xc5, 0x53, 0x70, 0xce, 0x99, 0x88, 0x59, 0xe6, 0x35, 0x7c, 0xd2, 0xef, 0x0c, 0x8f, 0x8b,
	0x0b, 0x06, 0x45, 0x91, 0x9a, 0x26, 0x9e, 0x00, 0x8c, 0x37, 0x92, 0x0b, 0x26, 0xe5, 0x2c, 0xf6,
	0x1c, 0xbd, 0xa3, 0x56, 0x09, 0x06, 0xf0, 0x70, 0x97, 0x89, 0x2f, 0xc9, 0xc1, 0x3c, 0xf9, 0x67,
	0xfe, 0x37, 0x01, 0xf7, 0x3a, 0x92, 0x32, 0x5f, 0x65, 0xc9, 0x66, 0xb9, 0xc2, 0x27, 0xd0, 0x08,
	0xb7, 0xe9, 0x6e, 0xb4, 0x48, 0xb0, 0x07, 0xad, 0xf1, 0xe6, 0xf3, 0x57, 0x96, 0xcf, 0x0a, 0x9e,
	0x6d, 0xba, 0xcb, 0x0b, 0x04, 0x15, 0x5f, 0x45, 0x6b, 0x66, 0xf8, 0xd6, 0x2a, 0xa5, 0x3c, 0x76,
	0x25, 0x4f, 0x29, 0x62, 0xa3, 0x26, 0xe2, 0x73, 0x68, 0xcf, 0x59, 0x76, 0xc7, 0x3e, 0xf2, 0x35,
	0xd3, 0xb4, 0x2c, 0x5a, 0x15, 0x82, 0x1b, 0x80, 0xd1, 0x78, 0x6a, 0xac, 0x51, 0x88, 0x23, 0x29,
	0xf9, 0x52, 0xac, 0x99, 0xc8, 0x4b, 0x4e, 0x55, 0x05, 0xdf, 0xee, 0x51, 0xd2, 0x07, 0xbb, 0x43,
	0x34, 0x7a, 0xd6, 0x3a, 0xb4, 0x3e, 0x16, 0xfc, 0x20, 0x0a, 0x84, 0x32, 0x99, 0x26, 0x42, 0x32,
	0x7c, 0x07, 0x9d, 0x32, 0x9e, 0xe7, 0x51, 0xbe, 0x91, 0x1a, 0xa8, 0xf2, 0xa5, 0x28, 0xd2, 0x83,
	0x21, 0x65, 0xf0, 0x65, 0x24, 0x65, 0xb4, 0x64, 0x46, 0xa8, 0x32, 0xc5, 0x53, 0xb0, 0xa7, 0x51,
	0x1e, 0x79, 0x96, 0x6f, 0xf5, 0xdd, 0xe1, 0x23, 0xb3, 0xa6, 0xa2, 0x45, 0x75, 0x3b, 0xd8, 0x42,
	0x27, 0xdc, 0xa6, 0x2c, 0xe3, 0x8a, 0x8a, 0xb6, 0xf0, 0x7e, 0x4b, 0x3c, 0x68, 0x86, 0xdb, 0x54,
	0x6b, 0x6e, 0x80, 0x4c, 0xaa, 0xe6, 0xd5, 0x31, 0xa5, 0x17, 0x45, 0x82, 0x3e, 0xb8, 0x57, 0xec,
	0x9b, 0xfa, 0x88, 0x3f, 0x88, 0xdb, 0xef, 0xda, 0x8e, 0x16, 0xad, 0x97, 0x82, 0x5f, 0x04, 0x9e,
	0xed, 0x43, 0xff, 0x3f, 0x31, 0x5e, 0xed, 0x89, 0xf1, 0xd4, 0xac, 0x39, 0x40, 0xd7, 0x23, 0x67,
	0x2f, 0xc1, 0x31, 0xeb, 0x5c, 0x68, 0xce, 0x17, 0x93, 0x49, 0x18, 0x4e, 0xbb, 0x47, 0x08, 0xe0,
	0xbc, 0x1f, 0xcd, 0x2e, 0xc2, 0x69, 0x97, 0x9c, 0x9d, 0x94, 0xff, 0x0e, 0xb6, 0xc0, 0xbe, 0x1c,
	0x5d, 0x84, 0xa6, 0x1f, 0xea, 0x98, 0x0c, 0x7f, 0x12, 0x70, 0x77, 0x32, 0x5f, 0x4f, 0x30, 0x84,
	0xe3, 0x73, 0x96, 0x57, 0x68, 0xf8, 0xd8, 0x1c, 0x50, 0xff, 0x75, 0x7a, 0x2f, 0xee, 0xbf, 0xca,
	0x50, 0x0c, 0x8e, 0xf0, 0xb5, 0x5e, 0x53, 0xfb, 0x30, 0x5d, 0xf3, 0x86, 0x92, 0xb4, 0x57, 0x39,
	0x5c, 0xbd, 0x72, 0xe3, 0xe8, 0xda, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x6d, 0xfa,
	0x0a, 0x68, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DivisionRPCClient is the client API for DivisionRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DivisionRPCClient interface {
	GetExperiment(ctx context.Context, in *BusinessInfo, opts ...grpc.CallOption) (*ExperimentInfoResponse, error)
	GetABDivision(ctx context.Context, in *User, opts ...grpc.CallOption) (*ABResponse, error)
}

type divisionRPCClient struct {
	cc *grpc.ClientConn
}

func NewDivisionRPCClient(cc *grpc.ClientConn) DivisionRPCClient {
	return &divisionRPCClient{cc}
}

func (c *divisionRPCClient) GetExperiment(ctx context.Context, in *BusinessInfo, opts ...grpc.CallOption) (*ExperimentInfoResponse, error) {
	out := new(ExperimentInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.DivisionRPC/GetExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *divisionRPCClient) GetABDivision(ctx context.Context, in *User, opts ...grpc.CallOption) (*ABResponse, error) {
	out := new(ABResponse)
	err := c.cc.Invoke(ctx, "/proto.DivisionRPC/GetABDivision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DivisionRPCServer is the server API for DivisionRPC service.
type DivisionRPCServer interface {
	GetExperiment(context.Context, *BusinessInfo) (*ExperimentInfoResponse, error)
	GetABDivision(context.Context, *User) (*ABResponse, error)
}

func RegisterDivisionRPCServer(s *grpc.Server, srv DivisionRPCServer) {
	s.RegisterService(&_DivisionRPC_serviceDesc, srv)
}

func _DivisionRPC_GetExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivisionRPCServer).GetExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DivisionRPC/GetExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivisionRPCServer).GetExperiment(ctx, req.(*BusinessInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DivisionRPC_GetABDivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivisionRPCServer).GetABDivision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DivisionRPC/GetABDivision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivisionRPCServer).GetABDivision(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _DivisionRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DivisionRPC",
	HandlerType: (*DivisionRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExperiment",
			Handler:    _DivisionRPC_GetExperiment_Handler,
		},
		{
			MethodName: "GetABDivision",
			Handler:    _DivisionRPC_GetABDivision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "division.proto",
}
