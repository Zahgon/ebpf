//go:build (mips || mips64 || ppc64 || s390x) && linux

package main

import (
	_ "embed"
	"io"

	"github.com/cilium/ebpf"
)

const (
	bpfMapKprobeMap     = "kprobe_map"
	bpfProgKprobeExecve = "kprobe_execve"
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
	KprobeExecve *ebpf.ProgramSpec `ebpf:"kprobe_execve"`
}

type bpfMapSpecs struct {
	KprobeMap *ebpf.MapSpec `ebpf:"kprobe_map"`
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
	KprobeMap *ebpf.Map `ebpf:"kprobe_map"`
}

func (m *bpfMaps) Close() error { _ = "STUB: not implemented"; return nil }

type bpfVariables struct {
}

type bpfPrograms struct {
	KprobeExecve *ebpf.Program `ebpf:"kprobe_execve"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_bpfeb.o
var _BpfBytes []byte
