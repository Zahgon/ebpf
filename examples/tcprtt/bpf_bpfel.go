//go:build (386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64 || wasm) && linux

package main

import (
	_ "embed"
	"io"
	"structs"

	"github.com/cilium/ebpf"
)

type bpfEvent struct {
	_     structs.HostLayout
	Sport uint16
	Dport uint16
	Saddr uint32
	Daddr uint32
	Srtt  uint32
}

const (
	bpfMapEvents    = "events"
	bpfProgTcpClose = "tcp_close"
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
	TcpClose *ebpf.ProgramSpec `ebpf:"tcp_close"`
}

type bpfMapSpecs struct {
	Events *ebpf.MapSpec `ebpf:"events"`
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
	Events *ebpf.Map `ebpf:"events"`
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
}

type bpfPrograms struct {
	TcpClose *ebpf.Program `ebpf:"tcp_close"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfel.o
var _BpfBytes []byte
