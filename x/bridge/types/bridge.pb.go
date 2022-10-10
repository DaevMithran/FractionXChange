// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/v1/bridge.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Bridge struct {
	Index               []byte                                        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	Id                  string                                        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	BridgeAccount       string                                        `protobuf:"bytes,3,opt,name=bridge_account,json=bridgeAccount,proto3" json:"bridge_account,omitempty" yaml:"bridge_account"`
	Cw20ContractAddress string                                        `protobuf:"bytes,4,opt,name=cw20_contract_address,json=cw20ContractAddress,proto3" json:"cw20_contract_address,omitempty" yaml:"cw20_contract_address"`
	Owner               github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty" yaml:"owner"`
	Creator             github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
}

func (m *Bridge) Reset()         { *m = Bridge{} }
func (m *Bridge) String() string { return proto.CompactTextString(m) }
func (*Bridge) ProtoMessage()    {}
func (*Bridge) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfdc6e997219362c, []int{0}
}
func (m *Bridge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bridge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bridge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bridge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bridge.Merge(m, src)
}
func (m *Bridge) XXX_Size() int {
	return m.Size()
}
func (m *Bridge) XXX_DiscardUnknown() {
	xxx_messageInfo_Bridge.DiscardUnknown(m)
}

var xxx_messageInfo_Bridge proto.InternalMessageInfo

func (m *Bridge) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Bridge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bridge) GetBridgeAccount() string {
	if m != nil {
		return m.BridgeAccount
	}
	return ""
}

func (m *Bridge) GetCw20ContractAddress() string {
	if m != nil {
		return m.Cw20ContractAddress
	}
	return ""
}

func (m *Bridge) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Bridge) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func init() {
	proto.RegisterType((*Bridge)(nil), "LimeChain.mantrachain.bridge.v1.Bridge")
}

func init() { proto.RegisterFile("bridge/v1/bridge.proto", fileDescriptor_dfdc6e997219362c) }

var fileDescriptor_dfdc6e997219362c = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x87, 0x9b, 0xf6, 0xb6, 0x57, 0xb5, 0xda, 0xea, 0x2a, 0x97, 0xa2, 0x80, 0x20, 0xa9, 0x32,
	0xa0, 0x2e, 0x8d, 0x29, 0x6c, 0x4c, 0x24, 0x5d, 0x61, 0x89, 0x90, 0x90, 0x90, 0x50, 0x95, 0xda,
	0x51, 0x6a, 0x41, 0xe2, 0xca, 0x71, 0xff, 0xbd, 0x05, 0x6f, 0xc0, 0xca, 0xa3, 0x30, 0x76, 0x64,
	0x8a, 0x50, 0xfa, 0x06, 0x19, 0x99, 0x50, 0x6d, 0x03, 0x45, 0x62, 0x62, 0xca, 0xc9, 0xef, 0x7c,
	0xf9, 0x9c, 0x23, 0x1f, 0xb0, 0x3b, 0x62, 0x04, 0x47, 0x21, 0x9c, 0xf5, 0xa1, 0xac, 0x9c, 0x09,
	0xa3, 0x9c, 0xea, 0xd6, 0x05, 0x89, 0xc3, 0xc1, 0x38, 0x20, 0x89, 0x13, 0x07, 0x09, 0x67, 0x01,
	0x12, 0xb5, 0x62, 0x66, 0xfd, 0xfd, 0x9d, 0x88, 0x46, 0x54, 0xb0, 0x70, 0x53, 0xc9, 0xcf, 0xec,
	0xc7, 0x0a, 0xa8, 0x79, 0x82, 0xd1, 0x8f, 0x40, 0x95, 0x24, 0x38, 0x5c, 0x18, 0x5a, 0x47, 0xeb,
	0x36, 0xbc, 0x7f, 0x45, 0x66, 0x35, 0x96, 0x41, 0x7c, 0x7f, 0x66, 0x8b, 0xd8, 0xf6, 0x65, 0x5b,
	0x3f, 0x04, 0x65, 0x82, 0x8d, 0x72, 0x47, 0xeb, 0xd6, 0xbd, 0x66, 0x91, 0x59, 0x75, 0x05, 0x61,
	0xdb, 0x2f, 0x13, 0xac, 0x9f, 0x83, 0x96, 0x3c, 0x74, 0x18, 0x20, 0x44, 0xa7, 0x09, 0x37, 0x2a,
	0x02, 0xdd, 0x2b, 0x32, 0xab, 0x2d, 0xd1, 0xef, 0x7d, 0xdb, 0x6f, 0xca, 0xc0, 0x95, 0xef, 0xfa,
	0x15, 0x68, 0xa3, 0xf9, 0xc9, 0xf1, 0x10, 0x51, 0x31, 0x07, 0x1f, 0x06, 0x18, 0xb3, 0x30, 0x4d,
	0x8d, 0x3f, 0x42, 0xd4, 0x29, 0x32, 0xeb, 0x40, 0x8a, 0x7e, 0xc4, 0x6c, 0xff, 0xff, 0x26, 0x1f,
	0xa8, 0xd8, 0x95, 0xa9, 0x7e, 0x0d, 0xaa, 0x74, 0x9e, 0x84, 0xcc, 0xa8, 0x8a, 0xf1, 0xdc, 0xaf,
	0xf1, 0x44, 0x6c, 0xbf, 0x65, 0x56, 0x2f, 0x22, 0x7c, 0x3c, 0x1d, 0x39, 0x88, 0xc6, 0x10, 0xd1,
	0x34, 0xa6, 0xa9, 0x7a, 0xf4, 0x52, 0x7c, 0x07, 0xf9, 0x72, 0x12, 0xa6, 0x8e, 0x8b, 0x90, 0x32,
	0xfa, 0xd2, 0xa7, 0xdf, 0x82, 0xbf, 0x88, 0x85, 0x01, 0xa7, 0xcc, 0xa8, 0x09, 0xf5, 0xa0, 0xc8,
	0xac, 0x96, 0xfa, 0x41, 0xd9, 0xf8, 0x85, 0xfc, 0xc3, 0xe9, 0x5d, 0x3e, 0xe5, 0xa6, 0xf6, 0x9c,
	0x9b, 0xda, 0x2a, 0x37, 0xb5, 0xd7, 0xdc, 0xd4, 0x1e, 0xd6, 0x66, 0x69, 0xb5, 0x36, 0x4b, 0x2f,
	0x6b, 0xb3, 0x74, 0x03, 0xb7, 0xac, 0x9f, 0x1b, 0x00, 0xb7, 0x36, 0x00, 0x2e, 0xd4, 0x9e, 0xc8,
	0x23, 0x46, 0x35, 0x71, 0xef, 0xa7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0xae, 0x1c, 0x7c,
	0x48, 0x02, 0x00, 0x00,
}

func (this *Bridge) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bridge)
	if !ok {
		that2, ok := that.(Bridge)
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
	if this.BridgeAccount != that1.BridgeAccount {
		return false
	}
	if this.Cw20ContractAddress != that1.Cw20ContractAddress {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	return true
}
func (m *Bridge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bridge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bridge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Cw20ContractAddress) > 0 {
		i -= len(m.Cw20ContractAddress)
		copy(dAtA[i:], m.Cw20ContractAddress)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.Cw20ContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BridgeAccount) > 0 {
		i -= len(m.BridgeAccount)
		copy(dAtA[i:], m.BridgeAccount)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.BridgeAccount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBridge(dAtA []byte, offset int, v uint64) int {
	offset -= sovBridge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bridge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = len(m.BridgeAccount)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = len(m.Cw20ContractAddress)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	return n
}

func sovBridge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBridge(x uint64) (n int) {
	return sovBridge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bridge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridge
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
			return fmt.Errorf("proto: Bridge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bridge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
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
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
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
					return ErrIntOverflowBridge
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
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
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
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cw20ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
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
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cw20ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
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
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
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
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
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
			skippy, err := skipBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridge
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
func skipBridge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBridge
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
					return 0, ErrIntOverflowBridge
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
					return 0, ErrIntOverflowBridge
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
				return 0, ErrInvalidLengthBridge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBridge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBridge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBridge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBridge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBridge = fmt.Errorf("proto: unexpected end of group")
)
