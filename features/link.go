package features

import (
	"errors"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

func HaveBPFLinkUprobeMulti() error { _ = "STUB: not implemented"; return nil }

var haveBPFLinkUprobeMulti = internal.NewFeatureTest("bpf_link_uprobe_multi", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Name: "probe_upm_link",
		Type: ebpf.Kprobe,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
		AttachType: ebpf.AttachTraceUprobeMulti,
		License:    "MIT",
	})
	if errors.Is(err, unix.E2BIG) {

		return ebpf.ErrNotSupported
	}
	if err != nil {
		return err
	}
	defer prog.Close()

	fd, err := sys.LinkCreateUprobeMulti(&sys.LinkCreateUprobeMultiAttr{
		ProgFd:     uint32(prog.FD()),
		AttachType: sys.BPF_TRACE_UPROBE_MULTI,
		Path:       sys.NewStringPointer("/"),
		Offsets:    sys.SlicePointer([]uint64{0}),
		Count:      1,
	})
	switch {
	case errors.Is(err, unix.EBADF):
		return nil
	case errors.Is(err, unix.EINVAL):
		return ebpf.ErrNotSupported
	case err != nil:
		return err
	}

	fd.Close()
	return errors.New("successfully attached uprobe_multi to /, kernel bug?")
}, "6.6")

func HaveBPFLinkKprobeMulti() error { _ = "STUB: not implemented"; return nil }

var haveBPFLinkKprobeMulti = internal.NewFeatureTest("bpf_link_kprobe_multi", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Name: "probe_kpm_link",
		Type: ebpf.Kprobe,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
		AttachType: ebpf.AttachTraceKprobeMulti,
		License:    "MIT",
	})
	if errors.Is(err, unix.E2BIG) {

		return ebpf.ErrNotSupported
	}
	if err != nil {
		return err
	}
	defer prog.Close()

	fd, err := sys.LinkCreateKprobeMulti(&sys.LinkCreateKprobeMultiAttr{
		ProgFd:     uint32(prog.FD()),
		AttachType: sys.BPF_TRACE_KPROBE_MULTI,
		Count:      1,
		Syms:       sys.NewStringSlicePointer([]string{"vprintk"}),
	})
	switch {
	case errors.Is(err, unix.EINVAL):
		return ebpf.ErrNotSupported

	case errors.Is(err, unix.EOPNOTSUPP):
		return ebpf.ErrNotSupported
	case err != nil:
		return err
	}

	fd.Close()

	return nil
}, "5.18")

func HaveBPFLinkKprobeSession() error { _ = "STUB: not implemented"; return nil }

var haveBPFLinkKprobeSession = internal.NewFeatureTest("bpf_link_kprobe_session", func() error {
	prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		Name: "probe_kps_link",
		Type: ebpf.Kprobe,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.Return(),
		},
		AttachType: ebpf.AttachTraceKprobeSession,
		License:    "MIT",
	})
	if errors.Is(err, unix.E2BIG) {

		return ebpf.ErrNotSupported
	}
	if err != nil {
		return err
	}
	defer prog.Close()

	fd, err := sys.LinkCreateKprobeMulti(&sys.LinkCreateKprobeMultiAttr{
		ProgFd:     uint32(prog.FD()),
		AttachType: sys.BPF_TRACE_KPROBE_SESSION,
		Count:      1,
		Syms:       sys.NewStringSlicePointer([]string{"vprintk"}),
	})
	switch {
	case errors.Is(err, unix.EINVAL):
		return ebpf.ErrNotSupported

	case errors.Is(err, unix.EOPNOTSUPP):
		return ebpf.ErrNotSupported
	case err != nil:
		return err
	}

	fd.Close()

	return nil
}, "6.10")
