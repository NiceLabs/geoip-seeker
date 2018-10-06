package qqwry

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/NiceLabs/geoip-seeker/shared"
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

func New(data []byte) (seeker *IPSeeker, err error) {
	if len(data) == 0 {
		err = shared.ErrFileSize
		return
	}
	seeker = new(IPSeeker)
	seeker.data = data

	seeker.firstIndex = int(binary.LittleEndian.Uint32(data[0:4]))
	seeker.lastIndex = int(binary.LittleEndian.Uint32(data[4:8]))
	seeker.indexCount = (seeker.lastIndex - seeker.firstIndex) / indexItemSize

	seeker.beginIP, _ = seeker.locateIndex(0)
	seeker.endIP, _ = seeker.locateIndex(seeker.indexCount)

	return
}

func (seeker *IPSeeker) LookupByIP(address net.IP) (record *shared.Record, err error) {
	address = address.To4()
	if address == nil {
		err = shared.ErrInvalidIPv4
		return
	}

	ip := ip2int(address)
	beginIndex := 0
	endIndex := seeker.indexCount

	if ip < seeker.beginIP {
		err = shared.ErrDataNotExists
		return
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

	record, err = seeker.LookupByIndex(beginIndex)
	if err != nil {
		return
	}
	if ip2int(record.BeginIP) > ip {
		err = shared.ErrDataNotExists
		return
	}
	record.IP = address
	return
}

func (seeker *IPSeeker) LookupByIndex(index int) (*shared.Record, error) {
	if index > seeker.indexCount && index >= 0 {
		return nil, errors.New("index out of range")
	}

	beginIP, offset := seeker.locateIndex(index)

	record := new(shared.Record)
	record.CountryName, record.RegionName = seeker.readRecord(offset+4, false)
	record.BeginIP = int2ip(beginIP)
	record.EndIP = net.IP(seeker.data[offset : offset+4])

	if record.CountryName == " CZ88.NET" {
		record.CountryName = ""
	}
	if record.RegionName == " CZ88.NET" {
		record.RegionName = ""
	}

	return record, nil
}

func (seeker *IPSeeker) IPv4Support() bool {
	return true
}

func (seeker *IPSeeker) IPv6Support() bool {
	return false
}

func (seeker *IPSeeker) BuildTime() time.Time {
	record, _ := seeker.LookupByIndex(seeker.indexCount)

	formats := []string{
		"%d\xc4\xea%d\xd4\xc2%d\xc8\xd5",
		"%d年%d月%d日",
		"%4d%2d%2d",
	}
	location := time.FixedZone("CST", +8*3600)
	for _, format := range formats {
		var year, month, day int
		_, err := fmt.Sscanf(record.RegionName, format, &year, &month, &day)
		if err != nil {
			continue
		}
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
	}
	return time.Unix(0, 0)
}

func (seeker *IPSeeker) RecordCount() uint64 {
	return uint64(seeker.indexCount - 1)
}

func (seeker *IPSeeker) locateIndex(index int) (beginIP uint32, offset int) {
	indexOffset := seeker.firstIndex + (indexItemSize * index)

	fields := padding(seeker.data[indexOffset:indexOffset+7], 8)

	beginIP = binary.LittleEndian.Uint32(fields[:4])
	offset = int(binary.LittleEndian.Uint32(fields[4:]))
	return
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
	record := int(binary.LittleEndian.Uint32(padding(seeker.data[index:offset], 4)))
	country, area = seeker.readRecord(record, false)
	if !onlyOne && mode == redirectMode2 {
		area, _ = seeker.readRecord(offset, true)
	}
	return
}
