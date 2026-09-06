package ebpf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/platform"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/sysenc"
	"github.com/cilium/ebpf/internal/unix"
)

var ErrNotSupported = internal.ErrNotSupported

var ErrProgIncompatible = errors.New("program is incompatible")

var errBadRelocation = errors.New("bad CO-RE relocation")

var errUnknownKfunc = errors.New("unknown kfunc")

type ProgramID = sys.ProgramID

const (
	outputPad = 256 + 2
)

const minVerifierLogSize = 64 * 1024

const maxVerifierLogSize = math.MaxUint32 >> 2

const maxVerifierAttempts = 30

type ProgramOptions struct {
	LogLevel LogLevel

	LogSizeStart uint32

	LogDisabled bool

	KernelTypes *btf.Spec

	ExtraRelocationTargets []*btf.Spec
}

type ProgramSpec struct {
	Name string

	Type ProgramType

	Ifindex uint32

	AttachType AttachType

	AttachTo string

	AttachTarget *Program

	SectionName string

	Instructions asm.Instructions

	Flags uint32

	License string

	KernelVersion uint32

	ByteOrder binary.ByteOrder
}

func (ps *ProgramSpec) Copy() *ProgramSpec { _ = "STUB: not implemented"; return nil }

func (ps *ProgramSpec) Tag() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (ps *ProgramSpec) Compatible(info *ProgramInfo) error { _ = "STUB: not implemented"; return nil }

func (ps *ProgramSpec) targetsKernelModule() bool { _ = "STUB: not implemented"; return false }

type VerifierError = internal.VerifierError

type Program struct {
	VerifierLog string

	fd         *sys.FD
	name       string
	pinnedPath string
	typ        ProgramType

	btf *btf.Handle
}

func NewProgram(spec *ProgramSpec) (*Program, error) { _ = "STUB: not implemented"; return nil, nil }

func NewProgramWithOptions(spec *ProgramSpec, opts ProgramOptions) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	coreBadLoad = fmt.Appendf(nil, "(18) r10 = 0x%x\n", btf.COREBadRelocationSentinel)

	coreBadCall  = fmt.Appendf(nil, "invalid func unknown#%d\n", btf.COREBadRelocationSentinel)
	kfuncBadCall = fmt.Appendf(nil, "invalid func unknown#%d\n", kfuncCallPoisonBase)
)

func newProgramWithOptions(spec *ProgramSpec, opts ProgramOptions, c *btf.Cache) (result *Program, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func programLoadError(err error, logBuf []byte, spec *ProgramSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func retryLogAttrs(attr *sys.ProgLoadAttr, startSize uint32, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func NewProgramFromFD(fd int) (*Program, error) { _ = "STUB: not implemented"; return nil, nil }

func NewProgramFromID(id ProgramID) (*Program, error) { _ = "STUB: not implemented"; return nil, nil }

func newProgramFromFD(fd *sys.FD) (*Program, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Program) String() string { _ = "STUB: not implemented"; return "" }

func (p *Program) Type() ProgramType { _ = "STUB: not implemented"; return *new(ProgramType) }

func (p *Program) Info() (*ProgramInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Program) Stats() (*ProgramStats, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Program) Handle() (*btf.Handle, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Program) SetHandle(h *btf.Handle) error { _ = "STUB: not implemented"; return nil }

func (p *Program) FD() int { _ = "STUB: not implemented"; return 0 }

func (p *Program) Clone() (*Program, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Program) Pin(fileName string) error { _ = "STUB: not implemented"; return nil }

func (p *Program) Unpin() error { _ = "STUB: not implemented"; return nil }

func (p *Program) IsPinned() bool { _ = "STUB: not implemented"; return false }

func (p *Program) Close() error { _ = "STUB: not implemented"; return nil }

type RunOptions struct {
	Data []byte

	DataOut []byte

	Context any

	ContextOut any

	Repeat uint32

	Flags uint32

	CPU uint32

	BatchSize uint32

	Reset func()
}

func (p *Program) Test(in []byte) (uint32, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (p *Program) Run(opts *RunOptions) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *Program) Benchmark(in []byte, repeat int, reset func()) (uint32, time.Duration, error) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), nil
}

var haveProgRun = internal.NewFeatureTest("BPF_PROG_RUN", func() error {
	if platform.IsWindows {
		return nil
	}

	prog, err := NewProgram(&ProgramSpec{

		Type: SocketFilter,
		Instructions: asm.Instructions{
			asm.LoadImm(asm.R0, 0, asm.DWord),
			asm.Return(),
		},
		License: "MIT",
	})
	if err != nil {

		return err
	}
	defer prog.Close()

	in := internal.EmptyBPFContext
	attr := sys.ProgRunAttr{
		ProgFd:     uint32(prog.FD()),
		DataSizeIn: uint32(len(in)),
		DataIn:     sys.SlicePointer(in),
	}

	err = sys.ProgRun(&attr)
	switch {
	case errors.Is(err, unix.EINVAL):

		return internal.ErrNotSupported

	case errors.Is(err, unix.EINTR):

		return nil

	case errors.Is(err, sys.ENOTSUPP):

		return nil
	}

	return err
}, "4.12", "windows:0.20")

func (p *Program) run(opts *RunOptions) (uint32, time.Duration, error) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), nil
}

func unmarshalProgram(buf sysenc.Buffer) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalProgram(p *Program, length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadPinnedProgram(fileName string, opts *LoadPinOptions) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProgramGetNextID(startID ProgramID) (ProgramID, error) {
	_ = "STUB: not implemented"
	return *new(ProgramID), nil
}

func (p *Program) BindMap(m *Map) error { _ = "STUB: not implemented"; return nil }

var errUnrecognizedAttachType = errors.New("unrecognized attach type")

func findProgramTargetInKernel(name string, progType ProgramType, attachType AttachType, cache *btf.Cache) (*btf.Handle, btf.TypeID, error) {
	_ = "STUB: not implemented"
	return nil, *new(btf.TypeID), nil
}

func findTargetInKernel[T btf.Type](typeName string, target *T, cache *btf.Cache) (*btf.Spec, *btf.Handle, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func findTargetInModule[T btf.Type](typeName string, target *T, cache *btf.Cache) (*btf.Spec, *btf.Handle, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func findTargetInProgram(prog *Program, name string, progType ProgramType, attachType AttachType) (btf.TypeID, error) {
	_ = "STUB: not implemented"
	return *new(btf.TypeID), nil
}
