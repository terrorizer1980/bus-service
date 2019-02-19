// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relay_commit.proto

package force_relay_commit

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

type RelayCommitRequest struct {
	Block                *RelayBlock    `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Action               []*RelayAction `protobuf:"bytes,2,rep,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RelayCommitRequest) Reset()         { *m = RelayCommitRequest{} }
func (m *RelayCommitRequest) String() string { return proto.CompactTextString(m) }
func (*RelayCommitRequest) ProtoMessage()    {}
func (*RelayCommitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_commit_01051abe939b6f32, []int{0}
}
func (m *RelayCommitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelayCommitRequest.Unmarshal(m, b)
}
func (m *RelayCommitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelayCommitRequest.Marshal(b, m, deterministic)
}
func (dst *RelayCommitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayCommitRequest.Merge(dst, src)
}
func (m *RelayCommitRequest) XXX_Size() int {
	return xxx_messageInfo_RelayCommitRequest.Size(m)
}
func (m *RelayCommitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayCommitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelayCommitRequest proto.InternalMessageInfo

func (m *RelayCommitRequest) GetBlock() *RelayBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *RelayCommitRequest) GetAction() []*RelayAction {
	if m != nil {
		return m.Action
	}
	return nil
}

type RelayBlock struct {
	Producer             uint64   `protobuf:"varint,1,opt,name=producer,proto3" json:"producer,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Previous             string   `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	Confirmed            int32    `protobuf:"varint,4,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	TransactionMroot     string   `protobuf:"bytes,5,opt,name=transaction_mroot,json=transactionMroot,proto3" json:"transaction_mroot,omitempty"`
	ActionMroot          string   `protobuf:"bytes,6,opt,name=action_mroot,json=actionMroot,proto3" json:"action_mroot,omitempty"`
	Mroot                string   `protobuf:"bytes,7,opt,name=mroot,proto3" json:"mroot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelayBlock) Reset()         { *m = RelayBlock{} }
func (m *RelayBlock) String() string { return proto.CompactTextString(m) }
func (*RelayBlock) ProtoMessage()    {}
func (*RelayBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_commit_01051abe939b6f32, []int{1}
}
func (m *RelayBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelayBlock.Unmarshal(m, b)
}
func (m *RelayBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelayBlock.Marshal(b, m, deterministic)
}
func (dst *RelayBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayBlock.Merge(dst, src)
}
func (m *RelayBlock) XXX_Size() int {
	return xxx_messageInfo_RelayBlock.Size(m)
}
func (m *RelayBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RelayBlock proto.InternalMessageInfo

func (m *RelayBlock) GetProducer() uint64 {
	if m != nil {
		return m.Producer
	}
	return 0
}

func (m *RelayBlock) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RelayBlock) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *RelayBlock) GetConfirmed() int32 {
	if m != nil {
		return m.Confirmed
	}
	return 0
}

func (m *RelayBlock) GetTransactionMroot() string {
	if m != nil {
		return m.TransactionMroot
	}
	return ""
}

func (m *RelayBlock) GetActionMroot() string {
	if m != nil {
		return m.ActionMroot
	}
	return ""
}

func (m *RelayBlock) GetMroot() string {
	if m != nil {
		return m.Mroot
	}
	return ""
}

type RelayAction struct {
	Account              uint64                  `protobuf:"varint,1,opt,name=account,proto3" json:"account,omitempty"`
	ActionName           uint64                  `protobuf:"varint,2,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	Authorization        []*RelayPermissionLevel `protobuf:"bytes,3,rep,name=authorization,proto3" json:"authorization,omitempty"`
	Data                 []byte                  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RelayAction) Reset()         { *m = RelayAction{} }
func (m *RelayAction) String() string { return proto.CompactTextString(m) }
func (*RelayAction) ProtoMessage()    {}
func (*RelayAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_commit_01051abe939b6f32, []int{2}
}
func (m *RelayAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelayAction.Unmarshal(m, b)
}
func (m *RelayAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelayAction.Marshal(b, m, deterministic)
}
func (dst *RelayAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayAction.Merge(dst, src)
}
func (m *RelayAction) XXX_Size() int {
	return xxx_messageInfo_RelayAction.Size(m)
}
func (m *RelayAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayAction.DiscardUnknown(m)
}

var xxx_messageInfo_RelayAction proto.InternalMessageInfo

func (m *RelayAction) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *RelayAction) GetActionName() uint64 {
	if m != nil {
		return m.ActionName
	}
	return 0
}

func (m *RelayAction) GetAuthorization() []*RelayPermissionLevel {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *RelayAction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RelayPermissionLevel struct {
	Actor                uint64   `protobuf:"varint,1,opt,name=actor,proto3" json:"actor,omitempty"`
	Permission           uint64   `protobuf:"varint,2,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelayPermissionLevel) Reset()         { *m = RelayPermissionLevel{} }
func (m *RelayPermissionLevel) String() string { return proto.CompactTextString(m) }
func (*RelayPermissionLevel) ProtoMessage()    {}
func (*RelayPermissionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_commit_01051abe939b6f32, []int{3}
}
func (m *RelayPermissionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelayPermissionLevel.Unmarshal(m, b)
}
func (m *RelayPermissionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelayPermissionLevel.Marshal(b, m, deterministic)
}
func (dst *RelayPermissionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayPermissionLevel.Merge(dst, src)
}
func (m *RelayPermissionLevel) XXX_Size() int {
	return xxx_messageInfo_RelayPermissionLevel.Size(m)
}
func (m *RelayPermissionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayPermissionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_RelayPermissionLevel proto.InternalMessageInfo

func (m *RelayPermissionLevel) GetActor() uint64 {
	if m != nil {
		return m.Actor
	}
	return 0
}

func (m *RelayPermissionLevel) GetPermission() uint64 {
	if m != nil {
		return m.Permission
	}
	return 0
}

// The response message containing the greetings
type RelayCommitReply struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelayCommitReply) Reset()         { *m = RelayCommitReply{} }
func (m *RelayCommitReply) String() string { return proto.CompactTextString(m) }
func (*RelayCommitReply) ProtoMessage()    {}
func (*RelayCommitReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_relay_commit_01051abe939b6f32, []int{4}
}
func (m *RelayCommitReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelayCommitReply.Unmarshal(m, b)
}
func (m *RelayCommitReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelayCommitReply.Marshal(b, m, deterministic)
}
func (dst *RelayCommitReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayCommitReply.Merge(dst, src)
}
func (m *RelayCommitReply) XXX_Size() int {
	return xxx_messageInfo_RelayCommitReply.Size(m)
}
func (m *RelayCommitReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayCommitReply.DiscardUnknown(m)
}

var xxx_messageInfo_RelayCommitReply proto.InternalMessageInfo

func (m *RelayCommitReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *RelayCommitReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RelayCommitRequest)(nil), "force_relay_commit.RelayCommitRequest")
	proto.RegisterType((*RelayBlock)(nil), "force_relay_commit.RelayBlock")
	proto.RegisterType((*RelayAction)(nil), "force_relay_commit.RelayAction")
	proto.RegisterType((*RelayPermissionLevel)(nil), "force_relay_commit.RelayPermission_level")
	proto.RegisterType((*RelayCommitReply)(nil), "force_relay_commit.RelayCommitReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RelayCommitClient is the client API for RelayCommit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelayCommitClient interface {
	// Sends a greeting
	RpcSendaction(ctx context.Context, in *RelayCommitRequest, opts ...grpc.CallOption) (*RelayCommitReply, error)
}

type relayCommitClient struct {
	cc *grpc.ClientConn
}

func NewRelayCommitClient(cc *grpc.ClientConn) RelayCommitClient {
	return &relayCommitClient{cc}
}

func (c *relayCommitClient) RpcSendaction(ctx context.Context, in *RelayCommitRequest, opts ...grpc.CallOption) (*RelayCommitReply, error) {
	out := new(RelayCommitReply)
	err := c.cc.Invoke(ctx, "/force_relay_commit.relay_commit/rpc_sendaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelayCommitServer is the server API for RelayCommit service.
type RelayCommitServer interface {
	// Sends a greeting
	RpcSendaction(context.Context, *RelayCommitRequest) (*RelayCommitReply, error)
}

func RegisterRelayCommitServer(s *grpc.Server, srv RelayCommitServer) {
	s.RegisterService(&_RelayCommit_serviceDesc, srv)
}

func _RelayCommit_RpcSendaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelayCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayCommitServer).RpcSendaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/force_relay_commit.relay_commit/RpcSendaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayCommitServer).RpcSendaction(ctx, req.(*RelayCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RelayCommit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "force_relay_commit.relay_commit",
	HandlerType: (*RelayCommitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "rpc_sendaction",
			Handler:    _RelayCommit_RpcSendaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relay_commit.proto",
}

func init() { proto.RegisterFile("relay_commit.proto", fileDescriptor_relay_commit_01051abe939b6f32) }

var fileDescriptor_relay_commit_01051abe939b6f32 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x25, 0xfb, 0x55, 0x3a, 0xbb, 0x54, 0x65, 0x54, 0xa4, 0xa8, 0x42, 0x6d, 0x89, 0x10, 0x2a,
	0x02, 0xe5, 0x50, 0x90, 0x38, 0xb3, 0x5c, 0x38, 0x50, 0xa8, 0x7c, 0xe1, 0x18, 0x5c, 0xc7, 0x2d,
	0x11, 0x49, 0x1c, 0x6c, 0xa7, 0x52, 0xb9, 0xf2, 0x4f, 0xb8, 0xf3, 0x7f, 0xf8, 0x39, 0xd8, 0xe3,
	0xcd, 0x6e, 0x22, 0x58, 0xf5, 0xe6, 0x37, 0xf3, 0xe6, 0xe5, 0xcd, 0xb3, 0x03, 0xa8, 0x65, 0xc9,
	0x6f, 0x33, 0xa1, 0xaa, 0xaa, 0xb0, 0x69, 0xa3, 0x95, 0x55, 0x88, 0x57, 0x4a, 0x0b, 0x99, 0xf5,
	0x3b, 0xc9, 0xcf, 0x08, 0x90, 0xf9, 0xc2, 0x3b, 0xc2, 0x4c, 0x7e, 0x6f, 0xa5, 0xb1, 0xf8, 0x1a,
	0xa6, 0x97, 0xa5, 0x12, 0xdf, 0xe2, 0xe8, 0x24, 0x3a, 0x9d, 0x9f, 0x1d, 0xa5, 0xff, 0x8e, 0xa6,
	0x34, 0xb6, 0xf4, 0x2c, 0x16, 0xc8, 0xf8, 0x06, 0x66, 0x5c, 0xd8, 0x42, 0xd5, 0xf1, 0xe8, 0x64,
	0xec, 0xc6, 0x8e, 0xb7, 0x8e, 0xbd, 0x25, 0x1a, 0x5b, 0xd1, 0x93, 0x3f, 0x11, 0xc0, 0x46, 0x0e,
	0x0f, 0xe1, 0xbe, 0x73, 0x9c, 0xb7, 0x42, 0x6a, 0x32, 0x30, 0x61, 0x6b, 0x8c, 0x7b, 0x30, 0x2a,
	0x72, 0xa7, 0x1f, 0x9d, 0xee, 0x32, 0x77, 0x0a, 0x5c, 0x79, 0x53, 0xa8, 0xd6, 0xc4, 0x63, 0xaa,
	0xae, 0x31, 0x3e, 0x86, 0x5d, 0xa1, 0xea, 0xab, 0x42, 0x57, 0x32, 0x8f, 0x27, 0xae, 0x39, 0x65,
	0x9b, 0x02, 0xbe, 0x80, 0x87, 0x56, 0xf3, 0xda, 0x04, 0x0f, 0x59, 0xa5, 0x95, 0xb2, 0xf1, 0x94,
	0x24, 0xf6, 0x7b, 0x8d, 0x73, 0x5f, 0xc7, 0x27, 0xb0, 0x18, 0xf0, 0x66, 0xc4, 0x9b, 0xf7, 0x29,
	0x07, 0x30, 0x0d, 0xbd, 0x1d, 0xea, 0x05, 0x90, 0xfc, 0x8e, 0x60, 0xde, 0x5b, 0x19, 0x63, 0xd8,
	0xe1, 0x42, 0xa8, 0xb6, 0xb6, 0xab, 0xd5, 0x3a, 0x88, 0xc7, 0xb0, 0x92, 0xcb, 0x6a, 0x5e, 0x49,
	0x5a, 0x71, 0xc2, 0x20, 0x94, 0x3e, 0xba, 0x0a, 0x7e, 0x82, 0x07, 0xbc, 0xb5, 0x5f, 0x95, 0x2e,
	0x7e, 0x70, 0x4a, 0x79, 0x4c, 0x29, 0x3f, 0xdf, 0x9a, 0xf2, 0x85, 0xd4, 0x55, 0x61, 0x8c, 0x97,
	0x2c, 0xe5, 0x8d, 0x2c, 0xd9, 0x70, 0x1e, 0x11, 0x26, 0x39, 0xb7, 0x9c, 0xa2, 0x59, 0x30, 0x3a,
	0x27, 0xe7, 0xf0, 0xe8, 0xbf, 0xb3, 0x7e, 0x3d, 0xe7, 0x45, 0x75, 0x37, 0x12, 0x00, 0x1e, 0x01,
	0x34, 0x6b, 0x66, 0xe7, 0x79, 0x53, 0x49, 0x96, 0xb0, 0x3f, 0x78, 0x5e, 0x4d, 0x79, 0xeb, 0x95,
	0xb4, 0x3f, 0x90, 0x92, 0x0b, 0x8a, 0x80, 0x0f, 0xa6, 0x92, 0xc6, 0xf0, 0x6b, 0xb9, 0xba, 0xdd,
	0x0e, 0x9e, 0x35, 0xb0, 0xe8, 0xef, 0x86, 0x5f, 0x60, 0x4f, 0x37, 0x22, 0x33, 0xb2, 0xce, 0x43,
	0x3a, 0xf8, 0x6c, 0x6b, 0x04, 0x83, 0x67, 0x7d, 0xf8, 0xf4, 0x4e, 0x9e, 0x73, 0x92, 0xdc, 0x5b,
	0xbe, 0x84, 0x83, 0x42, 0xa5, 0xd7, 0xee, 0x2b, 0x69, 0x9f, 0xba, 0x1c, 0xf8, 0xb8, 0x88, 0x7e,
	0x8d, 0xc6, 0xef, 0x3f, 0x7c, 0xbe, 0x9c, 0xd1, 0xef, 0xf5, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x95, 0x87, 0xa6, 0x4d, 0x74, 0x03, 0x00, 0x00,
}