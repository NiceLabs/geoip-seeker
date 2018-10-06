package ipdb

import (
	"strings"

	"github.com/NiceLabs/geoip-seeker/shared"
)

func makeLocation(data string, language int, fields []string) *shared.Location {
	location := new(shared.Location)
	mapping := map[string]*string{
		"country_name":     &location.CountryName,
		"region_name":      &location.RegionName,
		"city_name":        &location.CityName,
		"owner_domain":     &location.OwnerDomain,
		"isp_domain":       &location.ISPDomain,
		"latitude":         &location.Latitude,
		"longitude":        &location.Longitude,
		"timezone":         &location.TimeZone,
		"utc_offset":       &location.UTCOffset,
		"china_admin_code": &location.ChinaAdminCode,
		"idd_code":         &location.IDDCode,
		"country_code":     &location.CountryCode,
		"continent_code":   &location.ContinentCode,
		"idc":              &location.IDC,
		"base_station":     &location.BaseStation,
		"country_code3":    &location.CountryCode3,
		"european_union":   &location.EuropeanUnion,
		"currency_code":    &location.CurrencyCode,
		"currency_name":    &location.CurrencyName,
		"anycast":          &location.AnyCast,
	}
	values := strings.Split(data, "\t")
	values = values[language : language+len(fields)]
	for index, value := range values {
		name := fields[index]
		*mapping[name] = value
	}
	return location
}
