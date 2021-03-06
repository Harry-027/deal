// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deal/new_deal.proto

package types

import (
	fmt "fmt"
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

type NewDeal struct {
	DealId     string `protobuf:"bytes,1,opt,name=dealId,proto3" json:"dealId,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Vendor     string `protobuf:"bytes,3,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Commission uint64 `protobuf:"varint,4,opt,name=commission,proto3" json:"commission,omitempty"`
}

func (m *NewDeal) Reset()         { *m = NewDeal{} }
func (m *NewDeal) String() string { return proto.CompactTextString(m) }
func (*NewDeal) ProtoMessage()    {}
func (*NewDeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26f39c3dcb2c0e1, []int{0}
}
func (m *NewDeal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NewDeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NewDeal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NewDeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewDeal.Merge(m, src)
}
func (m *NewDeal) XXX_Size() int {
	return m.Size()
}
func (m *NewDeal) XXX_DiscardUnknown() {
	xxx_messageInfo_NewDeal.DiscardUnknown(m)
}

var xxx_messageInfo_NewDeal proto.InternalMessageInfo

func (m *NewDeal) GetDealId() string {
	if m != nil {
		return m.DealId
	}
	return ""
}

func (m *NewDeal) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NewDeal) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *NewDeal) GetCommission() uint64 {
	if m != nil {
		return m.Commission
	}
	return 0
}

func init() {
	proto.RegisterType((*NewDeal)(nil), "Harry027.deal.deal.NewDeal")
}

func init() { proto.RegisterFile("deal/new_deal.proto", fileDescriptor_f26f39c3dcb2c0e1) }

var fileDescriptor_f26f39c3dcb2c0e1 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0x4d, 0xcc,
	0xd1, 0xcf, 0x4b, 0x2d, 0x8f, 0x07, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0x3c,
	0x12, 0x8b, 0x8a, 0x2a, 0x0d, 0x8c, 0xcc, 0xf5, 0xc0, 0x82, 0x20, 0x42, 0x29, 0x9f, 0x8b, 0xdd,
	0x2f, 0xb5, 0xdc, 0x25, 0x35, 0x31, 0x47, 0x48, 0x8c, 0x8b, 0x0d, 0x24, 0xe4, 0x99, 0x22, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x89, 0x70, 0xb1, 0xe6, 0x97, 0xe7, 0xa5, 0x16,
	0x49, 0x30, 0x81, 0x85, 0x21, 0x1c, 0x90, 0xea, 0xb2, 0xd4, 0xbc, 0x94, 0xfc, 0x22, 0x09, 0x66,
	0x88, 0x6a, 0x08, 0x4f, 0x48, 0x8e, 0x8b, 0x2b, 0x39, 0x3f, 0x37, 0x37, 0xb3, 0xb8, 0x38, 0x33,
	0x3f, 0x4f, 0x82, 0x45, 0x81, 0x51, 0x83, 0x25, 0x08, 0x49, 0xc4, 0xc9, 0xe1, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xc1, 0x2e, 0xd5, 0x35, 0x30, 0x32, 0xd7, 0x07, 0x7b, 0xa4, 0x02, 0x42, 0x95,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x7d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x0c, 0x25, 0x31, 0xe4, 0x00, 0x00, 0x00,
}

func (m *NewDeal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewDeal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NewDeal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Commission != 0 {
		i = encodeVarintNewDeal(dAtA, i, uint64(m.Commission))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Vendor) > 0 {
		i -= len(m.Vendor)
		copy(dAtA[i:], m.Vendor)
		i = encodeVarintNewDeal(dAtA, i, uint64(len(m.Vendor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNewDeal(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DealId) > 0 {
		i -= len(m.DealId)
		copy(dAtA[i:], m.DealId)
		i = encodeVarintNewDeal(dAtA, i, uint64(len(m.DealId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNewDeal(dAtA []byte, offset int, v uint64) int {
	offset -= sovNewDeal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NewDeal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DealId)
	if l > 0 {
		n += 1 + l + sovNewDeal(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNewDeal(uint64(l))
	}
	l = len(m.Vendor)
	if l > 0 {
		n += 1 + l + sovNewDeal(uint64(l))
	}
	if m.Commission != 0 {
		n += 1 + sovNewDeal(uint64(m.Commission))
	}
	return n
}

func sovNewDeal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNewDeal(x uint64) (n int) {
	return sovNewDeal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NewDeal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNewDeal
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
			return fmt.Errorf("proto: NewDeal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewDeal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DealId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewDeal
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
				return ErrInvalidLengthNewDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewDeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DealId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewDeal
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
				return ErrInvalidLengthNewDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewDeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewDeal
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
				return ErrInvalidLengthNewDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewDeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			m.Commission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewDeal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNewDeal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNewDeal
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
func skipNewDeal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNewDeal
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
					return 0, ErrIntOverflowNewDeal
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
					return 0, ErrIntOverflowNewDeal
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
				return 0, ErrInvalidLengthNewDeal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNewDeal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNewDeal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNewDeal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNewDeal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNewDeal = fmt.Errorf("proto: unexpected end of group")
)
