//go:build (mips || mips64 || ppc64 || s390x) && linux

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	bpfProgXdpProgPass = "xdp_prog_pass"
	bpfProgXdpProgTx   = "xdp_prog_tx"
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
	XdpProgPass *ebpf.ProgramSpec `ebpf:"xdp_prog_pass"`
	XdpProgTx   *ebpf.ProgramSpec `ebpf:"xdp_prog_tx"`
}

type bpfMapSpecs struct {
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
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
}

type bpfPrograms struct {
	XdpProgPass *ebpf.Program `ebpf:"xdp_prog_pass"`
	XdpProgTx   *ebpf.Program `ebpf:"xdp_prog_tx"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfeb.o
var _BpfBytes []byte
