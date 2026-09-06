package ringbuf

import (
	"errors"
	"sync/atomic"
	"time"

	"golang.org/x/sys/windows"
)

var ErrFlushed = errors.New("ring buffer flushed")

var _ poller = (*windowsPoller)(nil)

type windowsPoller struct {
	closed      atomic.Bool
	handle      windows.Handle
	flushHandle windows.Handle
	handles     []windows.Handle
}

func newPoller(fd int) (*windowsPoller, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *windowsPoller) Wait(deadline time.Time) error { _ = "STUB: not implemented"; return nil }

func (p *windowsPoller) Flush() error { _ = "STUB: not implemented"; return nil }

func (p *windowsPoller) Close() error { _ = "STUB: not implemented"; return nil }
