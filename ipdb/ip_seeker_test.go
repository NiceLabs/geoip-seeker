package ipdb

import (
	"crypto/rand"
	"io/ioutil"
	"net"
	"testing"
)

var seeker *Seeker

func init() {
	data, _ := ioutil.ReadFile("../testdata/ipiptest.ipdb")
	seeker, _ = New(data)
}

func TestMetadata(t *testing.T) {
	t.Log(seeker.String())
	t.Log(seeker.LanguageNames())
	t.Log(seeker.LookupByIP(net.IPv4(1, 1, 1, 1)))
	_ = seeker.LanguageCode("CN")
}

func BenchmarkIPSeeker_LookupByIP(b *testing.B) {
	var items []net.IP
	for i := 0; i < b.N; i++ {
		ip := make(net.IP, 4)
		_, _ = rand.Read(ip)
		items = append(items, ip)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for _, item := range items {
		_, _ = seeker.LookupByIP(item)
	}
}
