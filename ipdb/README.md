# IPIP-IPDB ip database seeker

## Notes

1. thread safe implementation
2. no cache (cache to be managed by yourself)

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/NiceLabs/geoip-seeker/ipdb"
)

func main() {
	data, _ := ioutil.ReadFile("testdata/ipiptest.ipdb")
	seeker, _ := ipdb.New(data)

	record, _ := seeker.LookupByIP(net.ParseIP("114.114.114.114"))

	encodedRecord, _ := json.MarshalIndent(record, "", "  ")

	fmt.Println(seeker.String())
	// IPIP(IPDB) 2018-08-31 385083 [IPv4]
	fmt.Println(seeker.RecordCount())
	// 385083
	fmt.Println(seeker.BuildTime())
	// 2018-08-31 00:00:00 +0800 CST
	fmt.Println(string(encodedRecord))
	// {
	//   "IP": "114.114.114.114",
	//   "CountryName": "114DNS.COM",
	//   "RegionName": "114DNS.COM"
	// }
}
```

## Benchmark

```
$ go test --bench .
goos: darwin
goarch: amd64
pkg: github.com/NiceLabs/geoip-seeker/ipdb
BenchmarkIPSeeker_LookupByIP-12    	 3452234	       354 ns/op
PASS
ok  	github.com/NiceLabs/geoip-seeker/ipdb	2.952s
```

# References

1. https://ipip.net
2. https://github.com/larryli/ipv4

