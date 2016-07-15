// Code generated by protoc-gen-gogo.
// source: flags.proto
// DO NOT EDIT!

/*
	Package flavortown_flags is a generated protocol buffer package.

	It is generated from these files:
		flags.proto

	It has these top-level messages:
		User
*/
package flavortown_flags

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/opsee/protobuf/opseeproto"

import database_sql_driver "database/sql/driver"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type User struct {
	ScopeA bool `protobuf:"varint,1,opt,name=scopeA,proto3" json:"scopeA,omitempty"`
	ScopeB bool `protobuf:"varint,2,opt,name=scopeB,proto3" json:"scopeB,omitempty"`
	ScopeC bool `protobuf:"varint,3,opt,name=scopeC,proto3" json:"scopeC,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorFlags, []int{0} }

func init() {
	proto.RegisterType((*User)(nil), "flavortown.flags.User")
}
func (this *User) UInt64() uint64 {
	b := uint64(0)
	if this.ScopeA {
		b |= uint64(1) << uint64(0)
	}
	if this.ScopeB {
		b |= uint64(1) << uint64(1)
	}
	if this.ScopeC {
		b |= uint64(1) << uint64(2)
	}

	return b
}
func (this *User) HighFlags() []string {
	var b []string
	if this.ScopeA {
		b = append(b, "scope_a")
	}
	if this.ScopeB {
		b = append(b, "scope_b")
	}
	if this.ScopeC {
		b = append(b, "scope_c")
	}
	return b
}

func (this *User) LowFlags() []string {
	var b []string
	if !this.ScopeA {
		b = append(b, "scope_a")
	}
	if !this.ScopeB {
		b = append(b, "scope_b")
	}
	if !this.ScopeC {
		b = append(b, "scope_c")
	}
	return b
}

func (this *User) SetFlag(flag string) error {
	switch flag {
	case "scope_a":
		this.ScopeA = true
	case "scope_b":
		this.ScopeB = true
	case "scope_c":
		this.ScopeC = true
	default:
		return fmt.Errorf("invalid flag: %v", flag)
	}
	return nil
}
func (this *User) ClearFlag(flag string) error {
	switch flag {
	case "scope_a":
		this.ScopeA = false
	case "scope_b":
		this.ScopeB = false
	case "scope_c":
		this.ScopeC = false
	default:
		return fmt.Errorf("invalid flag: %v", flag)
	}
	return nil
}
func (this *User) SetFlags(flags ...string) []error {
	var errs []error
	for _, f := range flags {
		if err := this.SetFlag(f); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
func (this *User) ClearFlags(flags ...string) []error {
	var errs []error
	for _, f := range flags {
		if err := this.ClearFlag(f); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
func (this *User) TestFlag(flag string) bool {
	switch flag {
	case "scope_a":
		return this.ScopeA
	case "scope_b":
		return this.ScopeB
	case "scope_c":
		return this.ScopeC
	}
	return false
}
func (this *User) TestFlags(flags ...string) bool {
	for _, f := range flags {
		if !this.TestFlag(f) {
			return false
		}
	}
	return true
}
func (this *User) FromUInt64(b uint64) error {
	bb := b
	bb = b
	if bb&(uint64(1)<<uint(0)) > 0 {
		this.ScopeA = true
	} else {
		this.ScopeA = false
	}
	bb = b
	if bb&(uint64(1)<<uint(1)) > 0 {
		this.ScopeB = true
	} else {
		this.ScopeB = false
	}
	bb = b
	if bb&(uint64(1)<<uint(2)) > 0 {
		this.ScopeC = true
	} else {
		this.ScopeC = false
	}

	return nil
}
func (this *User) Scan(i interface{}) error {
	switch v := i.(type) {
	case int:
		return this.FromUInt64(uint64(v))
	case int32:
		return this.FromUInt64(uint64(v))
	case int64:
		return this.FromUInt64(uint64(v))
	case float32:
		return this.FromUInt64(uint64(v))
	case float64:
		return this.FromUInt64(uint64(v))
	}

	return fmt.Errorf("invalid type: %T", i)
}
func (this *User) Value() (database_sql_driver.Value, error) {
	return int64(this.UInt64()), nil
}
func (this *User) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ScopeA != that1.ScopeA {
		return false
	}
	if this.ScopeB != that1.ScopeB {
		return false
	}
	if this.ScopeC != that1.ScopeC {
		return false
	}
	return true
}

type UserGetter interface {
	GetUser() *User
}

var GraphQLUserType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLUserType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "flavortown_flagsUser",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"scopeA": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*User)
						if ok {
							return obj.ScopeA, nil
						}
						inter, ok := p.Source.(UserGetter)
						if ok {
							face := inter.GetUser()
							if face == nil {
								return nil, nil
							}
							return face.ScopeA, nil
						}
						return nil, fmt.Errorf("field scopeA not resolved")
					},
				},
				"scopeB": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*User)
						if ok {
							return obj.ScopeB, nil
						}
						inter, ok := p.Source.(UserGetter)
						if ok {
							face := inter.GetUser()
							if face == nil {
								return nil, nil
							}
							return face.ScopeB, nil
						}
						return nil, fmt.Errorf("field scopeB not resolved")
					},
				},
				"scopeC": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*User)
						if ok {
							return obj.ScopeC, nil
						}
						inter, ok := p.Source.(UserGetter)
						if ok {
							face := inter.GetUser()
							if face == nil {
								return nil, nil
							}
							return face.ScopeC, nil
						}
						return nil, fmt.Errorf("field scopeC not resolved")
					},
				},
			}
		}),
	})
}
func (m *User) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *User) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ScopeA {
		data[i] = 0x8
		i++
		if m.ScopeA {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.ScopeB {
		data[i] = 0x10
		i++
		if m.ScopeB {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.ScopeC {
		data[i] = 0x18
		i++
		if m.ScopeC {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Flags(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Flags(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFlags(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUser(r randyFlags, easy bool) *User {
	this := &User{}
	this.ScopeA = bool(bool(r.Intn(2) == 0))
	this.ScopeB = bool(bool(r.Intn(2) == 0))
	this.ScopeC = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyFlags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFlags(r randyFlags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFlags(r randyFlags) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneFlags(r)
	}
	return string(tmps)
}
func randUnrecognizedFlags(r randyFlags, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFlags(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFlags(data []byte, r randyFlags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFlags(data, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		data = encodeVarintPopulateFlags(data, uint64(v2))
	case 1:
		data = encodeVarintPopulateFlags(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFlags(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFlags(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFlags(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFlags(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.ScopeA {
		n += 2
	}
	if m.ScopeB {
		n += 2
	}
	if m.ScopeC {
		n += 2
	}
	return n
}

func sovFlags(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFlags(x uint64) (n int) {
	return sovFlags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlags
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeA", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScopeA = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeB", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScopeB = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeC", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScopeC = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFlags(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlags
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
func skipFlags(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlags
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFlags
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFlags
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipFlags(data[start:])
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
	ErrInvalidLengthFlags = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlags   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorFlags = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0x49, 0x4c,
	0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00, 0x72, 0xca, 0xf2, 0x8b, 0x4a, 0xf2,
	0xcb, 0xf3, 0xf4, 0xc0, 0xe2, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0x85, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03,
	0x66, 0x41, 0x0c, 0x90, 0x32, 0x40, 0x52, 0x9e, 0x5f, 0x50, 0x9c, 0x9a, 0x8a, 0x50, 0x0f, 0xe6,
	0x42, 0x34, 0x80, 0x99, 0x10, 0x1d, 0x4a, 0x21, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x62,
	0x5c, 0x6c, 0xc5, 0xc9, 0xf9, 0x05, 0xa9, 0x8e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x50,
	0x1e, 0x5c, 0xdc, 0x49, 0x82, 0x09, 0x49, 0xdc, 0x09, 0x2e, 0xee, 0x2c, 0xc1, 0x8c, 0x24, 0xee,
	0x6c, 0xc5, 0x72, 0x62, 0x91, 0x3c, 0xa3, 0x93, 0xcc, 0x8f, 0x87, 0x72, 0x8c, 0x2b, 0x1e, 0xc9,
	0x31, 0xee, 0x00, 0xe2, 0x13, 0x40, 0x7c, 0x01, 0x88, 0x1f, 0x00, 0xf1, 0x01, 0xa0, 0x6c, 0x12,
	0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x99, 0xb0, 0xf8, 0xfc, 0x00,
	0x00, 0x00,
}
