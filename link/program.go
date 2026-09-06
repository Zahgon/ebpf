//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type RawAttachProgramOptions struct {
	Target int

	Program *ebpf.Program

	Attach ebpf.AttachType

	Anchor Anchor

	Flags uint32

	ExpectedRevision uint64
}

func RawAttachProgram(opts RawAttachProgramOptions) error { _ = "STUB: not implemented"; return nil }

type RawDetachProgramOptions RawAttachProgramOptions

func RawDetachProgram(opts RawDetachProgramOptions) error { _ = "STUB: not implemented"; return nil }
