package ipdb

import (
	"io/ioutil"
	"net"
	"testing"
)

var seeker *IPSeeker

func init() {
	data, _ := ioutil.ReadFile("../testdata/ipiptest.ipdb")
	seeker, _ = New(data)
}

func BenchmarkIPSeeker_LookupByIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = seeker.LookupByIP(net.IP{103, 57, 164, 0})
	}
}
