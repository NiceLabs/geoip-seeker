# QQWay ip database seeker

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
	"strings"

	"github.com/xtomcom/geoip-seeker/qqway"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	data, _ := ioutil.ReadFile(qqwry)
	seeker, _ := qqway.New(data)
	location, _ := seeker.LookupByIP(net.ParseIP("103.57.164.0"))

	fmt.Println("String()\t\t:", gbkToUTF8([]byte(location.String())))
	fmt.Println()
	fmt.Println("IP Range\t\t:", location.BeginIP, "-", location.EndIP)
	fmt.Println("Country\t\t\t:", gbkToUTF8(location.Country))
	fmt.Println("Area\t\t\t:", gbkToUTF8(location.Area))
	fmt.Println()
	fmt.Println("QQWay RecordCount\t:", seeker.RecordCount())
	fmt.Println("QQWay Version\t\t:", gbkToUTF8(seeker.Version()))
}

func gbkToUTF8(value string) string {
	reader := transform.NewReader(strings.NewReader(value), simplifiedchinese.GBK.NewDecoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return ""
	}
	return string(data)
}
```

# References

see http://sewm.pku.edu.cn/src/other/qqwry/qqwry_format_detail.pdf
