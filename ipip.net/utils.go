package ipip_net

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
	"time"
)

func makeLocation(data string) *Location {
	location := new(Location)

	mapping := []*string{
		&location.Country, &location.Province, &location.City, &location.Unit,
		&location.ISP,
		&location.Latitude, &location.Longitude,
		&location.TimeZoneCode, &location.TimeZoneUTC,
		&location.GB2260Code, &location.CallingCode, &location.ISO3166Code, &location.ContinentCode,
	}

	for index, field := range strings.Split(data, "\t") {
		*mapping[index] = field
	}

	return location
}

func ip2int(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func padding(data []byte, length int) []byte {
	payload := make([]byte, length)
	copy(payload[0:len(data)], data)
	return payload
}

func resolvePublishDate(version string) time.Time {
	var year, month, day, hour int64
	year, _ = strconv.ParseInt(version[0:4], 10, 32)
	month, _ = strconv.ParseInt(version[4:6], 10, 32)
	day, _ = strconv.ParseInt(version[6:8], 10, 32)
	if len(version) == 10 {
		hour, _ = strconv.ParseInt(version[8:10], 10, 32)
	}

	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(int(year), time.Month(month), int(day), int(hour), 0, 0, 0, location)
}
