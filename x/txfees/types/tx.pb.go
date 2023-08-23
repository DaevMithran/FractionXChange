// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/txfees/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateFeeToken struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	PairId  uint64 `protobuf:"varint,3,opt,name=pairId,proto3" json:"pairId,omitempty"`
}

func (m *MsgCreateFeeToken) Reset()         { *m = MsgCreateFeeToken{} }
func (m *MsgCreateFeeToken) String() string { return proto.CompactTextString(m) }
func (*MsgCreateFeeToken) ProtoMessage()    {}
func (*MsgCreateFeeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_f79719d7abe38d35, []int{0}
}
func (m *MsgCreateFeeToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateFeeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateFeeToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateFeeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateFeeToken.Merge(m, src)
}
func (m *MsgCreateFeeToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateFeeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateFeeToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateFeeToken proto.InternalMessageInfo

func (m *MsgCreateFeeToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateFeeToken) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgCreateFeeToken) GetPairId() uint64 {
	if m != nil {
		return m.PairId
	}
	return 0
}

type MsgCreateFeeTokenResponse struct {
}

func (m *MsgCreateFeeTokenResponse) Reset()         { *m = MsgCreateFeeTokenResponse{} }
func (m *MsgCreateFeeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateFeeTokenResponse) ProtoMessage()    {}
func (*MsgCreateFeeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f79719d7abe38d35, []int{1}
}
func (m *MsgCreateFeeTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateFeeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateFeeTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateFeeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateFeeTokenResponse.Merge(m, src)
}
func (m *MsgCreateFeeTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateFeeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateFeeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateFeeTokenResponse proto.InternalMessageInfo

type MsgUpdateFeeToken struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	PairId  uint64 `protobuf:"varint,3,opt,name=pairId,proto3" json:"pairId,omitempty"`
}

func (m *MsgUpdateFeeToken) Reset()         { *m = MsgUpdateFeeToken{} }
func (m *MsgUpdateFeeToken) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateFeeToken) ProtoMessage()    {}
func (*MsgUpdateFeeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_f79719d7abe38d35, []int{2}
}
func (m *MsgUpdateFeeToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateFeeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateFeeToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateFeeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateFeeToken.Merge(m, src)
}
func (m *MsgUpdateFeeToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateFeeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateFeeToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateFeeToken proto.InternalMessageInfo

func (m *MsgUpdateFeeToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateFeeToken) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgUpdateFeeToken) GetPairId() uint64 {
	if m != nil {
		return m.PairId
	}
	return 0
}

type MsgUpdateFeeTokenResponse struct {
}

func (m *MsgUpdateFeeTokenResponse) Reset()         { *m = MsgUpdateFeeTokenResponse{} }
func (m *MsgUpdateFeeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateFeeTokenResponse) ProtoMessage()    {}
func (*MsgUpdateFeeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f79719d7abe38d35, []int{3}
}
func (m *MsgUpdateFeeTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateFeeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateFeeTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateFeeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateFeeTokenResponse.Merge(m, src)
}
func (m *MsgUpdateFeeTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateFeeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateFeeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateFeeTokenResponse proto.InternalMessageInfo

type MsgDeleteFeeToken struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *MsgDeleteFeeToken) Reset()         { *m = MsgDeleteFeeToken{} }
func (m *MsgDeleteFeeToken) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteFeeToken) ProtoMessage()    {}
func (*MsgDeleteFeeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_f79719d7abe38d35, []int{4}
}
func (m *MsgDeleteFeeToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteFeeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteFeeToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteFeeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteFeeToken.Merge(m, src)
}
func (m *MsgDeleteFeeToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteFeeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteFeeToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteFeeToken proto.InternalMessageInfo

func (m *MsgDeleteFeeToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteFeeToken) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type MsgDeleteFeeTokenResponse struct {
}

func (m *MsgDeleteFeeTokenResponse) Reset()         { *m = MsgDeleteFeeTokenResponse{} }
func (m *MsgDeleteFeeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteFeeTokenResponse) ProtoMessage()    {}
func (*MsgDeleteFeeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f79719d7abe38d35, []int{5}
}
func (m *MsgDeleteFeeTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteFeeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteFeeTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteFeeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteFeeTokenResponse.Merge(m, src)
}
func (m *MsgDeleteFeeTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteFeeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteFeeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteFeeTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateFeeToken)(nil), "mantrachain.txfees.MsgCreateFeeToken")
	proto.RegisterType((*MsgCreateFeeTokenResponse)(nil), "mantrachain.txfees.MsgCreateFeeTokenResponse")
	proto.RegisterType((*MsgUpdateFeeToken)(nil), "mantrachain.txfees.MsgUpdateFeeToken")
	proto.RegisterType((*MsgUpdateFeeTokenResponse)(nil), "mantrachain.txfees.MsgUpdateFeeTokenResponse")
	proto.RegisterType((*MsgDeleteFeeToken)(nil), "mantrachain.txfees.MsgDeleteFeeToken")
	proto.RegisterType((*MsgDeleteFeeTokenResponse)(nil), "mantrachain.txfees.MsgDeleteFeeTokenResponse")
}

func init() { proto.RegisterFile("mantrachain/txfees/tx.proto", fileDescriptor_f79719d7abe38d35) }

var fileDescriptor_f79719d7abe38d35 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x4d, 0xcc, 0x2b,
	0x29, 0x4a, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0xa9, 0x48, 0x4b, 0x4d, 0x2d, 0xd6, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x92, 0xd4, 0x83, 0x48, 0x4a, 0x29,
	0x61, 0xd1, 0x90, 0x96, 0x9a, 0x1a, 0x5f, 0x92, 0x9f, 0x9d, 0x9a, 0x07, 0xd1, 0xa7, 0x14, 0xcd,
	0x25, 0xe8, 0x5b, 0x9c, 0xee, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0xea, 0x96, 0x9a, 0x1a, 0x02, 0x92,
	0x12, 0x92, 0xe0, 0x62, 0x4f, 0x06, 0x89, 0xe4, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0x12, 0x4c, 0x60, 0x71, 0x08,
	0x47, 0x48, 0x8c, 0x8b, 0xad, 0x20, 0x31, 0xb3, 0xc8, 0x33, 0x45, 0x82, 0x59, 0x81, 0x51, 0x83,
	0x25, 0x08, 0xca, 0x53, 0x92, 0xe6, 0x92, 0xc4, 0x30, 0x3c, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x15, 0x6a, 0x73, 0x68, 0x41, 0x0a, 0xed, 0x6c, 0x46, 0x35, 0x1c, 0x6e, 0xb3, 0x33, 0xd8,
	0x66, 0x97, 0xd4, 0x9c, 0x54, 0xf2, 0x6d, 0x86, 0xda, 0x80, 0x6a, 0x08, 0xcc, 0x06, 0xa3, 0xbd,
	0x4c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x69, 0x5c, 0x7c, 0x68, 0x41, 0xab, 0xaa, 0x87, 0x19,
	0x51, 0x7a, 0x18, 0x81, 0x24, 0xa5, 0x4b, 0x94, 0x32, 0x98, 0x7d, 0x20, 0x7b, 0xd0, 0x02, 0x12,
	0x97, 0x3d, 0xa8, 0xca, 0x70, 0xda, 0x83, 0x3d, 0xe4, 0x40, 0xf6, 0xa0, 0x05, 0x1b, 0x2e, 0x7b,
	0x50, 0x95, 0xe1, 0xb4, 0x07, 0x7b, 0xf8, 0x39, 0xf9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x49, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe,
	0xaf, 0xa3, 0x5f, 0x48, 0x90, 0xa3, 0xae, 0x5b, 0x66, 0x5e, 0x62, 0x5e, 0x72, 0xaa, 0x3e, 0x72,
	0x6a, 0xaf, 0x80, 0x67, 0x90, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x62, 0x37, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x94, 0x80, 0xfe, 0xd3, 0x43, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateFeeToken(ctx context.Context, in *MsgCreateFeeToken, opts ...grpc.CallOption) (*MsgCreateFeeTokenResponse, error)
	UpdateFeeToken(ctx context.Context, in *MsgUpdateFeeToken, opts ...grpc.CallOption) (*MsgUpdateFeeTokenResponse, error)
	DeleteFeeToken(ctx context.Context, in *MsgDeleteFeeToken, opts ...grpc.CallOption) (*MsgDeleteFeeTokenResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateFeeToken(ctx context.Context, in *MsgCreateFeeToken, opts ...grpc.CallOption) (*MsgCreateFeeTokenResponse, error) {
	out := new(MsgCreateFeeTokenResponse)
	err := c.cc.Invoke(ctx, "/mantrachain.txfees.Msg/CreateFeeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateFeeToken(ctx context.Context, in *MsgUpdateFeeToken, opts ...grpc.CallOption) (*MsgUpdateFeeTokenResponse, error) {
	out := new(MsgUpdateFeeTokenResponse)
	err := c.cc.Invoke(ctx, "/mantrachain.txfees.Msg/UpdateFeeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteFeeToken(ctx context.Context, in *MsgDeleteFeeToken, opts ...grpc.CallOption) (*MsgDeleteFeeTokenResponse, error) {
	out := new(MsgDeleteFeeTokenResponse)
	err := c.cc.Invoke(ctx, "/mantrachain.txfees.Msg/DeleteFeeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateFeeToken(context.Context, *MsgCreateFeeToken) (*MsgCreateFeeTokenResponse, error)
	UpdateFeeToken(context.Context, *MsgUpdateFeeToken) (*MsgUpdateFeeTokenResponse, error)
	DeleteFeeToken(context.Context, *MsgDeleteFeeToken) (*MsgDeleteFeeTokenResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateFeeToken(ctx context.Context, req *MsgCreateFeeToken) (*MsgCreateFeeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeeToken not implemented")
}
func (*UnimplementedMsgServer) UpdateFeeToken(ctx context.Context, req *MsgUpdateFeeToken) (*MsgUpdateFeeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeeToken not implemented")
}
func (*UnimplementedMsgServer) DeleteFeeToken(ctx context.Context, req *MsgDeleteFeeToken) (*MsgDeleteFeeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeeToken not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateFeeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateFeeToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateFeeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mantrachain.txfees.Msg/CreateFeeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateFeeToken(ctx, req.(*MsgCreateFeeToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateFeeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateFeeToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateFeeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mantrachain.txfees.Msg/UpdateFeeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateFeeToken(ctx, req.(*MsgUpdateFeeToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteFeeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteFeeToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteFeeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mantrachain.txfees.Msg/DeleteFeeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteFeeToken(ctx, req.(*MsgDeleteFeeToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mantrachain.txfees.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeeToken",
			Handler:    _Msg_CreateFeeToken_Handler,
		},
		{
			MethodName: "UpdateFeeToken",
			Handler:    _Msg_UpdateFeeToken_Handler,
		},
		{
			MethodName: "DeleteFeeToken",
			Handler:    _Msg_DeleteFeeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mantrachain/txfees/tx.proto",
}

func (m *MsgCreateFeeToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateFeeToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateFeeToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateFeeTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateFeeTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateFeeTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateFeeToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateFeeToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateFeeToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateFeeTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateFeeTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateFeeTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteFeeToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteFeeToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteFeeToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteFeeTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteFeeTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteFeeTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateFeeToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovTx(uint64(m.PairId))
	}
	return n
}

func (m *MsgCreateFeeTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateFeeToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovTx(uint64(m.PairId))
	}
	return n
}

func (m *MsgUpdateFeeTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteFeeToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteFeeTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateFeeToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateFeeToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateFeeToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateFeeTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateFeeTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateFeeTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateFeeToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateFeeToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateFeeToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateFeeTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateFeeTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateFeeTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteFeeToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteFeeToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteFeeToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteFeeTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteFeeTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteFeeTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
