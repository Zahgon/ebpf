package btf

import (
	"errors"
	"math"

	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/unix"
)

var haveBTF = internal.NewFeatureTest("BTF", func() error {

	err := probeBTF(&Int{})
	if errors.Is(err, unix.EINVAL) || errors.Is(err, unix.EPERM) {
		return internal.ErrNotSupported
	}
	return err
}, "4.18")

var haveMapBTF = internal.NewFeatureTest("Map BTF (Var/Datasec)", func() error {
	if err := haveBTF(); err != nil {
		return err
	}

	v := &Var{
		Name: "a",
		Type: &Pointer{(*Void)(nil)},
	}

	err := probeBTF(v)
	if errors.Is(err, unix.EINVAL) || errors.Is(err, unix.EPERM) {

		return internal.ErrNotSupported
	}
	return err
}, "5.2")

var haveProgBTF = internal.NewFeatureTest("Program BTF (func/line_info)", func() error {
	if err := haveBTF(); err != nil {
		return err
	}

	fn := &Func{
		Name: "a",
		Type: &FuncProto{Return: (*Void)(nil)},
	}

	err := probeBTF(fn)
	if errors.Is(err, unix.EINVAL) || errors.Is(err, unix.EPERM) {
		return internal.ErrNotSupported
	}
	return err
}, "5.0")

var haveFuncLinkage = internal.NewFeatureTest("BTF func linkage", func() error {
	if err := haveProgBTF(); err != nil {
		return err
	}

	fn := &Func{
		Name:    "a",
		Type:    &FuncProto{Return: (*Void)(nil)},
		Linkage: GlobalFunc,
	}

	err := probeBTF(fn)
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	return err
}, "5.6")

var haveDeclTags = internal.NewFeatureTest("BTF decl tags", func() error {
	if err := haveBTF(); err != nil {
		return err
	}

	t := &Typedef{
		Name: "a",
		Type: &Int{},
		Tags: []string{"a"},
	}

	err := probeBTF(t)
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	return err
}, "5.16")

var haveTypeTags = internal.NewFeatureTest("BTF type tags", func() error {
	if err := haveBTF(); err != nil {
		return err
	}

	t := &TypeTag{
		Type:  &Int{},
		Value: "a",
	}

	err := probeBTF(t)
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	return err
}, "5.17")

var haveEnum64 = internal.NewFeatureTest("ENUM64", func() error {
	if err := haveBTF(); err != nil {
		return err
	}

	enum := &Enum{
		Size: 8,
		Values: []EnumValue{
			{"TEST", math.MaxUint32 + 1},
		},
	}

	err := probeBTF(enum)
	if errors.Is(err, unix.EINVAL) {
		return internal.ErrNotSupported
	}
	return err
}, "6.0")

func probeBTF(typ Type) error { _ = "STUB: not implemented"; return nil }
