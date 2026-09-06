//go:build linux

package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"golang.org/x/sys/unix"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/internal/sys"
	"github.com/cilium/ebpf/link"
)

//go:generate go tool bpf2go -tags linux bpf xdp.c -- -I../headers

func main() {
	if len(os.Args) < 7 {
		log.Fatalf("Usage: %s <ifname> <repeat> <batch_size> <src_ip> <dst_ip> <dst_mac>", os.Args[0])
	}

	ifaceName := os.Args[1]
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		log.Fatalf("lookup network iface %q: %s", ifaceName, err)
	}

	repeat, err := strconv.ParseUint(os.Args[2], 10, 32)
	if err != nil {
		log.Fatalf("parsing repeat count %q: %s", os.Args[2], err)
	}

	batchSize, err := strconv.ParseUint(os.Args[3], 10, 32)
	if err != nil {
		log.Fatalf("parsing batch size %q: %s", os.Args[3], err)
	}

	srcIP := net.ParseIP(os.Args[4]).To4()
	if srcIP == nil {
		log.Fatalf("invalid source IP address: %s", os.Args[4])
	}

	dstIP := net.ParseIP(os.Args[5]).To4()
	if dstIP == nil {
		log.Fatalf("invalid destination IP address: %s", os.Args[5])
	}

	dstMAC, err := net.ParseMAC(os.Args[6])
	if err != nil {
		log.Fatalf("invalid destination MAC address %q: %s", os.Args[6], err)
	}

	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %s", err)
	}
	defer objs.Close()

	l, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.XdpProgPass,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatalf("could not attach XDP program: %s", err)
	}
	defer l.Close()

	log.Printf("Attached XDP program to iface %q (index %d)", iface.Name, iface.Index)
	log.Printf("Running XDP program in live frame mode with Repeat: %d, BatchSize: %d", repeat, batchSize)
	log.Printf("Src MAC: %s, Dst MAC: %s", iface.HardwareAddr, dstMAC)
	log.Printf("Src IP: %s, Dst IP: %s", srcIP, dstIP)

	pkt := buildUDPPacket(iface.HardwareAddr, dstMAC, srcIP, dstIP, 12345, 9999, []byte("Hello, XDP!"))

	xdpmd := &sys.XdpMd{
		DataEnd:        uint32(len(pkt)),
		IngressIfindex: uint32(iface.Index),
	}
	ret, err := objs.XdpProgTx.Run(&ebpf.RunOptions{
		Data:      pkt,
		Repeat:    uint32(repeat),
		Flags:     unix.BPF_F_TEST_XDP_LIVE_FRAMES,
		Context:   xdpmd,
		BatchSize: uint32(batchSize),
	})
	if err != nil {
		log.Fatalf("running XDP program with BPF_F_TEST_XDP_LIVE_FRAMES: %s", err)
	}

	log.Printf("XDP program completed with return value: %d", ret)
}

func buildUDPPacket(srcMAC, dstMAC net.HardwareAddr, srcIP, dstIP net.IP, srcPort, dstPort uint16, payload []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ipChecksum(header []byte) uint16 { _ = "STUB: not implemented"; return 0 }
