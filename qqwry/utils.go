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

// see https://github.com/shuax/LocateIP/blob/master/loci/cz_update.c#L23-L29
func dateToVersion(year, month, day uint32) uint32 {
	month = (month + 9) % 12
	year = year - month/10
	day = 365*year + year/4 - year/100 + year/400 + (month*153+2)/5 + day - 1
	return day
}

// see https://github.com/shuax/LocateIP/blob/master/loci/cz_update.c#L31-L41
func versionToDate(version uint32) (year, month, day int) {
	y := (version*33 + 999) / 12053
	t := version - y*365 - y/4 + y/100 - y/400
	m := (t*5+2)/153 + 2

	year = int(y + m/12)
	month = int(m%12 + 1)
	day = int(t - (m*153-304)/5 + 1)
	return
}
