//go:build mips || mips64 || ppc64 || s390x

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	variablesProgConstExample  = "const_example"
	variablesProgGlobalExample = "global_example"
	variablesProgHiddenExample = "hidden_example"
	variablesVarConstU32       = "const_u32"
	variablesVarGlobalU16      = "global_u16"
)

func loadVariables() (*ebpf.CollectionSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func loadVariablesObjects(obj any, opts *ebpf.CollectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type variablesSpecs struct {
	variablesProgramSpecs
	variablesMapSpecs
	variablesVariableSpecs
}

type variablesProgramSpecs struct {
	ConstExample  *ebpf.ProgramSpec `ebpf:"const_example"`
	GlobalExample *ebpf.ProgramSpec `ebpf:"global_example"`
	HiddenExample *ebpf.ProgramSpec `ebpf:"hidden_example"`
}

type variablesMapSpecs struct {
}

type variablesVariableSpecs struct {
	ConstU32  *ebpf.VariableSpec `ebpf:"const_u32"`
	GlobalU16 *ebpf.VariableSpec `ebpf:"global_u16"`
}

type variablesObjects struct {
	variablesPrograms
	variablesMaps
	variablesVariables
}

func (o *variablesObjects) Close() error { _ = "STUB: not implemented"; return nil }

type variablesMaps struct {
}

func (m *variablesMaps) Close() error { _ = "STUB: not implemented"; return nil }

type variablesVariables struct {
	ConstU32  *ebpf.Variable `ebpf:"const_u32"`
	GlobalU16 *ebpf.Variable `ebpf:"global_u16"`
}

type variablesPrograms struct {
	ConstExample  *ebpf.Program `ebpf:"const_example"`
	GlobalExample *ebpf.Program `ebpf:"global_example"`
	HiddenExample *ebpf.Program `ebpf:"hidden_example"`
}

func (p *variablesPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _VariablesClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed variables_bpfeb.o
var _VariablesBytes []byte
