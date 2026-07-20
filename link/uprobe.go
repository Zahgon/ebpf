//go:build !windows

package link

import (
	"errors"
	"os"
	"sync"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal"
)

var (
	uprobeRefCtrOffsetPMUPath = "/sys/bus/event_source/devices/uprobe/format/ref_ctr_offset"

	uprobeRefCtrOffsetShift = 32
	haveRefCtrOffsetPMU     = internal.NewFeatureTest("RefCtrOffsetPMU", func() error {
		_, err := os.Stat(uprobeRefCtrOffsetPMUPath)
		if errors.Is(err, os.ErrNotExist) {
			return internal.ErrNotSupported
		}
		if err != nil {
			return err
		}
		return nil
	}, "4.20")

	ErrNoSymbol = errors.New("not found")
)

type Executable struct {
	path string

	cachedSymbols map[string]symbol

	cachedSymbolsOnce sync.Once
}

type symbol struct {
	addr uint64
	size uint64
}

func (s symbol) contains(address uint64) bool { _ = "STUB: not implemented"; return false }

type UprobeOptions struct {
	Address uint64

	Offset uint64

	PID int

	RefCtrOffset uint64

	Cookie uint64

	TraceFSPrefix string
}

func (uo *UprobeOptions) cookie() uint64 { _ = "STUB: not implemented"; return 0 }

func OpenExecutable(path string) (*Executable, error) { _ = "STUB: not implemented"; return nil, nil }

func (ex *Executable) load(f *internal.SafeELFFile) error { _ = "STUB: not implemented"; return nil }

func (ex *Executable) lazyLoadSymbols() error { _ = "STUB: not implemented"; return nil }

func (ex *Executable) address(symbol string, address, offset uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type SymbolOffset struct {
	Symbol string
	Offset uint64
}

func (ex *Executable) Symbol(address uint64) (SymbolOffset, error) {
	_ = "STUB: not implemented"
	return *new(SymbolOffset), nil
}

func (ex *Executable) Uprobe(symbol string, prog *ebpf.Program, opts *UprobeOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (ex *Executable) Uretprobe(symbol string, prog *ebpf.Program, opts *UprobeOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (ex *Executable) uprobe(symbol string, prog *ebpf.Program, opts *UprobeOptions, ret bool) (*perfEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
