package ringbuf

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/cilium/ebpf"
)

type RawReader struct {
	poller poller

	ring       eventRing
	deadline   time.Time
	bufferSize int
	drainErr   error

	inflight, closed atomic.Bool
}

var errConcurrent = errors.New("concurrent use of RawReader")

func (rr *RawReader) enter() error { _ = "STUB: not implemented"; return nil }

func (rr *RawReader) leave() { _ = "STUB: not implemented"; return }

func NewRawReader(m *ebpf.Map) (*RawReader, error) { _ = "STUB: not implemented"; return nil, nil }

func (rr *RawReader) WithLease(fn func(Lease) error) error { _ = "STUB: not implemented"; return nil }

func (rr *RawReader) Close() error { _ = "STUB: not implemented"; return nil }

func (rr *RawReader) SetDeadline(t time.Time) { _ = "STUB: not implemented"; return }

type Lease struct {
	rr *RawReader
}

func (s Lease) ReadSample() (data []byte, remain int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s Lease) Commit() { _ = "STUB: not implemented"; return }

func (rr *RawReader) BufferSize() int { _ = "STUB: not implemented"; return 0 }

func (rr *RawReader) Flush() error { _ = "STUB: not implemented"; return nil }

func (rr *RawReader) AvailableBytes() int { _ = "STUB: not implemented"; return 0 }
