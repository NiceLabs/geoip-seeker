package ipdb

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"time"

	"github.com/NiceLabs/geoip-seeker/shared"
)

type IPSeeker struct {
	meta     *MetaData
	records  []byte
	fileSize int
	language int
	v4Offset int
}

func New(data []byte) (seeker *IPSeeker, err error) {
	if len(data) == 0 {
		err = errors.New("data is empty")
		return
	}
	seeker = new(IPSeeker)
	if seeker.meta, err = loadMetaData(data); err != nil {
		return nil, err
	}
	seeker.fileSize = len(data)
	seeker.records = data[seeker.fileSize-seeker.meta.TotalSize:]
	seeker.v4Offset = seeker.findV4Offset()
	return seeker, nil
}

func (seeker *IPSeeker) LookupByIP(address net.IP) (record *shared.Record, err error) {
	node, err := seeker.findNode(address)
	if err != nil {
		return
	}
	data, err := seeker.resolveNode(node)
	record = makeRecord(string(data), seeker.language, seeker.meta.Fields)
	record.IP = address
	return
}

func (seeker *IPSeeker) IPv4Support() bool {
	return seeker.meta.IPv4Support()
}

func (seeker *IPSeeker) IPv6Support() bool {
	return seeker.meta.IPv6Support()
}

func (seeker *IPSeeker) BuildTime() time.Time {
	return seeker.meta.BuildDate()
}

func (seeker *IPSeeker) RecordCount() int {
	return seeker.meta.NodeCount
}

func (seeker *IPSeeker) LanguageCode(code string) (err error) {
	if index, ok := seeker.meta.Languages[code]; ok {
		seeker.language = index
		return
	}
	return shared.ErrNoSupportLanguage
}

func (seeker *IPSeeker) LanguageNames() []string {
	return seeker.meta.LanguageNames()
}

func (seeker *IPSeeker) String() string {
	return shared.ShowLibraryInfo("IPIP(IPDB)", seeker)
}

func (seeker *IPSeeker) findNode(ip net.IP) (node int, err error) {
	if ip := ip.To4(); ip != nil {
		if !seeker.meta.IPv4Support() {
			err = shared.ErrNoSupportIPv4
			return
		}
		return seeker.searchNode(ip, len(ip)*8)
	}
	if ip := ip.To16(); ip != nil {
		if !seeker.meta.IPv6Support() {
			err = shared.ErrNoSupportIPv6
			return
		}
		return seeker.searchNode(ip, len(ip)*8)
	}
	err = shared.ErrIPFormat
	return
}

func (seeker *IPSeeker) searchNode(ip net.IP, bitCount int) (node int, err error) {
	node = 0
	if bitCount == 32 {
		node = seeker.v4Offset
	}
	for i := 0; i < bitCount; i++ {
		if node > seeker.meta.NodeCount {
			break
		}
		index := ((0xFF & int(ip[i>>3])) >> uint(7-(i%8))) & 1
		node = seeker.readNode(node, index)
	}
	if node < seeker.meta.NodeCount {
		err = shared.ErrDataNotExists
	}
	return
}

func (seeker *IPSeeker) readNode(node, index int) int {
	offset := node*8 + index*4
	return int(binary.BigEndian.Uint32(seeker.records[offset : offset+4]))
}

func (seeker *IPSeeker) resolveNode(node int) (record []byte, err error) {
	resolved := node - seeker.meta.NodeCount + seeker.meta.NodeCount*8
	if resolved >= seeker.fileSize {
		err = shared.ErrDatabaseError
		return
	}
	size := int(binary.BigEndian.Uint16(seeker.records[resolved : resolved+2]))
	if (resolved + 2 + size) > len(seeker.records) {
		err = shared.ErrDatabaseError
		return
	}
	record = seeker.records[resolved+2 : resolved+2+size]
	return
}

func (seeker *IPSeeker) findV4Offset() (node int) {
	for i := 0; i < 96 && node < seeker.meta.NodeCount; i++ {
		if i >= 80 {
			node = seeker.readNode(node, 1)
		} else {
			node = seeker.readNode(node, 0)
		}
	}
	return
}

func loadMetaData(data []byte) (meta *MetaData, err error) {
	metaLength := int(binary.BigEndian.Uint32(data[:4]))
	metaData := data[4 : 4+metaLength]

	meta = new(MetaData)
	err = json.Unmarshal(metaData, meta)

	if len(meta.Languages) == 0 || len(meta.Fields) == 0 {
		return nil, shared.ErrMetaData
	}
	if len(data) != (4 + metaLength + meta.TotalSize) {
		return nil, shared.ErrFileSize
	}
	return
}
