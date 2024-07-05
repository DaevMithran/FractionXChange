// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/coinfactory/v1beta1/authorityMetadata.proto

package types

import (
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

// DenomAuthorityMetadata specifies metadata for addresses that have specific
// capabilities over a token factory denom. Right now there is only one Admin
// permission, but is planned to be extended to the future.
type DenomAuthorityMetadata struct {
	// Can be empty for no admin, or a valid mantrachain address
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (m *DenomAuthorityMetadata) Reset()         { *m = DenomAuthorityMetadata{} }
func (m *DenomAuthorityMetadata) String() string { return proto.CompactTextString(m) }
func (*DenomAuthorityMetadata) ProtoMessage()    {}
func (*DenomAuthorityMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0920b16480dc7e7d, []int{0}
}
func (m *DenomAuthorityMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomAuthorityMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomAuthorityMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomAuthorityMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomAuthorityMetadata.Merge(m, src)
}
func (m *DenomAuthorityMetadata) XXX_Size() int {
	return m.Size()
}
func (m *DenomAuthorityMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomAuthorityMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DenomAuthorityMetadata proto.InternalMessageInfo

func (m *DenomAuthorityMetadata) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func init() {
	proto.RegisterType((*DenomAuthorityMetadata)(nil), "mantrachain.coinfactory.v1beta1.DenomAuthorityMetadata")
}

func init() {
	proto.RegisterFile("mantrachain/coinfactory/v1beta1/authorityMetadata.proto", fileDescriptor_0920b16480dc7e7d)
}

var fileDescriptor_0920b16480dc7e7d = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcf, 0x4d, 0xcc, 0x2b,
	0x29, 0x4a, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0xce, 0xcf, 0xcc, 0x4b, 0x4b, 0x4c, 0x2e,
	0xc9, 0x2f, 0xaa, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xca, 0x2c, 0xa9, 0xf4, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x47, 0xd2, 0xa8, 0x87, 0xa4, 0x51, 0x0f, 0xaa, 0x51, 0x4a, 0x32,
	0x39, 0xbf, 0x38, 0x37, 0xbf, 0x38, 0x1e, 0xac, 0x5c, 0x1f, 0xc2, 0x81, 0xe8, 0x95, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x87, 0x88, 0x83, 0x58, 0x10, 0x51, 0x25, 0x3f, 0x2e, 0x31, 0x97, 0xd4, 0xbc,
	0xfc, 0x5c, 0x47, 0x74, 0x1b, 0x85, 0xf4, 0xb8, 0x58, 0x13, 0x53, 0x72, 0x33, 0xf3, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x9d, 0x24, 0x2e, 0x6d, 0xd1, 0x15, 0x81, 0x1a, 0xe8, 0x98, 0x92, 0x52,
	0x94, 0x5a, 0x5c, 0x1c, 0x5c, 0x52, 0x94, 0x99, 0x97, 0x1e, 0x04, 0x51, 0x66, 0xc5, 0xf2, 0x62,
	0x81, 0x3c, 0xa3, 0x53, 0xf0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59,
	0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xfb, 0x3a, 0xfa, 0x85, 0x04,
	0x39, 0xea, 0xba, 0x65, 0xe6, 0x25, 0xe6, 0x25, 0xa7, 0xea, 0x23, 0x07, 0x47, 0x05, 0x4a, 0x80,
	0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xce, 0x37, 0x38, 0xa2, 0x38, 0x01, 0x00, 0x00,
}

func (this *DenomAuthorityMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DenomAuthorityMetadata)
	if !ok {
		that2, ok := that.(DenomAuthorityMetadata)
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
	if this.Admin != that1.Admin {
		return false
	}
	return true
}
func (m *DenomAuthorityMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomAuthorityMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomAuthorityMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintAuthorityMetadata(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthorityMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthorityMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomAuthorityMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovAuthorityMetadata(uint64(l))
	}
	return n
}

func sovAuthorityMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthorityMetadata(x uint64) (n int) {
	return sovAuthorityMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomAuthorityMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthorityMetadata
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
			return fmt.Errorf("proto: DenomAuthorityMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomAuthorityMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorityMetadata
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
				return ErrInvalidLengthAuthorityMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorityMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthorityMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthorityMetadata
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
func skipAuthorityMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthorityMetadata
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
					return 0, ErrIntOverflowAuthorityMetadata
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
					return 0, ErrIntOverflowAuthorityMetadata
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
				return 0, ErrInvalidLengthAuthorityMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthorityMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthorityMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthorityMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthorityMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthorityMetadata = fmt.Errorf("proto: unexpected end of group")
)
