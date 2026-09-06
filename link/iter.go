//go:build !windows

package link

import (
	"io"

	"github.com/cilium/ebpf"
)

type IterOptions struct {
	Program *ebpf.Program

	Map *ebpf.Map
}

func AttachIter(opts IterOptions) (*Iter, error) { _ = "STUB: not implemented"; return nil, nil }

type Iter struct {
	RawLink
}

func (it *Iter) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

func (it *Iter) Open() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type bpfIterLinkInfoMap struct {
	map_fd uint32
}
