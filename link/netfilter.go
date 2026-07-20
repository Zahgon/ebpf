//go:build !windows

package link

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal/sys"
)

const NetfilterIPDefrag NetfilterAttachFlags = 0

type NetfilterAttachFlags uint32

type NetfilterInetHook = sys.NetfilterInetHook

const (
	NetfilterInetPreRouting  = sys.NF_INET_PRE_ROUTING
	NetfilterInetLocalIn     = sys.NF_INET_LOCAL_IN
	NetfilterInetForward     = sys.NF_INET_FORWARD
	NetfilterInetLocalOut    = sys.NF_INET_LOCAL_OUT
	NetfilterInetPostRouting = sys.NF_INET_POST_ROUTING
)

type NetfilterProtocolFamily = sys.NetfilterProtocolFamily

const (
	NetfilterProtoUnspec = sys.NFPROTO_UNSPEC
	NetfilterProtoInet   = sys.NFPROTO_INET
	NetfilterProtoIPv4   = sys.NFPROTO_IPV4
	NetfilterProtoARP    = sys.NFPROTO_ARP
	NetfilterProtoNetdev = sys.NFPROTO_NETDEV
	NetfilterProtoBridge = sys.NFPROTO_BRIDGE
	NetfilterProtoIPv6   = sys.NFPROTO_IPV6
)

type NetfilterOptions struct {
	Program *ebpf.Program

	ProtocolFamily NetfilterProtocolFamily

	Hook NetfilterInetHook

	Priority int32

	Flags uint32

	NetfilterFlags NetfilterAttachFlags
}

type netfilterLink struct {
	RawLink
}

func AttachNetfilter(opts NetfilterOptions) (Link, error) {
	_ = "STUB: not implemented"
	return *new(Link), nil
}

func (*netfilterLink) Update(_ *ebpf.Program) error { _ = "STUB: not implemented"; return nil }

func (nf *netfilterLink) Info() (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

var _ Link = (*netfilterLink)(nil)
