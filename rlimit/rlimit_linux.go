package rlimit

import (
	"sync"

	"github.com/cilium/ebpf/internal"
)

var (
	unsupportedMemcgAccounting = &internal.UnsupportedFeatureError{
		MinimumVersion: internal.Version{5, 11, 0},
		Name:           "memcg-based accounting for BPF memory",
	}
	haveMemcgAccounting error

	rlimitMu sync.Mutex
)

func init() {

	haveMemcgAccounting = detectMemcgAccounting()
}

func detectMemcgAccounting() error { _ = "STUB: not implemented"; return nil }

func RemoveMemlock() error { _ = "STUB: not implemented"; return nil }
