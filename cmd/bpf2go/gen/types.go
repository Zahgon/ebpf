//go:build !windows

package gen

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/btf"
)

func CollectGlobalTypes(spec *ebpf.CollectionSpec) []btf.Type {
	_ = "STUB: not implemented"
	return nil
}

func collectMapTypes(types []btf.Type, maps map[string]*ebpf.MapSpec) []btf.Type {
	_ = "STUB: not implemented"
	return nil
}

func collectVariableTypes(types []btf.Type, vars map[string]*ebpf.VariableSpec) []btf.Type {
	_ = "STUB: not implemented"
	return nil
}

func addType(types []btf.Type, incoming btf.Type) []btf.Type { _ = "STUB: not implemented"; return nil }

func selectType(t btf.Type) btf.Type { _ = "STUB: not implemented"; return *new(btf.Type) }
