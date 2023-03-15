// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/guard/v1/common.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type MsgAccounts struct {
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *MsgAccounts) Reset()         { *m = MsgAccounts{} }
func (m *MsgAccounts) String() string { return proto.CompactTextString(m) }
func (*MsgAccounts) ProtoMessage()    {}
func (*MsgAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2ac908ba79d3e1, []int{0}
}
func (m *MsgAccounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAccounts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAccounts.Merge(m, src)
}
func (m *MsgAccounts) XXX_Size() int {
	return m.Size()
}
func (m *MsgAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAccounts proto.InternalMessageInfo

func (m *MsgAccounts) GetAccounts() []string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type MsgIndexes struct {
	Indexes [][]byte `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
}

func (m *MsgIndexes) Reset()         { *m = MsgIndexes{} }
func (m *MsgIndexes) String() string { return proto.CompactTextString(m) }
func (*MsgIndexes) ProtoMessage()    {}
func (*MsgIndexes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2ac908ba79d3e1, []int{1}
}
func (m *MsgIndexes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIndexes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIndexes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIndexes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIndexes.Merge(m, src)
}
func (m *MsgIndexes) XXX_Size() int {
	return m.Size()
}
func (m *MsgIndexes) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIndexes.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIndexes proto.InternalMessageInfo

func (m *MsgIndexes) GetIndexes() [][]byte {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgAccounts)(nil), "mantrachain.guard.v1.MsgAccounts")
	proto.RegisterType((*MsgIndexes)(nil), "mantrachain.guard.v1.MsgIndexes")
}

func init() { proto.RegisterFile("mantrachain/guard/v1/common.proto", fileDescriptor_ce2ac908ba79d3e1) }

var fileDescriptor_ce2ac908ba79d3e1 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x4d, 0xcc, 0x2b,
	0x29, 0x4a, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x2f, 0x4d, 0x2c, 0x4a, 0xd1, 0x2f, 0x33,
	0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41,
	0x52, 0xa2, 0x07, 0x56, 0xa2, 0x57, 0x66, 0x28, 0x25, 0x99, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x1c,
	0x0f, 0x56, 0xa3, 0x0f, 0xe1, 0x40, 0x34, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x43, 0xc4, 0x41,
	0x2c, 0x88, 0xa8, 0x92, 0x33, 0x17, 0xb7, 0x6f, 0x71, 0xba, 0x63, 0x72, 0x72, 0x7e, 0x69, 0x5e,
	0x49, 0xb1, 0x90, 0x09, 0x17, 0x47, 0x22, 0x94, 0x2d, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0xe9, 0x24,
	0x71, 0x69, 0x8b, 0xae, 0x08, 0xd4, 0x20, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92,
	0xa2, 0xcc, 0xbc, 0xf4, 0x20, 0xb8, 0x4a, 0x25, 0x35, 0x2e, 0x2e, 0xdf, 0xe2, 0x74, 0xcf, 0xbc,
	0x94, 0xd4, 0x8a, 0xd4, 0x62, 0x21, 0x09, 0x2e, 0xf6, 0x4c, 0x08, 0x13, 0x6c, 0x04, 0x4f, 0x10,
	0x8c, 0xeb, 0xe4, 0xb3, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xfb, 0x64,
	0xe6, 0xa6, 0x3a, 0x83, 0x7d, 0x8f, 0x1c, 0x12, 0x15, 0xd0, 0xb0, 0x28, 0xa9, 0x2c, 0x48, 0x2d,
	0x4e, 0x62, 0x03, 0xfb, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xde, 0x44, 0xdf, 0x2d,
	0x01, 0x00, 0x00,
}

func (this *MsgAccounts) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgAccounts)
	if !ok {
		that2, ok := that.(MsgAccounts)
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
	if len(this.Accounts) != len(that1.Accounts) {
		return false
	}
	for i := range this.Accounts {
		if this.Accounts[i] != that1.Accounts[i] {
			return false
		}
	}
	return true
}
func (this *MsgIndexes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgIndexes)
	if !ok {
		that2, ok := that.(MsgIndexes)
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
	if len(this.Indexes) != len(that1.Indexes) {
		return false
	}
	for i := range this.Indexes {
		if !bytes.Equal(this.Indexes[i], that1.Indexes[i]) {
			return false
		}
	}
	return true
}
func (m *MsgAccounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAccounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAccounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgIndexes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIndexes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIndexes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		for iNdEx := len(m.Indexes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Indexes[iNdEx])
			copy(dAtA[i:], m.Indexes[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Indexes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *MsgAccounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func (m *MsgIndexes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		for _, b := range m.Indexes {
			l = len(b)
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
func (m *MsgAccounts) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAccounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgIndexes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIndexes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIndexes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexes = append(m.Indexes, make([]byte, postIndex-iNdEx))
			copy(m.Indexes[len(m.Indexes)-1], dAtA[iNdEx:postIndex])
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
