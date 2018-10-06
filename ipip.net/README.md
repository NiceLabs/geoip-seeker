# IPIP.Net ip database seeker

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
	location, _ := seeker.LookupByIP(net.ParseIP("103.57.164.0"))

	encoded, _ := json.MarshalIndent(location, "", "  ")

	fmt.Println(string(encoded))
}
```

# References

see https://ipip.net
