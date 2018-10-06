package ipip_net

import (
	"encoding/binary"
	"net"
	"time"
)

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

func resolvePublishDate(version string) (date time.Time) {
	layout := "2006010215"
	zone := time.FixedZone("CST", +8*3600)
	date, _ = time.ParseInLocation(layout, version, zone)
	return
}
