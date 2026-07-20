package features

import (
	"errors"
	"fmt"
	"os"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

func HaveMapType(mt ebpf.MapType) error { _ = "STUB: not implemented"; return nil }

func probeCgroupStorageMap(mt sys.MapType) error { _ = "STUB: not implemented"; return nil }

func probeStorageMap(mt sys.MapType) error { _ = "STUB: not implemented"; return nil }

func probeNestedMap(mt sys.MapType) error { _ = "STUB: not implemented"; return nil }

func probeMap(attr *sys.MapCreateAttr) error { _ = "STUB: not implemented"; return nil }

func createMap(attr *sys.MapCreateAttr) error { _ = "STUB: not implemented"; return nil }

var haveMapTypeMatrix = internal.FeatureMatrix[ebpf.MapType]{
	ebpf.Hash:           {Version: "3.19"},
	ebpf.Array:          {Version: "3.19"},
	ebpf.ProgramArray:   {Version: "4.2"},
	ebpf.PerfEventArray: {Version: "4.3"},
	ebpf.PerCPUHash:     {Version: "4.6"},
	ebpf.PerCPUArray:    {Version: "4.6"},
	ebpf.StackTrace: {
		Version: "4.6",
		Fn: func() error {
			return probeMap(&sys.MapCreateAttr{
				MapType:   sys.BPF_MAP_TYPE_STACK_TRACE,
				ValueSize: 8,
			})
		},
	},
	ebpf.CGroupArray: {Version: "4.8"},
	ebpf.LRUHash:     {Version: "4.10"},
	ebpf.LRUCPUHash:  {Version: "4.10"},
	ebpf.LPMTrie: {
		Version: "4.11",
		Fn: func() error {

			return probeMap(&sys.MapCreateAttr{
				MapType:   sys.BPF_MAP_TYPE_LPM_TRIE,
				KeySize:   8,
				ValueSize: 8,
				MapFlags:  sys.BPF_F_NO_PREALLOC,
			})
		},
	},
	ebpf.ArrayOfMaps: {
		Version: "4.12",
		Fn:      func() error { return probeNestedMap(sys.BPF_MAP_TYPE_ARRAY_OF_MAPS) },
	},
	ebpf.HashOfMaps: {
		Version: "4.12",
		Fn:      func() error { return probeNestedMap(sys.BPF_MAP_TYPE_HASH_OF_MAPS) },
	},
	ebpf.DevMap:   {Version: "4.14"},
	ebpf.SockMap:  {Version: "4.14"},
	ebpf.CPUMap:   {Version: "4.15"},
	ebpf.XSKMap:   {Version: "4.18"},
	ebpf.SockHash: {Version: "4.18"},
	ebpf.CGroupStorage: {
		Version: "4.19",
		Fn:      func() error { return probeCgroupStorageMap(sys.BPF_MAP_TYPE_CGROUP_STORAGE) },
	},
	ebpf.ReusePortSockArray: {Version: "4.19"},
	ebpf.PerCPUCGroupStorage: {
		Version: "4.20",
		Fn:      func() error { return probeCgroupStorageMap(sys.BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE) },
	},
	ebpf.Queue: {
		Version: "4.20",
		Fn: func() error {
			return createMap(&sys.MapCreateAttr{
				MapType:    sys.BPF_MAP_TYPE_QUEUE,
				KeySize:    0,
				ValueSize:  4,
				MaxEntries: 1,
			})
		},
	},
	ebpf.Stack: {
		Version: "4.20",
		Fn: func() error {
			return createMap(&sys.MapCreateAttr{
				MapType:    sys.BPF_MAP_TYPE_STACK,
				KeySize:    0,
				ValueSize:  4,
				MaxEntries: 1,
			})
		},
	},
	ebpf.SkStorage: {
		Version: "5.2",
		Fn:      func() error { return probeStorageMap(sys.BPF_MAP_TYPE_SK_STORAGE) },
	},
	ebpf.DevMapHash: {Version: "5.4"},
	ebpf.StructOpsMap: {
		Version: "5.6",
		Fn: func() error {

			err := probeMap(&sys.MapCreateAttr{
				MapType:               sys.BPF_MAP_TYPE_STRUCT_OPS,
				BtfVmlinuxValueTypeId: 1,
			})
			if errors.Is(err, sys.ENOTSUPP) {

				return nil
			}
			return err
		},
	},
	ebpf.RingBuf: {
		Version: "5.8",
		Fn: func() error {

			return createMap(&sys.MapCreateAttr{
				MapType:    sys.BPF_MAP_TYPE_RINGBUF,
				KeySize:    0,
				ValueSize:  0,
				MaxEntries: uint32(os.Getpagesize()),
			})
		},
	},
	ebpf.InodeStorage: {
		Version: "5.10",
		Fn:      func() error { return probeStorageMap(sys.BPF_MAP_TYPE_INODE_STORAGE) },
	},
	ebpf.TaskStorage: {
		Version: "5.11",
		Fn:      func() error { return probeStorageMap(sys.BPF_MAP_TYPE_TASK_STORAGE) },
	},
	ebpf.BloomFilter: {
		Version: "5.16",
		Fn: func() error {
			return createMap(&sys.MapCreateAttr{
				MapType:    sys.BPF_MAP_TYPE_BLOOM_FILTER,
				KeySize:    0,
				ValueSize:  4,
				MaxEntries: 1,
			})
		},
	},
	ebpf.UserRingbuf: {
		Version: "6.1",
		Fn: func() error {

			return createMap(&sys.MapCreateAttr{
				MapType:    sys.BPF_MAP_TYPE_USER_RINGBUF,
				KeySize:    0,
				ValueSize:  0,
				MaxEntries: uint32(os.Getpagesize()),
			})
		},
	},
	ebpf.CgroupStorage: {
		Version: "6.2",
		Fn:      func() error { return probeStorageMap(sys.BPF_MAP_TYPE_CGRP_STORAGE) },
	},
	ebpf.Arena: {
		Version: "6.9",
		Fn: func() error {
			return createMap(&sys.MapCreateAttr{
				MapType:    sys.BPF_MAP_TYPE_ARENA,
				KeySize:    0,
				ValueSize:  0,
				MaxEntries: 1,
				MapExtra:   0,
				MapFlags:   sys.BPF_F_MMAPABLE,
			})
		},
	},
}

func init() {
	for mt, ft := range haveMapTypeMatrix {
		ft.Name = mt.String()
		if ft.Fn == nil {

			mt := sys.MapType(mt)
			ft.Fn = func() error { return probeMap(&sys.MapCreateAttr{MapType: mt}) }
		}
	}
}

type MapFlags uint32

const (
	BPF_F_NO_PREALLOC = sys.BPF_F_NO_PREALLOC
	BPF_F_RDONLY_PROG = sys.BPF_F_RDONLY_PROG
	BPF_F_WRONLY_PROG = sys.BPF_F_WRONLY_PROG
	BPF_F_MMAPABLE    = sys.BPF_F_MMAPABLE
	BPF_F_INNER_MAP   = sys.BPF_F_INNER_MAP
)

func HaveMapFlag(flag MapFlags) (err error) { _ = "STUB: not implemented"; return nil }

func probeMapFlag(attr *sys.MapCreateAttr) error { _ = "STUB: not implemented"; return nil }

var haveMapFlagsMatrix = internal.FeatureMatrix[MapFlags]{
	BPF_F_NO_PREALLOC: {
		Version: "4.6",
		Fn: func() error {
			return probeMapFlag(&sys.MapCreateAttr{
				MapType:  sys.BPF_MAP_TYPE_HASH,
				MapFlags: BPF_F_NO_PREALLOC,
			})
		},
	},
	BPF_F_RDONLY_PROG: {
		Version: "5.2",
		Fn: func() error {
			return probeMapFlag(&sys.MapCreateAttr{
				MapFlags: BPF_F_RDONLY_PROG,
			})
		},
	},
	BPF_F_WRONLY_PROG: {
		Version: "5.2",
		Fn: func() error {
			return probeMapFlag(&sys.MapCreateAttr{
				MapFlags: BPF_F_WRONLY_PROG,
			})
		},
	},
	BPF_F_MMAPABLE: {
		Version: "5.5",
		Fn: func() error {
			return probeMapFlag(&sys.MapCreateAttr{
				MapFlags: BPF_F_MMAPABLE,
			})
		},
	},
	BPF_F_INNER_MAP: {
		Version: "5.10",
		Fn: func() error {
			return probeMapFlag(&sys.MapCreateAttr{
				MapFlags: BPF_F_INNER_MAP,
			})
		},
	},
}

func init() {
	for mf, ft := range haveMapFlagsMatrix {
		ft.Name = fmt.Sprint(mf)
	}
}
