package shared

import (
	"net"
	"strings"
)

type Location struct {
	IP      net.IP `json:"ip,omitempty"`
	BeginIP net.IP `json:"begin_ip,omitempty"`
	EndIP   net.IP `json:"end_ip,omitempty"`

	CountryName    string `json:"country_name,omitempty"`
	RegionName     string `json:"region_name,omitempty"`
	CityName       string `json:"city_name,omitempty"`
	OwnerDomain    string `json:"owner_domain,omitempty"`
	ISPDomain      string `json:"isp_domain,omitempty"`
	Latitude       string `json:"latitude,omitempty"`
	Longitude      string `json:"longitude,omitempty"`
	TimeZone       string `json:"time_zone,omitempty"`
	UTCOffset      string `json:"utc_offset,omitempty"`
	ChinaAdminCode string `json:"china_admin_code,omitempty"`
	IDDCode        string `json:"idd_code,omitempty"`
	CountryCode    string `json:"country_code,omitempty"`
	ContinentCode  string `json:"continent_code,omitempty"`
	IDC            string `json:"idc,omitempty"`
	BaseStation    string `json:"base_station,omitempty"`
	CountryCode3   string `json:"country_code_3,omitempty"`
	EuropeanUnion  string `json:"european_union,omitempty"`
	CurrencyCode   string `json:"currency_code,omitempty"`
	CurrencyName   string `json:"currency_name,omitempty"`
	AnyCast        string `json:"anycast,omitempty"`
}

func (location *Location) String() string {
	fields := []string{
		location.CountryName, location.RegionName, location.CityName,
		location.OwnerDomain, location.ISPDomain, location.Longitude,
		location.Latitude, location.TimeZone, location.UTCOffset,
		location.ChinaAdminCode, location.IDDCode, location.CountryCode,
		location.ContinentCode, location.IDC, location.BaseStation,
		location.CountryCode3, location.EuropeanUnion, location.CurrencyCode,
		location.CurrencyName, location.AnyCast,
	}
	for index := range fields {
		if fields[index] == "" {
			fields[index] = "N/A"
		}
	}
	return strings.Join(fields, "\t")
}
