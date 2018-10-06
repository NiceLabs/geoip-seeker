# QQWay ip database seeker

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

	fmt.Println(seeker.RecordCount())
	// 385083
	fmt.Println(seeker.BuildTime())
	// 2018-08-31 14:17:20 +0800 CST
	fmt.Println(seeker.LanguageNames())
	// [CN]
	fmt.Println(string(encodedRecord))
	// {
	//   "ip": "114.114.114.114",
	//   "country_name": "114DNS.COM",
	//   "region_name": "114DNS.COM"
	// }
}
```

# References

see https://ipip.net
