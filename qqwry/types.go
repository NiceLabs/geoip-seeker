package qqwry

import "net"

type index struct {
	ip             uint32
	beginIP, endIP net.IP
	offset         uint64
}
