// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polygon/envoy/v1/types.proto

package envoy

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters of the module.
type Params struct {
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0810ae754dabb2, []int{0}
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

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Locks  []Lock `protobuf:"bytes,3,rep,name=locks,proto3" json:"locks"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0810ae754dabb2, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLocks() []Lock {
	if m != nil {
		return m.Locks
	}
	return nil
}

// the id is a named action that the lock owner will do within
// `expires` number of blocks from `block`, otherwise the lock is
// relinquished
type Lock struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Envoy     string `protobuf:"bytes,2,opt,name=envoy,proto3" json:"envoy,omitempty"`
	AtBlock   uint32 `protobuf:"varint,3,opt,name=at_block,json=atBlock,proto3" json:"at_block,omitempty"`
	NumBlocks uint32 `protobuf:"varint,4,opt,name=num_blocks,json=numBlocks,proto3" json:"num_blocks,omitempty"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0810ae754dabb2, []int{2}
}
func (m *Lock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return m.Size()
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Lock) GetEnvoy() string {
	if m != nil {
		return m.Envoy
	}
	return ""
}

func (m *Lock) GetAtBlock() uint32 {
	if m != nil {
		return m.AtBlock
	}
	return 0
}

func (m *Lock) GetNumBlocks() uint32 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "polygon.envoy.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "polygon.envoy.v1.GenesisState")
	proto.RegisterType((*Lock)(nil), "polygon.envoy.v1.Lock")
}

func init() { proto.RegisterFile("polygon/envoy/v1/types.proto", fileDescriptor_0f0810ae754dabb2) }

var fileDescriptor_0f0810ae754dabb2 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x4e, 0x02, 0x31,
	0x1c, 0xc6, 0xaf, 0x72, 0x20, 0x14, 0x4d, 0x4c, 0x43, 0x4c, 0x21, 0x7a, 0x12, 0x26, 0x16, 0x7b,
	0xe1, 0x4c, 0xdc, 0xbd, 0xc5, 0xc5, 0xc1, 0x1c, 0x9b, 0x0b, 0x29, 0x47, 0x73, 0x12, 0x68, 0xff,
	0x97, 0x6b, 0x21, 0xc1, 0xd9, 0x07, 0xf0, 0x61, 0x7c, 0x08, 0x46, 0xe2, 0xe4, 0x64, 0x0c, 0xbc,
	0x88, 0xb9, 0xf6, 0x16, 0x75, 0xfb, 0xff, 0xbf, 0xdf, 0xf7, 0xf5, 0x4b, 0x5b, 0x7c, 0x91, 0xc3,
	0x72, 0x93, 0x81, 0x0a, 0x85, 0x5a, 0xc3, 0x26, 0x5c, 0x8f, 0x42, 0xb3, 0xc9, 0x85, 0x66, 0x79,
	0x01, 0x06, 0xc8, 0x59, 0x45, 0x99, 0xa5, 0x6c, 0x3d, 0xea, 0x75, 0x53, 0xd0, 0x12, 0xf4, 0xc4,
	0xf2, 0xd0, 0x2d, 0xce, 0xdc, 0xeb, 0x64, 0x90, 0x81, 0xd3, 0xcb, 0xc9, 0xa9, 0x83, 0x26, 0x6e,
	0x3c, 0xf2, 0x82, 0x4b, 0x3d, 0x78, 0xc1, 0x27, 0xf7, 0x42, 0x09, 0x3d, 0xd7, 0x63, 0xc3, 0x8d,
	0x20, 0xb7, 0xb8, 0x91, 0x5b, 0x42, 0x51, 0x1f, 0x0d, 0xdb, 0x11, 0x65, 0x7f, 0xdb, 0x98, 0x4b,
	0xc6, 0xfe, 0xf6, 0xeb, 0xca, 0x4b, 0x2a, 0x37, 0x89, 0x70, 0x7d, 0x09, 0xe9, 0x42, 0xd3, 0x5a,
	0xbf, 0x36, 0x6c, 0x47, 0xe7, 0xff, 0x63, 0x0f, 0x90, 0x2e, 0xaa, 0x90, 0xb3, 0x0e, 0x5e, 0x11,
	0xf6, 0x4b, 0x95, 0x10, 0xec, 0x2b, 0x2e, 0x85, 0xad, 0x6c, 0x25, 0x76, 0x26, 0x0c, 0xd7, 0x6d,
	0x94, 0x1e, 0x95, 0x62, 0x4c, 0x3f, 0xde, 0xaf, 0x3b, 0xd5, 0xcd, 0xee, 0x66, 0xb3, 0x42, 0x68,
	0x3d, 0x36, 0xc5, 0x5c, 0x65, 0x89, 0xb3, 0x91, 0x2e, 0x6e, 0x72, 0x33, 0x99, 0x96, 0x27, 0xd3,
	0x5a, 0x1f, 0x0d, 0x4f, 0x93, 0x63, 0x6e, 0xe2, 0x72, 0x25, 0x97, 0x18, 0xab, 0x95, 0x74, 0x4c,
	0x53, 0xdf, 0xc2, 0x96, 0x5a, 0x49, 0x4b, 0x75, 0x1c, 0x6d, 0xf7, 0x01, 0xda, 0xed, 0x03, 0xf4,
	0xbd, 0x0f, 0xd0, 0xdb, 0x21, 0xf0, 0x76, 0x87, 0xc0, 0xfb, 0x3c, 0x04, 0xde, 0x13, 0xcd, 0xe6,
	0xe6, 0x79, 0x35, 0x65, 0x29, 0xc8, 0xf0, 0xd7, 0x97, 0x4c, 0x1b, 0xf6, 0x1d, 0x6f, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x6f, 0x6b, 0x93, 0xaa, 0x01, 0x00, 0x00,
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
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Locks) > 0 {
		for iNdEx := len(m.Locks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Locks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Lock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumBlocks != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NumBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.AtBlock != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AtBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Envoy) > 0 {
		i -= len(m.Envoy)
		copy(dAtA[i:], m.Envoy)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Envoy)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
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
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Locks) > 0 {
		for _, e := range m.Locks {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Lock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Envoy)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.AtBlock != 0 {
		n += 1 + sovTypes(uint64(m.AtBlock))
	}
	if m.NumBlocks != 0 {
		n += 1 + sovTypes(uint64(m.NumBlocks))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locks = append(m.Locks, Lock{})
			if err := m.Locks[len(m.Locks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Lock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Lock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envoy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Envoy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AtBlock", wireType)
			}
			m.AtBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AtBlock |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			m.NumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)