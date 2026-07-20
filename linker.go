package ebpf

import (
	"encoding/binary"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
)

type handles []*btf.Handle

func (hs *handles) add(h *btf.Handle) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (hs handles) fdArray() []int32 { _ = "STUB: not implemented"; return nil }

func (hs *handles) Close() error { _ = "STUB: not implemented"; return nil }

func hasFunctionReferences(insns asm.Instructions) bool { _ = "STUB: not implemented"; return false }

func applyRelocations(insns asm.Instructions, bo binary.ByteOrder, b *btf.Builder, c *btf.Cache, kernelOverride *btf.Spec, extraTargets []*btf.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func flattenPrograms(progs map[string]*ProgramSpec, names []string) {
	_ = "STUB: not implemented"
	return
}

func flattenInstructions(name string, progs map[string]*ProgramSpec, refs map[*ProgramSpec][]string) asm.Instructions {
	_ = "STUB: not implemented"
	return *new(asm.Instructions)
}

func fixupAndValidate(insns asm.Instructions) error { _ = "STUB: not implemented"; return nil }

const kfuncCallPoisonBase = 0xdedc0de

func fixupKfuncs(insns asm.Instructions, cache *btf.Cache) (_ handles, err error) {
	_ = "STUB: not implemented"
	return *new(handles), nil
}

type incompatibleKfuncError struct {
	name string
	err  error
}

func (ike *incompatibleKfuncError) Error() string { _ = "STUB: not implemented"; return "" }

func fixupProbeReadKernel(ins *asm.Instruction) { _ = "STUB: not implemented"; return }

func resolveKconfigReferences(insns asm.Instructions) (_ *Map, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ksymFixup struct {
	*asm.Instruction
	*ksymMeta
}

func resolveKsymReferences(insns asm.Instructions, cache *btf.Cache) (handles, error) {
	_ = "STUB: not implemented"
	return *new(handles), nil
}

func applyUntypedKsymFixups(fixups []ksymFixup) error { _ = "STUB: not implemented"; return nil }

func applyTypedKsymFixups(fixups []ksymFixup, cache *btf.Cache) (modules handles, err error) {
	_ = "STUB: not implemented"
	return *new(handles), nil
}
