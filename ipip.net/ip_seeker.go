package ipip_net

import (
	"encoding/binary"
	"net"
	"time"

	"github.com/NiceLabs/geoip-seeker/shared"
)

const (
	ModeDAT  Mode = 0
	ModeDATX Mode = 1

	datIndexSpace  = 0x400
	datRecordSize  = 8
	datxIndexSpace = 0x40000
	datxRecordSize = 9
)

type Mode int

type IPSeeker struct {
	headerIndex []byte
	recordIndex []byte
	records     []byte

	recordSize int

	locateRecordIndex func(address net.IP) int
	getRecordLength   func(record []byte) int
}

func New(data []byte, mode Mode) (seeker *IPSeeker, err error) {
	if len(data) == 0 {
		err = shared.ErrFileSize
		return
	}
	seeker = new(IPSeeker)
	switch mode {
	case ModeDAT:
		seeker.init(data, datIndexSpace, datRecordSize)
		seeker.locateRecordIndex = func(address net.IP) int {
			return int(address[0])
		}
		seeker.getRecordLength = func(record []byte) int {
			return int(record[7])
		}
	case ModeDATX:
		seeker.init(data, datxIndexSpace, datxRecordSize)
		seeker.locateRecordIndex = func(address net.IP) int {
			return int(binary.BigEndian.Uint16(address[:2]))
		}
		seeker.getRecordLength = func(record []byte) int {
			return int(binary.BigEndian.Uint16(record[7:9]))
		}
	default:
		err = shared.ErrModeError
	}
	return
}

func (seeker *IPSeeker) init(data []byte, indexSpace, recordSize int) {
	seeker.recordSize = recordSize

	indexOffset := int(binary.BigEndian.Uint32(data[:4]))

	seeker.headerIndex = data[4 : 4+indexSpace]
	seeker.recordIndex = data[4+indexSpace : 4+indexOffset-indexSpace]
	seeker.records = data[indexOffset-indexSpace:]
}

func (seeker *IPSeeker) LookupByIP(address net.IP) (record *shared.Record, err error) {
	address = address.To4()
	if address == nil {
		err = shared.ErrInvalidIPv4
		return
	}

	record = seeker.locate(address)
	record.IP = address
	return
}

func (seeker *IPSeeker) LookupByIndex(index int) (record *shared.Record, err error) {
	if index > int(seeker.RecordCount()) && index >= 0 {
		err = shared.ErrIndexOutOfRange
		return
	}

	recordItem := seeker.locateRecord(index)
	record = seeker.getRecord(recordItem)
	if record == nil {
		return
	}
	if index == 0 {
		record.BeginIP = net.IPv4zero
	} else {
		record.BeginIP = shared.Int2IP(shared.IP2Int(seeker.locateRecord(index - 1)[:4]) + 1)
	}
	record.EndIP = recordItem[:4]
	return
}

func (seeker *IPSeeker) IPv4Support() bool {
	return true
}

func (seeker *IPSeeker) IPv6Support() bool {
	return false
}

func (seeker *IPSeeker) BuildTime() time.Time {
	recordCount := int(seeker.RecordCount())
	recordIndex := seeker.locateRecord(recordCount)
	record := seeker.getRecord(recordIndex)
	return resolveBuildTime(record.RegionName)
}

func (seeker *IPSeeker) RecordCount() int {
	return len(seeker.recordIndex)/seeker.recordSize - 1
}

func (seeker *IPSeeker) String() string {
	return shared.ShowLibraryInfo("IPIP(DAT/DATX)", seeker)
}

func (seeker *IPSeeker) locate(address net.IP) *shared.Record {
	beginIndex := seeker.locateBeginIndex(address)
	endIndex := int(seeker.RecordCount())

	ip := shared.IP2Int(address)
	for beginIndex <= endIndex {
		middleIndex := (beginIndex + endIndex) / 2
		middleIP := seeker.locateRecord(middleIndex)[:4]
		if shared.IP2Int(middleIP) < ip {
			beginIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}
	}
	record, _ := seeker.LookupByIndex(beginIndex)
	return record
}

func (seeker *IPSeeker) locateRecord(index int) []byte {
	offset := index * seeker.recordSize
	return seeker.recordIndex[offset : offset+seeker.recordSize]
}

func (seeker *IPSeeker) locateBeginIndex(address net.IP) int {
	offset := seeker.locateRecordIndex(address) * 4
	return int(binary.LittleEndian.Uint32(seeker.headerIndex[offset : offset+4]))
}

func (seeker *IPSeeker) getRecord(record []byte) *shared.Record {
	offset := int(binary.LittleEndian.Uint32(shared.Padding(record[4:7], 4)))
	length := seeker.getRecordLength(record)
	return makeRecord(string(seeker.records[offset : offset+length]))
}

func resolveBuildTime(version string) time.Time {
	layouts := []string{"2006010215", "20060102"}
	zone := time.FixedZone("CST", +8*3600)
	for _, layout := range layouts {
		date, err := time.ParseInLocation(layout, version, zone)
		if err != nil {
			continue
		}
		return date
	}
	return time.Unix(0, 0)
}
