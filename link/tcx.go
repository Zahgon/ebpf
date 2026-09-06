//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type TCXOptions struct {
	Interface int

	Program *ebpf.Program

	Attach ebpf.AttachType

	Anchor Anchor

	ExpectedRevision uint64

	Flags uint32
}

func AttachTCX(opts TCXOptions) (Link, error) { _ = "STUB: not implemented"; return *new(Link), nil }

type tcxLink struct {
	RawLink
}

var _ Link = (*tcxLink)(nil)

func (tcx *tcxLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
