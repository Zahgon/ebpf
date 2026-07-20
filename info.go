package ebpf

import (
	"errors"
	"io"
	"time"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/platform"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/unix"
)

type MapInfo struct {
	Type MapType

	KeySize uint32

	ValueSize uint32

	MaxEntries uint32

	Flags uint32

	Name string

	id       MapID
	btf      btf.ID
	mapExtra uint64
	memlock  uint64
	frozen   bool
}

func minimalMapInfoFromFd(fd *sys.FD) (*MapInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func newMapInfoFromFd(fd *sys.FD) (*MapInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func readMapInfoFromProc(fd *sys.FD, mi *MapInfo) error { _ = "STUB: not implemented"; return nil }

func (mi *MapInfo) ID() (MapID, bool) { _ = "STUB: not implemented"; return *new(MapID), false }

func (mi *MapInfo) BTFID() (btf.ID, bool) { _ = "STUB: not implemented"; return *new(btf.ID), false }

func (mi *MapInfo) MapExtra() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (mi *MapInfo) Memlock() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (mi *MapInfo) Frozen() bool { _ = "STUB: not implemented"; return false }

type ProgramStats struct {
	Runtime time.Duration

	RunCount uint64

	RecursionMisses uint64
}

func newProgramStatsFromFd(fd *sys.FD) (*ProgramStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type programJitedInfo struct {
	ksyms    []uint64
	numKsyms uint32

	insns    []byte
	numInsns uint32

	lineInfos    []uint64
	numLineInfos uint32

	lineInfoRecSize uint32

	funcLens    []uint32
	numFuncLens uint32
}

type ProgramInfo struct {
	Type ProgramType
	id   ProgramID

	Tag string

	Name string

	createdByUID     uint32
	haveCreatedByUID bool
	btf              btf.ID
	loadTime         time.Duration

	restricted bool

	maps                 []MapID
	insns                []byte
	numInsns             uint32
	jitedSize            uint32
	verifiedInstructions uint32

	jitedInfo programJitedInfo

	lineInfos    []byte
	numLineInfos uint32
	funcInfos    []byte
	numFuncInfos uint32

	memlock uint64
}

func minimalProgramInfoFromFd(fd *sys.FD) (*ProgramInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newProgramInfoFromFd(fd *sys.FD) (*ProgramInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readNameFromFunc(pi *ProgramInfo) (string, error) { _ = "STUB: not implemented"; return "", nil }

func readProgramInfoFromProc(fd *sys.FD, pi *ProgramInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (pi *ProgramInfo) ID() (ProgramID, bool) {
	_ = "STUB: not implemented"
	return *new(ProgramID), false
}

func (pi *ProgramInfo) CreatedByUID() (uint32, bool) { _ = "STUB: not implemented"; return 0, false }

func (pi *ProgramInfo) BTFID() (btf.ID, bool) {
	_ = "STUB: not implemented"
	return *new(btf.ID), false
}

func (pi *ProgramInfo) btfSpec() (*btf.Spec, error) { _ = "STUB: not implemented"; return nil, nil }

var ErrRestrictedKernel = internal.ErrRestrictedKernel

func (pi *ProgramInfo) LineInfos() (btf.LineOffsets, error) {
	_ = "STUB: not implemented"
	return *new(btf.LineOffsets), nil
}

func (pi *ProgramInfo) Instructions() (asm.Instructions, error) {
	_ = "STUB: not implemented"
	return *new(asm.Instructions), nil
}

func (pi *ProgramInfo) JitedSize() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (pi *ProgramInfo) TranslatedSize() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (pi *ProgramInfo) MapIDs() ([]MapID, bool) { _ = "STUB: not implemented"; return nil, false }

func (pi *ProgramInfo) LoadTime() (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

func (pi *ProgramInfo) VerifiedInstructions() (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (pi *ProgramInfo) JitedKsymAddrs() ([]uintptr, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pi *ProgramInfo) JitedInsns() ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func (pi *ProgramInfo) JitedLineInfos() ([]uint64, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pi *ProgramInfo) JitedFuncLens() ([]uint32, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pi *ProgramInfo) FuncInfos() (btf.FuncOffsets, error) {
	_ = "STUB: not implemented"
	return *new(btf.FuncOffsets), nil
}

func (pi *ProgramInfo) Memlock() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func scanFdInfo(fd *sys.FD, fields map[string]any) error { _ = "STUB: not implemented"; return nil }

func scanFdInfoReader(r io.Reader, fields map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func zero(arg any) bool { _ = "STUB: not implemented"; return false }

func EnableStats(which uint32) (io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), nil
}

var haveProgramInfoMapIDs = internal.NewFeatureTest("map IDs in program info", func() error {
	if platform.IsWindows {

		return nil
	}

	prog, err := progLoad(asm.Instructions{
		asm.LoadImm(asm.R0, 0, asm.DWord),
		asm.Return(),
	}, SocketFilter, "MIT")
	if err != nil {
		return err
	}
	defer prog.Close()

	err = sys.ObjInfo(prog, &sys.ProgInfo{

		NrMapIds: 1,
	})
	if errors.Is(err, unix.EINVAL) {

		return internal.ErrNotSupported
	}
	if errors.Is(err, unix.E2BIG) {

		return internal.ErrNotSupported
	}

	return err
}, "4.15", "windows:0.21.0")
