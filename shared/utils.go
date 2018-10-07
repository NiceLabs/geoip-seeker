package shared

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

func Int2IP(ip uint32) net.IP {
	address := make(net.IP, 4)
	binary.BigEndian.PutUint32(address, ip)
	return address
}

func IP2Int(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func Padding(data []byte, length int) []byte {
	payload := make([]byte, length)
	copy(payload[0:len(data)], data)
	return payload
}

func ShowLibraryInfo(name string, seeker IPSeeker) string {
	items := []string{
		name,
		seeker.BuildTime().Format("2006-01-02"),
		strconv.Itoa(seeker.RecordCount()),
	}
	if seeker.IPv4Support() {
		items = append(items, "[IPv4]")
	}
	if seeker.IPv6Support() {
		items = append(items, "[IPv6]")
	}
	return strings.Join(items, " ")
}
