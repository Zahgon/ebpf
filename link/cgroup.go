//go:build !windows

package link

import (
	"os"

	"github.com/cilium/ebpf"
)

type cgroupAttachFlags uint32

const (
	flagAllowOverride cgroupAttachFlags = 1 << iota

	flagAllowMulti

	flagReplace
)

type CgroupOptions struct {
	Path string

	Attach ebpf.AttachType

	Program *ebpf.Program
}

func AttachCgroup(opts CgroupOptions) (cg Link, err error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type progAttachCgroup struct {
	cgroup     *os.File
	current    *ebpf.Program
	attachType ebpf.AttachType
	flags      cgroupAttachFlags
}

var _ Link = (*progAttachCgroup)(nil)

func (cg *progAttachCgroup) isLink() { _ = "STUB: not implemented"; return }

func newProgAttachCgroup(cgroup *os.File, attach ebpf.AttachType, prog *ebpf.Program, flags cgroupAttachFlags) (*progAttachCgroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cg *progAttachCgroup) Close() error { _ = "STUB: not implemented"; return nil }

func (cg *progAttachCgroup) Update(prog *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (cg *progAttachCgroup) Pin(string) error { _ = "STUB: not implemented"; return nil }

func (cg *progAttachCgroup) Unpin() error { _ = "STUB: not implemented"; return nil }

func (cg *progAttachCgroup) Detach() error { _ = "STUB: not implemented"; return nil }

func (cg *progAttachCgroup) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

type linkCgroup struct {
	RawLink
}

var _ Link = (*linkCgroup)(nil)

func newLinkCgroup(cgroup *os.File, attach ebpf.AttachType, prog *ebpf.Program) (*linkCgroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cg *linkCgroup) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }
