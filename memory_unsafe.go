package ebpf

import (
	"errors"
	"reflect"
	"unsafe"
)

//go:linkname heapObjectsCanMove runtime.heapObjectsCanMove
func heapObjectsCanMove() bool

var unsafeMemory = false

var ErrInvalidType = errors.New("invalid type")

func newUnsafeMemory(fd, size int) (*Memory, error) { _ = "STUB: not implemented"; return nil, nil }

//go:nocheckptr
func allocate(size int) (unsafe.Pointer, error) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), nil
}

func mapmap(fd int, addr unsafe.Pointer, size, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

func unmap(size int) func(*byte) { _ = "STUB: not implemented"; return nil }

func checkUnsafeMemory[T comparable](mm *Memory, off uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func checkType(name string, t reflect.Type) error { _ = "STUB: not implemented"; return nil }

func memoryPointer[T comparable](mm *Memory, off uint32) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
