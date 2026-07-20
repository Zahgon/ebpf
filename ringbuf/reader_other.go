//go:build !windows

package ringbuf

import (
	"time"

	"github.com/cilium/ebpf/internal/epoll"
	"github.com/cilium/ebpf/internal/unix"
)

var ErrFlushed = epoll.ErrFlushed

var _ poller = (*epollPoller)(nil)

type epollPoller struct {
	*epoll.Poller
	events []unix.EpollEvent
}

func newPoller(fd int) (*epollPoller, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *epollPoller) Wait(deadline time.Time) error { _ = "STUB: not implemented"; return nil }
