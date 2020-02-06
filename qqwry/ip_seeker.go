package qqwry

import (
	"encoding/binary"
	"net"

	. "github.com/NiceLabs/geoip-seeker/shared"
)

type Seeker struct {
	data            []byte
	firstIP, lastIP uint32
	indexes         []*index
}

func New(data []byte) (*Seeker, error) {
	if len(data) == 0 {
		return nil, ErrFileSize
	}
	indexes := expandIndexes(data)
	seeker := &Seeker{
		data:    data,
		indexes: indexes,
		firstIP: indexes[0].ip,
		lastIP:  indexes[len(indexes)-1].ip,
	}
	return seeker, nil
}

func (s *Seeker) LookupByIP(address net.IP) (record *Record, err error) {
	if address = address.To4(); address == nil {
		err = ErrInvalidIPv4
		return
	}
	ip := binary.BigEndian.Uint32(address)
	head := uint64(0)
	tail := s.RecordCount() - 1
	if ip >= s.lastIP {
		head = tail
	} else {
		for (head + 1) < tail {
			index := (head + tail) / 2
			if s.indexes[index].ip <= ip {
				head = index
			} else {
				tail = index
			}
		}
	}
	record = s.index(head)
	record.IP = address
	return
}

func (s *Seeker) LookupByIndex(index uint64) (record *Record, err error) {
	record = s.index(index)
	return
}

func (s *Seeker) index(index uint64) (record *Record) {
	item := s.indexes[index]
	country, area := s.readRecord(item.offset)
	record = &Record{
		BeginIP:     item.beginIP,
		EndIP:       item.endIP,
		CountryName: country,
		RegionName:  area,
	}
	return
}

func (s *Seeker) readRecord(offset uint64) (country, area string) {
	switch mode := s.data[offset]; mode {
	case 1:
		return s.readRecord(ReadUInt24(s.data, offset+1))
	case 2:
		country = s.readCString(ReadUInt24(s.data, offset+1))
		area = s.readArea(offset + 4)
	default:
		country = s.readCString(offset)
		area = s.readArea(offset + uint64(len(country)) + 1)
	}
	return
}

func (s *Seeker) readArea(offset uint64) string {
	if s.data[offset] == 2 {
		offset = ReadUInt24(s.data, offset+1)
	}
	return s.readCString(offset)
}

func (s *Seeker) readCString(offset uint64) string {
	index := offset
	for s.data[index] != 0 {
		index++
	}
	return string(s.data[offset:index])
}
