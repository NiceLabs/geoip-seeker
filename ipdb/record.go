package ipdb

import (
	"strings"

	"github.com/NiceLabs/geoip-seeker/shared"
)

func makeRecord(data string, language int, fields []string) *shared.Record {
	record := new(shared.Record)
	mapping := map[string]*string{
		"country_name":     &record.CountryName,
		"region_name":      &record.RegionName,
		"city_name":        &record.CityName,
		"owner_domain":     &record.OwnerDomain,
		"isp_domain":       &record.ISPDomain,
		"latitude":         &record.Latitude,
		"longitude":        &record.Longitude,
		"timezone":         &record.TimeZone,
		"utc_offset":       &record.UTCOffset,
		"china_admin_code": &record.ChinaAdminCode,
		"idd_code":         &record.IDDCode,
		"country_code":     &record.CountryCode,
		"continent_code":   &record.ContinentCode,
		"idc":              &record.IDC,
		"base_station":     &record.BaseStation,
		"country_code3":    &record.CountryCode3,
		"european_union":   &record.EuropeanUnion,
		"currency_code":    &record.CurrencyCode,
		"currency_name":    &record.CurrencyName,
		"anycast":          &record.AnyCast,
	}
	values := strings.Split(data, "\t")
	values = values[language : language+len(fields)]
	for index, value := range values {
		name := fields[index]
		*mapping[name] = value
	}
	return record
}