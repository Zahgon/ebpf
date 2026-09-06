//go:build (386 || amd64) && linux

package main

import (
	_ "embed"
	"io"
	"structs"

	"github.com/cilium/ebpf"
)

type bpfEvent struct {
	_    structs.HostLayout
	Pid  uint32
	Line [80]uint8
}

const (
	bpfMapEvents                 = "events"
	bpfProgUretprobeBashReadline = "uretprobe_bash_readline"
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
	UretprobeBashReadline *ebpf.ProgramSpec `ebpf:"uretprobe_bash_readline"`
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
	UretprobeBashReadline *ebpf.Program `ebpf:"uretprobe_bash_readline"`
}

func (p *bpfPrograms) Close() error { _ = "STUB: not implemented"; return nil }

func _BpfClose(closers ...io.Closer) error { _ = "STUB: not implemented"; return nil }

//go:embed bpf_x86_bpfel.o
var _BpfBytes []byte
