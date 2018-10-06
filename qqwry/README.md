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
	toUTF8(&record.CountryName)
	toUTF8(&record.RegionName)

	encodedRecord, _ := json.MarshalIndent(record, "", "  ")

	fmt.Println(seeker.RecordCount())
	// 470237
	fmt.Println(seeker.BuildTime())
	// 2018-10-05 00:00:00 +0800 CST
	fmt.Println(string(encodedRecord))
	// {
	//   "IP": "114.114.114.114",
	//   "BeginIP": "114.114.114.114",
	//   "EndIP": "114.114.114.114",
	//   "CountryName": "江苏省南京市",
	//   "RegionName": "南京信风网络科技有限公司GreatbitDNS服务器"
	// }
}

func toUTF8(value *string) {
	reader := transform.NewReader(
		strings.NewReader(*value),
		simplifiedchinese.GBK.NewDecoder(),
	)
	data, _ := ioutil.ReadAll(reader)
	*value = string(data)
}
```

# References

1. https://web.archive.org/web/20140423114336/http://lumaqq.linuxsir.org/article/qqwry_format_detail.html
2. http://sewm.pku.edu.cn/src/other/qqwry/qqwry_format_detail.pdf
