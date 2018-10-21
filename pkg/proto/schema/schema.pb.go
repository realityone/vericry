// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

/*
	Package schema is a generated protocol buffer package.

	It is generated from these files:
		schema.proto

	It has these top-level messages:
		Token
		Stat
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Token struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PlainText string `protobuf:"bytes,2,opt,name=plain_text,json=plainText,proto3" json:"content"`
	HashData  string `protobuf:"bytes,3,opt,name=hash_data,json=hashData,proto3" json:"hash_data"`
	Trusty    bool   `protobuf:"varint,4,opt,name=trusty,proto3" json:"trusty"`
	TrustyAt  uint64 `protobuf:"varint,5,opt,name=trusty_at,json=trustyAt,proto3" json:"trusty_at"`
	CreatedAt uint64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt uint64 `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	IsDeleted bool   `protobuf:"varint,8,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{0} }

func (m *Token) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Token) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

func (m *Token) GetHashData() string {
	if m != nil {
		return m.HashData
	}
	return ""
}

func (m *Token) GetTrusty() bool {
	if m != nil {
		return m.Trusty
	}
	return false
}

func (m *Token) GetTrustyAt() uint64 {
	if m != nil {
		return m.TrustyAt
	}
	return 0
}

func (m *Token) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Token) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Token) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

type Stat struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Succeed   uint64 `protobuf:"varint,2,opt,name=succeed,proto3" json:"succeed"`
	Failed    uint64 `protobuf:"varint,3,opt,name=failed,proto3" json:"failed"`
	CreatedAt uint64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt uint64 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (m *Stat) Reset()                    { *m = Stat{} }
func (m *Stat) String() string            { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()               {}
func (*Stat) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{1} }

func (m *Stat) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Stat) GetSucceed() uint64 {
	if m != nil {
		return m.Succeed
	}
	return 0
}

func (m *Stat) GetFailed() uint64 {
	if m != nil {
		return m.Failed
	}
	return 0
}

func (m *Stat) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Stat) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Token)(nil), "Token")
	proto.RegisterType((*Stat)(nil), "Stat")
}
func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.Id))
	}
	if len(m.PlainText) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.PlainText)))
		i += copy(dAtA[i:], m.PlainText)
	}
	if len(m.HashData) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.HashData)))
		i += copy(dAtA[i:], m.HashData)
	}
	if m.Trusty {
		dAtA[i] = 0x20
		i++
		if m.Trusty {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.TrustyAt != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.TrustyAt))
	}
	if m.CreatedAt != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.UpdatedAt))
	}
	if m.IsDeleted {
		dAtA[i] = 0x40
		i++
		if m.IsDeleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Stat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.Id))
	}
	if m.Succeed != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.Succeed))
	}
	if m.Failed != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.Failed))
	}
	if m.CreatedAt != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.UpdatedAt))
	}
	return i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Token) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSchema(uint64(m.Id))
	}
	l = len(m.PlainText)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.HashData)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.Trusty {
		n += 2
	}
	if m.TrustyAt != 0 {
		n += 1 + sovSchema(uint64(m.TrustyAt))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovSchema(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovSchema(uint64(m.UpdatedAt))
	}
	if m.IsDeleted {
		n += 2
	}
	return n
}

func (m *Stat) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSchema(uint64(m.Id))
	}
	if m.Succeed != 0 {
		n += 1 + sovSchema(uint64(m.Succeed))
	}
	if m.Failed != 0 {
		n += 1 + sovSchema(uint64(m.Failed))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovSchema(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovSchema(uint64(m.UpdatedAt))
	}
	return n
}

func sovSchema(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlainText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlainText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trusty", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Trusty = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustyAt", wireType)
			}
			m.TrustyAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustyAt |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsDeleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsDeleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *Stat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
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
			return fmt.Errorf("proto: Stat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Succeed", wireType)
			}
			m.Succeed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Succeed |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			m.Failed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Failed |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
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
				return 0, ErrInvalidLengthSchema
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchema
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
				next, err := skipSchema(dAtA[start:])
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
	ErrInvalidLengthSchema = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("schema.proto", fileDescriptorSchema) }

var fileDescriptorSchema = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0xef, 0xe4, 0xa6, 0x69, 0x32, 0xf7, 0x0f, 0x92, 0x85, 0x04, 0x17, 0x49, 0x29, 0x08,
	0xa5, 0xd0, 0x76, 0xe1, 0x13, 0x18, 0xfa, 0x04, 0x63, 0xf7, 0x61, 0x9a, 0x99, 0x36, 0x83, 0x6d,
	0xa6, 0x34, 0x27, 0x50, 0xdf, 0xc4, 0x37, 0xd2, 0xa5, 0x0b, 0xd7, 0x41, 0xea, 0x2e, 0x4f, 0x21,
	0x39, 0x93, 0x1a, 0x11, 0x44, 0x5c, 0xcd, 0x37, 0xbf, 0xfc, 0x18, 0xce, 0xf9, 0x08, 0xfd, 0x5b,
	0xa4, 0x99, 0xdc, 0xf2, 0xe9, 0x6e, 0xaf, 0x41, 0x5f, 0x4c, 0xd6, 0x0a, 0xb2, 0x72, 0x39, 0x4d,
	0xf5, 0x76, 0xb6, 0xd6, 0x6b, 0x3d, 0x43, 0xbc, 0x2c, 0x57, 0x78, 0xc3, 0x0b, 0x26, 0xa3, 0x0f,
	0x9f, 0x2d, 0xda, 0x5b, 0xe8, 0x5b, 0x99, 0xfb, 0xe7, 0xd4, 0x52, 0x22, 0x20, 0x03, 0x32, 0xb2,
	0x63, 0xa7, 0xae, 0x22, 0x4b, 0x09, 0x66, 0x29, 0xe1, 0x8f, 0x29, 0xdd, 0x6d, 0xb8, 0xca, 0x13,
	0x90, 0x07, 0x08, 0xac, 0x01, 0x19, 0x79, 0xf1, 0x9f, 0xba, 0x8a, 0xfa, 0xa9, 0xce, 0x41, 0xe6,
	0xc0, 0x3c, 0xfc, 0xbc, 0x90, 0x07, 0xf0, 0xc7, 0xd4, 0xcb, 0x78, 0x91, 0x25, 0x82, 0x03, 0x0f,
	0x7e, 0xa3, 0xfa, 0xaf, 0xae, 0xa2, 0x0e, 0x32, 0xb7, 0x89, 0x73, 0x0e, 0xdc, 0x1f, 0x52, 0x07,
	0xf6, 0x65, 0x01, 0x77, 0x81, 0x3d, 0x20, 0x23, 0x37, 0xa6, 0x75, 0x15, 0xb5, 0x84, 0xb5, 0x67,
	0xf3, 0x9e, 0x49, 0x09, 0x87, 0xa0, 0x87, 0xa3, 0xe1, 0x7b, 0xef, 0x90, 0xb9, 0x26, 0x5e, 0x83,
	0x3f, 0xa1, 0x34, 0xdd, 0x4b, 0x0e, 0x52, 0x34, 0xb2, 0x83, 0xf2, 0xff, 0xba, 0x8a, 0x3e, 0x50,
	0xe6, 0xb5, 0xd9, 0xe8, 0xe5, 0x4e, 0x9c, 0xf4, 0x7e, 0xa7, 0x77, 0x94, 0x79, 0x6d, 0x36, 0xba,
	0x2a, 0x12, 0x21, 0x37, 0x12, 0xa4, 0x08, 0x5c, 0x9c, 0x18, 0xf5, 0x8e, 0x32, 0x4f, 0x15, 0x73,
	0x13, 0x87, 0x0f, 0x84, 0xda, 0x37, 0xc0, 0xe1, 0xcb, 0x56, 0x2f, 0x69, 0xbf, 0x28, 0xd3, 0x54,
	0x4a, 0x81, 0x95, 0xda, 0xa6, 0xd2, 0x16, 0xb1, 0x53, 0x68, 0x4a, 0x5a, 0x71, 0xb5, 0x91, 0x02,
	0xdb, 0xb4, 0x4d, 0x49, 0x86, 0xb0, 0xf6, 0xfc, 0xb4, 0xb8, 0xfd, 0xb3, 0xc5, 0x7b, 0xdf, 0x2c,
	0x1e, 0x9f, 0x3d, 0x1e, 0x43, 0xf2, 0x74, 0x0c, 0xc9, 0xcb, 0x31, 0x24, 0xf7, 0xaf, 0xe1, 0xaf,
	0xa5, 0x83, 0x7f, 0xce, 0xd5, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x92, 0xf9, 0xd6, 0x78,
	0x02, 0x00, 0x00,
}