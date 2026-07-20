package btf

import (
	"errors"
	"io"
	"iter"

	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

const btfMagic = 0xeB9F

var (
	ErrNotSupported    = internal.ErrNotSupported
	ErrNotFound        = errors.New("not found")
	ErrNoExtendedInfo  = errors.New("no extended info")
	ErrMultipleMatches = errors.New("multiple matching types")
)

type ID = sys.BTFID

type elfData struct {
	sectionSizes  map[string]uint32
	symbolOffsets map[elfSymbol]uint32
	fixups        map[Type]bool
}

type elfSymbol struct {
	section string
	name    string
}

type Spec struct {
	*decoder

	elf *elfData
}

func LoadSpec(file string) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadSpecFromReader(rd io.ReaderAt) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadSpecAndExtInfosFromReader(rd io.ReaderAt) (*Spec, *ExtInfos, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func symbolOffsets(file *internal.SafeELFFile) (map[elfSymbol]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadSpecFromELF(file *internal.SafeELFFile) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadRawSpec(btf []byte, base *Spec) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (elf *elfData) fixupDatasec(typ Type) error { _ = "STUB: not implemented"; return nil }

func fixupDatasecLayout(ds *Datasec) error { _ = "STUB: not implemented"; return nil }

func (s *Spec) Copy() *Spec { _ = "STUB: not implemented"; return nil }

func (s *Spec) TypeByID(id TypeID) (Type, error) { _ = "STUB: not implemented"; return *new(Type), nil }

func (s *Spec) TypeID(typ Type) (TypeID, error) {
	_ = "STUB: not implemented"
	return *new(TypeID), nil
}

func (s *Spec) AnyTypesByName(name string) ([]Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Spec) AnyTypeByName(name string) (Type, error) {
	_ = "STUB: not implemented"
	return *new(Type), nil
}

func (s *Spec) TypeByName(name string, typ any) error { _ = "STUB: not implemented"; return nil }

func LoadSplitSpec(file string, base *Spec) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadSplitSpecFromReader(r io.ReaderAt, base *Spec) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Spec) All() iter.Seq2[Type, error] { _ = "STUB: not implemented"; return nil }
