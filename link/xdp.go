//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type XDPAttachFlags uint32

const (
	XDPGenericMode XDPAttachFlags = 1 << (iota + 1)

	XDPDriverMode

	XDPOffloadMode
)

type XDPOptions struct {
	Program *ebpf.Program

	Interface int

	Flags XDPAttachFlags
}

func AttachXDP(opts XDPOptions) (Link, error) { _ = "STUB: not implemented"; return *new(Link), nil }

type xdpLink struct {
	RawLink
}

func (xdp *xdpLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
