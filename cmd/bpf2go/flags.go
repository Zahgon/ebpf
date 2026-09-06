//go:build !windows

package main

import (
	"flag"
	"go/build/constraint"
)

type buildTags struct {
	Expr constraint.Expr
}

var _ flag.Value = (*buildTags)(nil)

func (bt *buildTags) String() string { _ = "STUB: not implemented"; return "" }

func (bt *buildTags) Set(value string) error { _ = "STUB: not implemented"; return nil }

func andConstraints(x, y constraint.Expr) constraint.Expr {
	_ = "STUB: not implemented"
	return *new(constraint.Expr)
}
