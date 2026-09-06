package ebpf

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/internal/sysenc"
	"github.com/cilium/ebpf/internal/unix"
)

var (
	ErrKeyNotExist      = errors.New("key does not exist")
	ErrKeyExist         = errors.New("key already exists")
	ErrIterationAborted = errors.New("iteration aborted")
	ErrMapIncompatible  = errors.New("map spec is incompatible with existing map")

	errMapLookupKeyNotExist = fmt.Errorf("lookup: %w", sysErrKeyNotExist)
)

type MapOptions struct {
	PinPath        string
	LoadPinOptions LoadPinOptions
}

type MapID = sys.MapID

type MapSpec struct {
	Name       string
	Type       MapType
	KeySize    uint32
	ValueSize  uint32
	MaxEntries uint32

	Flags uint32

	Pinning PinType

	NumaNode uint32

	Contents []MapKV

	InnerMap *MapSpec

	MapExtra uint64

	Extra *bytes.Reader

	Key, Value btf.Type

	Tags []string
}

func (ms *MapSpec) String() string { _ = "STUB: not implemented"; return "" }

func (ms *MapSpec) Copy() *MapSpec { _ = "STUB: not implemented"; return nil }

func (spec *MapSpec) fixupMagicFields() (*MapSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ms *MapSpec) dataSection() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms *MapSpec) updateDataSection(vars map[string]*VariableSpec, sectionName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *MapSpec) readOnly() bool { _ = "STUB: not implemented"; return false }

func (ms *MapSpec) writeOnly() bool { _ = "STUB: not implemented"; return false }

type MapKV struct {
	Key   any
	Value any
}

func (ms *MapSpec) Compatible(m *Map) error { _ = "STUB: not implemented"; return nil }

type Map struct {
	name       string
	fd         *sys.FD
	typ        MapType
	keySize    uint32
	valueSize  uint32
	maxEntries uint32
	flags      uint32
	pinnedPath string

	fullValueSize int

	memory *Memory
}

func NewMapFromFD(fd int) (*Map, error) { _ = "STUB: not implemented"; return nil, nil }

func newMapFromFD(fd *sys.FD) (*Map, error) { _ = "STUB: not implemented"; return nil, nil }

func NewMap(spec *MapSpec) (*Map, error) { _ = "STUB: not implemented"; return nil, nil }

func NewMapWithOptions(spec *MapSpec, opts MapOptions) (*Map, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newMapWithOptions(spec *MapSpec, opts MapOptions, c *btf.Cache) (_ *Map, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Map) Memory() (*Memory, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) unsafeMemory() (*Memory, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) memorySize() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (spec *MapSpec) createMap(inner *sys.FD, c *btf.Cache) (_ *Map, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleMapCreateError(attr sys.MapCreateAttr, spec *MapSpec, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func newMapFromParts(fd *sys.FD, name string, typ MapType, keySize, valueSize, maxEntries, flags uint32) (*Map, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Map) String() string { _ = "STUB: not implemented"; return "" }

func (m *Map) Type() MapType { _ = "STUB: not implemented"; return *new(MapType) }

func (m *Map) KeySize() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Map) ValueSize() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Map) MaxEntries() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Map) Flags() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Map) Info() (*MapInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) Handle() (*btf.Handle, error) { _ = "STUB: not implemented"; return nil, nil }

type MapLookupFlags uint64

const LookupLock MapLookupFlags = sys.BPF_F_LOCK

func (m *Map) Lookup(key, valueOut any) error { _ = "STUB: not implemented"; return nil }

func (m *Map) LookupWithFlags(key, valueOut any, flags MapLookupFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) LookupAndDelete(key, valueOut any) error { _ = "STUB: not implemented"; return nil }

func (m *Map) LookupAndDeleteWithFlags(key, valueOut any, flags MapLookupFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) LookupBytes(key any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) lookupPerCPU(key, valueOut any, flags MapLookupFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) lookup(key any, valueOut sys.Pointer, flags MapLookupFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) lookupAndDeletePerCPU(key, valueOut any, flags MapLookupFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func ensurePerCPUSlice(sliceOrPtr any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (m *Map) lookupAndDelete(key any, valuePtr sys.Pointer, flags MapLookupFlags) error {
	_ = "STUB: not implemented"
	return nil
}

type MapUpdateFlags uint64

const (
	UpdateAny MapUpdateFlags = iota

	UpdateNoExist MapUpdateFlags = 1 << (iota - 1)

	UpdateExist

	UpdateLock
)

func (m *Map) Put(key, value any) error { _ = "STUB: not implemented"; return nil }

func (m *Map) Update(key, value any, flags MapUpdateFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) updatePerCPU(key, value any, flags MapUpdateFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) update(key any, valuePtr sys.Pointer, flags MapUpdateFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map) Delete(key any) error { _ = "STUB: not implemented"; return nil }

func (m *Map) NextKey(key, nextKeyOut any) error { _ = "STUB: not implemented"; return nil }

func (m *Map) NextKeyBytes(key any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) nextKey(key any, nextKeyOut sys.Pointer) error { _ = "STUB: not implemented"; return nil }

var mmapProtectedPage = sync.OnceValues(func() ([]byte, error) {
	return unix.Mmap(-1, 0, os.Getpagesize(), unix.PROT_NONE, unix.MAP_ANON|unix.MAP_SHARED)
})

func (m *Map) guessNonExistentKey() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) BatchLookup(cursor *MapBatchCursor, keysOut, valuesOut any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) BatchLookupAndDelete(cursor *MapBatchCursor, keysOut, valuesOut any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type MapBatchCursor struct {
	m      *Map
	opaque []byte
}

func (m *Map) batchLookup(cmd sys.Cmd, cursor *MapBatchCursor, keysOut, valuesOut any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) batchLookupPerCPU(cmd sys.Cmd, cursor *MapBatchCursor, keysOut, valuesOut any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) batchLookupCmd(cmd sys.Cmd, cursor *MapBatchCursor, count int, keysOut any, valuePtr sys.Pointer, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) BatchUpdate(keys, values any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) batchUpdate(count int, keys any, valuePtr sys.Pointer, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) batchUpdatePerCPU(keys, values any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Map) BatchDelete(keys any, opts *BatchOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func batchCount(keys, values any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Map) Iterate() *MapIterator { _ = "STUB: not implemented"; return nil }

func (m *Map) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Map) FD() int { _ = "STUB: not implemented"; return 0 }

func (m *Map) Clone() (*Map, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map) Pin(fileName string) error { _ = "STUB: not implemented"; return nil }

func (m *Map) Unpin() error { _ = "STUB: not implemented"; return nil }

func (m *Map) IsPinned() bool { _ = "STUB: not implemented"; return false }

func (m *Map) Freeze() error { _ = "STUB: not implemented"; return nil }

func (m *Map) finalize(spec *MapSpec) error { _ = "STUB: not implemented"; return nil }

func (m *Map) marshalKey(data any) (sys.Pointer, error) {
	_ = "STUB: not implemented"
	return *new(sys.Pointer), nil
}

func (m *Map) marshalValue(data any) (sys.Pointer, error) {
	_ = "STUB: not implemented"
	return *new(sys.Pointer), nil
}

func (m *Map) unmarshalValue(value any, buf sysenc.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadPinnedMap(fileName string, opts *LoadPinOptions) (*Map, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalMap(buf sysenc.Buffer) (*Map, error) { _ = "STUB: not implemented"; return nil, nil }

func marshalMap(m *Map, length int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type MapIterator struct {
	target *Map

	cursor            any
	count, maxEntries uint32
	done              bool
	err               error
}

func newMapIterator(target *Map) *MapIterator { _ = "STUB: not implemented"; return nil }

func (mi *MapIterator) Next(keyOut, valueOut any) bool { _ = "STUB: not implemented"; return false }

func (mi *MapIterator) Err() error { _ = "STUB: not implemented"; return nil }

func MapGetNextID(startID MapID) (MapID, error) { _ = "STUB: not implemented"; return *new(MapID), nil }

func NewMapFromID(id MapID) (*Map, error) { _ = "STUB: not implemented"; return nil, nil }

func sliceLen(slice any) (int, error) { _ = "STUB: not implemented"; return 0, nil }
