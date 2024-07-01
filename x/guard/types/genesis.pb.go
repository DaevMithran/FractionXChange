// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/guard/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the guard module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params                 Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	AccountPrivilegesList  []*AccountPrivileges  `protobuf:"bytes,2,rep,name=account_privileges_list,json=accountPrivilegesList,proto3" json:"account_privileges_list,omitempty"`
	GuardTransferCoins     []byte                `protobuf:"bytes,3,opt,name=guard_transfer_coins,json=guardTransferCoins,proto3" json:"guard_transfer_coins,omitempty"`
	RequiredPrivilegesList []*RequiredPrivileges `protobuf:"bytes,4,rep,name=required_privileges_list,json=requiredPrivilegesList,proto3" json:"required_privileges_list,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea5a8529cb82408, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAccountPrivilegesList() []*AccountPrivileges {
	if m != nil {
		return m.AccountPrivilegesList
	}
	return nil
}

func (m *GenesisState) GetGuardTransferCoins() []byte {
	if m != nil {
		return m.GuardTransferCoins
	}
	return nil
}

func (m *GenesisState) GetRequiredPrivilegesList() []*RequiredPrivileges {
	if m != nil {
		return m.RequiredPrivilegesList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mantrachain.guard.v1.GenesisState")
}

func init() {
	proto.RegisterFile("mantrachain/guard/v1/genesis.proto", fileDescriptor_0ea5a8529cb82408)
}

var fileDescriptor_0ea5a8529cb82408 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4a, 0xf3, 0x40,
	0x14, 0xc7, 0x33, 0xed, 0x47, 0xe1, 0x4b, 0xbb, 0x31, 0x54, 0x0d, 0x45, 0x62, 0xed, 0xc6, 0x20,
	0x98, 0xb1, 0xed, 0x01, 0xa4, 0x15, 0x74, 0x63, 0xa5, 0xc4, 0xae, 0xdc, 0x84, 0xe9, 0x74, 0x4c,
	0x07, 0x9a, 0x99, 0x38, 0x33, 0x29, 0x7a, 0x0b, 0x8f, 0xe1, 0xd2, 0x53, 0x48, 0x97, 0x5d, 0xba,
	0x12, 0x69, 0x17, 0x5e, 0x43, 0x3a, 0x09, 0x12, 0x6c, 0x70, 0x33, 0x3c, 0xde, 0xfb, 0xfd, 0xdf,
	0xfc, 0xe0, 0x99, 0xad, 0x08, 0x31, 0x25, 0x10, 0x9e, 0x22, 0xca, 0x60, 0x98, 0x20, 0x31, 0x81,
	0xf3, 0x36, 0x0c, 0x09, 0x23, 0x92, 0x4a, 0x2f, 0x16, 0x5c, 0x71, 0xab, 0x9e, 0x63, 0x3c, 0xcd,
	0x78, 0xf3, 0x76, 0x63, 0x07, 0x45, 0x94, 0x71, 0xa8, 0xdf, 0x14, 0x6c, 0xd4, 0x43, 0x1e, 0x72,
	0x5d, 0xc2, 0x4d, 0x95, 0x75, 0x8f, 0x0a, 0xbf, 0x88, 0x91, 0x40, 0x91, 0xfc, 0x13, 0xc1, 0x3c,
	0x8a, 0x38, 0x4b, 0x91, 0xd6, 0x5b, 0xc9, 0xac, 0x5d, 0xa5, 0x5a, 0xb7, 0x0a, 0x29, 0x62, 0x9d,
	0x9b, 0x95, 0x74, 0x87, 0x0d, 0x9a, 0xc0, 0xad, 0x76, 0x0e, 0xbc, 0x22, 0x4d, 0x6f, 0xa8, 0x99,
	0xfe, 0xff, 0xc5, 0xc7, 0xa1, 0xf1, 0xf2, 0xf5, 0x7a, 0x02, 0xfc, 0x2c, 0x66, 0x05, 0xe6, 0x3e,
	0xc2, 0x98, 0x27, 0x4c, 0x05, 0xb1, 0xa0, 0x73, 0x3a, 0x23, 0x21, 0x91, 0xc1, 0x8c, 0x4a, 0x65,
	0x97, 0x9a, 0x65, 0xb7, 0xda, 0x39, 0x2e, 0xde, 0xd8, 0x4b, 0x43, 0xc3, 0x9f, 0x8c, 0xbf, 0x8b,
	0x7e, 0xb7, 0xae, 0xa9, 0x54, 0xd6, 0x99, 0x59, 0xd7, 0xa1, 0x40, 0x09, 0xc4, 0xe4, 0x3d, 0x11,
	0x01, 0xe6, 0x94, 0x49, 0xbb, 0xdc, 0x04, 0x6e, 0xcd, 0xb7, 0xf4, 0x6c, 0x94, 0x8d, 0x2e, 0x36,
	0x13, 0x6b, 0x6c, 0xda, 0x82, 0x3c, 0x24, 0x54, 0x90, 0xc9, 0x96, 0xd3, 0x3f, 0xed, 0xe4, 0x16,
	0x3b, 0xf9, 0x59, 0x2a, 0x27, 0xb5, 0x27, 0xb6, 0x7a, 0x1b, 0xab, 0xfe, 0x60, 0xb1, 0x72, 0xc0,
	0x72, 0xe5, 0x80, 0xcf, 0x95, 0x03, 0x9e, 0xd7, 0x8e, 0xb1, 0x5c, 0x3b, 0xc6, 0xfb, 0xda, 0x31,
	0xee, 0xba, 0x21, 0x55, 0xd3, 0x64, 0xec, 0x61, 0x1e, 0xc1, 0x41, 0xef, 0x66, 0xe4, 0xf7, 0x4e,
	0x2f, 0x29, 0x43, 0x0c, 0x13, 0x98, 0xbf, 0xcf, 0x63, 0x76, 0x21, 0xf5, 0x14, 0x13, 0x39, 0xae,
	0xe8, 0xf3, 0x74, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x0f, 0x7d, 0xd4, 0x49, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequiredPrivilegesList) > 0 {
		for iNdEx := len(m.RequiredPrivilegesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequiredPrivilegesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.GuardTransferCoins) > 0 {
		i -= len(m.GuardTransferCoins)
		copy(dAtA[i:], m.GuardTransferCoins)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.GuardTransferCoins)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountPrivilegesList) > 0 {
		for iNdEx := len(m.AccountPrivilegesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountPrivilegesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.AccountPrivilegesList) > 0 {
		for _, e := range m.AccountPrivilegesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.GuardTransferCoins)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.RequiredPrivilegesList) > 0 {
		for _, e := range m.RequiredPrivilegesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountPrivilegesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountPrivilegesList = append(m.AccountPrivilegesList, &AccountPrivileges{})
			if err := m.AccountPrivilegesList[len(m.AccountPrivilegesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardTransferCoins", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardTransferCoins = append(m.GuardTransferCoins[:0], dAtA[iNdEx:postIndex]...)
			if m.GuardTransferCoins == nil {
				m.GuardTransferCoins = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredPrivilegesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequiredPrivilegesList = append(m.RequiredPrivilegesList, &RequiredPrivileges{})
			if err := m.RequiredPrivilegesList[len(m.RequiredPrivilegesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
