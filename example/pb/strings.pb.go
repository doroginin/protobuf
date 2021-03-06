// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/strings.proto

/*
	Package strings is a generated protocol buffer package.

	It is generated from these files:
		pb/strings.proto

	It has these top-level messages:
		StringRequest
		StringResponse
*/
package strings

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StringRequest struct {
	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (m *StringRequest) Reset()                    { *m = StringRequest{} }
func (m *StringRequest) String() string            { return proto.CompactTextString(m) }
func (*StringRequest) ProtoMessage()               {}
func (*StringRequest) Descriptor() ([]byte, []int) { return fileDescriptorStrings, []int{0} }

func (m *StringRequest) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

type StringResponse struct {
	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (m *StringResponse) Reset()                    { *m = StringResponse{} }
func (m *StringResponse) String() string            { return proto.CompactTextString(m) }
func (*StringResponse) ProtoMessage()               {}
func (*StringResponse) Descriptor() ([]byte, []int) { return fileDescriptorStrings, []int{1} }

func (m *StringResponse) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func init() {
	proto.RegisterType((*StringRequest)(nil), "StringRequest")
	proto.RegisterType((*StringResponse)(nil), "StringResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Strings service

type StringsClient interface {
	ToUpper(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
	ToLower(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
}

type stringsClient struct {
	cc *grpc.ClientConn
}

func NewStringsClient(cc *grpc.ClientConn) StringsClient {
	return &stringsClient{cc}
}

func (c *stringsClient) ToUpper(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := grpc.Invoke(ctx, "/Strings/ToUpper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringsClient) ToLower(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := grpc.Invoke(ctx, "/Strings/ToLower", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Strings service

type StringsServer interface {
	ToUpper(context.Context, *StringRequest) (*StringResponse, error)
	ToLower(context.Context, *StringRequest) (*StringResponse, error)
}

func RegisterStringsServer(s *grpc.Server, srv StringsServer) {
	s.RegisterService(&_Strings_serviceDesc, srv)
}

func _Strings_ToUpper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringsServer).ToUpper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Strings/ToUpper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringsServer).ToUpper(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Strings_ToLower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringsServer).ToLower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Strings/ToLower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringsServer).ToLower(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Strings_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Strings",
	HandlerType: (*StringsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToUpper",
			Handler:    _Strings_ToUpper_Handler,
		},
		{
			MethodName: "ToLower",
			Handler:    _Strings_ToLower_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/strings.proto",
}

func (m *StringRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Str) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStrings(dAtA, i, uint64(len(m.Str)))
		i += copy(dAtA[i:], m.Str)
	}
	return i, nil
}

func (m *StringResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Str) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStrings(dAtA, i, uint64(len(m.Str)))
		i += copy(dAtA[i:], m.Str)
	}
	return i, nil
}

func encodeVarintStrings(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StringRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Str)
	if l > 0 {
		n += 1 + l + sovStrings(uint64(l))
	}
	return n
}

func (m *StringResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Str)
	if l > 0 {
		n += 1 + l + sovStrings(uint64(l))
	}
	return n
}

func sovStrings(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStrings(x uint64) (n int) {
	return sovStrings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStrings
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
			return fmt.Errorf("proto: StringRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStrings
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
				return ErrInvalidLengthStrings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStrings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStrings
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
func (m *StringResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStrings
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
			return fmt.Errorf("proto: StringResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStrings
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
				return ErrInvalidLengthStrings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStrings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStrings
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
func skipStrings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStrings
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
					return 0, ErrIntOverflowStrings
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
					return 0, ErrIntOverflowStrings
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
				return 0, ErrInvalidLengthStrings
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStrings
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
				next, err := skipStrings(dAtA[start:])
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
	ErrInvalidLengthStrings = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStrings   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/strings.proto", fileDescriptorStrings) }

var fileDescriptorStrings = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x2f,
	0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x49, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x83, 0xca, 0x2a, 0x29, 0x72, 0xf1, 0x06, 0x83, 0x95, 0x07, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x09, 0x70, 0x31, 0x17, 0x97, 0x14, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0x98, 0x4a, 0x4a, 0x5c, 0x7c, 0x30, 0x25, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x98,
	0x6a, 0x8c, 0xf6, 0x31, 0x72, 0xb1, 0x43, 0x14, 0x15, 0x0b, 0x65, 0x72, 0xb1, 0x87, 0xe4, 0x87,
	0x16, 0x14, 0xa4, 0x16, 0x09, 0xf1, 0xe9, 0xa1, 0x18, 0x2e, 0xc5, 0xaf, 0x87, 0x6a, 0x92, 0x92,
	0x6d, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xcc, 0x85, 0xc4, 0x61, 0x8e, 0xd6, 0x2f, 0xc9, 0x8f, 0x2f,
	0x05, 0xe9, 0xd5, 0xaf, 0x2e, 0x2e, 0x29, 0xaa, 0x8d, 0x92, 0x11, 0x92, 0xd2, 0x2f, 0x33, 0xc4,
	0x21, 0x2b, 0xe4, 0x01, 0xb2, 0xca, 0x27, 0xbf, 0x9c, 0x18, 0xab, 0x64, 0xc0, 0x56, 0x89, 0x29,
	0x09, 0x22, 0x1b, 0x96, 0x03, 0xd2, 0x6b, 0xc5, 0xa8, 0xe5, 0x24, 0x79, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x0e, 0x55,
	0x96, 0xc4, 0x06, 0x0e, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xf3, 0x88, 0xe2,
	0x5b, 0x01, 0x00, 0x00,
}
