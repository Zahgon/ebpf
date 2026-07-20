package ebpf

import (
	"github.com/cilium/ebpf/internal/platform"
	"github.com/cilium/ebpf/internal/sys"
)

//go:generate go tool stringer -output types_string.go -type=MapType,ProgramType,PinType

type MapType uint32

const (
	UnspecifiedMap MapType = MapType(platform.LinuxTag | iota)

	Hash

	Array

	ProgramArray

	PerfEventArray

	PerCPUHash

	PerCPUArray

	StackTrace

	CGroupArray

	LRUHash

	LRUCPUHash

	LPMTrie

	ArrayOfMaps

	HashOfMaps

	DevMap

	SockMap

	CPUMap

	XSKMap

	SockHash

	CGroupStorage

	ReusePortSockArray

	PerCPUCGroupStorage

	Queue

	Stack

	SkStorage

	DevMapHash

	StructOpsMap

	RingBuf

	InodeStorage

	TaskStorage

	BloomFilter

	UserRingbuf

	CgroupStorage

	Arena
)

const (
	WindowsHash MapType = MapType(platform.WindowsTag | iota + 1)
	WindowsArray
	WindowsProgramArray
	WindowsPerCPUHash
	WindowsPerCPUArray
	WindowsHashOfMaps
	WindowsArrayOfMaps
	WindowsLRUHash
	WindowsLPMTrie
	WindowsQueue
	WindowsLRUCPUHash
	WindowsStack
	WindowsRingBuf
)

func MapTypeForPlatform(plat string, typ uint32) (MapType, error) {
	_ = "STUB: not implemented"
	return *new(MapType), nil
}

func (mt MapType) hasPerCPUValue() bool { _ = "STUB: not implemented"; return false }

func (mt MapType) canStoreMapOrProgram() bool { _ = "STUB: not implemented"; return false }

func (mt MapType) canStoreMap() bool { _ = "STUB: not implemented"; return false }

func (mt MapType) canStoreProgram() bool { _ = "STUB: not implemented"; return false }

func (mt MapType) canHaveValueSize() bool { _ = "STUB: not implemented"; return false }

func (mt MapType) mustHaveNoPrealloc() bool { _ = "STUB: not implemented"; return false }

func (mt MapType) mustHaveZeroMaxEntries() bool { _ = "STUB: not implemented"; return false }

type ProgramType uint32

const (
	UnspecifiedProgram    = ProgramType(sys.BPF_PROG_TYPE_UNSPEC)
	SocketFilter          = ProgramType(sys.BPF_PROG_TYPE_SOCKET_FILTER)
	Kprobe                = ProgramType(sys.BPF_PROG_TYPE_KPROBE)
	SchedCLS              = ProgramType(sys.BPF_PROG_TYPE_SCHED_CLS)
	SchedACT              = ProgramType(sys.BPF_PROG_TYPE_SCHED_ACT)
	TracePoint            = ProgramType(sys.BPF_PROG_TYPE_TRACEPOINT)
	XDP                   = ProgramType(sys.BPF_PROG_TYPE_XDP)
	PerfEvent             = ProgramType(sys.BPF_PROG_TYPE_PERF_EVENT)
	CGroupSKB             = ProgramType(sys.BPF_PROG_TYPE_CGROUP_SKB)
	CGroupSock            = ProgramType(sys.BPF_PROG_TYPE_CGROUP_SOCK)
	LWTIn                 = ProgramType(sys.BPF_PROG_TYPE_LWT_IN)
	LWTOut                = ProgramType(sys.BPF_PROG_TYPE_LWT_OUT)
	LWTXmit               = ProgramType(sys.BPF_PROG_TYPE_LWT_XMIT)
	SockOps               = ProgramType(sys.BPF_PROG_TYPE_SOCK_OPS)
	SkSKB                 = ProgramType(sys.BPF_PROG_TYPE_SK_SKB)
	CGroupDevice          = ProgramType(sys.BPF_PROG_TYPE_CGROUP_DEVICE)
	SkMsg                 = ProgramType(sys.BPF_PROG_TYPE_SK_MSG)
	RawTracepoint         = ProgramType(sys.BPF_PROG_TYPE_RAW_TRACEPOINT)
	CGroupSockAddr        = ProgramType(sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR)
	LWTSeg6Local          = ProgramType(sys.BPF_PROG_TYPE_LWT_SEG6LOCAL)
	LircMode2             = ProgramType(sys.BPF_PROG_TYPE_LIRC_MODE2)
	SkReuseport           = ProgramType(sys.BPF_PROG_TYPE_SK_REUSEPORT)
	FlowDissector         = ProgramType(sys.BPF_PROG_TYPE_FLOW_DISSECTOR)
	CGroupSysctl          = ProgramType(sys.BPF_PROG_TYPE_CGROUP_SYSCTL)
	RawTracepointWritable = ProgramType(sys.BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE)
	CGroupSockopt         = ProgramType(sys.BPF_PROG_TYPE_CGROUP_SOCKOPT)
	Tracing               = ProgramType(sys.BPF_PROG_TYPE_TRACING)
	StructOps             = ProgramType(sys.BPF_PROG_TYPE_STRUCT_OPS)
	Extension             = ProgramType(sys.BPF_PROG_TYPE_EXT)
	LSM                   = ProgramType(sys.BPF_PROG_TYPE_LSM)
	SkLookup              = ProgramType(sys.BPF_PROG_TYPE_SK_LOOKUP)
	Syscall               = ProgramType(sys.BPF_PROG_TYPE_SYSCALL)
	Netfilter             = ProgramType(sys.BPF_PROG_TYPE_NETFILTER)
)

const (
	WindowsXDP ProgramType = ProgramType(platform.WindowsTag) | (iota + 1)
	WindowsBind
	WindowsCGroupSockAddr
	WindowsSockOps
	WindowsXDPTest ProgramType = ProgramType(platform.WindowsTag) | 998
	WindowsSample  ProgramType = ProgramType(platform.WindowsTag) | 999
)

func ProgramTypeForPlatform(plat string, value uint32) (ProgramType, error) {
	_ = "STUB: not implemented"
	return *new(ProgramType), nil
}

type AttachType uint32

//go:generate go tool stringer -type AttachType -trimprefix Attach

const AttachNone AttachType = 0

const (
	AttachCGroupInetIngress          = AttachType(sys.BPF_CGROUP_INET_INGRESS)
	AttachCGroupInetEgress           = AttachType(sys.BPF_CGROUP_INET_EGRESS)
	AttachCGroupInetSockCreate       = AttachType(sys.BPF_CGROUP_INET_SOCK_CREATE)
	AttachCGroupSockOps              = AttachType(sys.BPF_CGROUP_SOCK_OPS)
	AttachSkSKBStreamParser          = AttachType(sys.BPF_SK_SKB_STREAM_PARSER)
	AttachSkSKBStreamVerdict         = AttachType(sys.BPF_SK_SKB_STREAM_VERDICT)
	AttachCGroupDevice               = AttachType(sys.BPF_CGROUP_DEVICE)
	AttachSkMsgVerdict               = AttachType(sys.BPF_SK_MSG_VERDICT)
	AttachCGroupInet4Bind            = AttachType(sys.BPF_CGROUP_INET4_BIND)
	AttachCGroupInet6Bind            = AttachType(sys.BPF_CGROUP_INET6_BIND)
	AttachCGroupInet4Connect         = AttachType(sys.BPF_CGROUP_INET4_CONNECT)
	AttachCGroupInet6Connect         = AttachType(sys.BPF_CGROUP_INET6_CONNECT)
	AttachCGroupInet4PostBind        = AttachType(sys.BPF_CGROUP_INET4_POST_BIND)
	AttachCGroupInet6PostBind        = AttachType(sys.BPF_CGROUP_INET6_POST_BIND)
	AttachCGroupUDP4Sendmsg          = AttachType(sys.BPF_CGROUP_UDP4_SENDMSG)
	AttachCGroupUDP6Sendmsg          = AttachType(sys.BPF_CGROUP_UDP6_SENDMSG)
	AttachLircMode2                  = AttachType(sys.BPF_LIRC_MODE2)
	AttachFlowDissector              = AttachType(sys.BPF_FLOW_DISSECTOR)
	AttachCGroupSysctl               = AttachType(sys.BPF_CGROUP_SYSCTL)
	AttachCGroupUDP4Recvmsg          = AttachType(sys.BPF_CGROUP_UDP4_RECVMSG)
	AttachCGroupUDP6Recvmsg          = AttachType(sys.BPF_CGROUP_UDP6_RECVMSG)
	AttachCGroupGetsockopt           = AttachType(sys.BPF_CGROUP_GETSOCKOPT)
	AttachCGroupSetsockopt           = AttachType(sys.BPF_CGROUP_SETSOCKOPT)
	AttachTraceRawTp                 = AttachType(sys.BPF_TRACE_RAW_TP)
	AttachTraceFEntry                = AttachType(sys.BPF_TRACE_FENTRY)
	AttachTraceFExit                 = AttachType(sys.BPF_TRACE_FEXIT)
	AttachModifyReturn               = AttachType(sys.BPF_MODIFY_RETURN)
	AttachLSMMac                     = AttachType(sys.BPF_LSM_MAC)
	AttachTraceIter                  = AttachType(sys.BPF_TRACE_ITER)
	AttachCgroupInet4GetPeername     = AttachType(sys.BPF_CGROUP_INET4_GETPEERNAME)
	AttachCgroupInet6GetPeername     = AttachType(sys.BPF_CGROUP_INET6_GETPEERNAME)
	AttachCgroupInet4GetSockname     = AttachType(sys.BPF_CGROUP_INET4_GETSOCKNAME)
	AttachCgroupInet6GetSockname     = AttachType(sys.BPF_CGROUP_INET6_GETSOCKNAME)
	AttachXDPDevMap                  = AttachType(sys.BPF_XDP_DEVMAP)
	AttachCgroupInetSockRelease      = AttachType(sys.BPF_CGROUP_INET_SOCK_RELEASE)
	AttachXDPCPUMap                  = AttachType(sys.BPF_XDP_CPUMAP)
	AttachSkLookup                   = AttachType(sys.BPF_SK_LOOKUP)
	AttachXDP                        = AttachType(sys.BPF_XDP)
	AttachSkSKBVerdict               = AttachType(sys.BPF_SK_SKB_VERDICT)
	AttachSkReuseportSelect          = AttachType(sys.BPF_SK_REUSEPORT_SELECT)
	AttachSkReuseportSelectOrMigrate = AttachType(sys.BPF_SK_REUSEPORT_SELECT_OR_MIGRATE)
	AttachPerfEvent                  = AttachType(sys.BPF_PERF_EVENT)
	AttachTraceKprobeMulti           = AttachType(sys.BPF_TRACE_KPROBE_MULTI)
	AttachTraceKprobeSession         = AttachType(sys.BPF_TRACE_KPROBE_SESSION)
	AttachLSMCgroup                  = AttachType(sys.BPF_LSM_CGROUP)
	AttachStructOps                  = AttachType(sys.BPF_STRUCT_OPS)
	AttachNetfilter                  = AttachType(sys.BPF_NETFILTER)
	AttachTCXIngress                 = AttachType(sys.BPF_TCX_INGRESS)
	AttachTCXEgress                  = AttachType(sys.BPF_TCX_EGRESS)
	AttachTraceUprobeMulti           = AttachType(sys.BPF_TRACE_UPROBE_MULTI)
	AttachCgroupUnixConnect          = AttachType(sys.BPF_CGROUP_UNIX_CONNECT)
	AttachCgroupUnixSendmsg          = AttachType(sys.BPF_CGROUP_UNIX_SENDMSG)
	AttachCgroupUnixRecvmsg          = AttachType(sys.BPF_CGROUP_UNIX_RECVMSG)
	AttachCgroupUnixGetpeername      = AttachType(sys.BPF_CGROUP_UNIX_GETPEERNAME)
	AttachCgroupUnixGetsockname      = AttachType(sys.BPF_CGROUP_UNIX_GETSOCKNAME)
	AttachNetkitPrimary              = AttachType(sys.BPF_NETKIT_PRIMARY)
	AttachNetkitPeer                 = AttachType(sys.BPF_NETKIT_PEER)
)

const (
	AttachWindowsXDP = AttachType(platform.WindowsTag | iota + 1)
	AttachWindowsBind
	AttachWindowsCGroupInet4Connect
	AttachWindowsCGroupInet6Connect
	AttachWindowsCgroupInet4RecvAccept
	AttachWindowsCgroupInet6RecvAccept
	AttachWindowsCGroupSockOps
	AttachWindowsSample
	AttachWindowsXDPTest
)

func AttachTypeForPlatform(plat string, value uint32) (AttachType, error) {
	_ = "STUB: not implemented"
	return *new(AttachType), nil
}

type AttachFlags uint32

type PinType uint32

const (
	PinNone PinType = iota

	PinByName
)

type LoadPinOptions struct {
	ReadOnly  bool
	WriteOnly bool

	Flags uint32
}

func (lpo *LoadPinOptions) Marshal() uint32 { _ = "STUB: not implemented"; return 0 }

type BatchOptions struct {
	ElemFlags uint64
	Flags     uint64
}

type LogLevel = sys.LogLevel

const (
	LogLevelBranch = sys.BPF_LOG_LEVEL1

	LogLevelInstruction = sys.BPF_LOG_LEVEL2

	LogLevelStats = sys.BPF_LOG_STATS
)
