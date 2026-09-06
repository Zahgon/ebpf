//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type NetkitOptions struct {
	Interface int

	Program *ebpf.Program

	Attach ebpf.AttachType

	Anchor Anchor

	ExpectedRevision uint64

	Flags uint32
}

func AttachNetkit(opts NetkitOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type netkitLink struct {
	RawLink
}

var _ Link = (*netkitLink)(nil)

func (netkit *netkitLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
