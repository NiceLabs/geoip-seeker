package ipdb

import "time"

type meta struct {
	Build     int64          `json:"build"`
	IPVersion uint16         `json:"ip_version"`
	Languages map[string]int `json:"languages"`
	NodeCount int            `json:"node_count"`
	TotalSize int            `json:"total_size"`
	Fields    []string       `json:"fields"`
}

func (m *meta) IPv4Support() bool {
	return (m.IPVersion & 0x01) == 0x01
}

func (m *meta) IPv6Support() bool {
	return (m.IPVersion & 0x02) == 0x02
}

func (m *meta) BuildDate() time.Time {
	zone := time.FixedZone("CST", +8*3600)
	return time.Unix(m.Build, 0).In(zone)
}

func (m *meta) LanguageNames() (languages []string) {
	for k := range m.Languages {
		languages = append(languages, k)
	}
	return
}
