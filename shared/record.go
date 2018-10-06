package shared

import (
	"net"
	"strings"
)

type Record struct {
	// IP
	IP      net.IP `json:",omitempty"`
	BeginIP net.IP `json:",omitempty"`
	EndIP   net.IP `json:",omitempty"`
	// GeoInfo
	CountryName string `json:",omitempty"`
	RegionName  string `json:",omitempty"`
	CityName    string `json:",omitempty"`
	// Owner
	OwnerDomain string `json:",omitempty"`
	ISPDomain   string `json:",omitempty"`
	// Geocoding
	Latitude  string `json:",omitempty"`
	Longitude string `json:",omitempty"`
	// Time zone
	TimeZone  string `json:",omitempty"`
	UTCOffset string `json:",omitempty"`
	// Country Code
	//   IDDCode = International Direct Dialing
	//   GB2260Code = GB/T 2260
	//   ISO3166Alpha2Code = ISO 3166-1 alpha-2
	//   ISO3166Alpha3Code = ISO 3166-1 alpha-3
	IDDCode           string `json:",omitempty"`
	GB2260Code        string `json:",omitempty"`
	ISO3166Alpha2Code string `json:",omitempty"`
	ISO3166Alpha3Code string `json:",omitempty"`
	ContinentCode     string `json:",omitempty"`
	// Service
	IDC           string `json:",omitempty"`
	BaseStation   string `json:",omitempty"`
	EuropeanUnion string `json:",omitempty"`
	CurrencyCode  string `json:",omitempty"`
	CurrencyName  string `json:",omitempty"`
	// BGP
	AnyCast string `json:",omitempty"`
}

func (record *Record) String() string {
	fields := []string{
		record.CountryName, record.RegionName, record.CityName,
		record.OwnerDomain, record.ISPDomain, record.Longitude,
		record.Latitude, record.TimeZone, record.UTCOffset,
		record.GB2260Code, record.IDDCode, record.ISO3166Alpha2Code,
		record.ContinentCode, record.IDC, record.BaseStation,
		record.ISO3166Alpha3Code, record.EuropeanUnion, record.CurrencyCode,
		record.CurrencyName, record.AnyCast,
	}
	for index := range fields {
		if fields[index] == "" {
			fields[index] = "N/A"
		}
	}
	return strings.Join(fields, "\t")
}
