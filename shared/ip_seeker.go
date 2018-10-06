package shared

import (
	"net"
	"time"
)

type IPSeeker interface {
	LookupByIP(address net.IP) (*Record, error)
	IPv4Support() bool
	IPv6Support() bool
	RecordCount() uint64
	BuildTime() time.Time
}
