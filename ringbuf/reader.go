package ringbuf

import (
	"errors"
	"os"
	"sync"
	"time"
	"unsafe"

	"github.com/cilium/ebpf"
)

var (
	ErrClosed = os.ErrClosed
	errEOR    = errors.New("end of ring")
)

type poller interface {
	Wait(deadline time.Time) error
	Flush() error
	Close() error
}

type eventRing interface {
	size() int
	available() int
	readSample() (data []byte, remain int, err error)
	commit()
	close() error
}

type ringbufHeader struct {
	Len uint32
	_   uint32
}

const ringbufHeaderSize = int(unsafe.Sizeof(ringbufHeader{}))

func (rh *ringbufHeader) isBusy() bool { _ = "STUB: not implemented"; return false }

func (rh *ringbufHeader) isDiscard() bool { _ = "STUB: not implemented"; return false }

func (rh *ringbufHeader) dataLen() int { _ = "STUB: not implemented"; return 0 }

func (rh *ringbufHeader) dataLenAligned() int { _ = "STUB: not implemented"; return 0 }

type Record struct {
	RawSample []byte

	Remaining int
}

type Reader struct {
	mu  sync.Mutex
	raw *RawReader
}

func NewReader(m *ebpf.Map) (*Reader, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Reader) SetDeadline(t time.Time) { _ = "STUB: not implemented"; return }

func (r *Reader) Read() (Record, error) { _ = "STUB: not implemented"; return *new(Record), nil }

func (r *Reader) ReadInto(rec *Record) error { _ = "STUB: not implemented"; return nil }

func (r *Reader) BufferSize() int { _ = "STUB: not implemented"; return 0 }

func (r *Reader) Flush() error { _ = "STUB: not implemented"; return nil }

func (r *Reader) AvailableBytes() int { _ = "STUB: not implemented"; return 0 }
