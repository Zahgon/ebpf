package btf

import (
	"os"
	"sync"
)

func LoadKernelSpec() (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadKernelModuleSpec(module string) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func loadKernelModuleSpec(module string, base *Spec) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findVMLinux() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

type Cache struct {
	mu            sync.RWMutex
	kernelTypes   *Spec
	moduleTypes   map[string]*Spec
	loadedModules []string
}

func NewCache() *Cache { _ = "STUB: not implemented"; return nil }

func (c *Cache) Kernel() (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cache) kernel() (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cache) Module(name string) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cache) Modules() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
