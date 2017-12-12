// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LandlordsProto.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	LandlordsProto.proto

It has these top-level messages:
	CS_PlayerLogin
	SC_PlayerLogin
	CS_PlayerRegister
	SC_PlayerRegister
*/
package msg

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

type SC_PlayerLogin_LoginResult int32

const (
	SC_PlayerLogin_SUCCESS       SC_PlayerLogin_LoginResult = 0
	SC_PlayerLogin_ACCOUNT_ERROR SC_PlayerLogin_LoginResult = 1
	SC_PlayerLogin_PWD_ERROR     SC_PlayerLogin_LoginResult = 2
)

var SC_PlayerLogin_LoginResult_name = map[int32]string{
	0: "SUCCESS",
	1: "ACCOUNT_ERROR",
	2: "PWD_ERROR",
}
var SC_PlayerLogin_LoginResult_value = map[string]int32{
	"SUCCESS":       0,
	"ACCOUNT_ERROR": 1,
	"PWD_ERROR":     2,
}

func (x SC_PlayerLogin_LoginResult) String() string {
	return proto.EnumName(SC_PlayerLogin_LoginResult_name, int32(x))
}
func (SC_PlayerLogin_LoginResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type SC_PlayerRegister_RegisterResult int32

const (
	SC_PlayerRegister_SUCCESS       SC_PlayerRegister_RegisterResult = 0
	SC_PlayerRegister_ACCOUNT_ERROR SC_PlayerRegister_RegisterResult = 1
)

var SC_PlayerRegister_RegisterResult_name = map[int32]string{
	0: "SUCCESS",
	1: "ACCOUNT_ERROR",
}
var SC_PlayerRegister_RegisterResult_value = map[string]int32{
	"SUCCESS":       0,
	"ACCOUNT_ERROR": 1,
}

func (x SC_PlayerRegister_RegisterResult) String() string {
	return proto.EnumName(SC_PlayerRegister_RegisterResult_name, int32(x))
}
func (SC_PlayerRegister_RegisterResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type CS_PlayerLogin struct {
	Account string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Pwd     string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *CS_PlayerLogin) Reset()                    { *m = CS_PlayerLogin{} }
func (m *CS_PlayerLogin) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerLogin) ProtoMessage()               {}
func (*CS_PlayerLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CS_PlayerLogin) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CS_PlayerLogin) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type SC_PlayerLogin struct {
	LoginResult SC_PlayerLogin_LoginResult `protobuf:"varint,1,opt,name=loginResult,enum=msg.SC_PlayerLogin_LoginResult" json:"loginResult,omitempty"`
}

func (m *SC_PlayerLogin) Reset()                    { *m = SC_PlayerLogin{} }
func (m *SC_PlayerLogin) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerLogin) ProtoMessage()               {}
func (*SC_PlayerLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SC_PlayerLogin) GetLoginResult() SC_PlayerLogin_LoginResult {
	if m != nil {
		return m.LoginResult
	}
	return SC_PlayerLogin_SUCCESS
}

type CS_PlayerRegister struct {
	Account string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Pwd     string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *CS_PlayerRegister) Reset()                    { *m = CS_PlayerRegister{} }
func (m *CS_PlayerRegister) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerRegister) ProtoMessage()               {}
func (*CS_PlayerRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CS_PlayerRegister) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CS_PlayerRegister) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type SC_PlayerRegister struct {
	RegisterResult SC_PlayerRegister_RegisterResult `protobuf:"varint,1,opt,name=registerResult,enum=msg.SC_PlayerRegister_RegisterResult" json:"registerResult,omitempty"`
}

func (m *SC_PlayerRegister) Reset()                    { *m = SC_PlayerRegister{} }
func (m *SC_PlayerRegister) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerRegister) ProtoMessage()               {}
func (*SC_PlayerRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SC_PlayerRegister) GetRegisterResult() SC_PlayerRegister_RegisterResult {
	if m != nil {
		return m.RegisterResult
	}
	return SC_PlayerRegister_SUCCESS
}

func init() {
	proto.RegisterType((*CS_PlayerLogin)(nil), "msg.CS_PlayerLogin")
	proto.RegisterType((*SC_PlayerLogin)(nil), "msg.SC_PlayerLogin")
	proto.RegisterType((*CS_PlayerRegister)(nil), "msg.CS_PlayerRegister")
	proto.RegisterType((*SC_PlayerRegister)(nil), "msg.SC_PlayerRegister")
	proto.RegisterEnum("msg.SC_PlayerLogin_LoginResult", SC_PlayerLogin_LoginResult_name, SC_PlayerLogin_LoginResult_value)
	proto.RegisterEnum("msg.SC_PlayerRegister_RegisterResult", SC_PlayerRegister_RegisterResult_name, SC_PlayerRegister_RegisterResult_value)
}

func init() { proto.RegisterFile("LandlordsProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf1, 0x49, 0xcc, 0x4b,
	0xc9, 0xc9, 0x2f, 0x4a, 0x29, 0x0e, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2b, 0x00, 0x91, 0x42, 0xcc,
	0xb9, 0xc5, 0xe9, 0x4a, 0x36, 0x5c, 0x7c, 0xce, 0xc1, 0xf1, 0x01, 0x39, 0x89, 0x95, 0xa9, 0x45,
	0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x42, 0x12, 0x5c, 0xec, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x00, 0x17, 0x73, 0x41, 0x79, 0x8a,
	0x04, 0x13, 0x58, 0x14, 0xc4, 0x54, 0x9a, 0xc8, 0xc8, 0xc5, 0x17, 0xec, 0x8c, 0xa2, 0xdd, 0x91,
	0x8b, 0x3b, 0x07, 0xc4, 0x08, 0x4a, 0x2d, 0x2e, 0xcd, 0x81, 0x18, 0xc1, 0x67, 0x24, 0xaf, 0x97,
	0x5b, 0x9c, 0xae, 0x87, 0xaa, 0x52, 0xcf, 0x07, 0xa1, 0x2c, 0x08, 0x59, 0x8f, 0x92, 0x0d, 0x17,
	0x37, 0x92, 0x9c, 0x10, 0x37, 0x17, 0x7b, 0x70, 0xa8, 0xb3, 0xb3, 0x6b, 0x70, 0xb0, 0x00, 0x83,
	0x90, 0x20, 0x17, 0xaf, 0xa3, 0xb3, 0xb3, 0x7f, 0xa8, 0x5f, 0x48, 0xbc, 0x6b, 0x50, 0x90, 0x7f,
	0x90, 0x00, 0xa3, 0x10, 0x2f, 0x17, 0x67, 0x40, 0xb8, 0x0b, 0x94, 0xcb, 0xa4, 0x64, 0xcf, 0x25,
	0x08, 0xf7, 0x51, 0x50, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x11, 0x49, 0x9e, 0x9a, 0xc2, 0xc8,
	0x25, 0x08, 0x77, 0x2a, 0xdc, 0x04, 0x5f, 0x2e, 0xbe, 0x22, 0x28, 0x1b, 0xc5, 0x6b, 0xaa, 0xa8,
	0x5e, 0x83, 0xa9, 0xd7, 0x0b, 0x42, 0x51, 0x1c, 0x84, 0xa6, 0x59, 0xc9, 0x80, 0x8b, 0x0f, 0x55,
	0x05, 0x21, 0x6f, 0x26, 0xb1, 0x81, 0x63, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x82,
	0x67, 0x3a, 0xcd, 0x01, 0x00, 0x00,
}