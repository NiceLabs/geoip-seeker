package shared

import (
	"fmt"
	"net"
	"time"
)

type IPSeeker interface {
	LookupByIP(address net.IP) (*Record, error)
	LookupByIndex(index uint64) (*Record, error)
	IPv4Support() bool
	IPv6Support() bool
	RecordCount() uint64
	BuildTime() time.Time
	fmt.Stringer
}

type Update interface {
	BuildTime() time.Time
	Size() uint64
	Download() (data []byte, err error)
}
