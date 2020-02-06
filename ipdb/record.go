package ipdb

import (
	"strings"

	. "github.com/NiceLabs/geoip-seeker/shared"
)

func makeRecord(data string, language uint16, fields []string) (record *Record) {
	record = new(Record)
	values := strings.Split(data, "\t")
	values = values[language : language+uint16(len(fields))]
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
		"idd_code":         &record.IDDCode,
		"china_admin_code": &record.GB2260Code,
		"country_code":     &record.ISO3166Alpha2Code,
		"country_code3":    &record.ISO3166Alpha3Code,
		"continent_code":   &record.ContinentCode,
		"idc":              &record.IDC,
		"base_station":     &record.BaseStation,
		"currency_code":    &record.CurrencyCode,
		"currency_name":    &record.CurrencyName,
		"european_union":   &record.EuropeanUnion,
		"anycast":          &record.AnyCast,
	}
	for index, end := language, language+uint16(len(fields)); index < end; index++ {
		if input, ok := mapping[fields[index]]; ok {
			*input = values[index]
		}
	}
	return record
}
