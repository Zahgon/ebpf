//go:build (mips || mips64 || ppc64 || s390x) && linux

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	bpfMapPktCount            = "pkt_count"
	bpfProgCountEgressPackets = "count_egress_packets"
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
	CountEgressPackets *ebpf.ProgramSpec `ebpf:"count_egress_packets"`
}

type bpfMapSpecs struct {
	PktCount *ebpf.MapSpec `ebpf:"pkt_count"`
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
	PktCount *ebpf.Map `ebpf:"pkt_count"`
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
}

type bpfPrograms struct {
	CountEgressPackets *ebpf.Program `ebpf:"count_egress_packets"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfeb.o
var _BpfBytes []byte
