package qqwry

import (
	"encoding/binary"
	"errors"
	"net"
)

const (
	indexItemSize = 7
	redirectMode1 = 0x01
	redirectMode2 = 0x02
)

type IPSeeker struct {
	data []byte

	firstIndex int
	lastIndex  int

	indexCount int

	beginIP uint32
	endIP   uint32
}

func New(data []byte) *IPSeeker {
	seeker := new(IPSeeker)
	seeker.data = data

	seeker.firstIndex = int(binary.LittleEndian.Uint32(data[0:4]))
	seeker.lastIndex = int(binary.LittleEndian.Uint32(data[4:8]))
	seeker.indexCount = (seeker.lastIndex - seeker.firstIndex) / indexItemSize

	seeker.beginIP, _ = seeker.locateIndex(0)
	seeker.endIP, _ = seeker.locateIndex(seeker.indexCount)

	return seeker
}

func (seeker *IPSeeker) LookupByIP(address net.IP) (*Location, error) {
	if address == nil {
		return nil, errors.New("invalid IP address")
	}
	ip := ip2int(address)
	beginIndex := 0
	endIndex := seeker.indexCount

	if ip < seeker.beginIP {
		return nil, errors.New("IP not found.")
	} else if ip >= seeker.endIP {
		beginIndex = endIndex
	} else {
		for (beginIndex + 1) < endIndex {
			middleIndex := (beginIndex + endIndex) / 2
			middleIP, _ := seeker.locateIndex(middleIndex)
			if middleIP <= ip {
				beginIndex = middleIndex
			} else {
				endIndex = middleIndex
			}
		}
	}

	location, err := seeker.locateRecord(beginIndex)
	if err != nil {
		return nil, err
	}
	if ip2int(location.BeginIP) > ip {
		return nil, errors.New("IP not found.")
	}
	location.IP = address
	return location, nil
}

func (seeker *IPSeeker) Version() string {
	location, _ := seeker.locateRecord(seeker.indexCount)
	return location.Area
}

func (seeker *IPSeeker) RecordCount() int {
	return seeker.indexCount - 1
}

func (seeker *IPSeeker) locateIndex(index int) (beginIP uint32, offset int) {
	indexOffset := seeker.firstIndex + (indexItemSize * index)

	fields := padding(seeker.data[indexOffset:indexOffset+7], 8)

	beginIP = binary.LittleEndian.Uint32(fields[:4])
	offset = int(binary.LittleEndian.Uint32(fields[4:]))
	return
}

func (seeker *IPSeeker) locateRecord(index int) (*Location, error) {
	if index > seeker.indexCount && index >= 0 {
		return nil, errors.New("index out of range.")
	}

	beginIP, offset := seeker.locateIndex(index)

	location := new(Location)
	location.Country, location.Area = seeker.readRecord(offset+4, false)
	location.BeginIP = int2ip(beginIP)
	location.EndIP = net.IP(seeker.data[offset : offset+4])

	if location.Area == " CZ88.NET" {
		location.Area = ""
	}

	return location, nil
}

func (seeker *IPSeeker) readRecord(index int, onlyOne bool) (country, area string) {
	mode := seeker.data[index]
	index += 1
	if mode != redirectMode1 && mode != redirectMode2 {
		country = readCString(seeker.data, index-1)
		if !onlyOne {
			index += len(country)
			area, _ = seeker.readRecord(index, true)
		}
		return
	}
	offset := index + 3
	offset_record := int(binary.LittleEndian.Uint32(padding(seeker.data[index:offset], 4)))
	country, area = seeker.readRecord(offset_record, true)
	if !onlyOne && mode == redirectMode2 {
		area, _ = seeker.readRecord(offset, true)
	}
	return
}
