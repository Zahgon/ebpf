package btf

import (
	"errors"
	"hash/maphash"
)

type deduper struct {
	visited   map[Type]struct{}
	hashCache map[hashCacheKey]uint64

	done map[Type]Type

	hashed  map[uint64][]Type
	eqCache map[typKey]bool

	seed maphash.Seed
}

func newDeduper() *deduper { _ = "STUB: not implemented"; return nil }

func (d *deduper) deduplicate(t Type) (Type, error) {
	_ = "STUB: not implemented"
	return *new(Type), nil
}

func (d *deduper) hashInsert(t Type) (Type, error) {
	_ = "STUB: not implemented"
	return *new(Type), nil
}

type hashCacheKey struct {
	t           Type
	depthBudget int
}

func (d *deduper) hash(t Type, depthBudget int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type typKey struct {
	a Type
	b Type
}

var errNotEquivalent = errors.New("types are not equivalent")

func (d *deduper) typesEquivalent(ta, tb Type, visited []Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *deduper) compositeEquivalent(at, bt Type, visited []Type) error {
	_ = "STUB: not implemented"
	return nil
}
