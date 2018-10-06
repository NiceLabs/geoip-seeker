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
	location, _ := seeker.LookupByIP(net.ParseIP("103.57.164.0"))

	encoded, _ := json.MarshalIndent(location, "", "  ")

	fmt.Println(string(encoded))
}
```

# References

see https://ipip.net
