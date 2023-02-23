// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/token/v1/soul_bonded_nfts_collection.proto

package types

import (
	fmt "fmt"
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

type SoulBondedNftsCollection struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *SoulBondedNftsCollection) Reset()         { *m = SoulBondedNftsCollection{} }
func (m *SoulBondedNftsCollection) String() string { return proto.CompactTextString(m) }
func (*SoulBondedNftsCollection) ProtoMessage()    {}
func (*SoulBondedNftsCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a002e287bfc1d5, []int{0}
}
func (m *SoulBondedNftsCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SoulBondedNftsCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SoulBondedNftsCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SoulBondedNftsCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoulBondedNftsCollection.Merge(m, src)
}
func (m *SoulBondedNftsCollection) XXX_Size() int {
	return m.Size()
}
func (m *SoulBondedNftsCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_SoulBondedNftsCollection.DiscardUnknown(m)
}

var xxx_messageInfo_SoulBondedNftsCollection proto.InternalMessageInfo

func (m *SoulBondedNftsCollection) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func init() {
	proto.RegisterType((*SoulBondedNftsCollection)(nil), "mantrachain.token.v1.SoulBondedNftsCollection")
}

func init() {
	proto.RegisterFile("mantrachain/token/v1/soul_bonded_nfts_collection.proto", fileDescriptor_a8a002e287bfc1d5)
}

var fileDescriptor_a8a002e287bfc1d5 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcb, 0x4d, 0xcc, 0x2b,
	0x29, 0x4a, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2f, 0x33,
	0xd4, 0x2f, 0xce, 0x2f, 0xcd, 0x89, 0x4f, 0xca, 0xcf, 0x4b, 0x49, 0x4d, 0x89, 0xcf, 0x4b, 0x2b,
	0x29, 0x8e, 0x4f, 0xce, 0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x41, 0xd2, 0xa7, 0x07, 0xd6, 0xa7, 0x57, 0x66, 0xa8, 0x64, 0xc0, 0x25,
	0x11, 0x9c, 0x5f, 0x9a, 0xe3, 0x04, 0xd6, 0xe9, 0x97, 0x56, 0x52, 0xec, 0x0c, 0xd7, 0x27, 0x24,
	0xc2, 0xc5, 0x9a, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1,
	0x38, 0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x4f, 0x66, 0x6e, 0xaa, 0x33, 0xd8, 0x89,
	0xc8, 0xce, 0xad, 0x80, 0x3a, 0xb8, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x30, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x12, 0x75, 0xb2, 0xd2, 0x00, 0x00, 0x00,
}

func (m *SoulBondedNftsCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SoulBondedNftsCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SoulBondedNftsCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSoulBondedNftsCollection(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSoulBondedNftsCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovSoulBondedNftsCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SoulBondedNftsCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSoulBondedNftsCollection(uint64(l))
	}
	return n
}

func sovSoulBondedNftsCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSoulBondedNftsCollection(x uint64) (n int) {
	return sovSoulBondedNftsCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SoulBondedNftsCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSoulBondedNftsCollection
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
			return fmt.Errorf("proto: SoulBondedNftsCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SoulBondedNftsCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSoulBondedNftsCollection
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
				return ErrInvalidLengthSoulBondedNftsCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSoulBondedNftsCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSoulBondedNftsCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSoulBondedNftsCollection
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
func skipSoulBondedNftsCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSoulBondedNftsCollection
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
					return 0, ErrIntOverflowSoulBondedNftsCollection
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
					return 0, ErrIntOverflowSoulBondedNftsCollection
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
				return 0, ErrInvalidLengthSoulBondedNftsCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSoulBondedNftsCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSoulBondedNftsCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSoulBondedNftsCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSoulBondedNftsCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSoulBondedNftsCollection = fmt.Errorf("proto: unexpected end of group")
)
