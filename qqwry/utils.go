package qqwry

import (
	"encoding/binary"
	"net"
)

func readCString(data []byte, offset int) string {
	for index := offset; index < len(data); index++ {
		if data[index] == 0 {
			return string(data[offset:index])
		}
	}
	return ""
}

func int2ip(ip uint32) net.IP {
	address := make(net.IP, 4)
	binary.BigEndian.PutUint32(address, ip)
	return address
}

func ip2int(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func padding(data []byte, length int) []byte {
	payload := make([]byte, length)
	copy(payload[0:len(data)], data)
	return payload
}
