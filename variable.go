package ebpf

import (
	"github.com/cilium/ebpf/btf"
)

type VariableSpec struct {
	Name string

	SectionName string

	Offset uint32

	Value []byte

	Type *btf.Var
}

func (s *VariableSpec) Set(in any) error { _ = "STUB: not implemented"; return nil }

func (s *VariableSpec) Get(out any) error { _ = "STUB: not implemented"; return nil }

func (s *VariableSpec) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (s *VariableSpec) Constant() bool { _ = "STUB: not implemented"; return false }

func (s *VariableSpec) String() string { _ = "STUB: not implemented"; return "" }

func (s *VariableSpec) Copy() *VariableSpec { _ = "STUB: not implemented"; return nil }

type Variable struct {
	name   string
	offset uint32
	size   uint32
	t      *btf.Var

	mm *Memory
}

func newVariable(name string, offset, size uint32, t *btf.Var, mm *Memory) (*Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Variable) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (v *Variable) ReadOnly() bool { _ = "STUB: not implemented"; return false }

func (v *Variable) Type() *btf.Var { _ = "STUB: not implemented"; return nil }

func (v *Variable) String() string { _ = "STUB: not implemented"; return "" }

func (v *Variable) Set(in any) error { _ = "STUB: not implemented"; return nil }

func (v *Variable) Get(out any) error { _ = "STUB: not implemented"; return nil }

func checkVariable[T any](v *Variable) error { _ = "STUB: not implemented"; return nil }

func VariablePointer[T comparable](v *Variable) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
