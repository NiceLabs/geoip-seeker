package qqwry

import (
	"encoding/binary"
	"net"

	"github.com/NiceLabs/geoip-seeker/shared"
)

type Seeker struct {
	data    []byte
	indexes []*index
}

func New(data []byte) (*Seeker, error) {
	if len(data) == 0 {
		return nil, shared.ErrFileSize
	}
	seeker := &Seeker{data: data}
	seeker.expandIndexes()
	return seeker, nil
}

func (s *Seeker) LookupByIP(address net.IP) (record *shared.Record, err error) {
	if address = address.To4(); address == nil {
		err = shared.ErrInvalidIPv4
		return
	}
	ip := uint(binary.BigEndian.Uint32(address))
	head := 0
	tail := len(s.indexes) - 1
	for (head + 1) < tail {
		index := (head + tail) / 2
		if s.indexes[index].ip <= ip {
			head = index
		} else {
			tail = index
		}
	}
	record = s.index(s.indexes[head])
	record.IP = address
	return
}

func (s *Seeker) index(index *index) *shared.Record {
	country, area := s.readRecord(index.offset)
	if area == " CZ88.NET" {
		area = ""
	}
	return &shared.Record{
		BeginIP:     index.begin,
		EndIP:       index.end,
		CountryName: country,
		RegionName:  area,
	}
}

func (s *Seeker) readRecord(offset uint) (country, area string) {
	switch mode := s.data[offset]; mode {
	case 1:
		return s.readRecord(s.readUInt24LE(offset + 1))
	case 2:
		country = s.readCString(s.readUInt24LE(offset + 1))
		area = s.readArea(offset + 4)
	default:
		country = s.readCString(offset)
		area = s.readArea(offset + 1 + uint(len(country)))
	}
	return
}

func (s *Seeker) readArea(offset uint) string {
	if s.data[offset] == 2 {
		offset = s.readUInt24LE(offset + 1)
	}
	return s.readCString(offset)
}

func (s *Seeker) readCString(offset uint) string {
	index := offset
	for s.data[index] != 0 {
		index++
	}
	return string(s.data[offset:index])
}
