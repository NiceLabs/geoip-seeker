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
	// Currency
	//   CurrencyCode = ISO 4217
	CurrencyCode string `json:",omitempty"`
	CurrencyName string `json:",omitempty"`
	// Service
	IDC         string `json:",omitempty"` // IDC | VPN
	BaseStation string `json:",omitempty"` // WiFi | BS (Base Station)
	// Other
	EuropeanUnion string `json:",omitempty"`
	AnyCast       string `json:",omitempty"`
}

func (r *Record) String() string {
	values := []string{
		r.BeginIP.String(),
		r.EndIP.String(),
		r.CountryName,
		r.RegionName,
	}
	for index := range values {
		if values[index] == "" {
			values[index] = "N/A"
		}
	}
	return strings.Join(values, "\t")
}
