// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/claim/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the claim module's parameters.
type Params struct {
	Airdrops []*Airdrop `protobuf:"bytes,1,rep,name=airdrops,proto3" json:"airdrops,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ac871d3875dc3, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAirdrops() []*Airdrop {
	if m != nil {
		return m.Airdrops
	}
	return nil
}

type Airdrop struct {
	AirdropIdentifier string `protobuf:"bytes,1,opt,name=airdrop_identifier,json=airdropIdentifier,proto3" json:"airdrop_identifier,omitempty" yaml:"airdrop_identifier"`
	// seconds
	AirdropStartTime time.Time `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	// seconds
	AirdropDuration time.Duration `protobuf:"bytes,3,opt,name=airdrop_duration,json=airdropDuration,proto3,stdduration" json:"airdrop_duration,omitempty" yaml:"airdrop_duration"`
	// denom of claimable asset
	ClaimDenom string `protobuf:"bytes,4,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
	// airdrop distribution account
	DistributorAddress string `protobuf:"bytes,5,opt,name=distributor_address,json=distributorAddress,proto3" json:"distributor_address,omitempty"`
	// ustrd tokens claimed so far in the current period
	ClaimedSoFar github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=claimed_so_far,json=claimedSoFar,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"claimed_so_far"`
}

func (m *Airdrop) Reset()         { *m = Airdrop{} }
func (m *Airdrop) String() string { return proto.CompactTextString(m) }
func (*Airdrop) ProtoMessage()    {}
func (*Airdrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7ac871d3875dc3, []int{1}
}
func (m *Airdrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Airdrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Airdrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Airdrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airdrop.Merge(m, src)
}
func (m *Airdrop) XXX_Size() int {
	return m.Size()
}
func (m *Airdrop) XXX_DiscardUnknown() {
	xxx_messageInfo_Airdrop.DiscardUnknown(m)
}

var xxx_messageInfo_Airdrop proto.InternalMessageInfo

func (m *Airdrop) GetAirdropIdentifier() string {
	if m != nil {
		return m.AirdropIdentifier
	}
	return ""
}

func (m *Airdrop) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Airdrop) GetAirdropDuration() time.Duration {
	if m != nil {
		return m.AirdropDuration
	}
	return 0
}

func (m *Airdrop) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func (m *Airdrop) GetDistributorAddress() string {
	if m != nil {
		return m.DistributorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "stride.claim.Params")
	proto.RegisterType((*Airdrop)(nil), "stride.claim.Airdrop")
}

func init() { proto.RegisterFile("stride/claim/params.proto", fileDescriptor_dd7ac871d3875dc3) }

var fileDescriptor_dd7ac871d3875dc3 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x3f, 0x8f, 0x13, 0x3f,
	0x10, 0xcd, 0xfe, 0xf2, 0x23, 0x80, 0x73, 0xe2, 0x8f, 0x01, 0xb1, 0x89, 0xc4, 0x6e, 0xb4, 0x12,
	0x28, 0x05, 0x67, 0x8b, 0xa3, 0x41, 0x50, 0x25, 0x3a, 0x21, 0x4e, 0xba, 0x02, 0x6d, 0xae, 0xa2,
	0x59, 0x79, 0x63, 0x67, 0xb1, 0x88, 0xe3, 0x95, 0xed, 0x20, 0xf2, 0x09, 0x68, 0xaf, 0xe4, 0x23,
	0x5d, 0x79, 0x25, 0xa2, 0x58, 0x50, 0xd2, 0x51, 0x5e, 0x49, 0x85, 0xec, 0xf5, 0xde, 0x85, 0x4b,
	0x95, 0xec, 0x7b, 0x6f, 0xde, 0xbc, 0x19, 0x0f, 0xe8, 0x69, 0xa3, 0x38, 0x65, 0x78, 0x3a, 0x27,
	0x5c, 0xe0, 0x92, 0x28, 0x22, 0x34, 0x2a, 0x95, 0x34, 0x12, 0xee, 0xd5, 0x14, 0x72, 0x54, 0xff,
	0x61, 0x21, 0x0b, 0xe9, 0x08, 0x6c, 0xff, 0xd5, 0x9a, 0x7e, 0x54, 0x48, 0x59, 0xcc, 0x19, 0x76,
	0x5f, 0xf9, 0x72, 0x86, 0xe9, 0x52, 0x11, 0xc3, 0xe5, 0xc2, 0xf3, 0xf1, 0x75, 0xde, 0x70, 0xc1,
	0xb4, 0x21, 0xa2, 0xac, 0x05, 0xc9, 0x1b, 0xd0, 0x79, 0xef, 0x9a, 0xc2, 0x17, 0xe0, 0x16, 0xe1,
	0x8a, 0x2a, 0x59, 0xea, 0x30, 0x18, 0xb4, 0x87, 0xdd, 0x83, 0x47, 0x68, 0x3b, 0x01, 0x1a, 0xd5,
	0x6c, 0x7a, 0x29, 0x4b, 0xfe, 0xb4, 0xc1, 0x4d, 0x8f, 0xc2, 0x63, 0x00, 0x3d, 0x9e, 0x71, 0xca,
	0x16, 0x86, 0xcf, 0x38, 0x53, 0x61, 0x30, 0x08, 0x86, 0xb7, 0xc7, 0x4f, 0x2e, 0xaa, 0xb8, 0xb7,
	0x22, 0x62, 0xfe, 0x3a, 0xd9, 0xd5, 0x24, 0xe9, 0x7d, 0x0f, 0x1e, 0x5d, 0x62, 0x50, 0x5e, 0xb9,
	0x69, 0x43, 0x94, 0xc9, 0x6c, 0xee, 0xf0, 0xbf, 0x41, 0x30, 0xec, 0x1e, 0xf4, 0x51, 0x3d, 0x14,
	0x6a, 0x86, 0x42, 0x27, 0xcd, 0x50, 0xe3, 0xa7, 0x67, 0x55, 0xdc, 0xda, 0xed, 0x76, 0xe5, 0x91,
	0x9c, 0xfe, 0x8c, 0x83, 0xf4, 0x9e, 0x27, 0x26, 0x16, 0xb7, 0xd5, 0xf0, 0x6b, 0x00, 0x1a, 0x30,
	0x6b, 0x76, 0x18, 0xb6, 0x5d, 0xbf, 0xde, 0x4e, 0xbf, 0x43, 0x2f, 0x18, 0x8f, 0x6c, 0xbb, 0xdf,
	0x55, 0xdc, 0xbf, 0x5e, 0xfa, 0x5c, 0x0a, 0x6e, 0x98, 0x28, 0xcd, 0xea, 0xa2, 0x8a, 0x1f, 0xff,
	0x1b, 0xa6, 0xd1, 0x24, 0xdf, 0x6c, 0x94, 0xbb, 0x1e, 0x6e, 0x3c, 0x61, 0x0c, 0xba, 0x6e, 0xdf,
	0x19, 0x65, 0x0b, 0x29, 0xc2, 0xff, 0xed, 0x06, 0x53, 0xe0, 0xa0, 0x43, 0x8b, 0x40, 0x0c, 0x1e,
	0x50, 0x6e, 0x5f, 0x26, 0x5f, 0x1a, 0xa9, 0x32, 0x42, 0xa9, 0x62, 0x5a, 0x87, 0x37, 0x9c, 0x10,
	0x6e, 0x51, 0xa3, 0x9a, 0x81, 0x27, 0xe0, 0x8e, 0x2b, 0x67, 0x34, 0xd3, 0x32, 0x9b, 0x11, 0x15,
	0x76, 0xdc, 0xb3, 0x20, 0x9b, 0xfe, 0x47, 0x15, 0x3f, 0x2b, 0xb8, 0xf9, 0xb8, 0xcc, 0xd1, 0x54,
	0x0a, 0x3c, 0x95, 0x5a, 0x48, 0xed, 0x7f, 0xf6, 0x35, 0xfd, 0x84, 0xcd, 0xaa, 0x64, 0x1a, 0x1d,
	0x2d, 0x4c, 0xba, 0xe7, 0x5d, 0x26, 0xf2, 0x2d, 0x51, 0xe3, 0x77, 0x67, 0xeb, 0x28, 0x38, 0x5f,
	0x47, 0xc1, 0xaf, 0x75, 0x14, 0x9c, 0x6e, 0xa2, 0xd6, 0xf9, 0x26, 0x6a, 0x7d, 0xdf, 0x44, 0xad,
	0x0f, 0x68, 0xcb, 0x6f, 0xe2, 0x2e, 0x68, 0xff, 0x98, 0xe4, 0x1a, 0xfb, 0x53, 0xff, 0xfc, 0x0a,
	0x7f, 0xf1, 0xf7, 0xee, 0xbc, 0xf3, 0x8e, 0x5b, 0xec, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd5, 0x53, 0xc7, 0xfd, 0x0c, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Airdrops) > 0 {
		for iNdEx := len(m.Airdrops) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Airdrops[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Airdrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Airdrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Airdrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClaimedSoFar.Size()
		i -= size
		if _, err := m.ClaimedSoFar.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.DistributorAddress) > 0 {
		i -= len(m.DistributorAddress)
		copy(dAtA[i:], m.DistributorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DistributorAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x22
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.AirdropDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.AirdropDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.AirdropIdentifier) > 0 {
		i -= len(m.AirdropIdentifier)
		copy(dAtA[i:], m.AirdropIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AirdropIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Airdrops) > 0 {
		for _, e := range m.Airdrops {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *Airdrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AirdropIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.AirdropDuration)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DistributorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ClaimedSoFar.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Airdrops", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Airdrops = append(m.Airdrops, &Airdrop{})
			if err := m.Airdrops[len(m.Airdrops)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Airdrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Airdrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Airdrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AirdropIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.AirdropDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedSoFar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimedSoFar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
