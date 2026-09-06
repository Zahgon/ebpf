package btf

import (
	"encoding/binary"
	"errors"

	"github.com/cilium/ebpf/asm"
)

const COREBadRelocationSentinel = 0xbad2310

type COREFixup struct {
	kind   coreKind
	local  uint64
	target uint64

	poison bool

	skipLocalValidation bool
}

func (f *COREFixup) equal(other COREFixup) bool { _ = "STUB: not implemented"; return false }

func (f *COREFixup) String() string { _ = "STUB: not implemented"; return "" }

func (f *COREFixup) Apply(ins *asm.Instruction) error { _ = "STUB: not implemented"; return nil }

func (f COREFixup) isNonExistant() bool { _ = "STUB: not implemented"; return false }

type coreKind uint32

const (
	reloFieldByteOffset coreKind = iota
	reloFieldByteSize
	reloFieldExists
	reloFieldSigned
	reloFieldLShiftU64
	reloFieldRShiftU64
	reloTypeIDLocal
	reloTypeIDTarget
	reloTypeExists
	reloTypeSize
	reloEnumvalExists
	reloEnumvalValue
	reloTypeMatches
)

func (k coreKind) checksForExistence() bool { _ = "STUB: not implemented"; return false }

func (k coreKind) String() string { _ = "STUB: not implemented"; return "" }

func CORERelocate(relos []*CORERelocation, targets []*Spec, bo binary.ByteOrder, resolveLocalTypeID func(Type) (TypeID, error)) ([]COREFixup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errAmbiguousRelocation = errors.New("ambiguous relocation")
var errImpossibleRelocation = errors.New("impossible relocation")
var errIncompatibleTypes = errors.New("incompatible types")

func coreCalculateFixups(relos []*CORERelocation, targets []Type, bo binary.ByteOrder, resolveTargetTypeID func(Type) (TypeID, error)) ([]COREFixup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errNoSignedness = errors.New("no signedness")

func coreCalculateFixup(relo *CORERelocation, target Type, bo binary.ByteOrder, resolveTargetTypeID func(Type) (TypeID, error)) (COREFixup, error) {
	_ = "STUB: not implemented"
	return *new(COREFixup), nil
}

func boolToUint64(val bool) uint64 { _ = "STUB: not implemented"; return 0 }

type coreAccessor []int

func parseCOREAccessor(accessor string) (coreAccessor, error) {
	_ = "STUB: not implemented"
	return *new(coreAccessor), nil
}

func (ca coreAccessor) String() string { _ = "STUB: not implemented"; return "" }

func (ca coreAccessor) enumValue(t Type) (*EnumValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type coreField struct {
	Type Type

	offset uint32

	bitfieldOffset Bits

	bitfieldSize Bits
}

func (cf *coreField) adjustOffsetToNthElement(n int) error { _ = "STUB: not implemented"; return nil }

func (cf *coreField) adjustOffsetBits(offset Bits) error { _ = "STUB: not implemented"; return nil }

func (cf *coreField) sizeBits() (Bits, error) { _ = "STUB: not implemented"; return *new(Bits), nil }

func coreFindField(localT Type, localAcc coreAccessor, targetT Type) (coreField, coreField, error) {
	_ = "STUB: not implemented"
	return *new(coreField), *new(coreField), nil
}

func coreFindMember(typ composite, name string) (Member, bool, error) {
	_ = "STUB: not implemented"
	return *new(Member), false, nil
}

func coreFindEnumValue(local Type, localAcc coreAccessor, target Type) (localValue, targetValue *EnumValue, _ error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CheckTypeCompatibility(localType Type, targetType Type) error {
	_ = "STUB: not implemented"
	return nil
}

type pair struct {
	A, B Type
}

func coreAreTypesCompatible(localType Type, targetType Type, visited map[pair]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func coreAreMembersCompatible(localType Type, targetType Type) error {
	_ = "STUB: not implemented"
	return nil
}

func coreEssentialNamesMatch(a, b string) bool { _ = "STUB: not implemented"; return false }

func coreTypesMatch(localType Type, targetType Type, visited map[pair]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func coreEncodingMatches(local, target *Int) bool { _ = "STUB: not implemented"; return false }

func coreEnumsMatch(local *Enum, target *Enum) error { _ = "STUB: not implemented"; return nil }
