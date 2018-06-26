// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alddle-matcher.proto

package samdasualddle

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Departure            string   `protobuf:"bytes,2,opt,name=departure,proto3" json:"departure,omitempty"`
	Destination          string   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	Expense              int32    `protobuf:"varint,4,opt,name=expense,proto3" json:"expense,omitempty"`
	Duration             int32    `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	FromDate             string   `protobuf:"bytes,6,opt,name=fromDate,proto3" json:"fromDate,omitempty"`
	ToDate               string   `protobuf:"bytes,7,opt,name=toDate,proto3" json:"toDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_alddle_matcher_4f555d82480e6da6, []int{0}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetDeparture() string {
	if m != nil {
		return m.Departure
	}
	return ""
}

func (m *RegisterRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *RegisterRequest) GetExpense() int32 {
	if m != nil {
		return m.Expense
	}
	return 0
}

func (m *RegisterRequest) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *RegisterRequest) GetFromDate() string {
	if m != nil {
		return m.FromDate
	}
	return ""
}

func (m *RegisterRequest) GetToDate() string {
	if m != nil {
		return m.ToDate
	}
	return ""
}

type RegisterReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_alddle_matcher_4f555d82480e6da6, []int{1}
}
func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (dst *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(dst, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

type UnregisterRequest struct {
	RegisterId           string   `protobuf:"bytes,1,opt,name=registerId,proto3" json:"registerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterRequest) Reset()         { *m = UnregisterRequest{} }
func (m *UnregisterRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterRequest) ProtoMessage()    {}
func (*UnregisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_alddle_matcher_4f555d82480e6da6, []int{2}
}
func (m *UnregisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterRequest.Unmarshal(m, b)
}
func (m *UnregisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterRequest.Marshal(b, m, deterministic)
}
func (dst *UnregisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterRequest.Merge(dst, src)
}
func (m *UnregisterRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterRequest.Size(m)
}
func (m *UnregisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterRequest proto.InternalMessageInfo

func (m *UnregisterRequest) GetRegisterId() string {
	if m != nil {
		return m.RegisterId
	}
	return ""
}

type UnregisterReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterReply) Reset()         { *m = UnregisterReply{} }
func (m *UnregisterReply) String() string { return proto.CompactTextString(m) }
func (*UnregisterReply) ProtoMessage()    {}
func (*UnregisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_alddle_matcher_4f555d82480e6da6, []int{3}
}
func (m *UnregisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterReply.Unmarshal(m, b)
}
func (m *UnregisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterReply.Marshal(b, m, deterministic)
}
func (dst *UnregisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterReply.Merge(dst, src)
}
func (m *UnregisterReply) XXX_Size() int {
	return xxx_messageInfo_UnregisterReply.Size(m)
}
func (m *UnregisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterReply proto.InternalMessageInfo

func (m *UnregisterReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MatchAndNotifyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchAndNotifyRequest) Reset()         { *m = MatchAndNotifyRequest{} }
func (m *MatchAndNotifyRequest) String() string { return proto.CompactTextString(m) }
func (*MatchAndNotifyRequest) ProtoMessage()    {}
func (*MatchAndNotifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_alddle_matcher_4f555d82480e6da6, []int{4}
}
func (m *MatchAndNotifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchAndNotifyRequest.Unmarshal(m, b)
}
func (m *MatchAndNotifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchAndNotifyRequest.Marshal(b, m, deterministic)
}
func (dst *MatchAndNotifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchAndNotifyRequest.Merge(dst, src)
}
func (m *MatchAndNotifyRequest) XXX_Size() int {
	return xxx_messageInfo_MatchAndNotifyRequest.Size(m)
}
func (m *MatchAndNotifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchAndNotifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MatchAndNotifyRequest proto.InternalMessageInfo

type MatchAndNotifyReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchAndNotifyReply) Reset()         { *m = MatchAndNotifyReply{} }
func (m *MatchAndNotifyReply) String() string { return proto.CompactTextString(m) }
func (*MatchAndNotifyReply) ProtoMessage()    {}
func (*MatchAndNotifyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_alddle_matcher_4f555d82480e6da6, []int{5}
}
func (m *MatchAndNotifyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchAndNotifyReply.Unmarshal(m, b)
}
func (m *MatchAndNotifyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchAndNotifyReply.Marshal(b, m, deterministic)
}
func (dst *MatchAndNotifyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchAndNotifyReply.Merge(dst, src)
}
func (m *MatchAndNotifyReply) XXX_Size() int {
	return xxx_messageInfo_MatchAndNotifyReply.Size(m)
}
func (m *MatchAndNotifyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchAndNotifyReply.DiscardUnknown(m)
}

var xxx_messageInfo_MatchAndNotifyReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "samdasualddle.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "samdasualddle.RegisterReply")
	proto.RegisterType((*UnregisterRequest)(nil), "samdasualddle.UnregisterRequest")
	proto.RegisterType((*UnregisterReply)(nil), "samdasualddle.UnregisterReply")
	proto.RegisterType((*MatchAndNotifyRequest)(nil), "samdasualddle.MatchAndNotifyRequest")
	proto.RegisterType((*MatchAndNotifyReply)(nil), "samdasualddle.MatchAndNotifyReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlddleClient is the client API for Alddle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlddleClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterReply, error)
	MatchAndNotify(ctx context.Context, in *MatchAndNotifyRequest, opts ...grpc.CallOption) (*MatchAndNotifyReply, error)
}

type alddleClient struct {
	cc *grpc.ClientConn
}

func NewAlddleClient(cc *grpc.ClientConn) AlddleClient {
	return &alddleClient{cc}
}

func (c *alddleClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/samdasualddle.Alddle/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alddleClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterReply, error) {
	out := new(UnregisterReply)
	err := c.cc.Invoke(ctx, "/samdasualddle.Alddle/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alddleClient) MatchAndNotify(ctx context.Context, in *MatchAndNotifyRequest, opts ...grpc.CallOption) (*MatchAndNotifyReply, error) {
	out := new(MatchAndNotifyReply)
	err := c.cc.Invoke(ctx, "/samdasualddle.Alddle/MatchAndNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlddleServer is the server API for Alddle service.
type AlddleServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterReply, error)
	MatchAndNotify(context.Context, *MatchAndNotifyRequest) (*MatchAndNotifyReply, error)
}

func RegisterAlddleServer(s *grpc.Server, srv AlddleServer) {
	s.RegisterService(&_Alddle_serviceDesc, srv)
}

func _Alddle_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlddleServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samdasualddle.Alddle/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlddleServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alddle_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlddleServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samdasualddle.Alddle/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlddleServer).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alddle_MatchAndNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchAndNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlddleServer).MatchAndNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samdasualddle.Alddle/MatchAndNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlddleServer).MatchAndNotify(ctx, req.(*MatchAndNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alddle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "samdasualddle.Alddle",
	HandlerType: (*AlddleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Alddle_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _Alddle_Unregister_Handler,
		},
		{
			MethodName: "MatchAndNotify",
			Handler:    _Alddle_MatchAndNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alddle-matcher.proto",
}

func init() {
	proto.RegisterFile("alddle-matcher.proto", fileDescriptor_alddle_matcher_4f555d82480e6da6)
}

var fileDescriptor_alddle_matcher_4f555d82480e6da6 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0xb5, 0x69, 0x3b, 0x52, 0x8b, 0x63, 0xab, 0x21, 0x94, 0x12, 0x16, 0x0f, 0x05,
	0xb1, 0x07, 0xfb, 0x0b, 0x0a, 0x5e, 0x14, 0x14, 0x09, 0x78, 0xf3, 0xb2, 0xba, 0xd3, 0x1a, 0xc8,
	0x97, 0xbb, 0x1b, 0x30, 0x77, 0xff, 0xa0, 0xff, 0x48, 0xb2, 0x49, 0xfa, 0x11, 0xbf, 0x8e, 0xef,
	0x3c, 0xef, 0x4c, 0xe6, 0xdd, 0x09, 0x0c, 0x79, 0x28, 0x44, 0x48, 0x97, 0x11, 0xd7, 0x2f, 0xaf,
	0x24, 0x67, 0xa9, 0x4c, 0x74, 0x82, 0x7d, 0xc5, 0x23, 0xc1, 0x55, 0x56, 0x42, 0xf6, 0x69, 0xc1,
	0xc0, 0xa7, 0x55, 0xa0, 0x34, 0x49, 0x9f, 0xde, 0x32, 0x52, 0x1a, 0x87, 0xd0, 0xa6, 0x88, 0x07,
	0xa1, 0x63, 0x79, 0xd6, 0xb4, 0xe7, 0x97, 0x02, 0xc7, 0xd0, 0x13, 0x94, 0x72, 0xa9, 0x33, 0x49,
	0x4e, 0xcb, 0x90, 0x4d, 0x01, 0x3d, 0x38, 0x14, 0xa4, 0x74, 0x10, 0x73, 0x1d, 0x24, 0xb1, 0xb3,
	0x6f, 0xf8, 0x76, 0x09, 0x1d, 0xe8, 0xd0, 0x7b, 0x4a, 0xb1, 0x22, 0xe7, 0xc0, 0xb3, 0xa6, 0x6d,
	0xbf, 0x96, 0xe8, 0x42, 0x57, 0x64, 0xb2, 0x6c, 0x6c, 0x1b, 0xb4, 0xd6, 0x05, 0x5b, 0xca, 0x24,
	0xba, 0xe6, 0x9a, 0x1c, 0xdb, 0x0c, 0x5d, 0x6b, 0x3c, 0x05, 0x5b, 0x27, 0x86, 0x74, 0x0c, 0xa9,
	0x14, 0x1b, 0x40, 0x7f, 0x13, 0x29, 0x0d, 0x73, 0x36, 0x87, 0xe3, 0xc7, 0x58, 0x36, 0x52, 0x4e,
	0x00, 0xea, 0xd2, 0x8d, 0xa8, 0xa2, 0x6e, 0x55, 0xd8, 0x05, 0x0c, 0xb6, 0x9b, 0xd2, 0x30, 0x2f,
	0x22, 0x44, 0xa4, 0x14, 0x5f, 0x51, 0xe5, 0xaf, 0x25, 0x3b, 0x83, 0xd1, 0x5d, 0xf1, 0xcc, 0x8b,
	0x58, 0xdc, 0x27, 0x3a, 0x58, 0xe6, 0xd5, 0x57, 0xd8, 0x08, 0x4e, 0x9a, 0x20, 0x0d, 0xf3, 0xab,
	0x8f, 0x16, 0xd8, 0x0b, 0x73, 0x01, 0xbc, 0x85, 0x6e, 0xbd, 0x2d, 0x4e, 0x66, 0x3b, 0xd7, 0x99,
	0x35, 0x2e, 0xe3, 0x8e, 0x7f, 0xe5, 0x45, 0xcc, 0x3d, 0x7c, 0x00, 0xd8, 0xec, 0x8c, 0x5e, 0xc3,
	0xfd, 0xed, 0x0d, 0xdc, 0xc9, 0x1f, 0x8e, 0x72, 0xe2, 0x13, 0x1c, 0xed, 0xee, 0x8f, 0xe7, 0x8d,
	0x9e, 0x1f, 0x73, 0xbb, 0xec, 0x1f, 0x97, 0x99, 0xfe, 0x6c, 0x9b, 0x7f, 0x72, 0xfe, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x8a, 0x88, 0x67, 0xab, 0x02, 0x00, 0x00,
}
