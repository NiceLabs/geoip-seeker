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

	location, _ := seeker.LookupByIP(net.ParseIP("103.57.164.0"))
	location.CountryName = convertGBKToUTF8(location.CountryName)
	location.RegionName = convertGBKToUTF8(location.RegionName)

	encoded, _ := json.MarshalIndent(location, "", "  ")

	fmt.Println(string(encoded))
}

func convertGBKToUTF8(value string) string {
	reader := transform.NewReader(
		strings.NewReader(value),
		simplifiedchinese.GBK.NewDecoder(),
	)
	data, _ := ioutil.ReadAll(reader)
	return string(data)
}
```

# References

see http://sewm.pku.edu.cn/src/other/qqwry/qqwry_format_detail.pdf
