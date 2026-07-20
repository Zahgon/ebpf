package btf

import (
	"encoding/binary"
	"unsafe"
)

//go:generate go tool stringer -linecomment -output=btf_types_string.go -type=FuncLinkage,VarLinkage,btfKind

type btfKind uint8

const (
	kindUnknown btfKind = iota
	kindInt
	kindPointer
	kindArray
	kindStruct
	kindUnion
	kindEnum
	kindForward
	kindTypedef
	kindVolatile
	kindConst
	kindRestrict

	kindFunc
	kindFuncProto

	kindVar
	kindDatasec

	kindFloat

	kindDeclTag

	kindTypeTag

	kindEnum64
)

type FuncLinkage int

const (
	StaticFunc FuncLinkage = iota
	GlobalFunc
	ExternFunc
)

type VarLinkage int

const (
	StaticVar VarLinkage = iota
	GlobalVar
	ExternVar
)

const (
	btfTypeKindShift     = 24
	btfTypeKindLen       = 5
	btfTypeVlenShift     = 0
	btfTypeVlenMask      = 16
	btfTypeKindFlagShift = 31
	btfTypeKindFlagMask  = 1
)

var btfHeaderLen = binary.Size(&btfHeader{})

type btfHeader struct {
	Magic   uint16
	Version uint8
	Flags   uint8
	HdrLen  uint32

	TypeOff   uint32
	TypeLen   uint32
	StringOff uint32
	StringLen uint32
}

type btfLayout struct {
	Off uint32
	Len uint32
}

func parseBTFHeader(buf []byte) (*btfHeader, *btfLayout, binary.ByteOrder, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(binary.ByteOrder), nil
}

type btfType struct {
	NameOff uint32

	Info uint32

	SizeType uint32
}

var btfTypeSize = int(unsafe.Sizeof(btfType{}))

func unmarshalBtfType(bt *btfType, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func mask(len uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func readBits(value, len, shift uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func writeBits(value, len, shift, new uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func (bt *btfType) info(len, shift uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func (bt *btfType) setInfo(value, len, shift uint32) { _ = "STUB: not implemented"; return }

func (bt *btfType) Kind() btfKind { _ = "STUB: not implemented"; return *new(btfKind) }

func (bt *btfType) SetKind(kind btfKind) { _ = "STUB: not implemented"; return }

func (bt *btfType) Vlen() int { _ = "STUB: not implemented"; return 0 }

func (bt *btfType) SetVlen(vlen int) { _ = "STUB: not implemented"; return }

func (bt *btfType) kindFlagBool() bool { _ = "STUB: not implemented"; return false }

func (bt *btfType) setKindFlagBool(set bool) { _ = "STUB: not implemented"; return }

func (bt *btfType) Bitfield() bool { _ = "STUB: not implemented"; return false }

func (bt *btfType) SetBitfield(isBitfield bool) { _ = "STUB: not implemented"; return }

func (bt *btfType) FwdKind() FwdKind { _ = "STUB: not implemented"; return *new(FwdKind) }

func (bt *btfType) SetFwdKind(kind FwdKind) { _ = "STUB: not implemented"; return }

func (bt *btfType) Signed() bool { _ = "STUB: not implemented"; return false }

func (bt *btfType) SetSigned(signed bool) { _ = "STUB: not implemented"; return }

func (bt *btfType) Linkage() FuncLinkage { _ = "STUB: not implemented"; return *new(FuncLinkage) }

func (bt *btfType) SetLinkage(linkage FuncLinkage) { _ = "STUB: not implemented"; return }

func (bt *btfType) Type() TypeID { _ = "STUB: not implemented"; return *new(TypeID) }

func (bt *btfType) SetType(id TypeID) { _ = "STUB: not implemented"; return }

func (bt *btfType) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (bt *btfType) SetSize(size uint32) { _ = "STUB: not implemented"; return }

func (bt *btfType) Encode(buf []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bt *btfType) DataLen() (int, error) { _ = "STUB: not implemented"; return 0, nil }

type btfInt struct {
	Raw uint32
}

const (
	btfIntEncodingLen   = 4
	btfIntEncodingShift = 24
	btfIntOffsetLen     = 8
	btfIntOffsetShift   = 16
	btfIntBitsLen       = 8
	btfIntBitsShift     = 0
)

var btfIntLen = int(unsafe.Sizeof(btfInt{}))

func unmarshalBtfInt(bi *btfInt, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bi btfInt) Encoding() IntEncoding { _ = "STUB: not implemented"; return *new(IntEncoding) }

func (bi *btfInt) SetEncoding(e IntEncoding) { _ = "STUB: not implemented"; return }

func (bi btfInt) Offset() Bits { _ = "STUB: not implemented"; return *new(Bits) }

func (bi *btfInt) SetOffset(offset uint32) { _ = "STUB: not implemented"; return }

func (bi btfInt) Bits() Bits { _ = "STUB: not implemented"; return *new(Bits) }

func (bi *btfInt) SetBits(bits byte) { _ = "STUB: not implemented"; return }

type btfArray struct {
	Type      TypeID
	IndexType TypeID
	Nelems    uint32
}

var btfArrayLen = int(unsafe.Sizeof(btfArray{}))

func unmarshalBtfArray(ba *btfArray, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfMember struct {
	NameOff uint32
	Type    TypeID
	Offset  uint32
}

var btfMemberLen = int(unsafe.Sizeof(btfMember{}))

func unmarshalBtfMember(bm *btfMember, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfVarSecinfo struct {
	Type   TypeID
	Offset uint32
	Size   uint32
}

var btfVarSecinfoLen = int(unsafe.Sizeof(btfVarSecinfo{}))

func unmarshalBtfVarSecInfo(bvsi *btfVarSecinfo, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfVariable struct {
	Linkage uint32
}

var btfVariableLen = int(unsafe.Sizeof(btfVariable{}))

func unmarshalBtfVariable(bv *btfVariable, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfEnum struct {
	NameOff uint32
	Val     uint32
}

var btfEnumLen = int(unsafe.Sizeof(btfEnum{}))

func unmarshalBtfEnum(be *btfEnum, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfEnum64 struct {
	NameOff uint32
	ValLo32 uint32
	ValHi32 uint32
}

var btfEnum64Len = int(unsafe.Sizeof(btfEnum64{}))

func unmarshalBtfEnum64(enum *btfEnum64, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfParam struct {
	NameOff uint32
	Type    TypeID
}

var btfParamLen = int(unsafe.Sizeof(btfParam{}))

func unmarshalBtfParam(param *btfParam, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type btfDeclTag struct {
	ComponentIdx uint32
}

var btfDeclTagLen = int(unsafe.Sizeof(btfDeclTag{}))

func unmarshalBtfDeclTag(bdt *btfDeclTag, b []byte, bo binary.ByteOrder) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
