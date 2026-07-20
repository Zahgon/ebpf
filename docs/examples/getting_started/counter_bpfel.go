//go:build (386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64 || wasm) && linux

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	counterMapPktCount      = "pkt_count"
	counterProgCountPackets = "count_packets"
)

func loadCounter() (*ebpf.CollectionSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func loadCounterObjects(obj any, opts *ebpf.CollectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type counterSpecs struct {
	counterProgramSpecs
	counterMapSpecs
	counterVariableSpecs
}

type counterProgramSpecs struct {
	CountPackets *ebpf.ProgramSpec `ebpf:"count_packets"`
}

type counterMapSpecs struct {
	PktCount *ebpf.MapSpec `ebpf:"pkt_count"`
}

type counterVariableSpecs struct {
}

type counterObjects struct {
	counterPrograms
	counterMaps
	counterVariables
}

func (o *counterObjects) Close() error { _ = "STUB: not implemented"; return nil }

type counterMaps struct {
	PktCount *ebpf.Map `ebpf:"pkt_count"`
}

func (m *counterMaps) Close() error { _ = "STUB: not implemented"; return nil }

type counterVariables struct {
}

type counterPrograms struct {
	CountPackets *ebpf.Program `ebpf:"count_packets"`
}

func (p *counterPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _CounterClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed counter_bpfel.o
var _CounterBytes []byte
