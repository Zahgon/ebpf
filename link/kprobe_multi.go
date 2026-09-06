//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type KprobeMultiOptions struct {
	Symbols []string

	Addresses []uintptr

	Cookies []uint64

	Session bool
}

func KprobeMulti(prog *ebpf.Program, opts KprobeMultiOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func KretprobeMulti(prog *ebpf.Program, opts KprobeMultiOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func kprobeMulti(prog *ebpf.Program, opts KprobeMultiOptions, flags uint32) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type kprobeMultiLink struct {
	RawLink
}

var _ Link = (*kprobeMultiLink)(nil)

func (kml *kprobeMultiLink) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (kml *kprobeMultiLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
