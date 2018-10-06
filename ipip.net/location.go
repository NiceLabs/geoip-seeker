package ipip_net

import (
	"strings"

	"github.com/NiceLabs/geoip-seeker/shared"
)

func makeLocation(data string) *shared.Location {
	if len(data) == 0 {
		return nil
	}

	location := new(shared.Location)

	mapping := []*string{
		&location.CountryName, &location.RegionName, &location.CityName,
		&location.OwnerDomain, &location.ISPDomain, &location.Longitude,
		&location.Latitude, &location.TimeZone, &location.UTCOffset,
		&location.ChinaAdminCode, &location.IDDCode, &location.CountryCode,
		&location.ContinentCode, &location.IDC, &location.BaseStation,
		&location.CountryCode3, &location.EuropeanUnion, &location.CurrencyCode,
		&location.CurrencyName, &location.AnyCast,
	}

	for index, field := range strings.Split(data, "\t") {
		*mapping[index] = field
	}

	return location
}
