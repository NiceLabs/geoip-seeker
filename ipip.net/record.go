package ipip_net

import (
	"strings"

	"github.com/NiceLabs/geoip-seeker/shared"
)

func makeRecord(data string) *shared.Record {
	if len(data) == 0 {
		return nil
	}

	record := new(shared.Record)

	mapping := []*string{
		&record.CountryName, &record.RegionName, &record.CityName,
		&record.OwnerDomain, &record.ISPDomain, &record.Longitude,
		&record.Latitude, &record.TimeZone, &record.UTCOffset,
		&record.GB2260Code, &record.IDDCode, &record.ISO3166Alpha2Code,
		&record.ContinentCode, &record.IDC, &record.BaseStation,
		&record.ISO3166Alpha3Code, &record.EuropeanUnion, &record.CurrencyCode,
		&record.CurrencyName, &record.AnyCast,
	}

	for index, field := range strings.Split(data, "\t") {
		*mapping[index] = field
	}

	return record
}
