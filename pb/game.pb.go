// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type C2SMsg struct {
	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *C2SMsg) Reset()         { *m = C2SMsg{} }
func (m *C2SMsg) String() string { return proto.CompactTextString(m) }
func (*C2SMsg) ProtoMessage()    {}
func (*C2SMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}
func (m *C2SMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C2SMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C2SMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *C2SMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SMsg.Merge(m, src)
}
func (m *C2SMsg) XXX_Size() int {
	return m.Size()
}
func (m *C2SMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SMsg.DiscardUnknown(m)
}

var xxx_messageInfo_C2SMsg proto.InternalMessageInfo

func (m *C2SMsg) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *C2SMsg) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type S2CMsg struct {
	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *S2CMsg) Reset()         { *m = S2CMsg{} }
func (m *S2CMsg) String() string { return proto.CompactTextString(m) }
func (*S2CMsg) ProtoMessage()    {}
func (*S2CMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}
func (m *S2CMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2CMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2CMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2CMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CMsg.Merge(m, src)
}
func (m *S2CMsg) XXX_Size() int {
	return m.Size()
}
func (m *S2CMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CMsg.DiscardUnknown(m)
}

var xxx_messageInfo_S2CMsg proto.InternalMessageInfo

func (m *S2CMsg) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *S2CMsg) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type UserLoginReq struct {
	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (m *UserLoginReq) Reset()         { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string { return proto.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()    {}
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}
func (m *UserLoginReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserLoginReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginReq.Merge(m, src)
}
func (m *UserLoginReq) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginReq proto.InternalMessageInfo

func (m *UserLoginReq) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type UserLoginResp struct {
	Code     ErrCode `protobuf:"varint,1,opt,name=Code,proto3,enum=pb.ErrCode" json:"Code,omitempty"`
	ServerID uint32  `protobuf:"varint,2,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
}

func (m *UserLoginResp) Reset()         { *m = UserLoginResp{} }
func (m *UserLoginResp) String() string { return proto.CompactTextString(m) }
func (*UserLoginResp) ProtoMessage()    {}
func (*UserLoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}
func (m *UserLoginResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserLoginResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserLoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResp.Merge(m, src)
}
func (m *UserLoginResp) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResp proto.InternalMessageInfo

func (m *UserLoginResp) GetCode() ErrCode {
	if m != nil {
		return m.Code
	}
	return SUCCESS
}

func (m *UserLoginResp) GetServerID() uint32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

func init() {
	proto.RegisterType((*C2SMsg)(nil), "pb.C2SMsg")
	proto.RegisterType((*S2CMsg)(nil), "pb.S2CMsg")
	proto.RegisterType((*UserLoginReq)(nil), "pb.UserLoginReq")
	proto.RegisterType((*UserLoginResp)(nil), "pb.UserLoginResp")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x4a, 0xce, 0x4f, 0x81,
	0xf2, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa, 0x64, 0xc2,
	0xc5, 0xe6, 0x6c, 0x14, 0xec, 0x5b, 0x9c, 0x2e, 0x24, 0xc6, 0xc5, 0x16, 0x5a, 0x9c, 0x5a, 0xe4,
	0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0x04, 0xe5, 0x09, 0x09, 0x71, 0xb1, 0x38, 0xe5,
	0xa7, 0x54, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x20, 0x5d, 0xc1, 0x46, 0xce,
	0xa4, 0xea, 0x52, 0xe3, 0xe2, 0x01, 0xc9, 0xfa, 0xe4, 0xa7, 0x67, 0xe6, 0x05, 0xa5, 0x16, 0xe2,
	0xd2, 0xab, 0xe4, 0xc3, 0xc5, 0x8b, 0xa4, 0xae, 0xb8, 0x40, 0x48, 0x9e, 0x8b, 0xc5, 0x39, 0x3f,
	0x25, 0x15, 0xac, 0x8c, 0xcf, 0x88, 0x5b, 0xaf, 0x20, 0x49, 0xcf, 0xb5, 0xa8, 0x08, 0x24, 0x14,
	0x04, 0x96, 0x10, 0x92, 0xe2, 0xe2, 0x08, 0x4e, 0x2d, 0x2a, 0x03, 0x9b, 0x05, 0xb2, 0x91, 0x37,
	0x08, 0xce, 0x77, 0x52, 0xb9, 0xf0, 0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x61, 0xc6, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x07, 0x87, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x24, 0x9f, 0xc3, 0xf6, 0x42, 0x01, 0x00, 0x00,
}

func (m *C2SMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2SMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *C2SMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserID != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *S2CMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2CMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2CMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserID != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserLoginReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserLoginReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UserID != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserLoginResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserLoginResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ServerID != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.ServerID))
		i--
		dAtA[i] = 0x10
	}
	if m.Code != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *C2SMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserID != 0 {
		n += 1 + sovGame(uint64(m.UserID))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *S2CMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserID != 0 {
		n += 1 + sovGame(uint64(m.UserID))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *UserLoginReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserID != 0 {
		n += 1 + sovGame(uint64(m.UserID))
	}
	return n
}

func (m *UserLoginResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovGame(uint64(m.Code))
	}
	if m.ServerID != 0 {
		n += 1 + sovGame(uint64(m.ServerID))
	}
	return n
}

func sovGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *C2SMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: C2SMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2SMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *S2CMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: S2CMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2CMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *UserLoginReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: UserLoginReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *UserLoginResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: UserLoginResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= ErrCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			m.ServerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGame = fmt.Errorf("proto: unexpected end of group")
)
