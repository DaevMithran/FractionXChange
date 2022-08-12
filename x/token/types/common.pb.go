// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/v1/common.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type TokenOption struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	SubValue string `protobuf:"bytes,3,opt,name=sub_value,json=subValue,proto3" json:"sub_value,omitempty"`
}

func (m *TokenOption) Reset()         { *m = TokenOption{} }
func (m *TokenOption) String() string { return proto.CompactTextString(m) }
func (*TokenOption) ProtoMessage()    {}
func (*TokenOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_1848076597c24b04, []int{0}
}
func (m *TokenOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenOption.Merge(m, src)
}
func (m *TokenOption) XXX_Size() int {
	return m.Size()
}
func (m *TokenOption) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenOption.DiscardUnknown(m)
}

var xxx_messageInfo_TokenOption proto.InternalMessageInfo

type TokenAttribute struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	SubValue string `protobuf:"bytes,3,opt,name=sub_value,json=subValue,proto3" json:"sub_value,omitempty"`
}

func (m *TokenAttribute) Reset()         { *m = TokenAttribute{} }
func (m *TokenAttribute) String() string { return proto.CompactTextString(m) }
func (*TokenAttribute) ProtoMessage()    {}
func (*TokenAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1848076597c24b04, []int{1}
}
func (m *TokenAttribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenAttribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAttribute.Merge(m, src)
}
func (m *TokenAttribute) XXX_Size() int {
	return m.Size()
}
func (m *TokenAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAttribute proto.InternalMessageInfo

type TokenImage struct {
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *TokenImage) Reset()         { *m = TokenImage{} }
func (m *TokenImage) String() string { return proto.CompactTextString(m) }
func (*TokenImage) ProtoMessage()    {}
func (*TokenImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1848076597c24b04, []int{2}
}
func (m *TokenImage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenImage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenImage.Merge(m, src)
}
func (m *TokenImage) XXX_Size() int {
	return m.Size()
}
func (m *TokenImage) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenImage.DiscardUnknown(m)
}

var xxx_messageInfo_TokenImage proto.InternalMessageInfo

type TokenLink struct {
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *TokenLink) Reset()         { *m = TokenLink{} }
func (m *TokenLink) String() string { return proto.CompactTextString(m) }
func (*TokenLink) ProtoMessage()    {}
func (*TokenLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_1848076597c24b04, []int{3}
}
func (m *TokenLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenLink.Merge(m, src)
}
func (m *TokenLink) XXX_Size() int {
	return m.Size()
}
func (m *TokenLink) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenLink.DiscardUnknown(m)
}

var xxx_messageInfo_TokenLink proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TokenOption)(nil), "LimeChain.mantrachain.token.v1.TokenOption")
	proto.RegisterType((*TokenAttribute)(nil), "LimeChain.mantrachain.token.v1.TokenAttribute")
	proto.RegisterType((*TokenImage)(nil), "LimeChain.mantrachain.token.v1.TokenImage")
	proto.RegisterType((*TokenLink)(nil), "LimeChain.mantrachain.token.v1.TokenLink")
}

func init() { proto.RegisterFile("token/v1/common.proto", fileDescriptor_1848076597c24b04) }

var fileDescriptor_1848076597c24b04 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xc9, 0xcf, 0x4e,
	0xcd, 0xd3, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xf3, 0xc9, 0xcc, 0x4d, 0x75, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcb, 0x4d, 0xcc,
	0x2b, 0x29, 0x4a, 0x4c, 0x06, 0xb3, 0xc1, 0x8a, 0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x4a, 0xf5, 0x41, 0x2c, 0x88, 0x2e, 0xa5, 0x38, 0x2e, 0xee, 0x10, 0x90, 0x0a, 0xff,
	0x82, 0x92, 0xcc, 0xfc, 0x3c, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82,
	0x09, 0x2c, 0x08, 0xe1, 0x08, 0x49, 0x73, 0x71, 0x16, 0x97, 0x26, 0xc5, 0x43, 0x64, 0x98, 0xc1,
	0x32, 0x1c, 0xc5, 0xa5, 0x49, 0x61, 0x20, 0xbe, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86, 0x17, 0x0b,
	0xe4, 0x19, 0x95, 0x12, 0xb9, 0xf8, 0xc0, 0xe6, 0x3b, 0x96, 0x94, 0x14, 0x65, 0x26, 0x95, 0x96,
	0xa4, 0x52, 0xdf, 0x0a, 0x1b, 0x2e, 0x2e, 0xb0, 0x15, 0x9e, 0xb9, 0x89, 0xe9, 0xd8, 0x8d, 0x17,
	0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x81, 0x1a, 0x0e, 0x62, 0x22, 0xe9, 0xb6, 0xe6, 0xe2, 0x04, 0xeb,
	0xf6, 0xc9, 0xcc, 0xcb, 0x26, 0x55, 0xb3, 0x93, 0xcf, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x0f, 0x8f, 0x1c, 0x7d, 0xa4, 0xc8, 0xd1, 0xaf, 0xd0, 0x87, 0xc4, 0x25, 0xc8,
	0xa2, 0xe2, 0x24, 0x36, 0x70, 0x94, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x66, 0xc1, 0x34,
	0x61, 0xe1, 0x01, 0x00, 0x00,
}

func (this *TokenOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenOption)
	if !ok {
		that2, ok := that.(TokenOption)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.SubValue != that1.SubValue {
		return false
	}
	return true
}
func (this *TokenAttribute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenAttribute)
	if !ok {
		that2, ok := that.(TokenAttribute)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.SubValue != that1.SubValue {
		return false
	}
	return true
}
func (this *TokenImage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenImage)
	if !ok {
		that2, ok := that.(TokenImage)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	return true
}
func (this *TokenLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenLink)
	if !ok {
		that2, ok := that.(TokenLink)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	return true
}
func (m *TokenOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubValue) > 0 {
		i -= len(m.SubValue)
		copy(dAtA[i:], m.SubValue)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.SubValue)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenAttribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenAttribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenAttribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubValue) > 0 {
		i -= len(m.SubValue)
		copy(dAtA[i:], m.SubValue)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.SubValue)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenImage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenImage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenImage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
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
func (m *TokenOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.SubValue)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *TokenAttribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.SubValue)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *TokenImage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *TokenLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenOption) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TokenOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubValue", wireType)
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
			m.SubValue = string(dAtA[iNdEx:postIndex])
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
func (m *TokenAttribute) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TokenAttribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenAttribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubValue", wireType)
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
			m.SubValue = string(dAtA[iNdEx:postIndex])
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
func (m *TokenImage) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TokenImage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenImage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
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
func (m *TokenLink) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TokenLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
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