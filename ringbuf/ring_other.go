//go:build !windows

package ringbuf

import (
	"runtime"
)

var _ eventRing = (*mmapEventRing)(nil)

type mmapEventRing struct {
	prod []byte
	cons []byte
	*ringReader
	cleanup runtime.Cleanup
}

func newRingBufEventRing(mapFD, size int) (*mmapEventRing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ring *mmapEventRing) close() error { _ = "STUB: not implemented"; return nil }
