// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_2b55273352d5f36f, []int{0}
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

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
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
	return fileDescriptor_auth_2b55273352d5f36f, []int{1}
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
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Status)(nil), "auth.Status")
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
	LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Status, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/auth.AuthChecker/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	LoginUser(context.Context, *User) (*Status, error)
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthChecker/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).LoginUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _AuthChecker_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_2b55273352d5f36f) }

var fileDescriptor_auth_2b55273352d5f36f = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xec, 0xb8, 0x58, 0x42, 0x8b,
	0x53, 0x8b, 0x84, 0xa4, 0xb8, 0x38, 0xfc, 0x32, 0x93, 0xb3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x40, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51,
	0x8a, 0x04, 0x13, 0x44, 0x0e, 0xc6, 0x57, 0x92, 0xe0, 0x62, 0x0b, 0x2e, 0x49, 0x2c, 0x29, 0x2d,
	0x16, 0xe2, 0xe3, 0x62, 0xf2, 0xcf, 0x06, 0xeb, 0xe5, 0x08, 0x02, 0xb2, 0x8c, 0xcc, 0xb8, 0xb8,
	0x1d, 0x81, 0x36, 0x38, 0x67, 0xa4, 0x26, 0x67, 0x03, 0x2d, 0x50, 0xe7, 0xe2, 0xf4, 0xc9, 0x4f,
	0xcf, 0xcc, 0x03, 0xdb, 0xc6, 0xa5, 0x07, 0x76, 0x08, 0x88, 0x2d, 0xc5, 0x03, 0x61, 0x43, 0x4c,
	0x51, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xe6, 0x4c,
	0x6c, 0xac, 0x00, 0x00, 0x00,
}
