package ipdb

type meta struct {
	Build     int64            `json:"build"`
	IPVersion uint8            `json:"ip_version"`
	Languages map[string]uint8 `json:"languages"`
	NodeCount int              `json:"node_count"`
	TotalSize int              `json:"total_size"`
	Fields    []string         `json:"fields"`
}
