package ebpf

import (
	"debug/elf"
	"encoding/binary"
	"errors"
	"io"
	"iter"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

type kconfigMetaKey struct{}

type kconfigMeta struct {
	Map    *MapSpec
	Offset uint32
}

type kfuncMetaKey struct{}

type kfuncMeta struct {
	Binding elf.SymBind
	Func    *btf.Func
}

type ksymMetaKey struct{}

type ksymMeta struct {
	Binding elf.SymBind
	Name    string
	Var     *btf.Var
}

type elfCode struct {
	*internal.SafeELFFile
	sections map[elf.SectionIndex]*elfSection
	license  string
	version  uint32
	btf      *btf.Spec
	extInfo  *btf.ExtInfos
	maps     map[string]*MapSpec
	vars     map[string]*VariableSpec
	kfuncs   map[string]*btf.Func
	ksyms    map[string]*btf.Var
	kconfig  *MapSpec
}

func LoadCollectionSpec(file string) (*CollectionSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadCollectionSpecFromReader(rd io.ReaderAt) (*CollectionSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadLicense(sec *elf.Section) (string, error) { _ = "STUB: not implemented"; return "", nil }

func loadVersion(sec *elf.Section, bo binary.ByteOrder) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func isDataSection(name string) bool { _ = "STUB: not implemented"; return false }

func isConstantDataSection(name string) bool { _ = "STUB: not implemented"; return false }

func isKconfigSection(name string) bool { _ = "STUB: not implemented"; return false }

type elfSectionKind int

const (
	undefSection elfSectionKind = iota
	mapSection
	btfMapSection
	programSection
	dataSection
	structOpsSection
)

type elfSection struct {
	*elf.Section
	kind elfSectionKind

	symbols map[uint64]elf.Symbol

	relocations map[uint64]elf.Symbol

	references int
}

func newElfSection(section *elf.Section, kind elfSectionKind) *elfSection {
	_ = "STUB: not implemented"
	return nil
}

func (es *elfSection) symbolsSorted() iter.Seq2[uint64, elf.Symbol] {
	_ = "STUB: not implemented"
	return nil
}

func (ec *elfCode) assignSymbols(symbols []elf.Symbol) { _ = "STUB: not implemented"; return }

func (ec *elfCode) loadRelocations(relSections map[elf.SectionIndex]*elf.Section, symbols []elf.Symbol) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *elfCode) loadProgramSections() (map[string]*ProgramSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ec *elfCode) loadFunctions(sec *elfSection) (map[string]asm.Instructions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func take[T any](q *[]T, f func(T) bool) *T { _ = "STUB: not implemented"; return nil }

func assignMetadata(ins *asm.Instruction, raw asm.RawInstructionOffset,
	fo *btf.FuncOffsets, lo *btf.LineOffsets, ro *btf.CORERelocationOffsets) {
	_ = "STUB: not implemented"
	return
}

func referenceRelativeJump(ins *asm.Instruction, offset uint64, symbols map[uint64]elf.Symbol) error {
	_ = "STUB: not implemented"
	return nil
}

func jumpTarget(offset uint64, ins asm.Instruction) uint64 { _ = "STUB: not implemented"; return 0 }

var errUnsupportedBinding = errors.New("unsupported binding")

func (ec *elfCode) relocateInstruction(ins *asm.Instruction, rel elf.Symbol) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *elfCode) loadMaps() error { _ = "STUB: not implemented"; return nil }

func (ec *elfCode) sectionVars(spec *btf.Spec, sec string) (map[string]*btf.Var, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ec *elfCode) loadBTFMaps() error { _ = "STUB: not implemented"; return nil }

func mapSpecFromBTF(es *elfSection, sym elf.Symbol, v *btf.Var, def *btf.Struct, spec *btf.Spec, name string, inner bool) (*MapSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func uintFromBTF(typ btf.Type) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func resolveBTFArrayMacro(typ btf.Type) (btf.Type, error) {
	_ = "STUB: not implemented"
	return *new(btf.Type), nil
}

func valuesRelocations(es *elfSection, sym elf.Symbol, member btf.Member) iter.Seq2[uint32, elf.Symbol] {
	_ = "STUB: not implemented"
	return nil
}

func resolveBTFValuesContents(es *elfSection, sym elf.Symbol, member btf.Member) ([]MapKV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ec *elfCode) loadDataSections() error { _ = "STUB: not implemented"; return nil }

func (ec *elfCode) loadKconfigSection() error { _ = "STUB: not implemented"; return nil }

func (ec *elfCode) loadKsymsSection() error { _ = "STUB: not implemented"; return nil }

func (ec *elfCode) associateStructOpsRelocs(progs map[string]*ProgramSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *elfCode) createStructOpsMap(vsi btf.VarSecinfo, userData []byte, flags uint32) (*btf.Struct, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

type libbpfElfSectionDef struct {
	pattern     string
	programType sys.ProgType
	attachType  sys.AttachType
	flags       libbpfElfSectionFlag
}

type libbpfElfSectionFlag uint32

const (
	_SEC_NONE libbpfElfSectionFlag = 0

	_SEC_EXP_ATTACH_OPT libbpfElfSectionFlag = 1 << (iota - 1)
	_SEC_ATTACHABLE
	_SEC_ATTACH_BTF
	_SEC_SLEEPABLE
	_SEC_XDP_FRAGS
	_SEC_USDT

	_SEC_ATTACHABLE_OPT = _SEC_ATTACHABLE | _SEC_EXP_ATTACH_OPT
)

func getProgType(sectionName string) (ProgramType, AttachType, uint32, string) {
	_ = "STUB: not implemented"
	return *new(ProgramType), *new(AttachType), 0, ""
}

func matchSectionName(sectionName, pattern string) (extra string, found bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ec *elfCode) loadSectionRelocations(sec *elf.Section, symbols []elf.Symbol) (map[uint64]elf.Symbol, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
