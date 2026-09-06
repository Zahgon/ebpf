//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
)

type QueryOptions struct {
	Target int

	Attach ebpf.AttachType

	QueryFlags uint32
}

type QueryResult struct {
	Programs []AttachedProgram

	Revision uint64
}

func (qr *QueryResult) HaveLinkInfo() bool { _ = "STUB: not implemented"; return false }

type AttachedProgram struct {
	ID     ebpf.ProgramID
	linkID ID
}

func (ap *AttachedProgram) LinkID() (ID, bool) { _ = "STUB: not implemented"; return *new(ID), false }

func QueryPrograms(opts QueryOptions) (*QueryResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
