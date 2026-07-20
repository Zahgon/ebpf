//go:build !windows

package link

import (
	"errors"
	"os"
	"unsafe"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/tracefs"
	"github.com/cilium/ebpf/internal/unix"
)

var (
	errInvalidInput = tracefs.ErrInvalidInput
)

const (
	perfAllThreads = -1
)

type perfEvent struct {
	tracefsEvent *tracefs.Event

	fd *sys.FD
}

func newPerfEvent(fd *sys.FD, event *tracefs.Event) *perfEvent {
	_ = "STUB: not implemented"
	return nil
}

func (pe *perfEvent) Close() error { _ = "STUB: not implemented"; return nil }

type PerfEvent interface {
	PerfEvent() (*os.File, error)
}

type perfEventLink struct {
	RawLink
	pe *perfEvent
}

func (pl *perfEventLink) isLink() { _ = "STUB: not implemented"; return }

func (pl *perfEventLink) Close() error { _ = "STUB: not implemented"; return nil }

func (pl *perfEventLink) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

var _ PerfEvent = (*perfEventLink)(nil)

func (pl *perfEventLink) PerfEvent() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func queryInfoWithString(fd *sys.FD, info sys.Info, stringField *sys.TypedPointer[byte], stringLengthField *uint32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (pl *perfEventLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

type perfEventIoctl struct {
	*perfEvent
}

func (pi *perfEventIoctl) isLink() { _ = "STUB: not implemented"; return }

func (pi *perfEventIoctl) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (pi *perfEventIoctl) Pin(string) error { _ = "STUB: not implemented"; return nil }

func (pi *perfEventIoctl) Unpin() error { _ = "STUB: not implemented"; return nil }

func (pi *perfEventIoctl) Detach() error { _ = "STUB: not implemented"; return nil }

func (pi *perfEventIoctl) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

var _ PerfEvent = (*perfEventIoctl)(nil)

func (pi *perfEventIoctl) PerfEvent() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func attachPerfEvent(pe *perfEvent, prog *ebpf.Program, cookie uint64) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func attachPerfEventIoctl(pe *perfEvent, prog *ebpf.Program) (*perfEventIoctl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func attachPerfEventLink(pe *perfEvent, prog *ebpf.Program, cookie uint64) (*perfEventLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unsafeStringPtr(str string) (unsafe.Pointer, error) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), nil
}

func openTracepointPerfEvent(tid uint64, pid int) (*sys.FD, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var haveBPFLinkPerfEvent = internal.NewFeatureTest("bpf_link_perf_event", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Name: "probe_bpf_perf_link",
		Type: ebpf.Kprobe,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
		License: "MIT",
	})
	if err != nil {
		return err
	}
	defer prog.Close()

	_, err = sys.LinkCreatePerfEvent(&sys.LinkCreatePerfEventAttr{
		ProgFd:     uint32(prog.FD()),
		AttachType: sys.BPF_PERF_EVENT,
	})
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	if errors.Is(err, unix.EBADF) {
		return nil
	}
	return err
}, "5.15")
