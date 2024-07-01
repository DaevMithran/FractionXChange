// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/lpfarm/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params defines the parameters for the module.
type Params struct {
	PrivatePlanCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=private_plan_creation_fee,json=privatePlanCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"private_plan_creation_fee"`
	FeeCollector           string                                   `protobuf:"bytes,2,opt,name=fee_collector,json=feeCollector,proto3" json:"fee_collector,omitempty"`
	MaxNumPrivatePlans     uint32                                   `protobuf:"varint,3,opt,name=max_num_private_plans,json=maxNumPrivatePlans,proto3" json:"max_num_private_plans,omitempty"`
	MaxBlockDuration       time.Duration                            `protobuf:"bytes,4,opt,name=max_block_duration,json=maxBlockDuration,proto3,stdduration" json:"max_block_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf8191677813348d, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPrivatePlanCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PrivatePlanCreationFee
	}
	return nil
}

func (m *Params) GetFeeCollector() string {
	if m != nil {
		return m.FeeCollector
	}
	return ""
}

func (m *Params) GetMaxNumPrivatePlans() uint32 {
	if m != nil {
		return m.MaxNumPrivatePlans
	}
	return 0
}

func (m *Params) GetMaxBlockDuration() time.Duration {
	if m != nil {
		return m.MaxBlockDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "mantrachain.lpfarm.v1beta1.Params")
}

func init() {
	proto.RegisterFile("mantrachain/lpfarm/v1beta1/params.proto", fileDescriptor_bf8191677813348d)
}

var fileDescriptor_bf8191677813348d = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0xcd, 0xde, 0xa1, 0x13, 0x18, 0x4e, 0x02, 0x0b, 0x50, 0x2e, 0x48, 0x4e, 0x04, 0x48, 0x44,
	0x48, 0xb7, 0x4b, 0x80, 0x8a, 0xee, 0x12, 0x74, 0x1d, 0x51, 0xb0, 0xa8, 0x68, 0xac, 0xf1, 0xde,
	0xd8, 0x67, 0x9d, 0x77, 0xc7, 0xf2, 0xae, 0x4f, 0xe1, 0x07, 0x28, 0xa8, 0x28, 0x29, 0xa9, 0xa9,
	0xf8, 0x8c, 0x2b, 0xaf, 0xa4, 0x22, 0x28, 0x29, 0xe0, 0x33, 0x90, 0xd7, 0x36, 0xa4, 0xb8, 0x66,
	0x77, 0x76, 0xe6, 0xcd, 0xbe, 0x37, 0x6f, 0xbc, 0x27, 0x0a, 0xb4, 0x2d, 0x41, 0x9e, 0x42, 0xa6,
	0x45, 0x5e, 0x24, 0x50, 0x2a, 0x71, 0x3e, 0x89, 0xd1, 0xc2, 0x44, 0x14, 0x50, 0x82, 0x32, 0xbc,
	0x28, 0xc9, 0x92, 0x3f, 0xd8, 0x02, 0xf2, 0x06, 0xc8, 0x5b, 0xe0, 0xe0, 0x0e, 0xa8, 0x4c, 0x93,
	0x70, 0x67, 0x03, 0x1f, 0xdc, 0x4d, 0x29, 0x25, 0x17, 0x8a, 0x3a, 0x6a, 0xb3, 0x81, 0x24, 0xa3,
	0xc8, 0x88, 0x18, 0x0c, 0xfe, 0xa3, 0x91, 0x94, 0xe9, 0xae, 0x9e, 0x12, 0xa5, 0x39, 0x0a, 0xf7,
	0x8a, 0xab, 0x44, 0x9c, 0x54, 0x25, 0xd8, 0x8c, 0xda, 0xfa, 0xc3, 0xf5, 0x8e, 0xb7, 0xb7, 0x70,
	0xaa, 0xfc, 0x8f, 0xcc, 0x3b, 0x28, 0xca, 0xec, 0x1c, 0x2c, 0x46, 0x45, 0x0e, 0x3a, 0x92, 0x25,
	0x3a, 0x68, 0x94, 0x20, 0xf6, 0xd9, 0x68, 0x77, 0x7c, 0xf3, 0xf9, 0x01, 0x6f, 0xf8, 0x78, 0xcd,
	0xd7, 0xa9, 0xe5, 0x33, 0xca, 0xf4, 0xf4, 0xd9, 0xc5, 0xcf, 0x61, 0xef, 0xdb, 0x6a, 0x38, 0x4e,
	0x33, 0x7b, 0x5a, 0xc5, 0x5c, 0x92, 0x12, 0xad, 0xb8, 0xe6, 0x3a, 0x34, 0x27, 0x67, 0xc2, 0x7e,
	0x28, 0xd0, 0xb8, 0x06, 0x13, 0xde, 0x6f, 0xd9, 0x16, 0x39, 0xe8, 0x59, 0xcb, 0x75, 0x8c, 0xe8,
	0x3f, 0xf2, 0xf6, 0x13, 0xc4, 0x48, 0x52, 0x9e, 0xa3, 0xb4, 0x54, 0xf6, 0x77, 0x46, 0x6c, 0x7c,
	0x23, 0xbc, 0x95, 0x20, 0xce, 0xba, 0x9c, 0x3f, 0xf1, 0xee, 0x29, 0x58, 0x46, 0xba, 0x52, 0xd1,
	0xb6, 0x68, 0xd3, 0xdf, 0x1d, 0xb1, 0xf1, 0x7e, 0xe8, 0x2b, 0x58, 0xce, 0x2b, 0xb5, 0xf8, 0xcf,
	0x60, 0xfc, 0xb7, 0x5e, 0x9d, 0x8d, 0xe2, 0x9c, 0xe4, 0x59, 0xd4, 0xf9, 0xd0, 0xbf, 0x36, 0x62,
	0x6e, 0xb0, 0xc6, 0x28, 0xde, 0x19, 0xc5, 0x5f, 0xb7, 0x80, 0xe9, 0xf5, 0x7a, 0xb0, 0x2f, 0xab,
	0x21, 0x0b, 0x6f, 0x2b, 0x58, 0x4e, 0xeb, 0xee, 0xae, 0xf6, 0xea, 0xf1, 0x9f, 0xaf, 0x43, 0xf6,
	0xe9, 0xf7, 0xf7, 0xa7, 0x0f, 0xb6, 0xb7, 0xbe, 0xec, 0xf6, 0xde, 0x38, 0x3b, 0x9d, 0x5f, 0xac,
	0x03, 0x76, 0xb9, 0x0e, 0xd8, 0xaf, 0x75, 0xc0, 0x3e, 0x6f, 0x82, 0xde, 0xe5, 0x26, 0xe8, 0xfd,
	0xd8, 0x04, 0xbd, 0xf7, 0x2f, 0xb7, 0xcc, 0x7a, 0x73, 0x34, 0x7f, 0x17, 0x1e, 0x1d, 0x1e, 0x67,
	0x1a, 0xb4, 0x44, 0x71, 0xe5, 0x87, 0xce, 0xbe, 0x78, 0xcf, 0x89, 0x7c, 0xf1, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x25, 0x0b, 0x26, 0x09, 0x6b, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if len(this.PrivatePlanCreationFee) != len(that1.PrivatePlanCreationFee) {
		return false
	}
	for i := range this.PrivatePlanCreationFee {
		if !this.PrivatePlanCreationFee[i].Equal(&that1.PrivatePlanCreationFee[i]) {
			return false
		}
	}
	if this.FeeCollector != that1.FeeCollector {
		return false
	}
	if this.MaxNumPrivatePlans != that1.MaxNumPrivatePlans {
		return false
	}
	if this.MaxBlockDuration != that1.MaxBlockDuration {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxBlockDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxBlockDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.MaxNumPrivatePlans != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxNumPrivatePlans))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FeeCollector) > 0 {
		i -= len(m.FeeCollector)
		copy(dAtA[i:], m.FeeCollector)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollector)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PrivatePlanCreationFee) > 0 {
		for iNdEx := len(m.PrivatePlanCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PrivatePlanCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PrivatePlanCreationFee) > 0 {
		for _, e := range m.PrivatePlanCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.FeeCollector)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxNumPrivatePlans != 0 {
		n += 1 + sovParams(uint64(m.MaxNumPrivatePlans))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxBlockDuration)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivatePlanCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivatePlanCreationFee = append(m.PrivatePlanCreationFee, types.Coin{})
			if err := m.PrivatePlanCreationFee[len(m.PrivatePlanCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCollector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumPrivatePlans", wireType)
			}
			m.MaxNumPrivatePlans = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumPrivatePlans |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxBlockDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
