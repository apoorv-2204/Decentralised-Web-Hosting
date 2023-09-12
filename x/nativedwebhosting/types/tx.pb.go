// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nativedwebhosting/nativedwebhosting/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgHostManifest struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Version      string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	HashFunction string `protobuf:"bytes,3,opt,name=hashFunction,proto3" json:"hashFunction,omitempty"`
}

func (m *MsgHostManifest) Reset()         { *m = MsgHostManifest{} }
func (m *MsgHostManifest) String() string { return proto.CompactTextString(m) }
func (*MsgHostManifest) ProtoMessage()    {}
func (*MsgHostManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b00cb2cbfc507cd1, []int{0}
}
func (m *MsgHostManifest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgHostManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgHostManifest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgHostManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHostManifest.Merge(m, src)
}
func (m *MsgHostManifest) XXX_Size() int {
	return m.Size()
}
func (m *MsgHostManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHostManifest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHostManifest proto.InternalMessageInfo

func (m *MsgHostManifest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgHostManifest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MsgHostManifest) GetHashFunction() string {
	if m != nil {
		return m.HashFunction
	}
	return ""
}

type MsgHostManifestResponse struct {
}

func (m *MsgHostManifestResponse) Reset()         { *m = MsgHostManifestResponse{} }
func (m *MsgHostManifestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgHostManifestResponse) ProtoMessage()    {}
func (*MsgHostManifestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b00cb2cbfc507cd1, []int{1}
}
func (m *MsgHostManifestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgHostManifestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgHostManifestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgHostManifestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHostManifestResponse.Merge(m, src)
}
func (m *MsgHostManifestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgHostManifestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHostManifestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHostManifestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgHostManifest)(nil), "nativedwebhosting.nativedwebhosting.MsgHostManifest")
	proto.RegisterType((*MsgHostManifestResponse)(nil), "nativedwebhosting.nativedwebhosting.MsgHostManifestResponse")
}

func init() {
	proto.RegisterFile("nativedwebhosting/nativedwebhosting/tx.proto", fileDescriptor_b00cb2cbfc507cd1)
}

var fileDescriptor_b00cb2cbfc507cd1 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc9, 0x4b, 0x2c, 0xc9,
	0x2c, 0x4b, 0x4d, 0x29, 0x4f, 0x4d, 0xca, 0xc8, 0x2f, 0x2e, 0xc9, 0xcc, 0x4b, 0xd7, 0xc7, 0x14,
	0x29, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xc6, 0x90, 0xd3, 0xc3, 0x10, 0x51,
	0xca, 0xe4, 0xe2, 0xf7, 0x2d, 0x4e, 0xf7, 0xc8, 0x2f, 0x2e, 0xf1, 0x4d, 0xcc, 0xcb, 0x4c, 0x4b,
	0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x41, 0x32, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12,
	0x4c, 0x10, 0x19, 0x28, 0x57, 0x48, 0x89, 0x8b, 0x27, 0x23, 0xb1, 0x38, 0xc3, 0xad, 0x34, 0x2f,
	0xb9, 0x04, 0x24, 0xcd, 0x0c, 0x96, 0x46, 0x11, 0x53, 0x92, 0xe4, 0x12, 0x47, 0xb3, 0x2a, 0x28,
	0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xa8, 0x8b, 0x91, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0xa8,
	0x89, 0x91, 0x8b, 0x07, 0xc5, 0x2d, 0x26, 0x7a, 0x44, 0x78, 0x42, 0x0f, 0xcd, 0x58, 0x29, 0x1b,
	0x72, 0x74, 0xc1, 0x1c, 0xe3, 0x14, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x2e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x89, 0x05, 0xf9,
	0xf9, 0x45, 0x65, 0xba, 0x46, 0x46, 0x06, 0x26, 0xfa, 0x7e, 0x60, 0xb3, 0x75, 0x5d, 0xc2, 0x53,
	0x93, 0x74, 0x3d, 0xa0, 0xd1, 0x50, 0x81, 0x2d, 0x6a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0,
	0xd1, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x9e, 0xa0, 0x7b, 0xce, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	HostManifest(ctx context.Context, in *MsgHostManifest, opts ...grpc.CallOption) (*MsgHostManifestResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) HostManifest(ctx context.Context, in *MsgHostManifest, opts ...grpc.CallOption) (*MsgHostManifestResponse, error) {
	out := new(MsgHostManifestResponse)
	err := c.cc.Invoke(ctx, "/nativedwebhosting.nativedwebhosting.Msg/HostManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	HostManifest(context.Context, *MsgHostManifest) (*MsgHostManifestResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) HostManifest(ctx context.Context, req *MsgHostManifest) (*MsgHostManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostManifest not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_HostManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgHostManifest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).HostManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nativedwebhosting.nativedwebhosting.Msg/HostManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).HostManifest(ctx, req.(*MsgHostManifest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nativedwebhosting.nativedwebhosting.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HostManifest",
			Handler:    _Msg_HostManifest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nativedwebhosting/nativedwebhosting/tx.proto",
}

func (m *MsgHostManifest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgHostManifest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgHostManifest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HashFunction) > 0 {
		i -= len(m.HashFunction)
		copy(dAtA[i:], m.HashFunction)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HashFunction)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgHostManifestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgHostManifestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgHostManifestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgHostManifest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HashFunction)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgHostManifestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgHostManifest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgHostManifest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgHostManifest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashFunction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashFunction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgHostManifestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgHostManifestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgHostManifestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
