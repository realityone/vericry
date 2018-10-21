// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1.proto

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		v1.proto

	It has these top-level messages:
		QuestionRequest
		TokenReply
		QuestionReply
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

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

type QuestionRequest struct {
}

func (m *QuestionRequest) Reset()                    { *m = QuestionRequest{} }
func (m *QuestionRequest) String() string            { return proto.CompactTextString(m) }
func (*QuestionRequest) ProtoMessage()               {}
func (*QuestionRequest) Descriptor() ([]byte, []int) { return fileDescriptorV1, []int{0} }

type TokenReply struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PlainText string `protobuf:"bytes,2,opt,name=plain_text,json=plainText,proto3" json:"content"`
}

func (m *TokenReply) Reset()                    { *m = TokenReply{} }
func (m *TokenReply) String() string            { return proto.CompactTextString(m) }
func (*TokenReply) ProtoMessage()               {}
func (*TokenReply) Descriptor() ([]byte, []int) { return fileDescriptorV1, []int{1} }

func (m *TokenReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TokenReply) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

type QuestionReply struct {
	Tokens []*TokenReply `protobuf:"bytes,1,rep,name=tokens" json:"tokens"`
}

func (m *QuestionReply) Reset()                    { *m = QuestionReply{} }
func (m *QuestionReply) String() string            { return proto.CompactTextString(m) }
func (*QuestionReply) ProtoMessage()               {}
func (*QuestionReply) Descriptor() ([]byte, []int) { return fileDescriptorV1, []int{2} }

func (m *QuestionReply) GetTokens() []*TokenReply {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func init() {
	proto.RegisterType((*QuestionRequest)(nil), "QuestionRequest")
	proto.RegisterType((*TokenReply)(nil), "TokenReply")
	proto.RegisterType((*QuestionReply)(nil), "QuestionReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Vericry service

type VericryClient interface {
	GetQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionReply, error)
}

type vericryClient struct {
	cc *grpc.ClientConn
}

func NewVericryClient(cc *grpc.ClientConn) VericryClient {
	return &vericryClient{cc}
}

func (c *vericryClient) GetQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionReply, error) {
	out := new(QuestionReply)
	err := grpc.Invoke(ctx, "/Vericry/GetQuestion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Vericry service

type VericryServer interface {
	GetQuestion(context.Context, *QuestionRequest) (*QuestionReply, error)
}

func RegisterVericryServer(s *grpc.Server, srv VericryServer) {
	s.RegisterService(&_Vericry_serviceDesc, srv)
}

func _Vericry_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VericryServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Vericry/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VericryServer).GetQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vericry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Vericry",
	HandlerType: (*VericryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestion",
			Handler:    _Vericry_GetQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1.proto",
}

func (m *QuestionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *TokenReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintV1(dAtA, i, uint64(m.Id))
	}
	if len(m.PlainText) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintV1(dAtA, i, uint64(len(m.PlainText)))
		i += copy(dAtA[i:], m.PlainText)
	}
	return i, nil
}

func (m *QuestionReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestionReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for _, msg := range m.Tokens {
			dAtA[i] = 0xa
			i++
			i = encodeVarintV1(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintV1(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QuestionRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *TokenReply) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovV1(uint64(m.Id))
	}
	l = len(m.PlainText)
	if l > 0 {
		n += 1 + l + sovV1(uint64(l))
	}
	return n
}

func (m *QuestionReply) Size() (n int) {
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovV1(uint64(l))
		}
	}
	return n
}

func sovV1(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozV1(x uint64) (n int) {
	return sovV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuestionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV1
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
			return fmt.Errorf("proto: QuestionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV1
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
func (m *TokenReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV1
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
			return fmt.Errorf("proto: TokenReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
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
					return ErrIntOverflowV1
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
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlainText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV1
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
func (m *QuestionReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV1
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
			return fmt.Errorf("proto: QuestionReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestionReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV1
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthV1
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, &TokenReply{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV1
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
func skipV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowV1
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
					return 0, ErrIntOverflowV1
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
					return 0, ErrIntOverflowV1
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
				return 0, ErrInvalidLengthV1
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowV1
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
				next, err := skipV1(dAtA[start:])
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
	ErrInvalidLengthV1 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowV1   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1.proto", fileDescriptorV1) }

var fileDescriptorV1 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x33, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x97, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x0b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98,
	0x05, 0x51, 0xae, 0x24, 0xc8, 0xc5, 0x1f, 0x58, 0x9a, 0x5a, 0x5c, 0x92, 0x99, 0x9f, 0x17, 0x94,
	0x5a, 0x08, 0x62, 0x29, 0x05, 0x70, 0x71, 0x85, 0xe4, 0x67, 0xa7, 0xe6, 0x05, 0xa5, 0x16, 0xe4,
	0x54, 0x0a, 0x89, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x38, 0xb1, 0xbd,
	0xba, 0x27, 0xcf, 0x94, 0x99, 0x12, 0xc4, 0x94, 0x99, 0x22, 0xa4, 0xc5, 0xc5, 0x55, 0x90, 0x93,
	0x98, 0x99, 0x17, 0x5f, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xfd,
	0xea, 0x9e, 0x3c, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x49, 0x10, 0x27, 0x58, 0x3a, 0x24,
	0xb5, 0xa2, 0x44, 0xc9, 0x81, 0x8b, 0x17, 0x61, 0x09, 0xc8, 0x50, 0x7d, 0x2e, 0xb6, 0x12, 0x90,
	0x15, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xdc, 0x7a, 0x08, 0x1b, 0x9d, 0xb8, 0x5e,
	0xdd, 0x93, 0x87, 0x4a, 0x07, 0x41, 0x69, 0x23, 0x2b, 0x2e, 0xf6, 0xb0, 0xd4, 0xa2, 0xcc, 0xe4,
	0x22, 0x90, 0x5e, 0x6e, 0xf7, 0xd4, 0x12, 0x98, 0x79, 0x42, 0x02, 0x7a, 0x68, 0xee, 0x97, 0xe2,
	0xd3, 0x43, 0xb1, 0xcc, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x77, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x04, 0x37, 0xb9, 0x03, 0x36, 0x01, 0x00, 0x00,
}
