package ipip_net

import (
	"encoding/binary"
	"errors"
	"net"
	"time"
)

type IPSeeker struct {
	headerIndex []byte
	recordIndex []byte
	records     []byte

	indexSpace int
	recordSize int

	locateRecordOffset func(address net.IP) int
	getRecordLength    func(record []byte) int
}

func NewDAT(data []byte) (*IPSeeker, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	seeker := new(IPSeeker)
	seeker.init(data, 0x400, 8)
	seeker.locateRecordOffset = func(address net.IP) int {
		return int(address[0]) * 0x4
	}
	seeker.getRecordLength = func(record []byte) int {
		return int(record[7])
	}
	return seeker, nil
}

func NewDATX(data []byte) (*IPSeeker, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	seeker := new(IPSeeker)
	seeker.init(data, 0x40000, 9)
	seeker.locateRecordOffset = func(address net.IP) int {
		return (int(address[0])*0x100 + int(address[1])) * 0x4
	}
	seeker.getRecordLength = func(record []byte) int {
		return int(binary.BigEndian.Uint16(record[7:9]))
	}
	return seeker, nil
}

func (seeker *IPSeeker) init(data []byte, indexSpace, recordSize int) {
	seeker.recordSize = recordSize
	seeker.indexSpace = indexSpace

	indexOffset := int(binary.BigEndian.Uint32(data[:4]))

	seeker.headerIndex = data[4:seeker.indexSpace]
	seeker.recordIndex = data[4+seeker.indexSpace : indexOffset-indexSpace]
	seeker.records = data[indexOffset-seeker.indexSpace:]
}

func (seeker *IPSeeker) LookupByIP(address net.IP) (location *Location, err error) {
	address = address.To4()
	if address == nil {
		err = errors.New("invalid IPv4 address")
		return
	}

	location = seeker.locate(address)
	location.IP = address
	return
}

func (seeker *IPSeeker) LookupByIndex(index int) (location *Location, err error) {
	if index > seeker.RecordCount() && index >= 0 {
		err = errors.New("index out of range.")
		return
	}

	record := seeker.locateRecord(index)
	location = seeker.getRecord(record)
	if location == nil {
		return
	}
	if index == 0 {
		location.BeginIP = net.IPv4zero
	} else {
		location.BeginIP = int2ip(ip2int(seeker.locateRecord(index - 1)[:4]) + 1)
	}
	location.EndIP = record[:4]
	return
}

func (seeker *IPSeeker) PublishDate() time.Time {
	recordCount := seeker.RecordCount()
	recordIndex := seeker.locateRecord(recordCount)
	location := seeker.getRecord(recordIndex)
	return resolvePublishDate(location.Province)
}

func (seeker *IPSeeker) RecordCount() int {
	return (len(seeker.recordIndex) / seeker.recordSize) - 1
}

func (seeker *IPSeeker) locate(address net.IP) *Location {
	beginIndex := seeker.locateBeginIndex(address)
	endIndex := seeker.RecordCount()

	ip := ip2int(address)
	for beginIndex <= endIndex {
		middleIndex := (beginIndex + endIndex) / 2
		middleIP := seeker.locateRecord(middleIndex)[:4]
		if ip2int(middleIP) <= ip {
			beginIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}
	}
	location, _ := seeker.LookupByIndex(beginIndex)
	return location
}

func (seeker *IPSeeker) locateRecord(index int) []byte {
	offset := index * seeker.recordSize
	return seeker.recordIndex[offset : offset+seeker.recordSize]
}

func (seeker *IPSeeker) locateBeginIndex(address net.IP) int {
	offset := seeker.locateRecordOffset(address)
	return int(binary.LittleEndian.Uint32(seeker.headerIndex[offset : offset+4]))
}

func (seeker *IPSeeker) getRecord(record []byte) *Location {
	offset := int(binary.LittleEndian.Uint32(padding(record[4:7], 4)))
	length := seeker.getRecordLength(record)
	return makeLocation(string(seeker.records[offset : offset+length]))
}
