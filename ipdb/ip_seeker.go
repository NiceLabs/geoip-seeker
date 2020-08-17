package ipdb

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"time"

	. "github.com/NiceLabs/geoip-seeker/shared"
)

type Seeker struct {
	meta     *meta
	records  []byte
	fileSize int
	language uint8
	v4Offset int
}

func New(data []byte) (*Seeker, error) {
	if len(data) == 0 {
		return nil, ErrFileSize
	}
	seeker := new(Seeker)
	if meta, err := loadMetadata(data); err != nil {
		return nil, err
	} else {
		seeker.meta = meta
	}
	seeker.fileSize = len(data)
	seeker.records = data[seeker.fileSize-seeker.meta.TotalSize:]
	seeker.v4Offset = seeker.findV4Offset()
	return seeker, nil
}

func (s *Seeker) LookupByIP(address net.IP) (record *Record, err error) {
	node, err := s.findNode(address)
	if err != nil {
		return
	}
	data, err := s.resolveNode(node)
	record = makeRecord(string(data), s.language, s.meta.Fields)
	record.IP = address
	return
}

func (s *Seeker) LookupByIndex(index uint64) (record *Record, err error) {
	return
}

func (s *Seeker) IPv4Support() bool   { return (s.meta.IPVersion & 0x01) == 0x01 }
func (s *Seeker) IPv6Support() bool   { return (s.meta.IPVersion & 0x02) == 0x02 }
func (s *Seeker) RecordCount() uint64 { return uint64(s.meta.NodeCount) }
func (s *Seeker) BuildTime() time.Time {
	location := time.FixedZone("CST", +8*3600)
	return time.Unix(s.meta.Build, 0).In(location)
}

func (s *Seeker) LanguageCode(code string) (err error) {
	if index, ok := s.meta.Languages[code]; ok {
		s.language = index
		return
	}
	return ErrNoSupportLanguage
}

func (s *Seeker) LanguageNames() (names []string) {
	for name := range s.meta.Languages {
		names = append(names, name)
	}
	return
}

func (s *Seeker) String() string {
	return ShowLibraryInfo("IPIP(IPDB)", s)
}

func (s *Seeker) findNode(ip net.IP) (node int, err error) {
	if ip := ip.To4(); ip != nil {
		if !s.IPv4Support() {
			err = ErrNoSupportIPv4
			return
		}
		return s.searchNode(ip, len(ip)*8)
	}
	if ip := ip.To16(); ip != nil {
		if !s.IPv6Support() {
			err = ErrNoSupportIPv6
			return
		}
		return s.searchNode(ip, len(ip)*8)
	}
	err = ErrIPFormat
	return
}

func (s *Seeker) searchNode(ip net.IP, bitCount int) (node int, err error) {
	node = 0
	if bitCount == 32 {
		node = s.v4Offset
	}
	for i := 0; i < bitCount; i++ {
		if node > s.meta.NodeCount {
			break
		}
		index := ((0xFF & int(ip[i>>3])) >> uint(7-(i%8))) & 1
		node = s.readNode(node, index)
	}
	if node < s.meta.NodeCount {
		err = ErrDataNotExists
	}
	return
}

func (s *Seeker) readNode(node, index int) int {
	offset := node*8 + index*4
	return int(binary.BigEndian.Uint32(s.records[offset : offset+4]))
}

func (s *Seeker) resolveNode(node int) (record []byte, err error) {
	resolved := node - s.meta.NodeCount + s.meta.NodeCount*8
	if resolved >= s.fileSize {
		err = ErrDatabaseError
		return
	}
	size := int(binary.BigEndian.Uint16(s.records[resolved : resolved+2]))
	if (resolved + 2 + size) > len(s.records) {
		err = ErrDatabaseError
		return
	}
	record = s.records[resolved+2 : resolved+2+size]
	return
}

func (s *Seeker) findV4Offset() (node int) {
	for i := 0; i < 96 && node < s.meta.NodeCount; i++ {
		if i >= 80 {
			node = s.readNode(node, 1)
		} else {
			node = s.readNode(node, 0)
		}
	}
	return
}

func loadMetadata(data []byte) (parsed *meta, err error) {
	length := int(binary.BigEndian.Uint32(data[:4]))
	original := data[4 : 4+length]

	parsed = new(meta)
	err = json.Unmarshal(original, parsed)
	if err != nil {
		return
	}
	if len(parsed.Languages) == 0 || len(parsed.Fields) == 0 {
		err = ErrMetaData
		return
	}
	if len(data) != (4 + length + parsed.TotalSize) {
		err = ErrFileSize
		return
	}
	return
}
