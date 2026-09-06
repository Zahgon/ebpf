//go:build (386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64 || wasm) && linux

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	bpfMapCountingMap  = "counting_map"
	bpfProgMmPageAlloc = "mm_page_alloc"
)

func loadBpf() (*ebpf.CollectionSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func loadBpfObjects(obj any, opts *ebpf.CollectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
	bpfVariableSpecs
}

type bpfProgramSpecs struct {
	MmPageAlloc *ebpf.ProgramSpec `ebpf:"mm_page_alloc"`
}

type bpfMapSpecs struct {
	CountingMap *ebpf.MapSpec `ebpf:"counting_map"`
}

type bpfVariableSpecs struct {
}

type bpfObjects struct {
	bpfPrograms
	bpfMaps
	bpfVariables
}

func (o *bpfObjects) Close() error { _ = "STUB: not implemented"; return nil }

type bpfMaps struct {
	CountingMap *ebpf.Map `ebpf:"counting_map"`
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
}

type bpfPrograms struct {
	MmPageAlloc *ebpf.Program `ebpf:"mm_page_alloc"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfel.o
var _BpfBytes []byte
