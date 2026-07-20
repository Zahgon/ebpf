//go:build (386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64 || wasm) && linux

package test

import (
	_ "embed"
	"io"
	"structs"

	"github.com/cilium/ebpf"
)

type testBar struct {
	_ structs.HostLayout
	A uint64
	B uint32
	_ [4]byte
}

type testBarfoo struct {
	_   structs.HostLayout
	Bar int64
	Baz bool
	_   [3]byte
	Boo testE
}

type testBaz struct {
	_ structs.HostLayout
	A uint64
}

type testE uint32

const (
	testEHOOPY testE = 0
	testEFROOD testE = 1
)

type testUbar struct {
	_ structs.HostLayout
	A uint32
	_ [4]byte
}

const (
	testMapMap1        = "map1"
	testProgFilter     = "filter"
	testVarAnInt       = "an_int"
	testVarIntArray    = "int_array"
	testVarMyConstant  = "my_constant"
	testVarStructArray = "struct_array"
	testVarStructConst = "struct_const"
	testVarStructVar   = "struct_var"
	testVarUnionVar    = "union_var"
)

func loadTest() (*ebpf.CollectionSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func loadTestObjects(obj any, opts *ebpf.CollectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type testSpecs struct {
	testProgramSpecs
	testMapSpecs
	testVariableSpecs
}

type testProgramSpecs struct {
	Filter *ebpf.ProgramSpec `ebpf:"filter"`
}

type testMapSpecs struct {
	Map1 *ebpf.MapSpec `ebpf:"map1"`
}

type testVariableSpecs struct {
	AnInt       *ebpf.VariableSpec `ebpf:"an_int"`
	IntArray    *ebpf.VariableSpec `ebpf:"int_array"`
	MyConstant  *ebpf.VariableSpec `ebpf:"my_constant"`
	StructArray *ebpf.VariableSpec `ebpf:"struct_array"`
	StructConst *ebpf.VariableSpec `ebpf:"struct_const"`
	StructVar   *ebpf.VariableSpec `ebpf:"struct_var"`
	UnionVar    *ebpf.VariableSpec `ebpf:"union_var"`
}

type testObjects struct {
	testPrograms
	testMaps
	testVariables
}

func (o *testObjects) Close() error { _ = "STUB: not implemented"; return nil }

type testMaps struct {
	Map1 *ebpf.Map `ebpf:"map1"`
}

func (m *testMaps) Close() error { _ = "STUB: not implemented"; return nil }

type testVariables struct {
	AnInt       *ebpf.Variable `ebpf:"an_int"`
	IntArray    *ebpf.Variable `ebpf:"int_array"`
	MyConstant  *ebpf.Variable `ebpf:"my_constant"`
	StructArray *ebpf.Variable `ebpf:"struct_array"`
	StructConst *ebpf.Variable `ebpf:"struct_const"`
	StructVar   *ebpf.Variable `ebpf:"struct_var"`
	UnionVar    *ebpf.Variable `ebpf:"union_var"`
}

type testPrograms struct {
	Filter *ebpf.Program `ebpf:"filter"`
}

func (p *testPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _TestClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed test_bpfel.o
var _TestBytes []byte
