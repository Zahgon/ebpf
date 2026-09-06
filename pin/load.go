package pin

import (
	"io"

	"github.com/cilium/ebpf"
)

type Pinner interface {
	io.Closer
	Pin(string) error
}

func Load(path string, opts *ebpf.LoadPinOptions) (Pinner, error) {
	_ = "STUB: not implemented"
	return *new(Pinner), nil
}
