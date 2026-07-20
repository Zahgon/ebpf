//go:build !windows

package ebpf

func loadCollectionFromNativeImage(_ string) (*Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
