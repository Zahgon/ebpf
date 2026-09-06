//go:build !windows

package pin

import (
	"iter"

	"github.com/cilium/ebpf"
)

func WalkDir(root string, opts *ebpf.LoadPinOptions) iter.Seq2[*Pin, error] {
	_ = "STUB: not implemented"
	return nil
}
