package ipdb

import "time"

type MetaData struct {
	Build     int64          `json:"build"`
	IPVersion uint16         `json:"ip_version"`
	Languages map[string]int `json:"languages"`
	NodeCount int            `json:"node_count"`
	TotalSize int            `json:"total_size"`
	Fields    []string       `json:"fields"`
}

func (m *MetaData) IPv4Support() bool {
	version := uint16(0x01)
	return (m.IPVersion & version) == version
}

func (m *MetaData) IPv6Support() bool {
	version := uint16(0x02)
	return (m.IPVersion & version) == version
}

func (m *MetaData) BuildDate() time.Time {
	zone := time.FixedZone("CST", +8*3600)
	return time.Unix(m.Build, 0).In(zone)
}

func (m *MetaData) LanguageNames() (languages []string) {
	for k := range m.Languages {
		languages = append(languages, k)
	}
	return
}