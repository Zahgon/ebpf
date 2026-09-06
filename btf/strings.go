package btf

import (
	"io"
	"sync"
)

type stringTable struct {
	base  *stringTable
	bytes []byte

	mu    sync.Mutex
	cache map[uint32]string
}

type sizedReader interface {
	io.Reader
	Size() int64
}

func readStringTable(r sizedReader, base *stringTable) (*stringTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newStringTable(bytes []byte, base *stringTable) (*stringTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (st *stringTable) Lookup(offset uint32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (st *stringTable) LookupBytes(offset uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (st *stringTable) lookupSlow(offset uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cst *stringTable) LookupCached(offset uint32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type stringTableBuilder struct {
	length  uint32
	strings map[string]uint32
}

func newStringTableBuilder(capacity int) *stringTableBuilder { _ = "STUB: not implemented"; return nil }

func (stb *stringTableBuilder) Add(str string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (stb *stringTableBuilder) append(str string) uint32 { _ = "STUB: not implemented"; return 0 }

func (stb *stringTableBuilder) Lookup(str string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (stb *stringTableBuilder) Length() int { _ = "STUB: not implemented"; return 0 }

func (stb *stringTableBuilder) AppendEncoded(buf []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (stb *stringTableBuilder) Copy() *stringTableBuilder { _ = "STUB: not implemented"; return nil }
