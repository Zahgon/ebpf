package btf

import (
	"encoding/binary"
	"hash/maphash"
	"iter"
	"sync"
)

type sharedBuf struct {
	raw []byte
}

type decoder struct {
	base      *decoder
	byteOrder binary.ByteOrder
	*sharedBuf
	strings *stringTable

	firstTypeID TypeID

	offsets  []int
	declTags map[TypeID][]TypeID

	namedTypes *fuzzyStringIndex

	mu              sync.Mutex
	types           map[TypeID]Type
	typeIDs         map[Type]TypeID
	legacyBitfields map[TypeID][2]Bits
}

func newDecoder(raw []byte, bo binary.ByteOrder, strings *stringTable, base *decoder) (*decoder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func allBtfTypeOffsets(buf []byte, bo binary.ByteOrder, header *btfType) iter.Seq2[int, error] {
	_ = "STUB: not implemented"
	return nil
}

func (d *decoder) Copy() *decoder { _ = "STUB: not implemented"; return nil }

func (d *decoder) copy(copiedTypes map[Type]Type) *decoder { _ = "STUB: not implemented"; return nil }

func (d *decoder) TypeID(typ Type) (TypeID, error) {
	_ = "STUB: not implemented"
	return *new(TypeID), nil
}

func (d *decoder) TypesByName(name essentialName) ([]Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *decoder) TypeByID(id TypeID) (Type, error) {
	_ = "STUB: not implemented"
	return *new(Type), nil
}

func (d *decoder) inflateType(id TypeID) (typ Type, err error) {
	_ = "STUB: not implemented"
	return *new(Type), nil
}

type fuzzyStringIndex struct {
	seed    maphash.Seed
	entries []fuzzyStringIndexEntry
}

func newFuzzyStringIndex(capacity int) *fuzzyStringIndex { _ = "STUB: not implemented"; return nil }

func (idx *fuzzyStringIndex) Add(name []byte, id TypeID) { _ = "STUB: not implemented"; return }

func (idx *fuzzyStringIndex) Build() { _ = "STUB: not implemented"; return }

func (idx *fuzzyStringIndex) Find(name string) iter.Seq[TypeID] {
	_ = "STUB: not implemented"
	return nil
}

type fuzzyStringIndexEntry uint64

func newFuzzyStringIndexEntry(hash uint32, id TypeID) fuzzyStringIndexEntry {
	_ = "STUB: not implemented"
	return *new(fuzzyStringIndexEntry)
}

func (e fuzzyStringIndexEntry) hash() uint32 { _ = "STUB: not implemented"; return 0 }

func (e fuzzyStringIndexEntry) id() TypeID { _ = "STUB: not implemented"; return *new(TypeID) }
