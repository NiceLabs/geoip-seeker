# IPIP-DAT/DATX ip database seeker

## Notes

1. thread safe implementation
2. no cache (cache to be managed by your-self)

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/NiceLabs/geoip-seeker/ipip.net"
)

func main() {
	data, _ := ioutil.ReadFile("testdata/17monipdb.datx")
	seeker, _ := ipip_net.New(data, ipip_net.ModeDATX)
	record, _ := seeker.LookupByIP(net.ParseIP("114.114.114.114"))

	encodedRecord, _ := json.MarshalIndent(record, "", "  ")

	fmt.Println(seeker.RecordCount())
	// 251008
	fmt.Println(seeker.BuildTime())
	// 2018-07-02 01:00:00 +0800 CST
	fmt.Println(string(encodedRecord))
	// {
	//   "ip": "114.114.114.114",
	//   "begin_ip": "114.114.112.0",
	//   "end_ip": "114.114.119.255",
	//   "country_name": "114DNS.COM",
	//   "region_name": "114DNS.COM"
	// }
}
```

# References

1. https://ipip.net
2. https://github.com/Moowei/ip-seeker
3. https://github.com/larryli/ipv4

