//go:build !windows

package link

import (
	"syscall"

	"github.com/cilium/ebpf"
)

func AttachSocketFilter(conn syscall.Conn, program *ebpf.Program) error {
	_ = "STUB: not implemented"
	return nil
}

func DetachSocketFilter(conn syscall.Conn) error { _ = "STUB: not implemented"; return nil }
