package ipdb

type meta struct {
	Build     int64          `json:"build"`
	IPVersion uint16         `json:"ip_version"`
	Languages map[string]int `json:"languages"`
	NodeCount uint32         `json:"node_count"`
	TotalSize uint32         `json:"total_size"`
	Fields    []string       `json:"fields"`
}
