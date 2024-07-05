// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package lpfarmv1beta1

import (
	_ "cosmossdk.io/api/amino"
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Params_1_list)(nil)

type _Params_1_list struct {
	list *[]*v1beta1.Coin
}

func (x *_Params_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Params_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Params_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Params_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Params_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Params_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Params_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Params_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Params                           protoreflect.MessageDescriptor
	fd_Params_private_plan_creation_fee protoreflect.FieldDescriptor
	fd_Params_fee_collector             protoreflect.FieldDescriptor
	fd_Params_max_num_private_plans     protoreflect.FieldDescriptor
	fd_Params_max_block_duration        protoreflect.FieldDescriptor
)

func init() {
	file_mantrachain_lpfarm_v1beta1_params_proto_init()
	md_Params = File_mantrachain_lpfarm_v1beta1_params_proto.Messages().ByName("Params")
	fd_Params_private_plan_creation_fee = md_Params.Fields().ByName("private_plan_creation_fee")
	fd_Params_fee_collector = md_Params.Fields().ByName("fee_collector")
	fd_Params_max_num_private_plans = md_Params.Fields().ByName("max_num_private_plans")
	fd_Params_max_block_duration = md_Params.Fields().ByName("max_block_duration")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_mantrachain_lpfarm_v1beta1_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.PrivatePlanCreationFee) != 0 {
		value := protoreflect.ValueOfList(&_Params_1_list{list: &x.PrivatePlanCreationFee})
		if !f(fd_Params_private_plan_creation_fee, value) {
			return
		}
	}
	if x.FeeCollector != "" {
		value := protoreflect.ValueOfString(x.FeeCollector)
		if !f(fd_Params_fee_collector, value) {
			return
		}
	}
	if x.MaxNumPrivatePlans != uint32(0) {
		value := protoreflect.ValueOfUint32(x.MaxNumPrivatePlans)
		if !f(fd_Params_max_num_private_plans, value) {
			return
		}
	}
	if x.MaxBlockDuration != nil {
		value := protoreflect.ValueOfMessage(x.MaxBlockDuration.ProtoReflect())
		if !f(fd_Params_max_block_duration, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee":
		return len(x.PrivatePlanCreationFee) != 0
	case "mantrachain.lpfarm.v1beta1.Params.fee_collector":
		return x.FeeCollector != ""
	case "mantrachain.lpfarm.v1beta1.Params.max_num_private_plans":
		return x.MaxNumPrivatePlans != uint32(0)
	case "mantrachain.lpfarm.v1beta1.Params.max_block_duration":
		return x.MaxBlockDuration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mantrachain.lpfarm.v1beta1.Params"))
		}
		panic(fmt.Errorf("message mantrachain.lpfarm.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee":
		x.PrivatePlanCreationFee = nil
	case "mantrachain.lpfarm.v1beta1.Params.fee_collector":
		x.FeeCollector = ""
	case "mantrachain.lpfarm.v1beta1.Params.max_num_private_plans":
		x.MaxNumPrivatePlans = uint32(0)
	case "mantrachain.lpfarm.v1beta1.Params.max_block_duration":
		x.MaxBlockDuration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mantrachain.lpfarm.v1beta1.Params"))
		}
		panic(fmt.Errorf("message mantrachain.lpfarm.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee":
		if len(x.PrivatePlanCreationFee) == 0 {
			return protoreflect.ValueOfList(&_Params_1_list{})
		}
		listValue := &_Params_1_list{list: &x.PrivatePlanCreationFee}
		return protoreflect.ValueOfList(listValue)
	case "mantrachain.lpfarm.v1beta1.Params.fee_collector":
		value := x.FeeCollector
		return protoreflect.ValueOfString(value)
	case "mantrachain.lpfarm.v1beta1.Params.max_num_private_plans":
		value := x.MaxNumPrivatePlans
		return protoreflect.ValueOfUint32(value)
	case "mantrachain.lpfarm.v1beta1.Params.max_block_duration":
		value := x.MaxBlockDuration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mantrachain.lpfarm.v1beta1.Params"))
		}
		panic(fmt.Errorf("message mantrachain.lpfarm.v1beta1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee":
		lv := value.List()
		clv := lv.(*_Params_1_list)
		x.PrivatePlanCreationFee = *clv.list
	case "mantrachain.lpfarm.v1beta1.Params.fee_collector":
		x.FeeCollector = value.Interface().(string)
	case "mantrachain.lpfarm.v1beta1.Params.max_num_private_plans":
		x.MaxNumPrivatePlans = uint32(value.Uint())
	case "mantrachain.lpfarm.v1beta1.Params.max_block_duration":
		x.MaxBlockDuration = value.Message().Interface().(*durationpb.Duration)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mantrachain.lpfarm.v1beta1.Params"))
		}
		panic(fmt.Errorf("message mantrachain.lpfarm.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee":
		if x.PrivatePlanCreationFee == nil {
			x.PrivatePlanCreationFee = []*v1beta1.Coin{}
		}
		value := &_Params_1_list{list: &x.PrivatePlanCreationFee}
		return protoreflect.ValueOfList(value)
	case "mantrachain.lpfarm.v1beta1.Params.max_block_duration":
		if x.MaxBlockDuration == nil {
			x.MaxBlockDuration = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.MaxBlockDuration.ProtoReflect())
	case "mantrachain.lpfarm.v1beta1.Params.fee_collector":
		panic(fmt.Errorf("field fee_collector of message mantrachain.lpfarm.v1beta1.Params is not mutable"))
	case "mantrachain.lpfarm.v1beta1.Params.max_num_private_plans":
		panic(fmt.Errorf("field max_num_private_plans of message mantrachain.lpfarm.v1beta1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mantrachain.lpfarm.v1beta1.Params"))
		}
		panic(fmt.Errorf("message mantrachain.lpfarm.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_Params_1_list{list: &list})
	case "mantrachain.lpfarm.v1beta1.Params.fee_collector":
		return protoreflect.ValueOfString("")
	case "mantrachain.lpfarm.v1beta1.Params.max_num_private_plans":
		return protoreflect.ValueOfUint32(uint32(0))
	case "mantrachain.lpfarm.v1beta1.Params.max_block_duration":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mantrachain.lpfarm.v1beta1.Params"))
		}
		panic(fmt.Errorf("message mantrachain.lpfarm.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in mantrachain.lpfarm.v1beta1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.PrivatePlanCreationFee) > 0 {
			for _, e := range x.PrivatePlanCreationFee {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.FeeCollector)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxNumPrivatePlans != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxNumPrivatePlans))
		}
		if x.MaxBlockDuration != nil {
			l = options.Size(x.MaxBlockDuration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.MaxBlockDuration != nil {
			encoded, err := options.Marshal(x.MaxBlockDuration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.MaxNumPrivatePlans != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxNumPrivatePlans))
			i--
			dAtA[i] = 0x18
		}
		if len(x.FeeCollector) > 0 {
			i -= len(x.FeeCollector)
			copy(dAtA[i:], x.FeeCollector)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FeeCollector)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.PrivatePlanCreationFee) > 0 {
			for iNdEx := len(x.PrivatePlanCreationFee) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.PrivatePlanCreationFee[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PrivatePlanCreationFee", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PrivatePlanCreationFee = append(x.PrivatePlanCreationFee, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PrivatePlanCreationFee[len(x.PrivatePlanCreationFee)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FeeCollector", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FeeCollector = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxNumPrivatePlans", wireType)
				}
				x.MaxNumPrivatePlans = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxNumPrivatePlans |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxBlockDuration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MaxBlockDuration == nil {
					x.MaxBlockDuration = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxBlockDuration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: mantrachain/lpfarm/v1beta1/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the parameters for the module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivatePlanCreationFee []*v1beta1.Coin      `protobuf:"bytes,1,rep,name=private_plan_creation_fee,json=privatePlanCreationFee,proto3" json:"private_plan_creation_fee,omitempty"`
	FeeCollector           string               `protobuf:"bytes,2,opt,name=fee_collector,json=feeCollector,proto3" json:"fee_collector,omitempty"`
	MaxNumPrivatePlans     uint32               `protobuf:"varint,3,opt,name=max_num_private_plans,json=maxNumPrivatePlans,proto3" json:"max_num_private_plans,omitempty"`
	MaxBlockDuration       *durationpb.Duration `protobuf:"bytes,4,opt,name=max_block_duration,json=maxBlockDuration,proto3" json:"max_block_duration,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mantrachain_lpfarm_v1beta1_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_mantrachain_lpfarm_v1beta1_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetPrivatePlanCreationFee() []*v1beta1.Coin {
	if x != nil {
		return x.PrivatePlanCreationFee
	}
	return nil
}

func (x *Params) GetFeeCollector() string {
	if x != nil {
		return x.FeeCollector
	}
	return ""
}

func (x *Params) GetMaxNumPrivatePlans() uint32 {
	if x != nil {
		return x.MaxNumPrivatePlans
	}
	return 0
}

func (x *Params) GetMaxBlockDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxBlockDuration
	}
	return nil
}

var File_mantrachain_lpfarm_v1beta1_params_proto protoreflect.FileDescriptor

var file_mantrachain_lpfarm_v1beta1_params_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x70,
	0x66, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x61, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69,
	0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2,
	0x02, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x86, 0x01, 0x0a, 0x19, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf,
	0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x16, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x65, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f, 0x6e,
	0x75, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x12, 0x6d, 0x61,
	0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x10, 0x6d, 0x61, 0x78,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x24, 0xe8,
	0xa0, 0x1f, 0x01, 0x8a, 0xe7, 0xb0, 0x2a, 0x1b, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x78, 0x2f, 0x6c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x42, 0xf2, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x6c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x4d, 0x4c, 0x58, 0xaa, 0x02, 0x1a, 0x4d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x4d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5c, 0x4c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x26, 0x4d, 0x61, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x4c,
	0x70, 0x66, 0x61, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x4d, 0x61, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x4c, 0x70, 0x66, 0x61, 0x72, 0x6d, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mantrachain_lpfarm_v1beta1_params_proto_rawDescOnce sync.Once
	file_mantrachain_lpfarm_v1beta1_params_proto_rawDescData = file_mantrachain_lpfarm_v1beta1_params_proto_rawDesc
)

func file_mantrachain_lpfarm_v1beta1_params_proto_rawDescGZIP() []byte {
	file_mantrachain_lpfarm_v1beta1_params_proto_rawDescOnce.Do(func() {
		file_mantrachain_lpfarm_v1beta1_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_mantrachain_lpfarm_v1beta1_params_proto_rawDescData)
	})
	return file_mantrachain_lpfarm_v1beta1_params_proto_rawDescData
}

var file_mantrachain_lpfarm_v1beta1_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mantrachain_lpfarm_v1beta1_params_proto_goTypes = []interface{}{
	(*Params)(nil),              // 0: mantrachain.lpfarm.v1beta1.Params
	(*v1beta1.Coin)(nil),        // 1: cosmos.base.v1beta1.Coin
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_mantrachain_lpfarm_v1beta1_params_proto_depIdxs = []int32{
	1, // 0: mantrachain.lpfarm.v1beta1.Params.private_plan_creation_fee:type_name -> cosmos.base.v1beta1.Coin
	2, // 1: mantrachain.lpfarm.v1beta1.Params.max_block_duration:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mantrachain_lpfarm_v1beta1_params_proto_init() }
func file_mantrachain_lpfarm_v1beta1_params_proto_init() {
	if File_mantrachain_lpfarm_v1beta1_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mantrachain_lpfarm_v1beta1_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mantrachain_lpfarm_v1beta1_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mantrachain_lpfarm_v1beta1_params_proto_goTypes,
		DependencyIndexes: file_mantrachain_lpfarm_v1beta1_params_proto_depIdxs,
		MessageInfos:      file_mantrachain_lpfarm_v1beta1_params_proto_msgTypes,
	}.Build()
	File_mantrachain_lpfarm_v1beta1_params_proto = out.File
	file_mantrachain_lpfarm_v1beta1_params_proto_rawDesc = nil
	file_mantrachain_lpfarm_v1beta1_params_proto_goTypes = nil
	file_mantrachain_lpfarm_v1beta1_params_proto_depIdxs = nil
}
