package ringbuf

import (
	"runtime"

	"github.com/cilium/ebpf/internal/sys"
)

var _ eventRing = (*windowsEventRing)(nil)

type windowsEventRing struct {
	mapFd            *sys.FD
	cons, prod, data *uint8
	*ringReader

	cleanup runtime.Cleanup
}

func newRingBufEventRing(mapFD, size int) (*windowsEventRing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ring *windowsEventRing) close() error { _ = "STUB: not implemented"; return nil }
