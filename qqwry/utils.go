package qqwry

import (
	"encoding/binary"
	"net"
)

type index struct {
	ip, offset uint
	begin, end net.IP
}

func (s *Seeker) expandIndexes() {
	first := uint(binary.LittleEndian.Uint32(s.data[0:4]))
	last := uint(binary.LittleEndian.Uint32(s.data[4:8]))
	for i := first; i < last+7; i += 7 {
		offset := s.readUInt24LE(i + 4)
		s.indexes = append(s.indexes, &index{
			ip:     uint(binary.LittleEndian.Uint32(s.data[i : i+4])),
			offset: offset + 4,
			begin:  s.readIP(i),
			end:    s.readIP(offset),
		})
	}
	return
}

func (s *Seeker) readIP(offset uint) net.IP {
	return net.IP{s.data[offset+3], s.data[offset+2], s.data[offset+1], s.data[offset]}
}

func (s *Seeker) readUInt24LE(offset uint) uint {
	return uint(s.data[offset]) | uint(s.data[offset+1])<<8 | uint(s.data[offset+2])<<16
}
