package ipip_net

import (
	"net"
	"strings"
)

type Location struct {
	IP            net.IP `json:"ip"`
	BeginIP       net.IP `json:"begin_ip"`
	EndIP         net.IP `json:"end_ip"`
	Country       string `json:"country"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Unit          string `json:"unit"`
	ISP           string `json:"isp"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
	TimeZoneCode  string `json:"time_zone_code"`
	TimeZoneUTC   string `json:"time_zone_utc"`
	GB2260Code    string `json:"gb2260_code"`
	CallingCode   string `json:"calling_code"`
	ISO3166Code   string `json:"iso3166_code"`
	ContinentCode string `json:"continent_code"`
}

func (location *Location) StringDAT() string {
	fields := []string{
		location.Country,
		location.Province,
		location.City,
		location.Unit,
	}
	for index := range fields {
		if fields[index] == "" {
			fields[index] = "N/A"
		}
	}
	return strings.Join(fields, "\t")
}

func (location *Location) StringDATX() string {
	fields := []string{
		location.Country,
		location.Province,
		location.City,
		location.Unit,
		location.ISP,
		location.Longitude,
		location.Latitude,
		location.TimeZoneCode,
		location.TimeZoneUTC,
		location.GB2260Code,
		location.CallingCode,
		location.ISO3166Code,
		location.ContinentCode,
	}
	for index := range fields {
		if fields[index] == "" {
			fields[index] = "N/A"
		}
	}
	return strings.Join(fields, "\t")
}

func (location *Location) String() string {
	return location.StringDATX()
}
