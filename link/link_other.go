//go:build !windows

package link

import (
	"github.com/cilium/ebpf/internal/sys"
)

const (
	UnspecifiedType   = sys.BPF_LINK_TYPE_UNSPEC
	RawTracepointType = sys.BPF_LINK_TYPE_RAW_TRACEPOINT
	TracingType       = sys.BPF_LINK_TYPE_TRACING
	CgroupType        = sys.BPF_LINK_TYPE_CGROUP
	IterType          = sys.BPF_LINK_TYPE_ITER
	NetNsType         = sys.BPF_LINK_TYPE_NETNS
	XDPType           = sys.BPF_LINK_TYPE_XDP
	PerfEventType     = sys.BPF_LINK_TYPE_PERF_EVENT
	KprobeMultiType   = sys.BPF_LINK_TYPE_KPROBE_MULTI
	TCXType           = sys.BPF_LINK_TYPE_TCX
	UprobeMultiType   = sys.BPF_LINK_TYPE_UPROBE_MULTI
	NetfilterType     = sys.BPF_LINK_TYPE_NETFILTER
	NetkitType        = sys.BPF_LINK_TYPE_NETKIT
	StructOpsType     = sys.BPF_LINK_TYPE_STRUCT_OPS
)

func AttachRawLink(opts RawLinkOptions) (*RawLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapRawLink(raw *RawLink) (_ Link, err error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

type TracingInfo struct {
	AttachType     sys.AttachType
	TargetObjectId uint32
	TargetBtfId    sys.TypeID
}

type CgroupInfo struct {
	CgroupId   uint64
	AttachType sys.AttachType
	_          [4]byte
}

type NetNsInfo struct {
	NetnsInode uint32
	AttachType sys.AttachType
}

type TCXInfo struct {
	Ifindex    uint32
	AttachType sys.AttachType
}

type XDPInfo struct {
	Ifindex uint32
}

type NetfilterInfo struct {
	ProtocolFamily NetfilterProtocolFamily
	Hook           NetfilterInetHook
	Priority       int32
	Flags          uint32
}

type NetkitInfo struct {
	Ifindex    uint32
	AttachType sys.AttachType
}

type RawTracepointInfo struct {
	Name string
}

type IterInfo struct {
	TargetName string
}

type KprobeMultiInfo struct {
	Count   uint32
	Flags   uint32
	Missed  uint64
	addrs   []uint64
	cookies []uint64
}

type KprobeMultiAddress struct {
	Address uint64
	Cookie  uint64
}

func (kpm *KprobeMultiInfo) Addresses() ([]KprobeMultiAddress, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type UprobeMultiInfo struct {
	Count         uint32
	Flags         uint32
	Missed        uint64
	offsets       []uint64
	cookies       []uint64
	refCtrOffsets []uint64

	File string
	pid  uint32
}

type UprobeMultiOffset struct {
	Offset         uint64
	Cookie         uint64
	ReferenceCount uint64
}

func (umi *UprobeMultiInfo) Offsets() ([]UprobeMultiOffset, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (umi *UprobeMultiInfo) Pid() (uint32, bool) { _ = "STUB: not implemented"; return 0, false }

const (
	PerfEventUnspecified = sys.BPF_PERF_EVENT_UNSPEC
	PerfEventUprobe      = sys.BPF_PERF_EVENT_UPROBE
	PerfEventUretprobe   = sys.BPF_PERF_EVENT_URETPROBE
	PerfEventKprobe      = sys.BPF_PERF_EVENT_KPROBE
	PerfEventKretprobe   = sys.BPF_PERF_EVENT_KRETPROBE
	PerfEventTracepoint  = sys.BPF_PERF_EVENT_TRACEPOINT
	PerfEventEvent       = sys.BPF_PERF_EVENT_EVENT
)

type PerfEventInfo struct {
	Type  sys.PerfEventType
	extra any
}

func (r *PerfEventInfo) Kprobe() *KprobeInfo { _ = "STUB: not implemented"; return nil }

func (r *PerfEventInfo) Uprobe() *UprobeInfo { _ = "STUB: not implemented"; return nil }

func (r *PerfEventInfo) Tracepoint() *TracepointInfo { _ = "STUB: not implemented"; return nil }

func (r *PerfEventInfo) Event() *EventInfo { _ = "STUB: not implemented"; return nil }

type KprobeInfo struct {
	Address  uint64
	Missed   uint64
	Function string
	Offset   uint32
}

type UprobeInfo struct {
	File                 string
	Offset               uint32
	Cookie               uint64
	OffsetReferenceCount uint64
}

type TracepointInfo struct {
	Tracepoint string
	Cookie     uint64
}

type EventInfo struct {
	Config uint64
	Type   uint32
	Cookie uint64
}

func (r Info) Tracing() *TracingInfo { _ = "STUB: not implemented"; return nil }

func (r Info) Cgroup() *CgroupInfo { _ = "STUB: not implemented"; return nil }

func (r Info) NetNs() *NetNsInfo { _ = "STUB: not implemented"; return nil }

func (r Info) XDP() *XDPInfo { _ = "STUB: not implemented"; return nil }

func (r Info) TCX() *TCXInfo { _ = "STUB: not implemented"; return nil }

func (r Info) Netfilter() *NetfilterInfo { _ = "STUB: not implemented"; return nil }

func (r Info) Netkit() *NetkitInfo { _ = "STUB: not implemented"; return nil }

func (r Info) KprobeMulti() *KprobeMultiInfo { _ = "STUB: not implemented"; return nil }

func (r Info) UprobeMulti() *UprobeMultiInfo { _ = "STUB: not implemented"; return nil }

func (r Info) PerfEvent() *PerfEventInfo { _ = "STUB: not implemented"; return nil }

func (r Info) RawTracepoint() *RawTracepointInfo { _ = "STUB: not implemented"; return nil }

func (r Info) Iter() *IterInfo { _ = "STUB: not implemented"; return nil }
