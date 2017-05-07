package ipip_net

import (
	"encoding/binary"
	"errors"
	"net"
	"time"
)

const (
	dataOffset = 4
)

type IPSeeker struct {
	data []byte

	indexOffset, indexSpace int
	recordSize, recordSpace int

	locateRecordOffset func(address net.IP) int
	getRecordLength    func(record []byte) int
}

func NewDAT(data []byte) *IPSeeker {
	seeker := new(IPSeeker)
	seeker.init(data, 0x400, 8)
	seeker.locateRecordOffset = func(address net.IP) int { return int(address[0]) * 0x4 }
	seeker.getRecordLength = func(record []byte) int { return int(record[7]) }
	return seeker
}

func NewDATX(data []byte) *IPSeeker {
	seeker := new(IPSeeker)
	seeker.init(data, 0x40000, 9)
	seeker.locateRecordOffset = func(address net.IP) int { return (int(address[0])*0x100 + int(address[1])) * 0x4 }
	seeker.getRecordLength = func(record []byte) int { return int(binary.BigEndian.Uint16(record[7:])) }
	return seeker
}

func (seeker *IPSeeker) init(data []byte, indexSpace, recordSize int) {
	seeker.data = data[dataOffset:]
	seeker.indexSpace = indexSpace
	seeker.indexOffset = int(binary.BigEndian.Uint32(data[:dataOffset]))
	seeker.recordSize = recordSize
	seeker.recordSpace = seeker.indexOffset - seeker.indexSpace - dataOffset
}

func (seeker *IPSeeker) LookupByIP(address net.IP) (*Location, error) {
	if address.To4() == nil {
		return nil, errors.New("invalid IP address")
	}
	ip, record := seeker.locate(address.To4())
	return makeRecord(int2ip(ip), record)
}

func (seeker *IPSeeker) PublishDate() time.Time {
	location, _ := seeker.LookupByIP(net.IPv4bcast)
	return resolvePublishDate(location.Province)
}

func (seeker *IPSeeker) RecordCount() int {
	recordSpace := seeker.recordSpace - seeker.indexSpace
	recordSize := seeker.recordSize
	return (recordSpace / recordSize) - 1
}

func (seeker *IPSeeker) locate(address net.IP) (ip uint32, record string) {
	beginIndex := seeker.locateBeginIndex(address)
	endIndex := seeker.locateEndIndex(address)

	ip = ip2int(address)
	for beginIndex < endIndex {
		middleIndex := (beginIndex + endIndex) / 2
		middleIP := seeker.locateRecord(middleIndex)[:4]
		if ip2int(middleIP) <= ip {
			beginIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}
	}

	record = seeker.getRecord(seeker.locateRecord(beginIndex))
	return
}

func (seeker *IPSeeker) locateRecord(index int) (record []byte) {
	indexSpace := seeker.indexSpace
	recordSize := seeker.recordSize
	offset := index*recordSize + indexSpace

	record = seeker.data[offset : offset+recordSize]
	return
}

func (seeker *IPSeeker) locateBeginIndex(address net.IP) int {
	offset := seeker.locateRecordOffset(address)
	return int(binary.LittleEndian.Uint32(seeker.data[offset : offset+4]))
}

func (seeker *IPSeeker) locateEndIndex(address net.IP) int {
	if address[0] == 255 {
		return seeker.RecordCount()
	}
	return seeker.locateBeginIndex([]byte{address[0] + 1, 0, 0, 0})
}

func (seeker *IPSeeker) getRecord(record []byte) string {
	offset := int(binary.LittleEndian.Uint32(padding(record[4:7], 4)))
	length := seeker.getRecordLength(record)
	offset = seeker.indexOffset + offset - seeker.indexSpace - dataOffset
	return string(seeker.data[offset : offset+length])
}
