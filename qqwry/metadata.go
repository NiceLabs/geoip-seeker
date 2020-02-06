package qqwry

import (
	"fmt"
	"time"

	"github.com/NiceLabs/geoip-seeker/shared"
)

func (s *Seeker) IPv4Support() bool   { return true }
func (s *Seeker) IPv6Support() bool   { return false }
func (s *Seeker) RecordCount() uint64 { return uint64(len(s.indexes)) }
func (s *Seeker) String() string      { return shared.ShowLibraryInfo("QQWry", s) }

func (s *Seeker) BuildTime() time.Time {
	record := s.index(s.RecordCount() - 1)
	formats := []string{
		"%d\xe5\xb9\xb4%d\xe6\x9c\x88%d\xe6\x97\xa5",
		"%d\xc4\xea%d\xd4\xc2%d\xc8\xd5",
		"%4d%2d%2d",
	}
	location := time.FixedZone("CST", +8*3600)
	var year, month, day int
	for _, format := range formats {
		_, err := fmt.Sscanf(record.RegionName, format, &year, &month, &day)
		if err == nil {
			break
		}
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}
