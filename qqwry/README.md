# QQWay ip database seeker

## Notes

1. thread safe implementation
2. no cache (cache to be managed by yourself)
3. no encoding convert

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"github.com/NiceLabs/geoip-seeker/qqwry"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	data, _ := ioutil.ReadFile("testdata/qqwry.dat")
	seeker, _ := qqwry.New(data)

	record, _ := seeker.LookupByIP(net.ParseIP("114.114.114.114"))
	fromGBKtoUTF8(&record.CountryName)
	fromGBKtoUTF8(&record.RegionName)

	encodedRecord, _ := json.MarshalIndent(record, "", "  ")

	fmt.Println(seeker.String())
	// QQWry 2020-07-30 525793 [IPv4]
	fmt.Println(seeker.BuildTime())
	// 2020-07-30 00:00:00 +0800 CST
	fmt.Println(seeker.RecordCount())
	// 525793
	fmt.Println(string(encodedRecord))
	// {
	//   "IP": "114.114.114.114",
	//   "BeginIP": "114.114.114.114",
	//   "EndIP": "114.114.114.114",
	//   "CountryName": "江苏省南京市",
	//   "RegionName": "南京信风网络科技有限公司GreatbitDNS服务器"
	// }
}

func fromGBKtoUTF8(value *string) {
	reader := transform.NewReader(
		strings.NewReader(*value),
		simplifiedchinese.GBK.NewDecoder(),
	)
	data, _ := ioutil.ReadAll(reader)
	*value = string(data)
}
```

## Benchmark

```
$ go test --bench .
goos: darwin
goarch: amd64
pkg: github.com/NiceLabs/geoip-seeker/qqwry
BenchmarkIPSeeker_LookupByIP-12    	 2537727	       470 ns/op
PASS
ok  	github.com/NiceLabs/geoip-seeker/qqwry	2.325s
```

# References

1. https://web.archive.org/web/20140423114336/http://lumaqq.linuxsir.org/article/qqwry_format_detail.html
2. http://sewm.pku.edu.cn/src/other/qqwry/qqwry_format_detail.pdf
