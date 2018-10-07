package shared

import (
	"fmt"
	"net"
	"time"
)

type IPSeeker interface {
	LookupByIP(address net.IP) (*Record, error)
	IPv4Support() bool
	IPv6Support() bool
	RecordCount() int
	BuildTime() time.Time
	fmt.Stringer
}

type Update interface {
	BuildTime() time.Time
	Size() uint32
	Download() (data []byte, err error)
}
