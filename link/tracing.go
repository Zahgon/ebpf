//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type tracing struct {
	RawLink
}

func (f *tracing) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (f *tracing) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

func AttachFreplace(targetProg *ebpf.Program, name string, prog *ebpf.Program) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type TracingOptions struct {
	Program *ebpf.Program

	AttachType ebpf.AttachType

	Cookie uint64
}

type LSMOptions struct {
	Program *ebpf.Program

	Cookie uint64
}

func attachBTFID(program *ebpf.Program, at ebpf.AttachType, cookie uint64) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func AttachTracing(opts TracingOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func AttachLSM(opts LSMOptions) (Link, error) { _ = "STUB: not implemented"; return *new(Link), nil }
