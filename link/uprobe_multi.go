//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type UprobeMultiOptions struct {
	Addresses []uint64

	Offsets []uint64

	RefCtrOffsets []uint64

	Cookies []uint64

	PID uint32
}

func (ex *Executable) UprobeMulti(symbols []string, prog *ebpf.Program, opts *UprobeMultiOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (ex *Executable) UretprobeMulti(symbols []string, prog *ebpf.Program, opts *UprobeMultiOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (ex *Executable) uprobeMulti(symbols []string, prog *ebpf.Program, opts *UprobeMultiOptions, flags uint32) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (ex *Executable) addresses(symbols []string, addresses, offsets []uint64) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type uprobeMultiLink struct {
	RawLink
}

var _ Link = (*uprobeMultiLink)(nil)

func (kml *uprobeMultiLink) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (kml *uprobeMultiLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
