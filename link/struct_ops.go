package link

import (
	"github.com/cilium/ebpf"
)

type structOpsLink struct {
	RawLink
}

func (*structOpsLink) Update(*ebpf.Program) error { _ = "STUB: not implemented"; return nil }

type StructOpsOptions struct {
	Map *ebpf.Map
}

func AttachStructOps(opts StructOpsOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}
