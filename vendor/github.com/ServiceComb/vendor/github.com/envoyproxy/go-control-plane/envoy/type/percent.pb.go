// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/percent.proto

/*
	Package envoy_type is a generated protocol buffer package.

	It is generated from these files:
		envoy/type/percent.proto
		envoy/type/range.proto
		envoy/type/string_match.proto

	It has these top-level messages:
		Percent
		FractionalPercent
		Int64Range
		StringMatch
*/
package envoy_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Fraction percentages support several fixed denominator values.
type FractionalPercent_DenominatorType int32

const (
	// 100.
	//
	// **Example**: 1/100 = 1%.
	FractionalPercent_HUNDRED FractionalPercent_DenominatorType = 0
	// 10,000.
	//
	// **Example**: 1/10000 = 0.01%.
	FractionalPercent_TEN_THOUSAND FractionalPercent_DenominatorType = 1
	// 1,000,000.
	//
	// **Example**: 1/1000000 = 0.0001%.
	FractionalPercent_MILLION FractionalPercent_DenominatorType = 2
)

var FractionalPercent_DenominatorType_name = map[int32]string{
	0: "HUNDRED",
	1: "TEN_THOUSAND",
	2: "MILLION",
}
var FractionalPercent_DenominatorType_value = map[string]int32{
	"HUNDRED":      0,
	"TEN_THOUSAND": 1,
	"MILLION":      2,
}

func (x FractionalPercent_DenominatorType) String() string {
	return proto.EnumName(FractionalPercent_DenominatorType_name, int32(x))
}
func (FractionalPercent_DenominatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorPercent, []int{1, 0}
}

// Identifies a percentage, in the range [0.0, 100.0].
type Percent struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Percent) Reset()                    { *m = Percent{} }
func (m *Percent) String() string            { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()               {}
func (*Percent) Descriptor() ([]byte, []int) { return fileDescriptorPercent, []int{0} }

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// A fractional percentage is used in cases in which for performance reasons performing floating
// point to integer conversions during randomness calculations is undesirable. The message includes
// both a numerator and denominator that together determine the final fractional value.
//
// * **Example**: 1/100 = 1%.
// * **Example**: 3/10000 = 0.03%.
type FractionalPercent struct {
	// Specifies the numerator. Defaults to 0.
	Numerator uint32 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// Specifies the denominator. If the denominator specified is less than the numerator, the final
	// fractional percentage is capped at 1 (100%).
	Denominator FractionalPercent_DenominatorType `protobuf:"varint,2,opt,name=denominator,proto3,enum=envoy.type.FractionalPercent_DenominatorType" json:"denominator,omitempty"`
}

func (m *FractionalPercent) Reset()                    { *m = FractionalPercent{} }
func (m *FractionalPercent) String() string            { return proto.CompactTextString(m) }
func (*FractionalPercent) ProtoMessage()               {}
func (*FractionalPercent) Descriptor() ([]byte, []int) { return fileDescriptorPercent, []int{1} }

func (m *FractionalPercent) GetNumerator() uint32 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *FractionalPercent) GetDenominator() FractionalPercent_DenominatorType {
	if m != nil {
		return m.Denominator
	}
	return FractionalPercent_HUNDRED
}

func init() {
	proto.RegisterType((*Percent)(nil), "envoy.type.Percent")
	proto.RegisterType((*FractionalPercent)(nil), "envoy.type.FractionalPercent")
	proto.RegisterEnum("envoy.type.FractionalPercent_DenominatorType", FractionalPercent_DenominatorType_name, FractionalPercent_DenominatorType_value)
}
func (this *Percent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Percent)
	if !ok {
		that2, ok := that.(Percent)
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
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *FractionalPercent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FractionalPercent)
	if !ok {
		that2, ok := that.(FractionalPercent)
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
	if this.Numerator != that1.Numerator {
		return false
	}
	if this.Denominator != that1.Denominator {
		return false
	}
	return true
}
func (m *Percent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Percent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x9
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i += 8
	}
	return i, nil
}

func (m *FractionalPercent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FractionalPercent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Numerator != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPercent(dAtA, i, uint64(m.Numerator))
	}
	if m.Denominator != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPercent(dAtA, i, uint64(m.Denominator))
	}
	return i, nil
}

func encodeVarintPercent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Percent) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *FractionalPercent) Size() (n int) {
	var l int
	_ = l
	if m.Numerator != 0 {
		n += 1 + sovPercent(uint64(m.Numerator))
	}
	if m.Denominator != 0 {
		n += 1 + sovPercent(uint64(m.Denominator))
	}
	return n
}

func sovPercent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPercent(x uint64) (n int) {
	return sovPercent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Percent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPercent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Percent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Percent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipPercent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPercent
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
func (m *FractionalPercent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPercent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FractionalPercent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FractionalPercent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Numerator", wireType)
			}
			m.Numerator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPercent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Numerator |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denominator", wireType)
			}
			m.Denominator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPercent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Denominator |= (FractionalPercent_DenominatorType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPercent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPercent
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
func skipPercent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPercent
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
					return 0, ErrIntOverflowPercent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPercent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPercent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPercent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPercent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPercent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPercent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/type/percent.proto", fileDescriptorPercent) }

var fileDescriptorPercent = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0xcb, 0xe8, 0x81, 0x64, 0xa4, 0xc4, 0xcb, 0x12, 0x73,
	0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x22, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c,
	0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x59, 0x71, 0xb1, 0x07, 0x40, 0xcc, 0x12, 0xd2, 0xe7,
	0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x74, 0x92, 0xdc, 0xf5,
	0xf2, 0x00, 0xb3, 0x88, 0x90, 0x90, 0x24, 0x03, 0x18, 0x44, 0x3a, 0x68, 0x32, 0x40, 0x41, 0x10,
	0x44, 0x9d, 0xd2, 0x59, 0x46, 0x2e, 0x41, 0xb7, 0xa2, 0xc4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xc4,
	0x1c, 0x98, 0x31, 0x32, 0x5c, 0x9c, 0x79, 0xa5, 0xb9, 0xa9, 0x45, 0x89, 0x25, 0xf9, 0x45, 0x60,
	0xa3, 0x78, 0x83, 0x10, 0x02, 0x42, 0xd1, 0x5c, 0xdc, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0x99, 0x79,
	0x60, 0x79, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x5d, 0x3d, 0x84, 0x07, 0xf4, 0x30, 0x4c, 0xd4,
	0x73, 0x41, 0x68, 0x08, 0xa9, 0x2c, 0x48, 0x75, 0xe2, 0x02, 0xb9, 0x8c, 0xb5, 0x89, 0x91, 0x49,
	0x80, 0x31, 0x08, 0xd9, 0x34, 0x25, 0x5b, 0x2e, 0x7e, 0x34, 0xb5, 0x42, 0xdc, 0x5c, 0xec, 0x1e,
	0xa1, 0x7e, 0x2e, 0x41, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0x02, 0x5c, 0x3c, 0x21, 0xae, 0x7e, 0xf1,
	0x21, 0x1e, 0xfe, 0xa1, 0xc1, 0x8e, 0x7e, 0x2e, 0x02, 0x8c, 0x20, 0x69, 0x5f, 0x4f, 0x1f, 0x1f,
	0x4f, 0x7f, 0x3f, 0x01, 0x26, 0x27, 0x81, 0x15, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x24, 0x36, 0x70, 0x20, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0x07, 0xea, 0x97, 0x7b, 0x01, 0x00, 0x00,
}
