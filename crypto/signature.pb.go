// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crypto/signature.proto

package crypto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SignatureRaw struct {
	Algorithm            int32    `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Sig                  []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	PubKey               []byte   `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureRaw) Reset()         { *m = SignatureRaw{} }
func (m *SignatureRaw) String() string { return proto.CompactTextString(m) }
func (*SignatureRaw) ProtoMessage()    {}
func (*SignatureRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_2e3c58efdb365d46, []int{0}
}
func (m *SignatureRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SignatureRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureRaw.Merge(dst, src)
}
func (m *SignatureRaw) XXX_Size() int {
	return m.Size()
}
func (m *SignatureRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureRaw.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureRaw proto.InternalMessageInfo

func (m *SignatureRaw) GetAlgorithm() int32 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *SignatureRaw) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *SignatureRaw) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterType((*SignatureRaw)(nil), "crypto.SignatureRaw")
}
func (m *SignatureRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureRaw) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Algorithm != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSignature(dAtA, i, uint64(m.Algorithm))
	}
	if len(m.Sig) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSignature(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	if len(m.PubKey) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSignature(dAtA, i, uint64(len(m.PubKey)))
		i += copy(dAtA[i:], m.PubKey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSignature(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SignatureRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Algorithm != 0 {
		n += 1 + sovSignature(uint64(m.Algorithm))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSignature(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSignature(x uint64) (n int) {
	return sovSignature(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignatureRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignature
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
			return fmt.Errorf("proto: SignatureRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Algorithm", wireType)
			}
			m.Algorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Algorithm |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSignature(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSignature
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSignature(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignature
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
					return 0, ErrIntOverflowSignature
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
					return 0, ErrIntOverflowSignature
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
				return 0, ErrInvalidLengthSignature
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSignature
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
				next, err := skipSignature(dAtA[start:])
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
	ErrInvalidLengthSignature = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignature   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("crypto/signature.proto", fileDescriptor_signature_2e3c58efdb365d46) }

var fileDescriptor_signature_2e3c58efdb365d46 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0xd7, 0x2f, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x83, 0x88, 0x2b, 0x85, 0x71, 0xf1, 0x04, 0xc3, 0xa4, 0x82, 0x12, 0xcb,
	0x85, 0x64, 0xb8, 0x38, 0x13, 0x73, 0xd2, 0xf3, 0x8b, 0x32, 0x4b, 0x32, 0x72, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0x10, 0x02, 0x42, 0x02, 0x5c, 0xcc, 0xc5, 0x99, 0xe9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x90, 0x18, 0x17, 0x5b, 0x41, 0x69, 0x92, 0x77, 0x6a, 0xa5,
	0x04, 0x33, 0x58, 0x10, 0xca, 0x73, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x5b, 0x6c, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x0d, 0x20, 0xdf, 0x92, 0x00, 0x00, 0x00,
}
