//go:build (386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64 || wasm) && linux

package main

import (
	_ "embed"
	"io"
	"structs"

	"github.com/cilium/ebpf"
)

type bpfRttEvent struct {
	_     structs.HostLayout
	Sport uint16
	Dport uint16
	Saddr uint32
	Daddr uint32
	Srtt  uint32
}

type bpfSkInfo struct {
	_      structs.HostLayout
	SkKey  bpfSkKey
	SkType uint8
	_      [3]byte
}

type bpfSkKey struct {
	_          structs.HostLayout
	LocalIp4   uint32
	RemoteIp4  uint32
	LocalPort  uint32
	RemotePort uint32
}

const (
	bpfMapMapEstabSk    = "map_estab_sk"
	bpfMapRttEvents     = "rtt_events"
	bpfProgBpfSockopsCb = "bpf_sockops_cb"
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
	BpfSockopsCb *ebpf.ProgramSpec `ebpf:"bpf_sockops_cb"`
}

type bpfMapSpecs struct {
	MapEstabSk *ebpf.MapSpec `ebpf:"map_estab_sk"`
	RttEvents  *ebpf.MapSpec `ebpf:"rtt_events"`
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
	MapEstabSk *ebpf.Map `ebpf:"map_estab_sk"`
	RttEvents  *ebpf.Map `ebpf:"rtt_events"`
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
}

type bpfPrograms struct {
	BpfSockopsCb *ebpf.Program `ebpf:"bpf_sockops_cb"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfel.o
var _BpfBytes []byte
