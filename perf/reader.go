//go:build !windows

package perf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal/epoll"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

var (
	ErrClosed  = os.ErrClosed
	ErrFlushed = epoll.ErrFlushed
	errEOR     = errors.New("end of ring")
)

var perfEventHeaderSize = binary.Size(perfEventHeader{})

type perfEventHeader struct {
	Type uint32
	Misc uint16
	Size uint16
}

type Record struct {
	CPU int

	RawSample []byte

	LostSamples uint64

	Remaining int
}

func readRecord(rd io.Reader, rec *Record, buf []byte, overwritable bool) error {
	_ = "STUB: not implemented"
	return nil
}

func readLostRecords(rd io.Reader) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

var perfEventSampleSize = binary.Size(uint32(0))

type perfEventSample struct {
	Size uint32
}

func readRawSample(rd io.Reader, buf, sampleBuf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Reader struct {
	poller *epoll.Poller

	mu           sync.Mutex
	array        *ebpf.Map
	rings        []*perfEventRing
	epollEvents  []unix.EpollEvent
	epollRings   []*perfEventRing
	eventHeader  []byte
	deadline     time.Time
	overwritable bool
	bufferSize   int
	pendingErr   error

	pauseMu  sync.Mutex
	eventFds []*sys.FD
	paused   bool
}

type ReaderOptions struct {
	WakeupEvents int

	Watermark int

	Overwritable bool
}

func NewReader(array *ebpf.Map, perCPUBuffer int) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReaderWithOptions(array *ebpf.Map, perCPUBuffer int, opts ReaderOptions) (pr *Reader, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pr *Reader) Close() error { _ = "STUB: not implemented"; return nil }

func (pr *Reader) SetDeadline(t time.Time) { _ = "STUB: not implemented"; return }

func (pr *Reader) Read() (Record, error) { _ = "STUB: not implemented"; return *new(Record), nil }

var errMustBePaused = fmt.Errorf("perf ringbuffer: must have been paused before reading overwritable buffer")

func (pr *Reader) ReadInto(rec *Record) error { _ = "STUB: not implemented"; return nil }

func (pr *Reader) Pause() error { _ = "STUB: not implemented"; return nil }

func (pr *Reader) Resume() error { _ = "STUB: not implemented"; return nil }

func (pr *Reader) BufferSize() int { _ = "STUB: not implemented"; return 0 }

func (pr *Reader) Flush() error { _ = "STUB: not implemented"; return nil }

func (pr *Reader) readRecordFromRing(rec *Record, ring *perfEventRing) error {
	_ = "STUB: not implemented"
	return nil
}

type unknownEventError struct {
	eventType uint32
}

func (uev *unknownEventError) Error() string { _ = "STUB: not implemented"; return "" }

func IsUnknownEvent(err error) bool { _ = "STUB: not implemented"; return false }
