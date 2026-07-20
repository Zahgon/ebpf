package asm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"io"
)

const InstructionSize = 8

type RawInstructionOffset uint64

var ErrUnreferencedSymbol = errors.New("unreferenced symbol")
var ErrUnsatisfiedMapReference = errors.New("unsatisfied map reference")
var ErrUnsatisfiedProgramReference = errors.New("unsatisfied program reference")

func (rio RawInstructionOffset) Bytes() uint64 { _ = "STUB: not implemented"; return 0 }

type Instruction struct {
	OpCode   OpCode
	Dst      Register
	Src      Register
	Offset   int16
	Constant int64

	Metadata Metadata
}

func (ins *Instruction) Width() RawInstructionOffset {
	_ = "STUB: not implemented"
	return *new(RawInstructionOffset)
}

func (ins *Instruction) Unmarshal(r io.Reader, bo binary.ByteOrder, platform string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ins Instruction) Marshal(w io.Writer, bo binary.ByteOrder) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ins *Instruction) AssociateMap(m FDer) error { _ = "STUB: not implemented"; return nil }

func (ins *Instruction) encodeMapFD(fd int) { _ = "STUB: not implemented"; return }

func (ins *Instruction) mapFd() int { _ = "STUB: not implemented"; return 0 }

func (ins *Instruction) RewriteMapOffset(offset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (ins *Instruction) mapOffset() uint32 { _ = "STUB: not implemented"; return 0 }

func (ins *Instruction) AssociateBTFID(id uint32, moduleFD int) error {
	_ = "STUB: not implemented"
	return nil
}

func (ins *Instruction) IsLoadFromMap() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsFunctionCall() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsKfuncCall() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsLoadOfFunctionPointer() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsLoadOfBTFID() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsFunctionReference() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsBuiltinCall() bool { _ = "STUB: not implemented"; return false }

func (ins *Instruction) IsConstantLoad(size Size) bool { _ = "STUB: not implemented"; return false }

func (ins Instruction) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (ins Instruction) equal(other Instruction) bool { _ = "STUB: not implemented"; return false }

func (ins Instruction) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (ins Instruction) WithMetadata(meta Metadata) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

type symbolMeta struct{}

func (ins Instruction) WithSymbol(name string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (ins Instruction) Symbol() string { _ = "STUB: not implemented"; return "" }

type referenceMeta struct{}

func (ins Instruction) WithReference(ref string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (ins Instruction) Reference() string { _ = "STUB: not implemented"; return "" }

type mapMeta struct{}

func (ins Instruction) Map() FDer { _ = "STUB: not implemented"; return *new(FDer) }

type sourceMeta struct{}

func (ins Instruction) WithSource(src fmt.Stringer) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (ins Instruction) Source() fmt.Stringer { _ = "STUB: not implemented"; return *new(fmt.Stringer) }

type Comment string

func (s Comment) String() string { _ = "STUB: not implemented"; return "" }

type FDer interface {
	FD() int
}

type Instructions []Instruction

func AppendInstructions(insns Instructions, r io.Reader, bo binary.ByteOrder, platform string) (Instructions, error) {
	_ = "STUB: not implemented"
	return *new(Instructions), nil
}

func (insns Instructions) Name() string { _ = "STUB: not implemented"; return "" }

func (insns Instructions) String() string { _ = "STUB: not implemented"; return "" }

func (insns Instructions) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (insns Instructions) AssociateMap(symbol string, m FDer) error {
	_ = "STUB: not implemented"
	return nil
}

func (insns Instructions) SymbolOffsets() (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (insns Instructions) FunctionReferences() []string { _ = "STUB: not implemented"; return nil }

func (insns Instructions) ReferenceOffsets() map[string][]int {
	_ = "STUB: not implemented"
	return nil
}

func (insns Instructions) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (insns Instructions) Marshal(w io.Writer, bo binary.ByteOrder) error {
	_ = "STUB: not implemented"
	return nil
}

func (insns Instructions) Tag(bo binary.ByteOrder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (insns Instructions) HasTag(tag string, bo binary.ByteOrder) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (insns Instructions) tagSha1(bo binary.ByteOrder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (insns Instructions) tagSha256(bo binary.ByteOrder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (insns Instructions) hash(h hash.Hash, bo binary.ByteOrder) error {
	_ = "STUB: not implemented"
	return nil
}

func (insns Instructions) encodeFunctionReferences() error { _ = "STUB: not implemented"; return nil }

func (insns Instructions) encodeMapPointers() error { _ = "STUB: not implemented"; return nil }

func (insns Instructions) Iterate() *InstructionIterator { _ = "STUB: not implemented"; return nil }

type InstructionIterator struct {
	insns Instructions

	Ins *Instruction

	Index int

	Offset RawInstructionOffset
}

func (iter *InstructionIterator) Next() bool { _ = "STUB: not implemented"; return false }

type bpfRegisters uint8

func newBPFRegisters(dst, src Register, bo binary.ByteOrder) (bpfRegisters, error) {
	_ = "STUB: not implemented"
	return *new(bpfRegisters), nil
}
