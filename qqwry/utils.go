package qqwry

import (
	"encoding/binary"
	"net"

	. "github.com/NiceLabs/geoip-seeker/shared"
)

func expandIndexes(data []byte) (indexes []*index) {
	first := binary.LittleEndian.Uint32(data[0:4])
	last := binary.LittleEndian.Uint32(data[4:8])
	for i := first; i < last+7; i += 7 {
		offset := ReadUInt24(data, uint64(i+4))
		indexes = append(indexes, &index{
			ip:      binary.LittleEndian.Uint32(data[i : i+4]),
			beginIP: readIP(data, uint64(i)),
			endIP:   readIP(data, offset),
			offset:  offset + 4,
		})
	}
	return
}

func readIP(data []byte, offset uint64) net.IP {
	return net.IP{data[offset+3], data[offset+2], data[offset+1], data[offset]}
}
