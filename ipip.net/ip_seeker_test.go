package ipip_net

import (
	"io/ioutil"
	"net"
	"testing"
)

var seeker *IPSeeker

func init() {
	data, _ := ioutil.ReadFile("../testdata/17monipdb.datx")
	seeker, _ = New(data, ModeDATX)
}

func BenchmarkIPSeeker_LookupByIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = seeker.LookupByIP(net.IP{103, 57, 164, 0})
	}
}

func BenchmarkIPSeeker_LookupByIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = seeker.LookupByIndex(0)
	}
}
