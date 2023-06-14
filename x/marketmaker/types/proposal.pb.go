// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/marketmaker/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MarketMakerProposal struct {
	// title specifies the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description specifies the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// set the market makers to eligible, refund deposit
	Inclusions []MarketMakerHandle `protobuf:"bytes,3,rep,name=inclusions,proto3" json:"inclusions" yaml:"inclusions"`
	// delete existing eligible market makers
	Exclusions []MarketMakerHandle `protobuf:"bytes,4,rep,name=exclusions,proto3" json:"exclusions" yaml:"exclusions"`
	// delete the not eligible market makers, refund deposit
	Rejections []MarketMakerHandle `protobuf:"bytes,5,rep,name=rejections,proto3" json:"rejections" yaml:"rejections"`
	// distribute claimable incentive to eligible market makers
	Distributions []IncentiveDistribution `protobuf:"bytes,6,rep,name=distributions,proto3" json:"distributions" yaml:"distributions"`
}

func (m *MarketMakerProposal) Reset()      { *m = MarketMakerProposal{} }
func (*MarketMakerProposal) ProtoMessage() {}
func (*MarketMakerProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffcfbb1b2495b7c9, []int{0}
}
func (m *MarketMakerProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMakerProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMakerProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMakerProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMakerProposal.Merge(m, src)
}
func (m *MarketMakerProposal) XXX_Size() int {
	return m.Size()
}
func (m *MarketMakerProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMakerProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMakerProposal proto.InternalMessageInfo

type MarketMakerHandle struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	PairId  uint64 `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
}

func (m *MarketMakerHandle) Reset()         { *m = MarketMakerHandle{} }
func (m *MarketMakerHandle) String() string { return proto.CompactTextString(m) }
func (*MarketMakerHandle) ProtoMessage()    {}
func (*MarketMakerHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffcfbb1b2495b7c9, []int{1}
}
func (m *MarketMakerHandle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMakerHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMakerHandle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMakerHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMakerHandle.Merge(m, src)
}
func (m *MarketMakerHandle) XXX_Size() int {
	return m.Size()
}
func (m *MarketMakerHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMakerHandle.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMakerHandle proto.InternalMessageInfo

type IncentiveDistribution struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	PairId  uint64                                   `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	Amount  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *IncentiveDistribution) Reset()         { *m = IncentiveDistribution{} }
func (m *IncentiveDistribution) String() string { return proto.CompactTextString(m) }
func (*IncentiveDistribution) ProtoMessage()    {}
func (*IncentiveDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffcfbb1b2495b7c9, []int{2}
}
func (m *IncentiveDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveDistribution.Merge(m, src)
}
func (m *IncentiveDistribution) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveDistribution proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MarketMakerProposal)(nil), "mantrachain.marketmaker.v1beta1.MarketMakerProposal")
	proto.RegisterType((*MarketMakerHandle)(nil), "mantrachain.marketmaker.v1beta1.MarketMakerHandle")
	proto.RegisterType((*IncentiveDistribution)(nil), "mantrachain.marketmaker.v1beta1.IncentiveDistribution")
}

func init() {
	proto.RegisterFile("mantrachain/marketmaker/v1beta1/proposal.proto", fileDescriptor_ffcfbb1b2495b7c9)
}

var fileDescriptor_ffcfbb1b2495b7c9 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0x93, 0xdf, 0xba, 0xee, 0x87, 0xa7, 0x21, 0x2d, 0x14, 0xa9, 0x9d, 0x50, 0x52, 0xe5,
	0x54, 0x09, 0x96, 0xa8, 0x43, 0x42, 0x62, 0xb7, 0x16, 0x84, 0xd8, 0xa1, 0x08, 0x05, 0x4e, 0x5c,
	0x26, 0x27, 0x31, 0x9d, 0x69, 0x6c, 0x47, 0xb6, 0x3b, 0x75, 0xbc, 0x02, 0x8e, 0x1c, 0x39, 0xf6,
	0xcc, 0x2b, 0xd9, 0x71, 0x47, 0x4e, 0x05, 0xb5, 0x17, 0xce, 0x7d, 0x05, 0x28, 0xb6, 0x69, 0x5d,
	0x81, 0xd4, 0xcb, 0x38, 0x25, 0xcf, 0x1f, 0x7f, 0x3f, 0x7e, 0xac, 0xe7, 0x0b, 0x22, 0x02, 0xa9,
	0xe4, 0x30, 0xbb, 0x80, 0x98, 0xc6, 0x04, 0xf2, 0x11, 0x92, 0x04, 0x8e, 0x10, 0x8f, 0x2f, 0xbb,
	0x29, 0x92, 0xb0, 0x1b, 0x97, 0x9c, 0x95, 0x4c, 0xc0, 0x22, 0x2a, 0x39, 0x93, 0xcc, 0x0b, 0xac,
	0xfe, 0xc8, 0xea, 0x8f, 0x4c, 0xff, 0x91, 0x9f, 0x31, 0x41, 0x98, 0x88, 0x53, 0x28, 0xd0, 0x4a,
	0x24, 0x63, 0x98, 0x6a, 0x81, 0xa3, 0x96, 0xae, 0x9f, 0xab, 0x28, 0xd6, 0x81, 0x29, 0x35, 0x86,
	0x6c, 0xc8, 0x74, 0xbe, 0xfa, 0x33, 0xd9, 0x60, 0xc8, 0xd8, 0xb0, 0x40, 0xb1, 0x8a, 0xd2, 0xf1,
	0xfb, 0x58, 0x62, 0x82, 0x84, 0x84, 0xa4, 0x34, 0x0d, 0xdd, 0x6d, 0x23, 0xd8, 0xd7, 0x54, 0x47,
	0xc2, 0x69, 0x0d, 0xdc, 0x1b, 0xa8, 0xec, 0xa0, 0xca, 0xbe, 0x36, 0x33, 0x7a, 0x0d, 0xb0, 0x2b,
	0xb1, 0x2c, 0x50, 0xd3, 0x6d, 0xbb, 0x9d, 0x3b, 0x89, 0x0e, 0xbc, 0x36, 0xd8, 0xcf, 0x91, 0xc8,
	0x38, 0x2e, 0x25, 0x66, 0xb4, 0xf9, 0x9f, 0xaa, 0xd9, 0x29, 0x8f, 0x00, 0x80, 0x69, 0x56, 0x8c,
	0x05, 0x66, 0x54, 0x34, 0x77, 0xda, 0x3b, 0x9d, 0xfd, 0x93, 0x93, 0x68, 0xcb, 0x53, 0x45, 0xd6,
	0x0d, 0x5e, 0x42, 0x9a, 0x17, 0xa8, 0xdf, 0xba, 0x9e, 0x05, 0xce, 0x72, 0x16, 0x1c, 0x5e, 0x41,
	0x52, 0x9c, 0x86, 0x6b, 0xcd, 0x30, 0xb1, 0x00, 0x15, 0x0e, 0x4d, 0x56, 0xb8, 0xda, 0x6d, 0xe1,
	0xd6, 0x9a, 0x61, 0x62, 0x01, 0x2a, 0x1c, 0x47, 0x1f, 0x50, 0x26, 0x15, 0x6e, 0xf7, 0xb6, 0x70,
	0x6b, 0xcd, 0x30, 0xb1, 0x00, 0xde, 0x47, 0x70, 0x90, 0x63, 0x21, 0x39, 0x4e, 0xc7, 0x9a, 0x58,
	0x57, 0xc4, 0x27, 0x5b, 0x89, 0x67, 0x34, 0x43, 0x54, 0xe2, 0x4b, 0xf4, 0xdc, 0x3a, 0xde, 0x7f,
	0x60, 0xa8, 0x0d, 0x4d, 0xdd, 0x90, 0x0e, 0x93, 0x4d, 0xd4, 0xe9, 0xff, 0x9f, 0xa6, 0x81, 0xf3,
	0x65, 0x1a, 0x38, 0xe1, 0x04, 0x1c, 0xfe, 0x31, 0x81, 0xf7, 0x08, 0xec, 0xc1, 0x3c, 0xe7, 0x48,
	0x08, 0xbd, 0x21, 0x7d, 0x6f, 0x39, 0x0b, 0xee, 0x6a, 0x61, 0x53, 0x08, 0x93, 0xdf, 0x2d, 0xde,
	0x43, 0xb0, 0x57, 0x42, 0xcc, 0xcf, 0x71, 0xae, 0x76, 0xa6, 0x66, 0x77, 0x9b, 0x42, 0x98, 0xd4,
	0xab, 0xbf, 0xb3, 0x5c, 0x93, 0x7f, 0x56, 0xe4, 0xa5, 0x0b, 0xee, 0xff, 0x75, 0x94, 0x7f, 0x88,
	0xf7, 0x24, 0xa8, 0x43, 0xc2, 0xc6, 0x54, 0x9a, 0xed, 0x6d, 0x45, 0xc6, 0x9a, 0x95, 0x8f, 0x57,
	0x2f, 0xfc, 0x8c, 0x61, 0xda, 0xef, 0x99, 0x07, 0x3d, 0x30, 0x60, 0x75, 0x2c, 0xfc, 0xfa, 0x3d,
	0xe8, 0x0c, 0xb1, 0xbc, 0x18, 0xa7, 0x51, 0xc6, 0x88, 0x31, 0xb6, 0xf9, 0x1c, 0x8b, 0x7c, 0x14,
	0xcb, 0xab, 0x12, 0x09, 0xa5, 0x20, 0x12, 0xc3, 0x5a, 0x0f, 0xdd, 0x7f, 0x73, 0x3d, 0xf7, 0xdd,
	0x9b, 0xb9, 0xef, 0xfe, 0x98, 0xfb, 0xee, 0xe7, 0x85, 0xef, 0xdc, 0x2c, 0x7c, 0xe7, 0xdb, 0xc2,
	0x77, 0xde, 0x3d, 0xb5, 0x54, 0x07, 0xbd, 0x57, 0x6f, 0x93, 0xde, 0xf1, 0x0b, 0x4c, 0x21, 0xcd,
	0x50, 0x6c, 0x1b, 0x7f, 0xb2, 0x61, 0x7d, 0x05, 0x4b, 0xeb, 0xca, 0xed, 0x8f, 0x7f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x49, 0x1d, 0x93, 0x90, 0xe5, 0x04, 0x00, 0x00,
}

func (m *MarketMakerProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMakerProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMakerProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Distributions) > 0 {
		for iNdEx := len(m.Distributions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Distributions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Rejections) > 0 {
		for iNdEx := len(m.Rejections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rejections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Exclusions) > 0 {
		for iNdEx := len(m.Exclusions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Exclusions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Inclusions) > 0 {
		for iNdEx := len(m.Inclusions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inclusions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketMakerHandle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMakerHandle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMakerHandle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PairId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketMakerProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Inclusions) > 0 {
		for _, e := range m.Inclusions {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Exclusions) > 0 {
		for _, e := range m.Exclusions {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Rejections) > 0 {
		for _, e := range m.Rejections {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Distributions) > 0 {
		for _, e := range m.Distributions {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *MarketMakerHandle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovProposal(uint64(m.PairId))
	}
	return n
}

func (m *IncentiveDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovProposal(uint64(m.PairId))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketMakerProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MarketMakerProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMakerProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inclusions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inclusions = append(m.Inclusions, MarketMakerHandle{})
			if err := m.Inclusions[len(m.Inclusions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exclusions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Exclusions = append(m.Exclusions, MarketMakerHandle{})
			if err := m.Exclusions[len(m.Exclusions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rejections = append(m.Rejections, MarketMakerHandle{})
			if err := m.Rejections[len(m.Rejections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distributions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distributions = append(m.Distributions, IncentiveDistribution{})
			if err := m.Distributions[len(m.Distributions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *MarketMakerHandle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MarketMakerHandle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMakerHandle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *IncentiveDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: IncentiveDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
