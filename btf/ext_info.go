package btf

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
)

type ExtInfos struct {
	Funcs     map[string]FuncOffsets
	Lines     map[string]LineOffsets
	CORERelos map[string]CORERelocationOffsets
}

func (ei *ExtInfos) Section(name string) (FuncOffsets, LineOffsets, CORERelocationOffsets) {
	_ = "STUB: not implemented"
	return *new(FuncOffsets), *new(LineOffsets), *new(CORERelocationOffsets)
}

func loadExtInfosFromELF(file *internal.SafeELFFile, spec *Spec) (*ExtInfos, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadExtInfos(r io.ReaderAt, bo binary.ByteOrder, spec *Spec) (*ExtInfos, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalExtInfos(insns asm.Instructions, b *Builder) (funcInfos, lineInfos []byte, _ error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type btfExtHeader struct {
	Magic   uint16
	Version uint8
	Flags   uint8

	HdrLen uint32

	FuncInfoOff uint32
	FuncInfoLen uint32
	LineInfoOff uint32
	LineInfoLen uint32
}

func parseBTFExtHeader(r io.Reader, bo binary.ByteOrder) (*btfExtHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *btfExtHeader) funcInfoStart() int64 { _ = "STUB: not implemented"; return 0 }

func (h *btfExtHeader) lineInfoStart() int64 { _ = "STUB: not implemented"; return 0 }

func (h *btfExtHeader) coreReloStart(ch *btfExtCOREHeader) int64 {
	_ = "STUB: not implemented"
	return 0
}

type btfExtCOREHeader struct {
	COREReloOff uint32
	COREReloLen uint32
}

func parseBTFExtCOREHeader(r io.Reader, bo binary.ByteOrder, extHeader *btfExtHeader) (*btfExtCOREHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type btfExtInfoSec struct {
	SecNameOff uint32
	NumInfo    uint32
}

func parseExtInfoSec(r io.Reader, bo binary.ByteOrder, strings *stringTable) (string, *btfExtInfoSec, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func parseExtInfoRecordSize(r io.Reader, bo binary.ByteOrder) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type FuncOffsets = []FuncOffset

var FuncInfoSize = uint32(binary.Size(bpfFuncInfo{}))

type FuncOffset struct {
	Offset asm.RawInstructionOffset
	Func   *Func
}

type bpfFuncInfo struct {
	InsnOff uint32
	TypeID  TypeID
}

func newFuncOffset(fi bpfFuncInfo, spec *Spec) (*FuncOffset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newFuncOffsets(bfis []bpfFuncInfo, spec *Spec) (FuncOffsets, error) {
	_ = "STUB: not implemented"
	return *new(FuncOffsets), nil
}

func LoadFuncInfos(reader io.Reader, bo binary.ByteOrder, recordNum uint32, spec *Spec) (FuncOffsets, error) {
	_ = "STUB: not implemented"
	return *new(FuncOffsets), nil
}

func (fi *FuncOffset) marshal(w *bytes.Buffer, b *Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func parseFuncInfos(r io.Reader, bo binary.ByteOrder, strings *stringTable) (map[string][]bpfFuncInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFuncInfoRecords(r io.Reader, bo binary.ByteOrder, recordSize uint32, recordNum uint32, offsetInBytes bool) ([]bpfFuncInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var LineInfoSize = uint32(binary.Size(bpfLineInfo{}))

type Line struct {
	fileName   string
	line       string
	lineNumber uint32
	lineColumn uint32
}

func (li *Line) FileName() string { _ = "STUB: not implemented"; return "" }

func (li *Line) Line() string { _ = "STUB: not implemented"; return "" }

func (li *Line) LineNumber() uint32 { _ = "STUB: not implemented"; return 0 }

func (li *Line) LineColumn() uint32 { _ = "STUB: not implemented"; return 0 }

func (li *Line) String() string { _ = "STUB: not implemented"; return "" }

type LineOffsets = []LineOffset

type LineOffset struct {
	Offset asm.RawInstructionOffset
	Line   *Line
}

const (
	bpfLineShift = 10
	bpfLineMax   = (1 << (32 - bpfLineShift)) - 1
	bpfColumnMax = (1 << bpfLineShift) - 1
)

type bpfLineInfo struct {
	InsnOff     uint32
	FileNameOff uint32
	LineOff     uint32
	LineCol     uint32
}

func LoadLineInfos(reader io.Reader, bo binary.ByteOrder, recordNum uint32, spec *Spec) (LineOffsets, error) {
	_ = "STUB: not implemented"
	return *new(LineOffsets), nil
}

func newLineInfo(li bpfLineInfo, strings *stringTable) (LineOffset, error) {
	_ = "STUB: not implemented"
	return *new(LineOffset), nil
}

func newLineInfos(blis []bpfLineInfo, strings *stringTable) (LineOffsets, error) {
	_ = "STUB: not implemented"
	return *new(LineOffsets), nil
}

func (li *LineOffset) marshal(w *bytes.Buffer, b *Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func parseLineInfos(r io.Reader, bo binary.ByteOrder, strings *stringTable) (map[string][]bpfLineInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseLineInfoRecords(r io.Reader, bo binary.ByteOrder, recordSize uint32, recordNum uint32, offsetInBytes bool) ([]bpfLineInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bpfCORERelo struct {
	InsnOff      uint32
	TypeID       TypeID
	AccessStrOff uint32
	Kind         coreKind
}

type CORERelocation struct {
	typ      Type
	accessor coreAccessor
	kind     coreKind

	id TypeID
}

func (cr *CORERelocation) String() string { _ = "STUB: not implemented"; return "" }

type coreRelocationMeta struct{}

func CORERelocationMetadata(ins *asm.Instruction) *CORERelocation {
	_ = "STUB: not implemented"
	return nil
}

func WithCORERelocationMetadata(ins asm.Instruction, relo *CORERelocation) asm.Instruction {
	_ = "STUB: not implemented"
	return *new(asm.Instruction)
}

type CORERelocationOffsets = []CORERelocationOffset

type CORERelocationOffset struct {
	Relo   *CORERelocation
	Offset asm.RawInstructionOffset
}

func newRelocationInfo(relo bpfCORERelo, spec *Spec, strings *stringTable) (*CORERelocationOffset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newRelocationInfos(brs []bpfCORERelo, spec *Spec, strings *stringTable) (CORERelocationOffsets, error) {
	_ = "STUB: not implemented"
	return *new(CORERelocationOffsets), nil
}

var extInfoReloSize = binary.Size(bpfCORERelo{})

func parseCORERelos(r io.Reader, bo binary.ByteOrder, strings *stringTable) (map[string][]bpfCORERelo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseCOREReloRecords(r io.Reader, bo binary.ByteOrder, recordNum uint32) ([]bpfCORERelo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
