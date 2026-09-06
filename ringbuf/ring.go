package ringbuf

import (
	"sync/atomic"
)

type ringReader struct {
	prod_pos, cons_pos *atomic.Uintptr

	localCons uintptr

	mask uintptr

	data []byte
}

func newRingReader(cons_ptr, prod_ptr *atomic.Uintptr, data []byte) *ringReader {
	_ = "STUB: not implemented"
	return nil
}

func (rr *ringReader) size() int { _ = "STUB: not implemented"; return 0 }

func (rr *ringReader) available() int { _ = "STUB: not implemented"; return 0 }

func (rr *ringReader) commit() { _ = "STUB: not implemented"; return }

func (rr *ringReader) readSample() (data []byte, remain int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
