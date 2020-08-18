package shared

import (
	"strconv"
	"strings"
)

func ShowLibraryInfo(name string, seeker IPSeeker) string {
	items := []string{
		name,
		seeker.BuildTime().Format("2006-01-02"),
		strconv.FormatUint(seeker.RecordCount(), 10),
	}
	if seeker.IPv4Support() {
		items = append(items, "[IPv4]")
	}
	if seeker.IPv6Support() {
		items = append(items, "[IPv6]")
	}
	return strings.Join(items, " ")
}
