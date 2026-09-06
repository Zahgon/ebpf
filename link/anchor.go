//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal/sys"
)

const anchorFlags = sys.BPF_F_REPLACE |
	sys.BPF_F_BEFORE |
	sys.BPF_F_AFTER |
	sys.BPF_F_ID |
	sys.BPF_F_LINK_MPROG

type Anchor interface {
	anchor() (fdOrID, flags uint32, _ error)
}

type firstAnchor struct{}

func (firstAnchor) anchor() (fdOrID, flags uint32, _ error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func Head() Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

type lastAnchor struct{}

func (lastAnchor) anchor() (fdOrID, flags uint32, _ error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func Tail() Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func BeforeLink(target Link) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func AfterLink(target Link) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func BeforeLinkByID(target ID) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func AfterLinkByID(target ID) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func BeforeProgram(target *ebpf.Program) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func AfterProgram(target *ebpf.Program) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func ReplaceProgram(target *ebpf.Program) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func BeforeProgramByID(target ebpf.ProgramID) Anchor {
	_ = "STUB: not implemented"
	return *new(Anchor)
}

func AfterProgramByID(target ebpf.ProgramID) Anchor { _ = "STUB: not implemented"; return *new(Anchor) }

func ReplaceProgramByID(target ebpf.ProgramID) Anchor {
	_ = "STUB: not implemented"
	return *new(Anchor)
}

type anchor struct {
	target   any
	position uint32
}

func (ap anchor) anchor() (fdOrID, flags uint32, _ error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
