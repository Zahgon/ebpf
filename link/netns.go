//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type NetNsLink struct {
	RawLink
}

func AttachNetNs(ns int, prog *ebpf.Program) (*NetNsLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NetNsLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
