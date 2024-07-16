// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/rewards/v1beta1/common.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

type ProviderPair struct {
	PairId                uint64                                      `protobuf:"varint,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	LastClaimedSnapshotId uint64                                      `protobuf:"varint,2,opt,name=lastClaimedSnapshotId,proto3" json:"lastClaimedSnapshotId,omitempty"`
	OwedRewards           github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,3,rep,name=owedRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"owedRewards" yaml:"owed_rewards"`
	Balances              github_com_cosmos_cosmos_sdk_types.Coins    `protobuf:"bytes,4,rep,name=balances,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balances" yaml:"owed_rewards"`
	PoolIdToIdx           map[uint64]uint64                           `protobuf:"bytes,5,rep,name=poolIdToIdx,proto3" json:"poolIdToIdx,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LastDepositTime       *time.Time                                  `protobuf:"bytes,6,opt,name=lastDepositTime,proto3,stdtime" json:"lastDepositTime,omitempty" yaml:"last_deposit_time"`
}

func (m *ProviderPair) Reset()         { *m = ProviderPair{} }
func (m *ProviderPair) String() string { return proto.CompactTextString(m) }
func (*ProviderPair) ProtoMessage()    {}
func (*ProviderPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_274771e038f60916, []int{0}
}
func (m *ProviderPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderPair.Merge(m, src)
}
func (m *ProviderPair) XXX_Size() int {
	return m.Size()
}
func (m *ProviderPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderPair.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderPair proto.InternalMessageInfo

type SnapshotPool struct {
	PoolId          uint64                                      `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty"`
	RewardsPerToken github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=rewardsPerToken,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewardsPerToken" yaml:"rewards_per_token"`
}

func (m *SnapshotPool) Reset()         { *m = SnapshotPool{} }
func (m *SnapshotPool) String() string { return proto.CompactTextString(m) }
func (*SnapshotPool) ProtoMessage()    {}
func (*SnapshotPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_274771e038f60916, []int{1}
}
func (m *SnapshotPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotPool.Merge(m, src)
}
func (m *SnapshotPool) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotPool) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotPool.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProviderPair)(nil), "mantrachain.rewards.legacy.ProviderPair")
	proto.RegisterMapType((map[uint64]uint64)(nil), "mantrachain.rewards.legacy.ProviderPair.PoolIdToIdxEntry")
	proto.RegisterType((*SnapshotPool)(nil), "mantrachain.rewards.legacy.SnapshotPool")
}

func init() {
	proto.RegisterFile("mantrachain/rewards/legacy/common.proto", fileDescriptor_274771e038f60916)
}

var fileDescriptor_274771e038f60916 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x4d, 0x1a, 0x64, 0x52, 0x68, 0x59, 0xab, 0x6c, 0xa3, 0xec, 0x86, 0xe0, 0x21,
	0x20, 0x9d, 0xa5, 0x55, 0x41, 0x72, 0x10, 0x9a, 0x56, 0x25, 0x82, 0x1a, 0xb6, 0x39, 0x89, 0xb0,
	0xcc, 0xee, 0x8e, 0xc9, 0x98, 0xdd, 0x9d, 0x65, 0x67, 0x92, 0x36, 0x17, 0xcf, 0xe2, 0xa9, 0x37,
	0xaf, 0x1e, 0xc5, 0x0f, 0x22, 0x39, 0xf6, 0xe8, 0x29, 0x95, 0xe4, 0x1b, 0xf4, 0xe6, 0x4d, 0x76,
	0x67, 0x12, 0x97, 0x50, 0x4a, 0x3d, 0xed, 0x3e, 0xe6, 0xbd, 0xff, 0xfb, 0xbf, 0x79, 0x3f, 0x06,
	0x36, 0x42, 0x1c, 0x89, 0x04, 0x7b, 0x7d, 0x4c, 0x23, 0x2b, 0x21, 0x27, 0x38, 0xf1, 0xb9, 0x35,
	0xda, 0x73, 0x89, 0xc0, 0x7b, 0x96, 0xc7, 0xc2, 0x90, 0x45, 0x28, 0x4e, 0x98, 0x60, 0xda, 0xbd,
	0x5c, 0x26, 0x52, 0x99, 0x48, 0x65, 0x56, 0xb7, 0x7b, 0xac, 0xc7, 0xb2, 0x3c, 0x2b, 0xfd, 0x93,
	0x25, 0x55, 0xb3, 0xc7, 0x58, 0x2f, 0x20, 0x56, 0x16, 0xb9, 0xc3, 0x0f, 0x96, 0xa0, 0x21, 0xe1,
	0x02, 0x87, 0xb1, 0x4a, 0xd8, 0xf1, 0x18, 0x0f, 0x19, 0x77, 0x64, 0xa5, 0x0c, 0xd4, 0x91, 0x21,
	0x23, 0xcb, 0xc5, 0x9c, 0xe4, 0x0c, 0x51, 0x65, 0xa7, 0xfe, 0xa7, 0x04, 0x37, 0x3a, 0x09, 0x1b,
	0x51, 0x9f, 0x24, 0x1d, 0x4c, 0x13, 0xed, 0x2e, 0x2c, 0xc7, 0x98, 0x26, 0x6d, 0x5f, 0x07, 0x35,
	0xd0, 0x28, 0xd9, 0x2a, 0xd2, 0x1e, 0xc3, 0x3b, 0x01, 0xe6, 0xe2, 0x30, 0xc0, 0x34, 0x24, 0xfe,
	0x71, 0x84, 0x63, 0xde, 0x67, 0xa2, 0xed, 0xeb, 0x6b, 0x59, 0xda, 0xd5, 0x87, 0xda, 0x17, 0x00,
	0x2b, 0xec, 0x84, 0xf8, 0xb6, 0x1c, 0x54, 0x2f, 0xd6, 0x8a, 0x8d, 0xca, 0xfe, 0x7d, 0xa4, 0x3c,
	0xa6, 0xae, 0x16, 0xc3, 0xa3, 0x23, 0xe2, 0x1d, 0x32, 0x1a, 0xb5, 0x5e, 0x4d, 0xa6, 0x66, 0xe1,
	0x72, 0x6a, 0xde, 0x1e, 0xe3, 0x30, 0x68, 0xd6, 0xd3, 0x72, 0x47, 0x5d, 0x54, 0xfd, 0xc7, 0x85,
	0xf9, 0xb0, 0x47, 0x45, 0x7f, 0xe8, 0x22, 0x8f, 0x85, 0x6a, 0x54, 0xf5, 0xd9, 0xe5, 0xfe, 0xc0,
	0x12, 0xe3, 0x98, 0xf0, 0x85, 0x14, 0xb7, 0xf3, 0xcd, 0xb5, 0x4f, 0xf0, 0x96, 0x8b, 0x03, 0x1c,
	0x79, 0x84, 0xeb, 0xa5, 0xcc, 0xc8, 0xce, 0x95, 0x46, 0x32, 0x17, 0x2f, 0xaf, 0x77, 0xd1, 0xb8,
	0x81, 0x0b, 0x69, 0x61, 0xd9, 0x53, 0x7b, 0x0f, 0x2b, 0x31, 0x63, 0x41, 0xdb, 0xef, 0xb2, 0xb6,
	0x7f, 0xaa, 0xaf, 0x67, 0x16, 0x9a, 0xe8, 0x1a, 0x20, 0x50, 0x7e, 0x35, 0xa8, 0xf3, 0xaf, 0xf8,
	0x79, 0x24, 0x92, 0xb1, 0x9d, 0x97, 0xd3, 0x3e, 0xc2, 0xcd, 0x74, 0x07, 0x47, 0x24, 0x66, 0x9c,
	0x8a, 0x2e, 0x0d, 0x89, 0x5e, 0xae, 0x81, 0x46, 0x65, 0xbf, 0x8a, 0x24, 0x3f, 0x68, 0xc1, 0x0f,
	0xea, 0x2e, 0xf8, 0x69, 0x3d, 0x98, 0x4c, 0x4d, 0x70, 0x39, 0x35, 0x75, 0x39, 0x65, 0x2a, 0xe0,
	0xf8, 0x52, 0xc1, 0x49, 0x29, 0xab, 0x9f, 0x5d, 0x98, 0xc0, 0x5e, 0x15, 0xae, 0x3e, 0x83, 0x5b,
	0xab, 0x66, 0xb4, 0x2d, 0x58, 0x1c, 0x90, 0xb1, 0xa2, 0x26, 0xfd, 0xd5, 0xb6, 0xe1, 0xfa, 0x08,
	0x07, 0x43, 0xa2, 0x10, 0x91, 0x41, 0x73, 0xed, 0x29, 0x68, 0x96, 0x3e, 0x7f, 0x33, 0x0b, 0xf5,
	0x9f, 0x00, 0x6e, 0x2c, 0x58, 0x49, 0xe5, 0x32, 0xf6, 0x32, 0xd9, 0x25, 0x7b, 0x59, 0xa4, 0x7d,
	0x05, 0x70, 0x53, 0xdd, 0x4c, 0x87, 0x24, 0x5d, 0x36, 0x20, 0x91, 0xbe, 0x76, 0x03, 0x92, 0xde,
	0xaa, 0x1d, 0xaa, 0xe9, 0x94, 0x84, 0x13, 0x93, 0xc4, 0x11, 0xa9, 0xc8, 0x7f, 0xe3, 0xb4, 0xea,
	0x42, 0x0e, 0xd2, 0x3a, 0xfe, 0x3e, 0x33, 0xc0, 0x64, 0x66, 0x80, 0xf3, 0x99, 0x01, 0x7e, 0xcf,
	0x0c, 0x70, 0x36, 0x37, 0x0a, 0xe7, 0x73, 0xa3, 0xf0, 0x6b, 0x6e, 0x14, 0xde, 0x3d, 0xc9, 0x75,
	0x78, 0x7d, 0xf0, 0xa6, 0x6b, 0x1f, 0xec, 0xbe, 0xa0, 0x51, 0x8a, 0x84, 0x95, 0x7f, 0x35, 0x4e,
	0x97, 0xef, 0x46, 0xd6, 0xd4, 0x2d, 0x67, 0xeb, 0x7a, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x25,
	0x71, 0xad, 0x54, 0x5b, 0x04, 0x00, 0x00,
}

func (this *ProviderPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProviderPair)
	if !ok {
		that2, ok := that.(ProviderPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PairId != that1.PairId {
		return false
	}
	if this.LastClaimedSnapshotId != that1.LastClaimedSnapshotId {
		return false
	}
	if len(this.OwedRewards) != len(that1.OwedRewards) {
		return false
	}
	for i := range this.OwedRewards {
		if !this.OwedRewards[i].Equal(&that1.OwedRewards[i]) {
			return false
		}
	}
	if len(this.Balances) != len(that1.Balances) {
		return false
	}
	for i := range this.Balances {
		if !this.Balances[i].Equal(&that1.Balances[i]) {
			return false
		}
	}
	if len(this.PoolIdToIdx) != len(that1.PoolIdToIdx) {
		return false
	}
	for i := range this.PoolIdToIdx {
		if this.PoolIdToIdx[i] != that1.PoolIdToIdx[i] {
			return false
		}
	}
	if that1.LastDepositTime == nil {
		if this.LastDepositTime != nil {
			return false
		}
	} else if !this.LastDepositTime.Equal(*that1.LastDepositTime) {
		return false
	}
	return true
}
func (this *SnapshotPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SnapshotPool)
	if !ok {
		that2, ok := that.(SnapshotPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolId != that1.PoolId {
		return false
	}
	if len(this.RewardsPerToken) != len(that1.RewardsPerToken) {
		return false
	}
	for i := range this.RewardsPerToken {
		if !this.RewardsPerToken[i].Equal(&that1.RewardsPerToken[i]) {
			return false
		}
	}
	return true
}
func (m *ProviderPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastDepositTime != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.LastDepositTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.LastDepositTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintCommon(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PoolIdToIdx) > 0 {
		for k := range m.PoolIdToIdx {
			v := m.PoolIdToIdx[k]
			baseI := i
			i = encodeVarintCommon(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintCommon(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintCommon(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.OwedRewards) > 0 {
		for iNdEx := len(m.OwedRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OwedRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.LastClaimedSnapshotId != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.LastClaimedSnapshotId))
		i--
		dAtA[i] = 0x10
	}
	if m.PairId != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardsPerToken) > 0 {
		for iNdEx := len(m.RewardsPerToken) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardsPerToken[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PoolId != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProviderPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairId != 0 {
		n += 1 + sovCommon(uint64(m.PairId))
	}
	if m.LastClaimedSnapshotId != 0 {
		n += 1 + sovCommon(uint64(m.LastClaimedSnapshotId))
	}
	if len(m.OwedRewards) > 0 {
		for _, e := range m.OwedRewards {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if len(m.PoolIdToIdx) > 0 {
		for k, v := range m.PoolIdToIdx {
			_ = k
			_ = v
			mapEntrySize := 1 + sovCommon(uint64(k)) + 1 + sovCommon(uint64(v))
			n += mapEntrySize + 1 + sovCommon(uint64(mapEntrySize))
		}
	}
	if m.LastDepositTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.LastDepositTime)
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *SnapshotPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovCommon(uint64(m.PoolId))
	}
	if len(m.RewardsPerToken) > 0 {
		for _, e := range m.RewardsPerToken {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: ProviderPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastClaimedSnapshotId", wireType)
			}
			m.LastClaimedSnapshotId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastClaimedSnapshotId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwedRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwedRewards = append(m.OwedRewards, types.DecCoin{})
			if err := m.OwedRewards[len(m.OwedRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, types.Coin{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIdToIdx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolIdToIdx == nil {
				m.PoolIdToIdx = make(map[uint64]uint64)
			}
			var mapkey uint64
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCommon(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCommon
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PoolIdToIdx[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastDepositTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastDepositTime == nil {
				m.LastDepositTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.LastDepositTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *SnapshotPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: SnapshotPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsPerToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardsPerToken = append(m.RewardsPerToken, types.DecCoin{})
			if err := m.RewardsPerToken[len(m.RewardsPerToken)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)