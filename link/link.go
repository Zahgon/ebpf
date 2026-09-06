package link

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

type Type = sys.LinkType

var ErrNotSupported = internal.ErrNotSupported

type Link interface {
	Update(*ebpf.Program) error

	Pin(string) error

	Unpin() error

	Close() error

	Detach() error

	Info() (*Info, error)

	isLink()
}

func NewFromFD(fd int) (Link, error) { _ = "STUB: not implemented"; return *new(Link), nil }

func NewFromID(id ID) (Link, error) { _ = "STUB: not implemented"; return *new(Link), nil }

func LoadPinnedLink(fileName string, opts *ebpf.LoadPinOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type ID = sys.LinkID

type RawLinkOptions struct {
	Target int

	Program *ebpf.Program

	Attach ebpf.AttachType

	BTF btf.TypeID

	Flags uint32
}

type Info struct {
	Type    Type
	ID      ID
	Program ebpf.ProgramID
	extra   any
}

type RawLink struct {
	fd         *sys.FD
	pinnedPath string
}

func loadPinnedRawLink(fileName string, opts *ebpf.LoadPinOptions) (*RawLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *RawLink) isLink() { _ = "STUB: not implemented"; return }

func (l *RawLink) FD() int { _ = "STUB: not implemented"; return 0 }

func (l *RawLink) Close() error { _ = "STUB: not implemented"; return nil }

func (l *RawLink) Pin(fileName string) error { _ = "STUB: not implemented"; return nil }

func (l *RawLink) Unpin() error { _ = "STUB: not implemented"; return nil }

func (l *RawLink) IsPinned() bool { _ = "STUB: not implemented"; return false }

func (l *RawLink) Update(new *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

type RawLinkUpdateOptions struct {
	New   *ebpf.Program
	Old   *ebpf.Program
	Flags uint32
}

func (l *RawLink) UpdateArgs(opts RawLinkUpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *RawLink) Detach() error { _ = "STUB: not implemented"; return nil }

func (l *RawLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

type Iterator struct {
	ID ID

	Link Link
	err  error
}

func (it *Iterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *Iterator) Take() Link { _ = "STUB: not implemented"; return *new(Link) }

func (it *Iterator) Err() error { _ = "STUB: not implemented"; return nil }

func (it *Iterator) Close() { _ = "STUB: not implemented"; return }
