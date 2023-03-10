// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mantrachain/token/v1/nft_collection.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type NftCollection struct {
	Index    []byte                                        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	Id       string                                        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Images   []*TokenImage                                 `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty" yaml:"images"`
	Url      string                                        `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty" yaml:"url"`
	Links    []*TokenLink                                  `protobuf:"bytes,5,rep,name=links,proto3" json:"links,omitempty" yaml:"links"`
	Options  []*TokenOption                                `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty" yaml:"options"`
	Category string                                        `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty" yaml:"category"`
	Creator  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,8,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
}

func (m *NftCollection) Reset()         { *m = NftCollection{} }
func (m *NftCollection) String() string { return proto.CompactTextString(m) }
func (*NftCollection) ProtoMessage()    {}
func (*NftCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e853694fa05167, []int{0}
}
func (m *NftCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftCollection.Merge(m, src)
}
func (m *NftCollection) XXX_Size() int {
	return m.Size()
}
func (m *NftCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_NftCollection.DiscardUnknown(m)
}

var xxx_messageInfo_NftCollection proto.InternalMessageInfo

func (m *NftCollection) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *NftCollection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NftCollection) GetImages() []*TokenImage {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *NftCollection) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *NftCollection) GetLinks() []*TokenLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *NftCollection) GetOptions() []*TokenOption {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NftCollection) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *NftCollection) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func init() {
	proto.RegisterType((*NftCollection)(nil), "mantrachain.token.v1.NftCollection")
}

func init() {
	proto.RegisterFile("mantrachain/token/v1/nft_collection.proto", fileDescriptor_e9e853694fa05167)
}

var fileDescriptor_e9e853694fa05167 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0x96, 0xb6, 0x9b, 0x59, 0x07, 0x98, 0x1d, 0xa2, 0x49, 0x24, 0x99, 0x0f, 0xa8,
	0x1c, 0xe6, 0x68, 0x70, 0xe3, 0xb6, 0xf4, 0x80, 0x10, 0x15, 0x93, 0x2c, 0x4e, 0x48, 0x08, 0x65,
	0x8e, 0x97, 0x59, 0x8d, 0xed, 0xca, 0x76, 0xa7, 0xf5, 0x2d, 0x78, 0x0c, 0x1e, 0x85, 0xe3, 0x8e,
	0x9c, 0x22, 0xd4, 0x3e, 0x00, 0x52, 0x8e, 0x9c, 0x50, 0x9c, 0x74, 0xeb, 0x61, 0xea, 0x29, 0x56,
	0xbe, 0xdf, 0xf7, 0xfb, 0xe7, 0x8b, 0x3f, 0xf0, 0x46, 0xa4, 0xd2, 0xea, 0x94, 0x5e, 0xa7, 0x5c,
	0xc6, 0x56, 0xcd, 0x98, 0x8c, 0x6f, 0xce, 0x62, 0x79, 0x65, 0xbf, 0x53, 0x55, 0x14, 0x8c, 0x5a,
	0xae, 0x24, 0x9e, 0x6b, 0x65, 0x15, 0x3c, 0xda, 0x42, 0xb1, 0x43, 0xf1, 0xcd, 0xd9, 0xf1, 0x51,
	0xae, 0x72, 0xe5, 0x80, 0xb8, 0x3e, 0x35, 0xec, 0xf1, 0xc9, 0xa3, 0x5a, 0xaa, 0x84, 0xd8, 0xe8,
	0xd0, 0xdf, 0x1e, 0x18, 0x7d, 0xbe, 0xb2, 0x93, 0xfb, 0x18, 0xf8, 0x1a, 0xf4, 0xb9, 0xcc, 0xd8,
	0xad, 0xef, 0x45, 0xde, 0xf8, 0x20, 0x79, 0x5e, 0x95, 0xe1, 0xc1, 0x32, 0x15, 0xc5, 0x7b, 0xe4,
	0x5e, 0x23, 0xd2, 0x94, 0xe1, 0x2b, 0xd0, 0xe5, 0x99, 0xdf, 0x8d, 0xbc, 0xf1, 0x7e, 0x32, 0xaa,
	0xca, 0x70, 0xbf, 0x85, 0x32, 0x44, 0xba, 0x3c, 0x83, 0x9f, 0xc0, 0x80, 0x8b, 0x34, 0x67, 0xc6,
	0xef, 0x45, 0xbd, 0xf1, 0xd3, 0xb7, 0x11, 0x7e, 0xec, 0xc3, 0xf1, 0x97, 0xfa, 0xf0, 0xb1, 0x06,
	0x93, 0x17, 0x55, 0x19, 0x8e, 0x5a, 0x89, 0xeb, 0x44, 0xa4, 0x55, 0xc0, 0x08, 0xf4, 0x16, 0xba,
	0xf0, 0x9f, 0xb8, 0xb0, 0xc3, 0xaa, 0x0c, 0x41, 0xc3, 0x2d, 0x74, 0x81, 0x48, 0x5d, 0x82, 0x1f,
	0x40, 0xbf, 0xe0, 0x72, 0x66, 0xfc, 0xbe, 0x4b, 0x0b, 0x77, 0xa4, 0x4d, 0xb9, 0x9c, 0x6d, 0x8f,
	0xe5, 0xfa, 0x10, 0x69, 0xfa, 0xe1, 0x05, 0x18, 0xaa, 0x79, 0xfd, 0x23, 0x8c, 0x3f, 0x70, 0xaa,
	0x93, 0x1d, 0xaa, 0x0b, 0x47, 0x26, 0xb0, 0x2a, 0xc3, 0xc3, 0x46, 0xd6, 0xf6, 0x22, 0xb2, 0xb1,
	0xc0, 0x18, 0xec, 0xd1, 0xd4, 0xb2, 0x5c, 0xe9, 0xa5, 0x3f, 0x74, 0x03, 0xbc, 0xac, 0xca, 0xf0,
	0x59, 0x83, 0x6f, 0x2a, 0x88, 0xdc, 0x43, 0xf0, 0x1b, 0x18, 0x52, 0xcd, 0x52, 0xab, 0xb4, 0xbf,
	0xe7, 0xae, 0x60, 0xf2, 0xa0, 0x6f, 0x0b, 0xe8, 0x5f, 0x19, 0x9e, 0xe6, 0xdc, 0x5e, 0x2f, 0x2e,
	0x31, 0x55, 0x22, 0xa6, 0xca, 0x08, 0x65, 0xda, 0xc7, 0xa9, 0xc9, 0x66, 0xb1, 0x5d, 0xce, 0x99,
	0xc1, 0xe7, 0x94, 0x9e, 0x67, 0x99, 0x66, 0xc6, 0x90, 0x8d, 0x33, 0x99, 0xfe, 0x5c, 0x05, 0xde,
	0xaf, 0x55, 0xe0, 0xdd, 0xad, 0x02, 0xef, 0xcf, 0x2a, 0xf0, 0x7e, 0xac, 0x83, 0xce, 0xdd, 0x3a,
	0xe8, 0xfc, 0x5e, 0x07, 0x9d, 0xaf, 0x78, 0xcb, 0x3a, 0xe5, 0x82, 0x4d, 0xdc, 0xee, 0x6c, 0xef,
	0xd1, 0x6d, 0xbb, 0x49, 0x2e, 0xe1, 0x72, 0xe0, 0xd6, 0xe8, 0xdd, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x05, 0xe9, 0xcd, 0x06, 0xc2, 0x02, 0x00, 0x00,
}

func (this *NftCollection) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NftCollection)
	if !ok {
		that2, ok := that.(NftCollection)
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
	if !bytes.Equal(this.Index, that1.Index) {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if len(this.Images) != len(that1.Images) {
		return false
	}
	for i := range this.Images {
		if !this.Images[i].Equal(that1.Images[i]) {
			return false
		}
	}
	if this.Url != that1.Url {
		return false
	}
	if len(this.Links) != len(that1.Links) {
		return false
	}
	for i := range this.Links {
		if !this.Links[i].Equal(that1.Links[i]) {
			return false
		}
	}
	if len(this.Options) != len(that1.Options) {
		return false
	}
	for i := range this.Options {
		if !this.Options[i].Equal(that1.Options[i]) {
			return false
		}
	}
	if this.Category != that1.Category {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	return true
}
func (m *NftCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNftCollection(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = encodeVarintNftCollection(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Options[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintNftCollection(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Images) > 0 {
		for iNdEx := len(m.Images) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Images[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNftCollection(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNftCollection(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNftCollection(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNftCollection(uint64(l))
	}
	if len(m.Images) > 0 {
		for _, e := range m.Images {
			l = e.Size()
			n += 1 + l + sovNftCollection(uint64(l))
		}
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovNftCollection(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovNftCollection(uint64(l))
		}
	}
	if len(m.Options) > 0 {
		for _, e := range m.Options {
			l = e.Size()
			n += 1 + l + sovNftCollection(uint64(l))
		}
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovNftCollection(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNftCollection(uint64(l))
	}
	return n
}

func sovNftCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftCollection(x uint64) (n int) {
	return sovNftCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftCollection
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
			return fmt.Errorf("proto: NftCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
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
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Images", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Images = append(m.Images, &TokenImage{})
			if err := m.Images[len(m.Images)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &TokenLink{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, &TokenOption{})
			if err := m.Options[len(m.Options)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftCollection
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
				return ErrInvalidLengthNftCollection
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftCollection
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
func skipNftCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftCollection
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
					return 0, ErrIntOverflowNftCollection
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
					return 0, ErrIntOverflowNftCollection
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
				return 0, ErrInvalidLengthNftCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftCollection = fmt.Errorf("proto: unexpected end of group")
)
