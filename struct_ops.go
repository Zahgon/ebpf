package ebpf

import (
	"github.com/cilium/ebpf/btf"
)

const structOpsValuePrefix = "bpf_struct_ops_"
const structOpsLinkSec = ".struct_ops.link"
const structOpsSec = ".struct_ops"
const structOpsKeySize = 4

type structOpsMemberLayout struct {
	member btf.Member
	off    int
	size   int
	typ    btf.Type
}

func newStructOpsMemberLayout(m btf.Member) (*structOpsMemberLayout, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ml *structOpsMemberLayout) bytes(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func structOpsFuncPtrMember(m btf.Member) error { _ = "STUB: not implemented"; return nil }

func structOpsFindInnerType(vType *btf.Struct) (*btf.Struct, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func structOpsFindTarget(userType *btf.Struct, cache *btf.Cache) (vType *btf.Struct, id btf.TypeID, module *btf.Handle, err error) {
	_ = "STUB: not implemented"
	return nil, *new(btf.TypeID), nil, nil
}

func structOpsPopulateValue(km btf.Member, kernVData []byte, p *Program) error {
	_ = "STUB: not implemented"
	return nil
}

func structOpsValidateMemberPair(m, km btf.Member) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func structOpsCopyMemberBytes(m, km btf.Member, data, kernVData []byte, size int) error {
	_ = "STUB: not implemented"
	return nil
}

func structOpsCopyMember(m, km btf.Member, data []byte, kernVData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func structOpsIsMemZeroed(data []byte) bool { _ = "STUB: not implemented"; return false }

func structOpsSetAttachTo(
	sec *elfSection,
	baseOff uint32,
	userSt *btf.Struct,
	progs map[string]*ProgramSpec) error {
	_ = "STUB: not implemented"
	return nil
}
