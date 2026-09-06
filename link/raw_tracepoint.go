//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal/sys"
)

type RawTracepointOptions struct {
	Name string

	Program *ebpf.Program
}

func AttachRawTracepoint(opts RawTracepointOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type simpleRawTracepoint struct {
	fd *sys.FD
}

var _ Link = (*simpleRawTracepoint)(nil)

func (frt *simpleRawTracepoint) isLink() { _ = "STUB: not implemented"; return }

func (frt *simpleRawTracepoint) Close() error { _ = "STUB: not implemented"; return nil }

func (frt *simpleRawTracepoint) Update(_ *ebpf.Program) error {
	_ = "STUB: not implemented"
	return nil
}

func (frt *simpleRawTracepoint) Pin(string) error { _ = "STUB: not implemented"; return nil }

func (frt *simpleRawTracepoint) Unpin() error { _ = "STUB: not implemented"; return nil }

func (frt *simpleRawTracepoint) Detach() error { _ = "STUB: not implemented"; return nil }

func (frt *simpleRawTracepoint) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

type rawTracepoint struct {
	RawLink
}

var _ Link = (*rawTracepoint)(nil)

func (rt *rawTracepoint) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (rt *rawTracepoint) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
