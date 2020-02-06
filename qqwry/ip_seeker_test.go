package qqwry

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"testing"
)

var seeker *Seeker

func init() {
	data, _ := ioutil.ReadFile("../testdata/qqwry.dat")
	seeker, _ = New(data)
}

func TestSeeker_Init(t *testing.T) {
	_, _ = New(nil)
}

func TestSeeker_Metadata(t *testing.T) {
	t.Log(seeker.String())
	t.Log(seeker.BuildTime())
	t.Log(seeker.RecordCount())
}

func TestSeeker_InvalidLookup(t *testing.T) {
	_, _ = seeker.LookupByIP(nil)
}

func TestSeeker_LookupByIP(t *testing.T) {
	cases := []net.IP{
		{0, 0, 0, 0},
		{127, 0, 0, 1},
		{172, 16, 0, 0},
		{192, 168, 0, 0},
		{100, 64, 0, 0},
		{156, 154, 114, 35},
		{157, 160, 206, 174},
		{220, 174, 130, 251},
		{255, 0, 0, 0},
		{255, 255, 255, 255},
	}
	for index, unit := range cases {
		record, err := seeker.LookupByIP(unit)
		if err != nil {
			t.Fatal(index, err)
		}
		fmt.Println(record, record.BeginIP, record.EndIP)
	}
}

func BenchmarkIPSeeker_LookupByIP(b *testing.B) {
	var ips []net.IP
	for i := 0; i < b.N; i++ {
		ip := make(net.IP, 4)
		_, _ = rand.Read(ip)
		ips = append(ips, ip)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for _, ip := range ips {
		_, _ = seeker.LookupByIP(ip)
	}
}
