//go:build !windows

package perf

import (
	"runtime"

	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

type perfEventRing struct {
	cpu  int
	mmap []byte
	ringReader
	cleanup runtime.Cleanup
}

func newPerfEventRing(cpu, perCPUBuffer int, opts ReaderOptions) (_ *sys.FD, _ *perfEventRing, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func perfBufferSize(perCPUBuffer int) int { _ = "STUB: not implemented"; return 0 }

func (ring *perfEventRing) Close() error { _ = "STUB: not implemented"; return nil }

func createPerfEvent(cpu int, opts ReaderOptions) (*sys.FD, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ringReader interface {
	loadHead()
	size() int
	remaining() int
	writeTail()
	Read(p []byte) (int, error)
}

type forwardReader struct {
	meta       *unix.PerfEventMmapPage
	head, tail uint64
	mask       uint64
	ring       []byte
}

func newForwardReader(meta *unix.PerfEventMmapPage, ring []byte) *forwardReader {
	_ = "STUB: not implemented"
	return nil
}

func (rr *forwardReader) loadHead() { _ = "STUB: not implemented"; return }

func (rr *forwardReader) size() int { _ = "STUB: not implemented"; return 0 }

func (rr *forwardReader) remaining() int { _ = "STUB: not implemented"; return 0 }

func (rr *forwardReader) writeTail() { _ = "STUB: not implemented"; return }

func (rr *forwardReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type reverseReader struct {
	meta *unix.PerfEventMmapPage

	head uint64

	read uint64

	tail uint64
	mask uint64
	ring []byte
}

func newReverseReader(meta *unix.PerfEventMmapPage, ring []byte) *reverseReader {
	_ = "STUB: not implemented"
	return nil
}

func (rr *reverseReader) loadHead() { _ = "STUB: not implemented"; return }

func (rr *reverseReader) size() int { _ = "STUB: not implemented"; return 0 }

func (rr *reverseReader) remaining() int { _ = "STUB: not implemented"; return 0 }

func (rr *reverseReader) writeTail() { _ = "STUB: not implemented"; return }

func (rr *reverseReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
