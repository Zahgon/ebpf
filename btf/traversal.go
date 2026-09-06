package btf

import (
	"iter"
)

func postorder(root Type, visited map[Type]struct{}) iter.Seq[Type] {
	_ = "STUB: not implemented"
	return nil
}

func visitInPostorder(root Type, visited map[Type]struct{}, yield func(typ Type) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func children(typ Type) iter.Seq[*Type] { _ = "STUB: not implemented"; return nil }
