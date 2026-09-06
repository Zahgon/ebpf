package ebpf

import (
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/sysenc"
)

func marshalMapSyscallInput(data any, length int) (sys.Pointer, error) {
	_ = "STUB: not implemented"
	return *new(sys.Pointer), nil
}

func makeMapSyscallOutput(dst any, length int) sysenc.Buffer {
	_ = "STUB: not implemented"
	return *new(sysenc.Buffer)
}

func appendPerCPUSlice(buf []byte, slice any, possibleCPUs, elemLength, alignedElemLength int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalPerCPUValue(slice any, elemLength int) (sys.Pointer, error) {
	_ = "STUB: not implemented"
	return *new(sys.Pointer), nil
}

func marshalBatchPerCPUValue(slice any, batchLen, elemLength int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalPerCPUValue(slice any, elemLength int, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalBatchPerCPUValue(slice any, batchLen, elemLength int, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}
