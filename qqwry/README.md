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
	"strings"

	"github.com/NiceLabs/geoip-seeker/qqwry"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	data, _ := ioutil.ReadFile("testdata/qqwry.dat")
	seeker, _ := qqwry.New(data)

	record, _ := seeker.LookupByIP(net.ParseIP("114.114.114.114"))
	record.CountryName = gbkToUTF8(record.CountryName)
	record.RegionName = gbkToUTF8(record.RegionName)

	encodedRecord, _ := json.MarshalIndent(record, "", "  ")

	fmt.Println(seeker.RecordCount())
	// 470237
	fmt.Println(seeker.BuildTime())
	// 2018-10-05 00:00:00 +0800 CST
	fmt.Println(string(encodedRecord))
	// {
	//   "ip": "114.114.114.114",
	//   "begin_ip": "114.114.114.114",
	//   "end_ip": "114.114.114.114",
	//   "country_name": "江苏省南京市",
	//   "region_name": "南京信风网络科技有限公司GreatbitDNS服务器"
	// }
}

func gbkToUTF8(value string) string {
	reader := transform.NewReader(
		strings.NewReader(value),
		simplifiedchinese.GBK.NewDecoder(),
	)
	data, _ := ioutil.ReadAll(reader)
	return string(data)
}
```

# References

1. https://web.archive.org/web/20140423114336/http://lumaqq.linuxsir.org/article/qqwry_format_detail.html
2. http://sewm.pku.edu.cn/src/other/qqwry/qqwry_format_detail.pdf
