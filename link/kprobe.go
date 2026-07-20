//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal/tracefs"
)

type KprobeOptions struct {
	Cookie uint64

	Offset uint64

	RetprobeMaxActive int

	TraceFSPrefix string
}

func (ko *KprobeOptions) cookie() uint64 { _ = "STUB: not implemented"; return 0 }

func SyscallWrapper(syscall string) string { _ = "STUB: not implemented"; return "" }

func Kprobe(symbol string, prog *ebpf.Program, opts *KprobeOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func Kretprobe(symbol string, prog *ebpf.Program, opts *KprobeOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func isValidKprobeSymbol(s string) bool { _ = "STUB: not implemented"; return false }

func kprobe(symbol string, prog *ebpf.Program, opts *KprobeOptions, ret bool) (*perfEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pmuProbe(args tracefs.ProbeArgs) (*perfEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tracefsProbe(args tracefs.ProbeArgs) (*perfEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
