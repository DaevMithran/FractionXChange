// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vault/v1/nft_stake.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NftStake struct {
	Index            []byte  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	MarketplaceIndex []byte  `protobuf:"bytes,2,opt,name=marketplace_index,json=marketplaceIndex,proto3" json:"marketplace_index,omitempty" yaml:"marketplace_index"`
	CollectionIndex  []byte  `protobuf:"bytes,3,opt,name=collection_index,json=collectionIndex,proto3" json:"collection_index,omitempty" yaml:"collection_index"`
	Staked           []Stake `protobuf:"bytes,4,rep,name=staked,proto3" json:"staked" yaml:"staked"`
}

func (m *NftStake) Reset()         { *m = NftStake{} }
func (m *NftStake) String() string { return proto.CompactTextString(m) }
func (*NftStake) ProtoMessage()    {}
func (*NftStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_76b576b5811bd884, []int{0}
}
func (m *NftStake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftStake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftStake.Merge(m, src)
}
func (m *NftStake) XXX_Size() int {
	return m.Size()
}
func (m *NftStake) XXX_DiscardUnknown() {
	xxx_messageInfo_NftStake.DiscardUnknown(m)
}

var xxx_messageInfo_NftStake proto.InternalMessageInfo

func (m *NftStake) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *NftStake) GetMarketplaceIndex() []byte {
	if m != nil {
		return m.MarketplaceIndex
	}
	return nil
}

func (m *NftStake) GetCollectionIndex() []byte {
	if m != nil {
		return m.CollectionIndex
	}
	return nil
}

func (m *NftStake) GetStaked() []Stake {
	if m != nil {
		return m.Staked
	}
	return nil
}

type Stake struct {
	Amount     string                                        `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty" yaml:"amount"`
	Staked     bool                                          `protobuf:"varint,2,opt,name=staked,proto3" json:"staked,omitempty" yaml:"staked"`
	Shares     string                                        `protobuf:"bytes,3,opt,name=shares,proto3" json:"shares,omitempty" yaml:"shares"`
	Validator  string                                        `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty" yaml:"validator"`
	Chain      string                                        `protobuf:"bytes,5,opt,name=chain,proto3" json:"chain,omitempty" yaml:"chain"`
	Data       *types.Any                                    `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty" yaml:"data"`
	Creator    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,7,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
	Owner      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,8,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty" yaml:"owner"`
	StakedAt   time.Time                                     `protobuf:"bytes,9,opt,name=staked_at,json=stakedAt,proto3,stdtime" json:"staked_at" yaml:"staked_at"`
	UnstakedAt *time.Time                                    `protobuf:"bytes,10,opt,name=unstaked_at,json=unstakedAt,proto3,stdtime" json:"unstaked_at,omitempty" yaml:"unstaked_at"`
}

func (m *Stake) Reset()         { *m = Stake{} }
func (m *Stake) String() string { return proto.CompactTextString(m) }
func (*Stake) ProtoMessage()    {}
func (*Stake) Descriptor() ([]byte, []int) {
	return fileDescriptor_76b576b5811bd884, []int{1}
}
func (m *Stake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stake.Merge(m, src)
}
func (m *Stake) XXX_Size() int {
	return m.Size()
}
func (m *Stake) XXX_DiscardUnknown() {
	xxx_messageInfo_Stake.DiscardUnknown(m)
}

var xxx_messageInfo_Stake proto.InternalMessageInfo

func (m *Stake) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Stake) GetStaked() bool {
	if m != nil {
		return m.Staked
	}
	return false
}

func (m *Stake) GetShares() string {
	if m != nil {
		return m.Shares
	}
	return ""
}

func (m *Stake) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *Stake) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Stake) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Stake) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *Stake) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Stake) GetStakedAt() time.Time {
	if m != nil {
		return m.StakedAt
	}
	return time.Time{}
}

func (m *Stake) GetUnstakedAt() *time.Time {
	if m != nil {
		return m.UnstakedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*NftStake)(nil), "LimeChain.mantrachain.vault.v1.NftStake")
	proto.RegisterType((*Stake)(nil), "LimeChain.mantrachain.vault.v1.Stake")
}

func init() { proto.RegisterFile("vault/v1/nft_stake.proto", fileDescriptor_76b576b5811bd884) }

var fileDescriptor_76b576b5811bd884 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x9b, 0xdd, 0xa6, 0xdb, 0x4e, 0x57, 0xad, 0xb1, 0x62, 0xac, 0x4b, 0x52, 0x06, 0x94,
	0x22, 0x34, 0x61, 0x57, 0xf0, 0x20, 0x78, 0x48, 0x17, 0xc4, 0x05, 0xf1, 0x10, 0x57, 0x04, 0x45,
	0xca, 0x34, 0x99, 0xb6, 0xa1, 0x49, 0xa6, 0x24, 0xd3, 0xba, 0xfd, 0x17, 0x7b, 0xf7, 0xe8, 0x9f,
	0xd9, 0xe3, 0x1e, 0x3d, 0x45, 0x69, 0xff, 0x41, 0x8f, 0x9e, 0x24, 0xdf, 0x4c, 0xb7, 0xa5, 0x05,
	0x85, 0x3d, 0x75, 0xfa, 0x7d, 0xef, 0xfb, 0xcc, 0xcc, 0x3b, 0x33, 0x41, 0xfa, 0x94, 0x4c, 0x42,
	0x6e, 0x4f, 0x8f, 0xed, 0xb8, 0xcf, 0xbb, 0x29, 0x27, 0x23, 0x6a, 0x8d, 0x13, 0xc6, 0x99, 0x66,
	0xbc, 0x0b, 0x22, 0x7a, 0x3a, 0x24, 0x41, 0x6c, 0x45, 0x24, 0xe6, 0x09, 0xf1, 0x60, 0x0c, 0x7a,
	0x6b, 0x7a, 0xdc, 0x78, 0x3c, 0x60, 0x6c, 0x10, 0x52, 0x1b, 0xd4, 0xbd, 0x49, 0xdf, 0x26, 0xf1,
	0x4c, 0x58, 0x1b, 0xe6, 0x76, 0x8b, 0x07, 0x11, 0x4d, 0x39, 0x89, 0xc6, 0x52, 0x50, 0x1f, 0xb0,
	0x01, 0x83, 0xa1, 0x9d, 0x8f, 0x44, 0x15, 0x7f, 0xdf, 0x43, 0xe5, 0xf7, 0x7d, 0xfe, 0x21, 0x5f,
	0x84, 0xf6, 0x0c, 0xa9, 0x41, 0xec, 0xd3, 0x0b, 0x5d, 0x69, 0x2a, 0xad, 0xc3, 0x4e, 0x6d, 0x99,
	0x99, 0x87, 0x33, 0x12, 0x85, 0xaf, 0x30, 0x94, 0xb1, 0x2b, 0xda, 0xda, 0x19, 0xba, 0x1f, 0x91,
	0x64, 0x44, 0xf9, 0x38, 0x24, 0x1e, 0xed, 0x0a, 0xcf, 0x1e, 0x78, 0x8e, 0x96, 0x99, 0xa9, 0x0b,
	0xcf, 0x8e, 0x04, 0xbb, 0xb5, 0x8d, 0xda, 0x19, 0xa0, 0xde, 0xa0, 0x9a, 0xc7, 0xc2, 0x90, 0x7a,
	0x3c, 0x60, 0xb1, 0x24, 0xed, 0x03, 0xe9, 0xc9, 0x32, 0x33, 0x1f, 0x09, 0xd2, 0xb6, 0x02, 0xbb,
	0xf7, 0xd6, 0x25, 0xc1, 0x39, 0x47, 0x25, 0x08, 0xd2, 0xd7, 0x8b, 0xcd, 0xfd, 0x56, 0xf5, 0xe4,
	0xa9, 0xf5, 0xef, 0x28, 0x2d, 0xd8, 0x71, 0xe7, 0xe1, 0x55, 0x66, 0x16, 0x96, 0x99, 0x79, 0x47,
	0x4c, 0x24, 0x10, 0xd8, 0x95, 0x2c, 0xfc, 0x43, 0x45, 0xaa, 0x88, 0xa6, 0x8d, 0x4a, 0x24, 0x62,
	0x93, 0x98, 0x43, 0x36, 0x15, 0x30, 0x2a, 0x6b, 0xa3, 0xe8, 0x61, 0x57, 0x8a, 0x72, 0xb9, 0x5c,
	0x4e, 0x1e, 0x4b, 0x79, 0x5b, 0xbe, 0x35, 0x0f, 0xc8, 0x87, 0x24, 0xa1, 0x29, 0xec, 0x7d, 0x87,
	0x2e, 0x7a, 0xb9, 0x1c, 0x06, 0xda, 0x4b, 0x54, 0x99, 0x92, 0x30, 0xf0, 0x09, 0x67, 0x89, 0x5e,
	0x04, 0x87, 0x2e, 0x1d, 0x35, 0xe1, 0xb8, 0x69, 0x63, 0x77, 0x2d, 0xd5, 0x9e, 0x23, 0x15, 0x52,
	0xd0, 0x55, 0xf0, 0xd4, 0xa5, 0x47, 0x9e, 0x31, 0xb4, 0xb0, 0x2b, 0x24, 0xda, 0x6b, 0x54, 0xf4,
	0x09, 0x27, 0x7a, 0xa9, 0xa9, 0xb4, 0xaa, 0x27, 0x75, 0x4b, 0x5c, 0x2f, 0x6b, 0x75, 0xbd, 0x2c,
	0x27, 0x9e, 0x75, 0x1e, 0x48, 0x40, 0x55, 0x00, 0x72, 0x3d, 0x76, 0xc1, 0xa6, 0x7d, 0x45, 0x07,
	0x5e, 0x42, 0x61, 0x81, 0x07, 0x70, 0x9c, 0xa7, 0xcb, 0xcc, 0xbc, 0x2b, 0x27, 0x12, 0x0d, 0xfc,
	0x27, 0x33, 0xdb, 0x83, 0x80, 0x0f, 0x27, 0x3d, 0xcb, 0x63, 0x91, 0xed, 0xb1, 0x34, 0x62, 0xa9,
	0xfc, 0x69, 0xa7, 0xfe, 0xc8, 0xe6, 0xb3, 0x31, 0x4d, 0x2d, 0xc7, 0xf3, 0x1c, 0xdf, 0x4f, 0x68,
	0x9a, 0xba, 0x2b, 0xa6, 0xf6, 0x09, 0xa9, 0xec, 0x5b, 0x4c, 0x13, 0xbd, 0x0c, 0x70, 0x67, 0xbd,
	0x0b, 0x28, 0xdf, 0x02, 0x2d, 0x78, 0xda, 0x47, 0x54, 0x11, 0x67, 0xd2, 0x25, 0x5c, 0xaf, 0xc0,
	0xde, 0x1b, 0x3b, 0x7b, 0x3f, 0x5f, 0x3d, 0xad, 0xce, 0x91, 0xbc, 0x3f, 0xb5, 0xcd, 0x73, 0xed,
	0x12, 0x8e, 0x2f, 0x7f, 0x99, 0x8a, 0x5b, 0x16, 0xff, 0x1d, 0xae, 0x7d, 0x41, 0xd5, 0x49, 0xbc,
	0x06, 0xa3, 0xff, 0x82, 0x0d, 0x19, 0xad, 0x26, 0xc0, 0x1b, 0x66, 0x81, 0x46, 0xab, 0x8a, 0xc3,
	0x3b, 0x6f, 0xaf, 0xe6, 0x86, 0x72, 0x3d, 0x37, 0x94, 0xdf, 0x73, 0x43, 0xb9, 0x5c, 0x18, 0x85,
	0xeb, 0x85, 0x51, 0xf8, 0xb9, 0x30, 0x0a, 0x9f, 0xad, 0x8d, 0x0c, 0x6e, 0xde, 0x83, 0xbd, 0xf1,
	0x1e, 0xec, 0x0b, 0x5b, 0x7c, 0x8c, 0x20, 0x8f, 0x5e, 0x09, 0x56, 0xf2, 0xe2, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x69, 0x74, 0xa2, 0xa2, 0x04, 0x00, 0x00,
}

func (m *NftStake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftStake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftStake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Staked) > 0 {
		for iNdEx := len(m.Staked) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Staked[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftStake(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CollectionIndex) > 0 {
		i -= len(m.CollectionIndex)
		copy(dAtA[i:], m.CollectionIndex)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.CollectionIndex)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MarketplaceIndex) > 0 {
		i -= len(m.MarketplaceIndex)
		copy(dAtA[i:], m.MarketplaceIndex)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.MarketplaceIndex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Stake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnstakedAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.UnstakedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.UnstakedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintNftStake(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x52
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StakedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StakedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintNftStake(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftStake(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Shares) > 0 {
		i -= len(m.Shares)
		copy(dAtA[i:], m.Shares)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Shares)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Staked {
		i--
		if m.Staked {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintNftStake(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftStake(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftStake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftStake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = len(m.MarketplaceIndex)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = len(m.CollectionIndex)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	if len(m.Staked) > 0 {
		for _, e := range m.Staked {
			l = e.Size()
			n += 1 + l + sovNftStake(uint64(l))
		}
	}
	return n
}

func (m *Stake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	if m.Staked {
		n += 2
	}
	l = len(m.Shares)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNftStake(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StakedAt)
	n += 1 + l + sovNftStake(uint64(l))
	if m.UnstakedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UnstakedAt)
		n += 1 + l + sovNftStake(uint64(l))
	}
	return n
}

func sovNftStake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftStake(x uint64) (n int) {
	return sovNftStake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftStake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftStake
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
			return fmt.Errorf("proto: NftStake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftStake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = append(m.Index[:0], dAtA[iNdEx:postIndex]...)
			if m.Index == nil {
				m.Index = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketplaceIndex", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketplaceIndex = append(m.MarketplaceIndex[:0], dAtA[iNdEx:postIndex]...)
			if m.MarketplaceIndex == nil {
				m.MarketplaceIndex = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionIndex", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionIndex = append(m.CollectionIndex[:0], dAtA[iNdEx:postIndex]...)
			if m.CollectionIndex == nil {
				m.CollectionIndex = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Staked = append(m.Staked, Stake{})
			if err := m.Staked[len(m.Staked)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftStake
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
func (m *Stake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftStake
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
			return fmt.Errorf("proto: Stake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Staked = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StakedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftStake
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
				return ErrInvalidLengthNftStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UnstakedAt == nil {
				m.UnstakedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.UnstakedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftStake
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
func skipNftStake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftStake
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
					return 0, ErrIntOverflowNftStake
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
					return 0, ErrIntOverflowNftStake
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
				return 0, ErrInvalidLengthNftStake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftStake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftStake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftStake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftStake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftStake = fmt.Errorf("proto: unexpected end of group")
)
