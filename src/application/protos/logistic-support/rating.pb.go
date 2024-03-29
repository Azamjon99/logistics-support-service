// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rating.proto

package logistic_support

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

type Rating struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=order_id,proto3" json:"order_id,omitempty"`
	Rating               int32    `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment              string   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rating) Reset()         { *m = Rating{} }
func (m *Rating) String() string { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()    {}
func (*Rating) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98441b067dcd23a, []int{0}
}

func (m *Rating) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rating.Unmarshal(m, b)
}
func (m *Rating) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rating.Marshal(b, m, deterministic)
}
func (m *Rating) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rating.Merge(m, src)
}
func (m *Rating) XXX_Size() int {
	return xxx_messageInfo_Rating.Size(m)
}
func (m *Rating) XXX_DiscardUnknown() {
	xxx_messageInfo_Rating.DiscardUnknown(m)
}

var xxx_messageInfo_Rating proto.InternalMessageInfo

func (m *Rating) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rating) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Rating) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Rating) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Rating) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Rating) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreateRatingRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,proto3" json:"order_id,omitempty"`
	RatingValue          int32    `protobuf:"varint,2,opt,name=rating_value,proto3" json:"rating_value,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRatingRequest) Reset()         { *m = CreateRatingRequest{} }
func (m *CreateRatingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRatingRequest) ProtoMessage()    {}
func (*CreateRatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98441b067dcd23a, []int{1}
}

func (m *CreateRatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRatingRequest.Unmarshal(m, b)
}
func (m *CreateRatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRatingRequest.Marshal(b, m, deterministic)
}
func (m *CreateRatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRatingRequest.Merge(m, src)
}
func (m *CreateRatingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRatingRequest.Size(m)
}
func (m *CreateRatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRatingRequest proto.InternalMessageInfo

func (m *CreateRatingRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *CreateRatingRequest) GetRatingValue() int32 {
	if m != nil {
		return m.RatingValue
	}
	return 0
}

func (m *CreateRatingRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type UpdateRatingRequest struct {
	RatingId             string   `protobuf:"bytes,1,opt,name=rating_id,proto3" json:"rating_id,omitempty"`
	RatingValue          int32    `protobuf:"varint,2,opt,name=rating_value,proto3" json:"rating_value,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRatingRequest) Reset()         { *m = UpdateRatingRequest{} }
func (m *UpdateRatingRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRatingRequest) ProtoMessage()    {}
func (*UpdateRatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98441b067dcd23a, []int{2}
}

func (m *UpdateRatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRatingRequest.Unmarshal(m, b)
}
func (m *UpdateRatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRatingRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRatingRequest.Merge(m, src)
}
func (m *UpdateRatingRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRatingRequest.Size(m)
}
func (m *UpdateRatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRatingRequest proto.InternalMessageInfo

func (m *UpdateRatingRequest) GetRatingId() string {
	if m != nil {
		return m.RatingId
	}
	return ""
}

func (m *UpdateRatingRequest) GetRatingValue() int32 {
	if m != nil {
		return m.RatingValue
	}
	return 0
}

func (m *UpdateRatingRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type GetRatingRequest struct {
	RatingId             string   `protobuf:"bytes,1,opt,name=rating_id,proto3" json:"rating_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRatingRequest) Reset()         { *m = GetRatingRequest{} }
func (m *GetRatingRequest) String() string { return proto.CompactTextString(m) }
func (*GetRatingRequest) ProtoMessage()    {}
func (*GetRatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98441b067dcd23a, []int{3}
}

func (m *GetRatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRatingRequest.Unmarshal(m, b)
}
func (m *GetRatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRatingRequest.Marshal(b, m, deterministic)
}
func (m *GetRatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRatingRequest.Merge(m, src)
}
func (m *GetRatingRequest) XXX_Size() int {
	return xxx_messageInfo_GetRatingRequest.Size(m)
}
func (m *GetRatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRatingRequest proto.InternalMessageInfo

func (m *GetRatingRequest) GetRatingId() string {
	if m != nil {
		return m.RatingId
	}
	return ""
}

type GetRatingByOrderRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRatingByOrderRequest) Reset()         { *m = GetRatingByOrderRequest{} }
func (m *GetRatingByOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetRatingByOrderRequest) ProtoMessage()    {}
func (*GetRatingByOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98441b067dcd23a, []int{4}
}

func (m *GetRatingByOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRatingByOrderRequest.Unmarshal(m, b)
}
func (m *GetRatingByOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRatingByOrderRequest.Marshal(b, m, deterministic)
}
func (m *GetRatingByOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRatingByOrderRequest.Merge(m, src)
}
func (m *GetRatingByOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetRatingByOrderRequest.Size(m)
}
func (m *GetRatingByOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRatingByOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRatingByOrderRequest proto.InternalMessageInfo

func (m *GetRatingByOrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type RatingResponse struct {
	Rating               *Rating  `protobuf:"bytes,1,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RatingResponse) Reset()         { *m = RatingResponse{} }
func (m *RatingResponse) String() string { return proto.CompactTextString(m) }
func (*RatingResponse) ProtoMessage()    {}
func (*RatingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98441b067dcd23a, []int{5}
}

func (m *RatingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RatingResponse.Unmarshal(m, b)
}
func (m *RatingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RatingResponse.Marshal(b, m, deterministic)
}
func (m *RatingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RatingResponse.Merge(m, src)
}
func (m *RatingResponse) XXX_Size() int {
	return xxx_messageInfo_RatingResponse.Size(m)
}
func (m *RatingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RatingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RatingResponse proto.InternalMessageInfo

func (m *RatingResponse) GetRating() *Rating {
	if m != nil {
		return m.Rating
	}
	return nil
}

func init() {
	proto.RegisterType((*Rating)(nil), "Rating")
	proto.RegisterType((*CreateRatingRequest)(nil), "CreateRatingRequest")
	proto.RegisterType((*UpdateRatingRequest)(nil), "UpdateRatingRequest")
	proto.RegisterType((*GetRatingRequest)(nil), "GetRatingRequest")
	proto.RegisterType((*GetRatingByOrderRequest)(nil), "GetRatingByOrderRequest")
	proto.RegisterType((*RatingResponse)(nil), "RatingResponse")
}

func init() { proto.RegisterFile("rating.proto", fileDescriptor_f98441b067dcd23a) }

var fileDescriptor_f98441b067dcd23a = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xe5, 0xf4, 0x4f, 0xfa, 0x77, 0x40, 0x15, 0x72, 0x11, 0x58, 0x08, 0x41, 0xe5, 0x55,
	0x37, 0x84, 0x9b, 0x78, 0x81, 0xb2, 0x60, 0x89, 0x64, 0x89, 0x0d, 0x9b, 0x28, 0xc4, 0x56, 0x65,
	0xa9, 0x8d, 0x5d, 0x7b, 0x82, 0xc4, 0x0b, 0xf1, 0x9c, 0x28, 0xce, 0xa5, 0x0d, 0x2b, 0x10, 0xcb,
	0xf9, 0xce, 0xcc, 0x9c, 0x33, 0x89, 0xe1, 0xd0, 0xe5, 0xa8, 0xcb, 0x55, 0x6a, 0x9d, 0x41, 0xc3,
	0x3f, 0x09, 0x24, 0x22, 0x00, 0x3a, 0x85, 0x48, 0x4b, 0x46, 0xe6, 0x64, 0x31, 0x11, 0x91, 0x96,
	0xf4, 0x0c, 0xfe, 0x1b, 0x27, 0x95, 0xcb, 0xb4, 0x64, 0x51, 0xa0, 0x7d, 0x4d, 0x4f, 0x20, 0x69,
	0xd6, 0xb0, 0xd1, 0x9c, 0x2c, 0x62, 0xd1, 0x56, 0x94, 0xc1, 0xb8, 0x30, 0x9b, 0x8d, 0x2a, 0x91,
	0xfd, 0x0b, 0x23, 0x5d, 0x49, 0x2f, 0x00, 0x0a, 0xa7, 0x72, 0x54, 0x32, 0xcb, 0x91, 0xc5, 0x41,
	0xdc, 0x23, 0xb5, 0x5e, 0x59, 0xd9, 0xe9, 0x49, 0xa3, 0xef, 0x08, 0x37, 0x30, 0x7b, 0x0c, 0xdd,
	0x4d, 0x5a, 0xa1, 0xb6, 0x95, 0xf2, 0x38, 0x08, 0x49, 0xbe, 0x85, 0xe4, 0xdd, 0xad, 0xd9, 0x7b,
	0xbe, 0xae, 0x54, 0x38, 0x22, 0x16, 0x03, 0xb6, 0x1f, 0x78, 0x34, 0x08, 0xcc, 0xb7, 0x30, 0x7b,
	0x09, 0xf6, 0x43, 0xc3, 0x73, 0x98, 0xb4, 0x0b, 0x7a, 0xc7, 0x1d, 0xf8, 0xa3, 0xe5, 0x0d, 0x1c,
	0x3d, 0x29, 0xfc, 0x85, 0x1f, 0x7f, 0x80, 0xd3, 0x7e, 0x62, 0xf9, 0xf1, 0x5c, 0x5f, 0xfe, 0x83,
	0x2f, 0xc3, 0x6f, 0x61, 0xda, 0xb9, 0x78, 0x6b, 0x4a, 0xaf, 0xe8, 0x65, 0xff, 0x43, 0xeb, 0xde,
	0x83, 0xbb, 0x71, 0xda, 0x36, 0xb4, 0x78, 0x79, 0xfc, 0x4a, 0xd3, 0xeb, 0xb5, 0x59, 0x69, 0x8f,
	0xba, 0xb8, 0xf2, 0x95, 0xb5, 0xc6, 0xe1, 0x5b, 0x12, 0x5e, 0xd1, 0xfd, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0xd6, 0x62, 0xd5, 0x55, 0x02, 0x00, 0x00,
}
