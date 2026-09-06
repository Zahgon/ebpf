package btf

import (
	"github.com/cilium/ebpf/internal/sys"
)

type Handle struct {
	fd *sys.FD

	size uint32

	needsKernelBase bool
}

func NewHandle(b *Builder) (*Handle, error) { _ = "STUB: not implemented"; return nil, nil }

func NewHandleFromRawBTF(btf []byte) (*Handle, error) { _ = "STUB: not implemented"; return nil, nil }

func NewHandleFromID(id ID) (*Handle, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *Handle) Spec(base *Spec) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *Handle) Close() error { _ = "STUB: not implemented"; return nil }

func (h *Handle) FD() int { _ = "STUB: not implemented"; return 0 }

func (h *Handle) Clone() (*Handle, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *Handle) Info() (*HandleInfo, error) { _ = "STUB: not implemented"; return nil, nil }

type HandleInfo struct {
	ID ID

	Name string

	IsKernel bool

	size uint32
}

func newHandleInfoFromFD(fd *sys.FD) (*HandleInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *HandleInfo) IsVmlinux() bool { _ = "STUB: not implemented"; return false }

func (i *HandleInfo) IsModule() bool { _ = "STUB: not implemented"; return false }

type HandleIterator struct {
	ID ID

	Handle *Handle
	err    error
}

func (it *HandleIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *HandleIterator) Take() *Handle { _ = "STUB: not implemented"; return nil }

func (it *HandleIterator) Err() error { _ = "STUB: not implemented"; return nil }

func FindHandle(predicate func(info *HandleInfo) bool) (*Handle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
