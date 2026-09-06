package btf

import (
	"encoding/binary"
	"sync"

	"github.com/cilium/ebpf/internal"
)

type MarshalOptions struct {
	Order binary.ByteOrder

	StripFuncLinkage bool

	ReplaceDeclTags bool

	ReplaceTypeTags bool

	ReplaceEnum64 bool

	PreventNoTypeFound bool
}

func KernelMarshalOptions() *MarshalOptions { _ = "STUB: not implemented"; return nil }

type encoder struct {
	MarshalOptions

	pending internal.Deque[Type]
	strings *stringTableBuilder
	ids     map[Type]TypeID
	visited map[Type]struct{}
	lastID  TypeID
}

var bufferPool = sync.Pool{
	New: func() any {
		buf := make([]byte, btfHeaderLen+128)
		return &buf
	},
}

func getByteSlice() *[]byte { _ = "STUB: not implemented"; return nil }

func putByteSlice(buf *[]byte) { _ = "STUB: not implemented"; return }

type Builder struct {
	types []Type

	stableIDs map[Type]TypeID

	strings *stringTableBuilder

	deduper *deduper
}

type BuilderOptions struct {
	Deduplicate bool
}

func NewBuilder(types []Type, opts *BuilderOptions) (*Builder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) Empty() bool { _ = "STUB: not implemented"; return false }

func (b *Builder) Add(typ Type) (TypeID, error) {
	_ = "STUB: not implemented"
	return *new(TypeID), nil
}

func (b *Builder) Spec() (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Builder) Marshal(buf []byte, opts *MarshalOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Builder) addString(str string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *encoder) allocateIDs(root Type) error { _ = "STUB: not implemented"; return nil }

func (e *encoder) id(typ Type) TypeID { _ = "STUB: not implemented"; return *new(TypeID) }

func (e *encoder) deflatePending(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateType(buf []byte, typ Type) (_ []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateInt(buf []byte, raw *btfType, i *Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateDeclTag(buf []byte, raw *btfType, tag *declTag) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateConst(raw *btfType, c *Const) { _ = "STUB: not implemented"; return }

func (e *encoder) deflateTypeTag(raw *btfType, tag *TypeTag) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *encoder) deflateUnion(buf []byte, raw *btfType, union *Union) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateMembers(buf []byte, header *btfType, members []Member) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateEnum(buf []byte, raw *btfType, enum *Enum) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateEnumValues(buf []byte, enum *Enum) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateEnum64(buf []byte, raw *btfType, enum *Enum) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateEnum64Values(buf []byte, values []EnumValue) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateFuncParams(buf []byte, params []FuncParam) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *encoder) deflateVarSecinfos(buf []byte, vars []VarSecinfo) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalMapKV(key, value Type) (_ *Handle, keyID, valueID TypeID, err error) {
	_ = "STUB: not implemented"
	return nil, *new(TypeID), *new(TypeID), nil
}
