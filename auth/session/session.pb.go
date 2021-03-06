// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package session

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

type User struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_0b31544a80ba5e48, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type Session struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_0b31544a80ba5e48, []int{1}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Status struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_0b31544a80ba5e48, []int{2}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "session.User")
	proto.RegisterType((*Session)(nil), "session.Session")
	proto.RegisterType((*Status)(nil), "session.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthCheckerClient interface {
	CreateSession(ctx context.Context, in *User, opts ...grpc.CallOption) (*Session, error)
	DestroySession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Status, error)
	ValidateSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Status, error)
	GetOwner(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) CreateSession(ctx context.Context, in *User, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) DestroySession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/DestroySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) ValidateSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/ValidateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetOwner(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	CreateSession(context.Context, *User) (*Session, error)
	DestroySession(context.Context, *Session) (*Status, error)
	ValidateSession(context.Context, *Session) (*Status, error)
	GetOwner(context.Context, *Session) (*User, error)
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).CreateSession(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_DestroySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).DestroySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/DestroySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).DestroySession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_ValidateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).ValidateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/ValidateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).ValidateSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetOwner(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _AuthChecker_CreateSession_Handler,
		},
		{
			MethodName: "DestroySession",
			Handler:    _AuthChecker_DestroySession_Handler,
		},
		{
			MethodName: "ValidateSession",
			Handler:    _AuthChecker_ValidateSession_Handler,
		},
		{
			MethodName: "GetOwner",
			Handler:    _AuthChecker_GetOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}

func init() { proto.RegisterFile("session.proto", fileDescriptor_session_0b31544a80ba5e48) }

var fileDescriptor_session_0b31544a80ba5e48 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x94, 0xb8,
	0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0xa4, 0xb8, 0x38, 0xfc, 0x32, 0x93, 0xb3, 0xf3, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x49, 0x2e, 0xf6, 0x60, 0x88,
	0x72, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0xa8, 0x02, 0x20, 0x4b, 0x49, 0x82, 0x8b, 0x2d, 0xb8,
	0x24, 0xb1, 0xa4, 0xb4, 0x18, 0x24, 0xe3, 0x9f, 0x0d, 0x96, 0xe1, 0x08, 0x02, 0xb2, 0x8c, 0xee,
	0x33, 0x72, 0x71, 0x3b, 0x96, 0x96, 0x64, 0x38, 0x67, 0xa4, 0x26, 0x67, 0x03, 0x2d, 0x30, 0xe2,
	0xe2, 0x75, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x85, 0x19, 0xc5, 0xab, 0x07, 0x73, 0x12, 0xc8, 0x01,
	0x52, 0x02, 0x70, 0x2e, 0x54, 0x81, 0x12, 0x83, 0x90, 0x29, 0x17, 0x9f, 0x4b, 0x6a, 0x71, 0x49,
	0x51, 0x7e, 0x25, 0x4c, 0x13, 0x86, 0x2a, 0x29, 0x7e, 0x84, 0x08, 0xd8, 0x21, 0x40, 0x6d, 0x66,
	0x5c, 0xfc, 0x61, 0x89, 0x39, 0x99, 0x29, 0x48, 0x96, 0x11, 0xa5, 0x4f, 0x97, 0x8b, 0xc3, 0x3d,
	0xb5, 0xc4, 0xbf, 0x3c, 0x0f, 0xe8, 0x5c, 0x4c, 0x0d, 0xa8, 0xee, 0x55, 0x62, 0x48, 0x62, 0x03,
	0x07, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xda, 0x62, 0x6e, 0x5b, 0x01, 0x00, 0x00,
}
