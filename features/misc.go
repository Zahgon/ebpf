package features

import (
	"errors"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

func HaveLargeInstructions() error { _ = "STUB: not implemented"; return nil }

var haveLargeInstructions = internal.NewFeatureTest(">4096 instructions", func() error {
	const maxInsns = 4096

	insns := make(asm.Instructions, maxInsns, maxInsns+1)
	for i := range insns {
		insns[i] = asm.Mov.Imm(asm.R0, 1)
	}
	insns = append(insns, asm.Return())

	return probeProgram(&ebpf.ProgramSpec{
		Type:         ebpf.SocketFilter,
		Instructions: insns,
	})
}, "5.2")

func HaveBoundedLoops() error { _ = "STUB: not implemented"; return nil }

var haveBoundedLoops = internal.NewFeatureTest("bounded loops", func() error {
	return probeProgram(&ebpf.ProgramSpec{
		Type: ebpf.SocketFilter,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 10),
			asm.Sub.Imm(asm.R0, 1).WithSymbol("loop"),
			asm.JNE.Imm(asm.R0, 0, "loop"),
			asm.Return(),
		},
	})
}, "5.3")

func HaveV2ISA() error { _ = "STUB: not implemented"; return nil }

var haveV2ISA = internal.NewFeatureTest("v2 ISA", func() error {
	err := probeProgram(&ebpf.ProgramSpec{
		Type: ebpf.SocketFilter,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.JLT.Imm(asm.R0, 0, "exit"),
			asm.Mov.Imm(asm.R0, 1),
			asm.Return().WithSymbol("exit"),
		},
	})

	if errors.Is(err, sys.ENOTSUPP) {
		return ebpf.ErrNotSupported
	}
	return err
}, "4.14")

func HaveV3ISA() error { _ = "STUB: not implemented"; return nil }

var haveV3ISA = internal.NewFeatureTest("v3 ISA", func() error {
	err := probeProgram(&ebpf.ProgramSpec{
		Type: ebpf.SocketFilter,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.JLT.Imm32(asm.R0, 0, "exit"),
			asm.Mov.Imm(asm.R0, 1),
			asm.Return().WithSymbol("exit"),
		},
	})

	if errors.Is(err, sys.ENOTSUPP) {
		return ebpf.ErrNotSupported
	}
	return err
}, "5.1")

func HaveV4ISA() error { _ = "STUB: not implemented"; return nil }

var haveV4ISA = internal.NewFeatureTest("v4 ISA", func() error {
	err := probeProgram(&ebpf.ProgramSpec{
		Type: ebpf.SocketFilter,
		Instructions: asm.Instructions{
			asm.Mov.Imm(asm.R0, 0),
			asm.JEq.Imm(asm.R0, 1, "error"),
			asm.LongJump("exit"),
			asm.Mov.Imm(asm.R0, 1).WithSymbol("error"),
			asm.Return().WithSymbol("exit"),
		},
	})

	if errors.Is(err, sys.ENOTSUPP) {
		return ebpf.ErrNotSupported
	}
	return err
}, "6.6")
