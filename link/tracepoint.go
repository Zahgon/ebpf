//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type TracepointOptions struct {
	Cookie uint64
}

func Tracepoint(group, name string, prog *ebpf.Program, opts *TracepointOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}
