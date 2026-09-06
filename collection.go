package ebpf

import (
	"encoding/binary"
	"reflect"

	"github.com/cilium/ebpf/btf"
)

type CollectionOptions struct {
	Maps     MapOptions
	Programs ProgramOptions

	MapReplacements map[string]*Map

	Cache *btf.Cache
}

type CollectionSpec struct {
	Maps     map[string]*MapSpec
	Programs map[string]*ProgramSpec

	Variables map[string]*VariableSpec

	Types *btf.Spec

	ByteOrder binary.ByteOrder
}

func (cs *CollectionSpec) Copy() *CollectionSpec { _ = "STUB: not implemented"; return nil }

func copyMapOfSpecs[T interface{ Copy() T }](m map[string]T) map[string]T {
	_ = "STUB: not implemented"
	return nil
}

func (cs *CollectionSpec) Assign(to any) error { _ = "STUB: not implemented"; return nil }

func (cs *CollectionSpec) LoadAndAssign(to any, opts *CollectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type Collection struct {
	Programs map[string]*Program
	Maps     map[string]*Map

	Variables map[string]*Variable
}

func NewCollection(spec *CollectionSpec) (*Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCollectionWithOptions(spec *CollectionSpec, opts CollectionOptions) (*Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type collectionLoader struct {
	coll     *CollectionSpec
	opts     *CollectionOptions
	maps     map[string]*Map
	programs map[string]*Program
	vars     map[string]*Variable
	types    *btf.Cache
}

func newCollectionLoader(coll *CollectionSpec, opts *CollectionOptions) (*collectionLoader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func populateKallsyms(progs map[string]*ProgramSpec) error { _ = "STUB: not implemented"; return nil }

func (cl *collectionLoader) close() { _ = "STUB: not implemented"; return }

func (cl *collectionLoader) loadMap(mapName string) (*Map, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *collectionLoader) loadProgram(progName string) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *collectionLoader) loadVariable(varName string) (*Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *collectionLoader) populateDeferredMaps() error { _ = "STUB: not implemented"; return nil }

func (cl *collectionLoader) populateStructOps(m *Map, mapSpec *MapSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func resolveKconfig(m *MapSpec) error { _ = "STUB: not implemented"; return nil }

func LoadCollection(file string) (*Collection, error) { _ = "STUB: not implemented"; return nil, nil }

func (coll *Collection) Assign(to any) error { _ = "STUB: not implemented"; return nil }

func (coll *Collection) Close() { _ = "STUB: not implemented"; return }

func (coll *Collection) DetachMap(name string) *Map { _ = "STUB: not implemented"; return nil }

func (coll *Collection) DetachProgram(name string) *Program { _ = "STUB: not implemented"; return nil }

type structField struct {
	reflect.StructField
	value reflect.Value
}

func ebpfFields(structVal reflect.Value, visited map[reflect.Type]bool) ([]structField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func assignValues(to any, getValue func(typ reflect.Type, name string) (any, error)) error {
	_ = "STUB: not implemented"
	return nil
}
