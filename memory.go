package ebpf

import (
	"errors"
	"runtime"
)

var ErrReadOnly = errors.New("resource is read-only")

type Memory struct {
	b    []byte
	ro   bool
	heap bool

	cleanup runtime.Cleanup
}

func newMemory(fd, size int) (*Memory, error) { _ = "STUB: not implemented"; return nil, nil }

func memoryCleanupFunc() func([]byte) { _ = "STUB: not implemented"; return nil }

func (mm *Memory) close() { _ = "STUB: not implemented"; return }

func (mm *Memory) Size() uint32 { _ = "STUB: not implemented"; return 0 }

func (mm *Memory) ReadOnly() bool { _ = "STUB: not implemented"; return false }

func (mm *Memory) bounds(off, size uint32) bool { _ = "STUB: not implemented"; return false }

func (mm *Memory) ReadAt(p []byte, off int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mm *Memory) WriteAt(p []byte, off int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
