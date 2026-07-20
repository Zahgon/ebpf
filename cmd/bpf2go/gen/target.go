//go:build !windows

package gen

import (
	"errors"
	"go/build/constraint"
)

var ErrInvalidTarget = errors.New("unsupported target")

var targetsByGoArch = map[GoArch]Target{
	"386":      {"bpfel", "x86", ""},
	"amd64":    {"bpfel", "x86", ""},
	"arm":      {"bpfel", "arm", ""},
	"arm64":    {"bpfel", "arm64", ""},
	"loong64":  {"bpfel", "loongarch", ""},
	"mips":     {"bpfeb", "mips", ""},
	"mipsle":   {"bpfel", "", ""},
	"mips64":   {"bpfeb", "", ""},
	"mips64le": {"bpfel", "", ""},
	"ppc64":    {"bpfeb", "powerpc", ""},
	"ppc64le":  {"bpfel", "powerpc", ""},
	"riscv64":  {"bpfel", "riscv", ""},
	"s390x":    {"bpfeb", "s390", ""},
	"wasm":     {"bpfel", "", "js"},
}

type Target struct {
	clang string

	linux string

	goos string
}

func TargetsByGoArch() map[GoArch]Target { _ = "STUB: not implemented"; return nil }

func (tgt *Target) IsGeneric() bool { _ = "STUB: not implemented"; return false }

func (tgt *Target) Suffix() string { _ = "STUB: not implemented"; return "" }

func (tgt *Target) ObsoleteSuffix() string { _ = "STUB: not implemented"; return "" }

type GoArch string

type GoArches []GoArch

func (arches GoArches) Constraint() constraint.Expr {
	_ = "STUB: not implemented"
	return *new(constraint.Expr)
}

func FindTarget(id string) (Target, GoArches, error) {
	_ = "STUB: not implemented"
	return *new(Target), *new(GoArches), nil
}

func orConstraints(x, y constraint.Expr) constraint.Expr {
	_ = "STUB: not implemented"
	return *new(constraint.Expr)
}
