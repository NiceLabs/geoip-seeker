package ipdb

import (
	"strings"

	"github.com/NiceLabs/geoip-seeker/shared"
)

func makeRecord(data string, language int, fields []string) *shared.Record {
	record := new(shared.Record)
	values := strings.Split(data, "\t")
	values = values[language : language+len(fields)]
	for index, value := range values {
		switch fields[index] {
		case "country_name":
			record.CountryName = value
		case "region_name":
			record.RegionName = value
		case "city_name":
			record.CityName = value
		case "owner_domain":
			record.OwnerDomain = value
		case "isp_domain":
			record.ISPDomain = value
		case "latitude":
			record.Latitude = value
		case "longitude":
			record.Longitude = value
		case "timezone":
			record.TimeZone = value
		case "utc_offset":
			record.UTCOffset = value
		case "idd_code":
			record.IDDCode = value
		case "china_admin_code":
			record.GB2260Code = value
		case "country_code":
			record.ISO3166Alpha2Code = value
		case "country_code3":
			record.ISO3166Alpha3Code = value
		case "continent_code":
			record.ContinentCode = value
		case "idc":
			record.IDC = value
		case "base_station":
			record.BaseStation = value
		case "currency_code":
			record.CurrencyCode = value
		case "currency_name":
			record.CurrencyName = value
		case "european_union":
			record.EuropeanUnion = value
		case "anycast":
			record.AnyCast = value
		}
	}
	return record
}
