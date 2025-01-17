// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fractionnft/v1/fracctionnft.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type TokenizedNFT struct {
	Owner         string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	TimeoutHeight int64  `protobuf:"varint,2,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
}

func (m *TokenizedNFT) Reset()         { *m = TokenizedNFT{} }
func (m *TokenizedNFT) String() string { return proto.CompactTextString(m) }
func (*TokenizedNFT) ProtoMessage()    {}
func (*TokenizedNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c59428ea112953, []int{0}
}
func (m *TokenizedNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenizedNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenizedNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenizedNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenizedNFT.Merge(m, src)
}
func (m *TokenizedNFT) XXX_Size() int {
	return m.Size()
}
func (m *TokenizedNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenizedNFT.DiscardUnknown(m)
}

var xxx_messageInfo_TokenizedNFT proto.InternalMessageInfo

func (m *TokenizedNFT) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TokenizedNFT) GetTimeoutHeight() int64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

type NFTData struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *NFTData) Reset()         { *m = NFTData{} }
func (m *NFTData) String() string { return proto.CompactTextString(m) }
func (*NFTData) ProtoMessage()    {}
func (*NFTData) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c59428ea112953, []int{1}
}
func (m *NFTData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTData.Merge(m, src)
}
func (m *NFTData) XXX_Size() int {
	return m.Size()
}
func (m *NFTData) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTData.DiscardUnknown(m)
}

var xxx_messageInfo_NFTData proto.InternalMessageInfo

func (m *NFTData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NFTData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NFTData) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenizedNFT)(nil), "fractionnft.v1.TokenizedNFT")
	proto.RegisterType((*NFTData)(nil), "fractionnft.v1.NFTData")
}

func init() { proto.RegisterFile("fractionnft/v1/fracctionnft.proto", fileDescriptor_90c59428ea112953) }

var fileDescriptor_90c59428ea112953 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3d, 0x4f, 0x02, 0x31,
	0x18, 0xc7, 0x39, 0x11, 0x0d, 0x55, 0x49, 0x6c, 0x18, 0x4e, 0x86, 0x0b, 0x92, 0x98, 0xb0, 0x70,
	0x0d, 0x71, 0x70, 0x06, 0x0d, 0x71, 0x91, 0xc4, 0xf3, 0x5c, 0x5c, 0x48, 0xb9, 0x2b, 0xbd, 0xc6,
	0xb4, 0x0f, 0x69, 0x0b, 0xbe, 0x7c, 0x0a, 0x3f, 0x8c, 0x1f, 0xc2, 0x91, 0x38, 0x39, 0x1a, 0xf8,
	0x22, 0x86, 0x16, 0x93, 0x5b, 0x9a, 0xff, 0xcb, 0xaf, 0x6d, 0x9e, 0x07, 0x9d, 0xcf, 0x34, 0xcd,
	0xac, 0x00, 0xa5, 0x66, 0x96, 0x2c, 0xfb, 0x64, 0x6b, 0xff, 0x7d, 0x3c, 0xd7, 0x60, 0x01, 0x37,
	0x4a, 0x48, 0xbc, 0xec, 0xb7, 0x9a, 0x1c, 0x38, 0xb8, 0x8a, 0x6c, 0x95, 0xa7, 0x5a, 0xa7, 0x54,
	0x0a, 0x05, 0xc4, 0x9d, 0xbb, 0xe8, 0x2c, 0x03, 0x23, 0xc1, 0x4c, 0x3c, 0xeb, 0x8d, 0xaf, 0x3a,
	0x0c, 0x1d, 0xa7, 0xf0, 0xcc, 0x94, 0x78, 0x67, 0xf9, 0x78, 0x94, 0xe2, 0x18, 0xd5, 0xe0, 0x45,
	0x31, 0x1d, 0x06, 0xed, 0xa0, 0x5b, 0x1f, 0x86, 0xdf, 0x9f, 0xbd, 0xe6, 0xee, 0xc2, 0x20, 0xcf,
	0x35, 0x33, 0xe6, 0xc1, 0x6a, 0xa1, 0x78, 0xe2, 0x31, 0x7c, 0x81, 0x1a, 0x56, 0x48, 0x06, 0x0b,
	0x3b, 0x29, 0x98, 0xe0, 0x85, 0x0d, 0xf7, 0xda, 0x41, 0xb7, 0x9a, 0x9c, 0xec, 0xd2, 0x5b, 0x17,
	0x76, 0x1e, 0xd1, 0xe1, 0x78, 0x94, 0xde, 0x50, 0x4b, 0x31, 0x46, 0xfb, 0x8a, 0x4a, 0xe6, 0x3f,
	0x48, 0x9c, 0xc6, 0x6d, 0x74, 0x94, 0x33, 0x93, 0x69, 0x31, 0xdf, 0x8e, 0xe7, 0x9e, 0xa8, 0x27,
	0xe5, 0x08, 0x37, 0x51, 0x4d, 0x48, 0xca, 0x59, 0x58, 0x75, 0x9d, 0x37, 0xc3, 0xfb, 0xaf, 0x75,
	0x14, 0xac, 0xd6, 0x51, 0xf0, 0xbb, 0x8e, 0x82, 0x8f, 0x4d, 0x54, 0x59, 0x6d, 0xa2, 0xca, 0xcf,
	0x26, 0xaa, 0x3c, 0x5d, 0x71, 0x61, 0x8b, 0xc5, 0x34, 0xce, 0x40, 0x92, 0xbb, 0xc1, 0x38, 0x4d,
	0x06, 0xbd, 0xeb, 0x82, 0x0a, 0x45, 0x24, 0x55, 0x56, 0xd3, 0xcc, 0xe9, 0x57, 0x52, 0x5e, 0xba,
	0x7d, 0x9b, 0x33, 0x33, 0x3d, 0x70, 0x7b, 0xb9, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xac,
	0x93, 0x44, 0x90, 0x01, 0x00, 0x00,
}

func (m *TokenizedNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenizedNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenizedNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutHeight != 0 {
		i = encodeVarintFracctionnft(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintFracctionnft(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFTData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Image) > 0 {
		i -= len(m.Image)
		copy(dAtA[i:], m.Image)
		i = encodeVarintFracctionnft(dAtA, i, uint64(len(m.Image)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintFracctionnft(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFracctionnft(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFracctionnft(dAtA []byte, offset int, v uint64) int {
	offset -= sovFracctionnft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenizedNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovFracctionnft(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovFracctionnft(uint64(m.TimeoutHeight))
	}
	return n
}

func (m *NFTData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFracctionnft(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovFracctionnft(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovFracctionnft(uint64(l))
	}
	return n
}

func sovFracctionnft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFracctionnft(x uint64) (n int) {
	return sovFracctionnft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenizedNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFracctionnft
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
			return fmt.Errorf("proto: TokenizedNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenizedNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFracctionnft
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
				return ErrInvalidLengthFracctionnft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFracctionnft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFracctionnft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFracctionnft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFracctionnft
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
func (m *NFTData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFracctionnft
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
			return fmt.Errorf("proto: NFTData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFracctionnft
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
				return ErrInvalidLengthFracctionnft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFracctionnft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFracctionnft
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
				return ErrInvalidLengthFracctionnft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFracctionnft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFracctionnft
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
				return ErrInvalidLengthFracctionnft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFracctionnft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFracctionnft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFracctionnft
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
func skipFracctionnft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFracctionnft
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
					return 0, ErrIntOverflowFracctionnft
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
					return 0, ErrIntOverflowFracctionnft
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
				return 0, ErrInvalidLengthFracctionnft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFracctionnft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFracctionnft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFracctionnft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFracctionnft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFracctionnft = fmt.Errorf("proto: unexpected end of group")
)
