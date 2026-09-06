//go:build (mips || mips64 || ppc64 || s390x) && linux

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	bpfProgEgressProgFunc  = "egress_prog_func"
	bpfProgIngressProgFunc = "ingress_prog_func"
	bpfVarEgressPktCount   = "egress_pkt_count"
	bpfVarIngressPktCount  = "ingress_pkt_count"
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
	EgressProgFunc  *ebpf.ProgramSpec `ebpf:"egress_prog_func"`
	IngressProgFunc *ebpf.ProgramSpec `ebpf:"ingress_prog_func"`
}

type bpfMapSpecs struct {
}

type bpfVariableSpecs struct {
	EgressPktCount  *ebpf.VariableSpec `ebpf:"egress_pkt_count"`
	IngressPktCount *ebpf.VariableSpec `ebpf:"ingress_pkt_count"`
}

type bpfObjects struct {
	bpfPrograms
	bpfMaps
	bpfVariables
}

func (o *bpfObjects) Close() error { _ = "STUB: not implemented"; return nil }

type bpfMaps struct {
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
	EgressPktCount  *ebpf.Variable `ebpf:"egress_pkt_count"`
	IngressPktCount *ebpf.Variable `ebpf:"ingress_pkt_count"`
}

type bpfPrograms struct {
	EgressProgFunc  *ebpf.Program `ebpf:"egress_prog_func"`
	IngressProgFunc *ebpf.Program `ebpf:"ingress_prog_func"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfeb.o
var _BpfBytes []byte
