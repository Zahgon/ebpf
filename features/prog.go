package features

import (
	"errors"
	"fmt"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

func HaveProgramType(pt ebpf.ProgramType) (err error) { _ = "STUB: not implemented"; return nil }

func probeProgram(spec *ebpf.ProgramSpec) error { _ = "STUB: not implemented"; return nil }

var haveProgramTypeMatrix = internal.FeatureMatrix[ebpf.ProgramType]{
	ebpf.SocketFilter:  {Version: "3.19"},
	ebpf.Kprobe:        {Version: "4.1"},
	ebpf.SchedCLS:      {Version: "4.1"},
	ebpf.SchedACT:      {Version: "4.1"},
	ebpf.TracePoint:    {Version: "4.7"},
	ebpf.XDP:           {Version: "4.8"},
	ebpf.PerfEvent:     {Version: "4.9"},
	ebpf.CGroupSKB:     {Version: "4.10"},
	ebpf.CGroupSock:    {Version: "4.10"},
	ebpf.LWTIn:         {Version: "4.10"},
	ebpf.LWTOut:        {Version: "4.10"},
	ebpf.LWTXmit:       {Version: "4.10"},
	ebpf.SockOps:       {Version: "4.13"},
	ebpf.SkSKB:         {Version: "4.14"},
	ebpf.CGroupDevice:  {Version: "4.15"},
	ebpf.SkMsg:         {Version: "4.17"},
	ebpf.RawTracepoint: {Version: "4.17"},
	ebpf.CGroupSockAddr: {
		Version: "4.17",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:       ebpf.CGroupSockAddr,
				AttachType: ebpf.AttachCGroupInet4Connect,
			})
		},
	},
	ebpf.LWTSeg6Local:          {Version: "4.18"},
	ebpf.LircMode2:             {Version: "4.18"},
	ebpf.SkReuseport:           {Version: "4.19"},
	ebpf.FlowDissector:         {Version: "4.20"},
	ebpf.CGroupSysctl:          {Version: "5.2"},
	ebpf.RawTracepointWritable: {Version: "5.2"},
	ebpf.CGroupSockopt: {
		Version: "5.3",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:       ebpf.CGroupSockopt,
				AttachType: ebpf.AttachCGroupGetsockopt,
			})
		},
	},
	ebpf.Tracing: {
		Version: "5.5",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:       ebpf.Tracing,
				AttachType: ebpf.AttachTraceFEntry,
				AttachTo:   "bpf_init",
			})
		},
	},
	ebpf.StructOps: {
		Version: "5.6",
		Fn: func() error {
			err := probeProgram(&ebpf.ProgramSpec{
				Type:    ebpf.StructOps,
				License: "GPL",
			})
			if errors.Is(err, sys.ENOTSUPP) {

				return nil
			}
			return err
		},
	},
	ebpf.Extension: {
		Version: "5.6",
		Fn: func() error {

			btfFn := btf.Func{
				Name: "a",
				Type: &btf.FuncProto{
					Return: &btf.Int{},
					Params: []btf.FuncParam{
						{Name: "ctx", Type: &btf.Pointer{Target: &btf.Struct{Name: "xdp_md"}}},
					},
				},
				Linkage: btf.GlobalFunc,
			}
			insns := asm.Instructions{
				btf.WithFuncMetadata(asm.Mov.Imm(asm.R0, 0), &btfFn),
				asm.Return(),
			}

			prog, err := ebpf.NewProgramWithOptions(
				&ebpf.ProgramSpec{
					Type:         ebpf.XDP,
					Instructions: insns,
				},
				ebpf.ProgramOptions{
					LogDisabled: true,
				},
			)
			if err != nil {
				return err
			}
			defer prog.Close()

			return probeProgram(&ebpf.ProgramSpec{
				Type:         ebpf.Extension,
				Instructions: insns,
				AttachTarget: prog,
				AttachTo:     btfFn.Name,
			})
		},
	},
	ebpf.LSM: {
		Version: "5.7",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:       ebpf.LSM,
				AttachType: ebpf.AttachLSMMac,
				AttachTo:   "file_mprotect",
				License:    "GPL",
			})
		},
	},
	ebpf.SkLookup: {
		Version: "5.9",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:       ebpf.SkLookup,
				AttachType: ebpf.AttachSkLookup,
			})
		},
	},
	ebpf.Syscall: {
		Version: "5.14",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:  ebpf.Syscall,
				Flags: sys.BPF_F_SLEEPABLE,
			})
		},
	},
	ebpf.Netfilter: {
		Version: "6.4",
		Fn: func() error {
			return probeProgram(&ebpf.ProgramSpec{
				Type:       ebpf.Netfilter,
				AttachType: ebpf.AttachNetfilter,
			})
		},
	},
}

func init() {
	for key, ft := range haveProgramTypeMatrix {
		ft.Name = key.String()
		if ft.Fn == nil {
			key := key
			ft.Fn = func() error { return probeProgram(&ebpf.ProgramSpec{Type: key}) }
		}
	}
}

type helperKey struct {
	typ    ebpf.ProgramType
	helper asm.BuiltinFunc
}

var helperCache = internal.NewFeatureCache(func(key helperKey) *internal.FeatureTest {
	return &internal.FeatureTest{
		Name: fmt.Sprintf("%s for program type %s", key.helper, key.typ),
		Fn: func() error {
			return haveProgramHelper(key.typ, key.helper)
		},
	}
})

func HaveProgramHelper(pt ebpf.ProgramType, helper asm.BuiltinFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func haveProgramHelper(pt ebpf.ProgramType, helper asm.BuiltinFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func logContainsAll(log []string, needles ...string) bool { _ = "STUB: not implemented"; return false }

func helperProbeNotImplemented(pt ebpf.ProgramType) bool { _ = "STUB: not implemented"; return false }
