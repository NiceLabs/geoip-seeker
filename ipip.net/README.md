# IPIP.Net ip database seeker

## Notes

1. thread safe implementation
2. no cache (cache to be managed by yourself)

## Example

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/xtomcom/geoip-seeker/ipip.net"
)

func main() {
	data, _ := ioutil.ReadFile("./17monipdb.datx")
	seeker := ipip_net.NewDATX(data)
	location, _ := seeker.LookupByIP(net.ParseIP("103.57.164.0"))

	fmt.Println("String()\t\t:", location.String())
	fmt.Println()
	fmt.Println("IP\t\t\t:", location.IP)
	fmt.Println("Country\t\t\t:", location.Country)
	fmt.Println("Province\t\t:", location.Province)
	fmt.Println("City\t\t\t:", location.City)
	fmt.Println("Unit\t\t\t:", location.Unit)
	fmt.Println()
	fmt.Println("ISP\t\t\t:", location.ISP)
	fmt.Println("Longitude\t\t:", location.Longitude)
	fmt.Println("Latitude\t\t:", location.Latitude)
	fmt.Println("TimeZoneCode\t\t:", location.TimeZoneCode)
	fmt.Println("TimeZoneUTC\t\t:", location.TimeZoneUTC)
	fmt.Println("GB2260Code\t\t:", location.GB2260Code)
	fmt.Println()
	fmt.Println("IPIP.Net RecordCount\t:", seeker.RecordCount())
	fmt.Println("IPIP.Net PublishDate\t:", seeker.PublishDate())
}

```

# References

see https://ipip.net
