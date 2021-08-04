// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/nft_transfer/v1/nft_transfer.proto

package types

import (
	fmt "fmt"
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

type NonFungibleTokenPacketData struct {
	// the class to which the NFT to be transferred belongs
	Class string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	// the nft id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// the address defined by NFT outside the chain
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// the nft sender
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the nft receiver
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// identify whether it is far away from the source chain
	AwayFromOrigin bool `protobuf:"varint,6,opt,name=away_from_origin,json=awayFromOrigin,proto3" json:"away_from_origin,omitempty"`
}

func (m *NonFungibleTokenPacketData) Reset()         { *m = NonFungibleTokenPacketData{} }
func (m *NonFungibleTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*NonFungibleTokenPacketData) ProtoMessage()    {}
func (*NonFungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4237993fda6e21, []int{0}
}
func (m *NonFungibleTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonFungibleTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonFungibleTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonFungibleTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonFungibleTokenPacketData.Merge(m, src)
}
func (m *NonFungibleTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *NonFungibleTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonFungibleTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_NonFungibleTokenPacketData proto.InternalMessageInfo

func (m *NonFungibleTokenPacketData) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetAwayFromOrigin() bool {
	if m != nil {
		return m.AwayFromOrigin
	}
	return false
}

func init() {
	proto.RegisterType((*NonFungibleTokenPacketData)(nil), "ibc.applications.nft_transfer.v1.NonFungibleTokenPacketData")
}

func init() {
	proto.RegisterFile("ibc/applications/nft_transfer/v1/nft_transfer.proto", fileDescriptor_0e4237993fda6e21)
}

var fileDescriptor_0e4237993fda6e21 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0x32, 0x41,
	0x14, 0x85, 0x19, 0xf8, 0x21, 0xfc, 0x53, 0x10, 0x32, 0x21, 0x66, 0x43, 0xb1, 0x21, 0x56, 0x34,
	0xee, 0x84, 0xd0, 0x58, 0x1b, 0x43, 0xa9, 0x86, 0x58, 0xd9, 0x90, 0x99, 0xd9, 0x61, 0xbc, 0x61,
	0x77, 0xee, 0x66, 0x66, 0x16, 0xc3, 0x5b, 0xf8, 0x2c, 0x3e, 0x85, 0x25, 0xa5, 0xa5, 0x81, 0x17,
	0x31, 0x3b, 0x18, 0x23, 0x76, 0xf7, 0xfb, 0xee, 0x39, 0xcd, 0xa1, 0x73, 0x90, 0x8a, 0x8b, 0xaa,
	0x2a, 0x40, 0x89, 0x00, 0x68, 0x3d, 0xb7, 0xeb, 0xb0, 0x0a, 0x4e, 0x58, 0xbf, 0xd6, 0x8e, 0x6f,
	0x67, 0x67, 0x9c, 0x55, 0x0e, 0x03, 0xb2, 0x09, 0x48, 0x95, 0xfd, 0x2e, 0x65, 0x67, 0xa1, 0xed,
	0x6c, 0x3c, 0x32, 0x68, 0x30, 0x86, 0x79, 0x73, 0x9d, 0x7a, 0x97, 0x6f, 0x84, 0x8e, 0xef, 0xd0,
	0x2e, 0x6a, 0x6b, 0x40, 0x16, 0xfa, 0x11, 0x37, 0xda, 0x3e, 0x08, 0xb5, 0xd1, 0xe1, 0x56, 0x04,
	0xc1, 0x46, 0xb4, 0xab, 0x0a, 0xe1, 0x7d, 0x42, 0x26, 0x64, 0xfa, 0x7f, 0x79, 0x02, 0x36, 0xa0,
	0x6d, 0xc8, 0x93, 0x76, 0x54, 0x6d, 0xc8, 0xd9, 0x90, 0x76, 0x6a, 0x07, 0x49, 0x27, 0x8a, 0xe6,
	0x64, 0x17, 0xb4, 0xe7, 0xb5, 0xcd, 0xb5, 0x4b, 0xfe, 0x45, 0xf9, 0x4d, 0x6c, 0x4c, 0xfb, 0x4e,
	0x2b, 0x0d, 0x5b, 0xed, 0x92, 0x6e, 0xfc, 0xfc, 0x30, 0x9b, 0xd2, 0xa1, 0x78, 0x11, 0xbb, 0xd5,
	0xda, 0x61, 0xb9, 0x42, 0x07, 0x06, 0x6c, 0xd2, 0x9b, 0x90, 0x69, 0x7f, 0x39, 0x68, 0xfc, 0xc2,
	0x61, 0x79, 0x1f, 0xed, 0xcd, 0xf2, 0xfd, 0x90, 0x92, 0xfd, 0x21, 0x25, 0x9f, 0x87, 0x94, 0xbc,
	0x1e, 0xd3, 0xd6, 0xfe, 0x98, 0xb6, 0x3e, 0x8e, 0x69, 0xeb, 0xe9, 0xda, 0x40, 0x78, 0xae, 0x65,
	0xa6, 0xb0, 0xe4, 0x0a, 0x7d, 0x89, 0x9e, 0x83, 0x54, 0x57, 0x06, 0x79, 0x89, 0x79, 0x5d, 0x68,
	0xdf, 0x0c, 0xfb, 0x67, 0xd0, 0xb0, 0xab, 0xb4, 0x97, 0xbd, 0xb8, 0xc7, 0xfc, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x24, 0xb0, 0x7f, 0x65, 0x7e, 0x01, 0x00, 0x00,
}

func (m *NonFungibleTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonFungibleTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonFungibleTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AwayFromOrigin {
		i--
		if m.AwayFromOrigin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Class) > 0 {
		i -= len(m.Class)
		copy(dAtA[i:], m.Class)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Class)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NonFungibleTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Class)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	if m.AwayFromOrigin {
		n += 2
	}
	return n
}

func sovNftTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftTransfer(x uint64) (n int) {
	return sovNftTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NonFungibleTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftTransfer
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
			return fmt.Errorf("proto: NonFungibleTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonFungibleTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Class = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwayFromOrigin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AwayFromOrigin = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftTransfer
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
func skipNftTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftTransfer
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
					return 0, ErrIntOverflowNftTransfer
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
					return 0, ErrIntOverflowNftTransfer
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
				return 0, ErrInvalidLengthNftTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftTransfer = fmt.Errorf("proto: unexpected end of group")
)
