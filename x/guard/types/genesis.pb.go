// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/guard/v1/genesis.proto

package types

import (
	fmt "fmt"
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
	Params                Params               `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	AccountPrivilegesList []*AccountPrivileges `protobuf:"bytes,2,rep,name=account_privileges_list,json=accountPrivilegesList,proto3" json:"account_privileges_list,omitempty"`
	GuardTransferCoins    []byte               `protobuf:"bytes,3,opt,name=guard_transfer_coins,json=guardTransferCoins,proto3" json:"guard_transfer_coins,omitempty"`
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

type AccountPrivileges struct {
	Account    []byte `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Privileges []byte `protobuf:"bytes,2,opt,name=privileges,proto3" json:"privileges,omitempty"`
}

func (m *AccountPrivileges) Reset()         { *m = AccountPrivileges{} }
func (m *AccountPrivileges) String() string { return proto.CompactTextString(m) }
func (*AccountPrivileges) ProtoMessage()    {}
func (*AccountPrivileges) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea5a8529cb82408, []int{1}
}
func (m *AccountPrivileges) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountPrivileges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountPrivileges.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountPrivileges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountPrivileges.Merge(m, src)
}
func (m *AccountPrivileges) XXX_Size() int {
	return m.Size()
}
func (m *AccountPrivileges) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountPrivileges.DiscardUnknown(m)
}

var xxx_messageInfo_AccountPrivileges proto.InternalMessageInfo

func (m *AccountPrivileges) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *AccountPrivileges) GetPrivileges() []byte {
	if m != nil {
		return m.Privileges
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mantrachain.guard.v1.GenesisState")
	proto.RegisterType((*AccountPrivileges)(nil), "mantrachain.guard.v1.AccountPrivileges")
}

func init() {
	proto.RegisterFile("mantrachain/guard/v1/genesis.proto", fileDescriptor_0ea5a8529cb82408)
}

var fileDescriptor_0ea5a8529cb82408 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0xaf, 0x60, 0x30, 0x29, 0x2c, 0x36, 0x18, 0x2f, 0xc4, 0x54, 0x64, 0x91, 0xa9, 0x27,
	0xb8, 0xb9, 0x09, 0x83, 0x0e, 0x98, 0x90, 0xd3, 0xc9, 0xe5, 0x52, 0xce, 0x5a, 0x9a, 0x70, 0xed,
	0xa5, 0x2d, 0x44, 0xbf, 0x85, 0x1f, 0x8b, 0x91, 0xd1, 0x45, 0x63, 0xb8, 0x2f, 0x62, 0xe8, 0x9d,
	0x7a, 0xd1, 0xdb, 0xda, 0xf7, 0xff, 0xfd, 0xff, 0xaf, 0x7d, 0x0f, 0xf6, 0x12, 0x2a, 0xad, 0xa6,
	0xf1, 0x9c, 0x0a, 0x19, 0xf0, 0x25, 0xd5, 0x8f, 0xc1, 0x6a, 0x10, 0x70, 0x26, 0x99, 0x11, 0x86,
	0xa4, 0x5a, 0x59, 0x85, 0xda, 0x25, 0x86, 0x38, 0x86, 0xac, 0x06, 0x9d, 0x36, 0x57, 0x5c, 0x39,
	0x20, 0xd8, 0x9d, 0x72, 0xb6, 0x73, 0x5a, 0x99, 0x97, 0x52, 0x4d, 0x93, 0x22, 0xae, 0xf7, 0x0e,
	0x60, 0xeb, 0x3a, 0x6f, 0x70, 0x67, 0xa9, 0x65, 0xe8, 0x12, 0x36, 0x72, 0xc0, 0x07, 0x5d, 0xd0,
	0x6f, 0x0e, 0x8f, 0x49, 0x55, 0x43, 0x32, 0x75, 0xcc, 0x68, 0x6f, 0xfd, 0x71, 0xe2, 0x85, 0x85,
	0x03, 0x45, 0xf0, 0x88, 0xc6, 0xb1, 0x5a, 0x4a, 0x1b, 0xa5, 0x5a, 0xac, 0xc4, 0x82, 0x71, 0x66,
	0xa2, 0x85, 0x30, 0xd6, 0xaf, 0x75, 0xeb, 0xfd, 0xe6, 0xf0, 0xac, 0x3a, 0xec, 0x2a, 0x37, 0x4d,
	0x7f, 0x3c, 0xe1, 0x21, 0xfd, 0x5b, 0x9a, 0x08, 0x63, 0xd1, 0x39, 0x6c, 0x3b, 0x53, 0x64, 0x35,
	0x95, 0xe6, 0x89, 0xe9, 0x28, 0x56, 0x42, 0x1a, 0xbf, 0xde, 0x05, 0xfd, 0x56, 0x88, 0x9c, 0x76,
	0x5f, 0x48, 0xe3, 0x9d, 0xd2, 0xbb, 0x85, 0x07, 0xff, 0xd2, 0x91, 0x0f, 0xf7, 0x8b, 0x7c, 0xf7,
	0xc9, 0x56, 0xf8, 0x7d, 0x45, 0x18, 0xc2, 0xdf, 0x97, 0xfb, 0x35, 0x27, 0x96, 0x2a, 0xa3, 0x9b,
	0xf5, 0x16, 0x83, 0xcd, 0x16, 0x83, 0xcf, 0x2d, 0x06, 0xaf, 0x19, 0xf6, 0x36, 0x19, 0xf6, 0xde,
	0x32, 0xec, 0x3d, 0x10, 0x2e, 0xec, 0x7c, 0x39, 0x23, 0xb1, 0x4a, 0x82, 0x89, 0x48, 0xd8, 0xd8,
	0x0d, 0xbd, 0xbc, 0x80, 0xe7, 0x62, 0x05, 0xf6, 0x25, 0x65, 0x66, 0xd6, 0x70, 0xf3, 0xbf, 0xf8,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x45, 0x2f, 0x1f, 0xfc, 0xf4, 0x01, 0x00, 0x00,
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

func (m *AccountPrivileges) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountPrivileges) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountPrivileges) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Privileges) > 0 {
		i -= len(m.Privileges)
		copy(dAtA[i:], m.Privileges)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Privileges)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
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
	return n
}

func (m *AccountPrivileges) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Privileges)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
func (m *AccountPrivileges) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AccountPrivileges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountPrivileges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = append(m.Account[:0], dAtA[iNdEx:postIndex]...)
			if m.Account == nil {
				m.Account = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
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
			m.Privileges = append(m.Privileges[:0], dAtA[iNdEx:postIndex]...)
			if m.Privileges == nil {
				m.Privileges = []byte{}
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
