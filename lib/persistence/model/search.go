package model

// Response :nodoc
type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int        `json:"total"`
		MaxScore float32    `json:"max_score"`
		Hits     []*HitInfo `json:"hits"`
	} `json:"hits"`
}

// HitInfo :nodoc:
type HitInfo struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	ID     string                 `json:"_id"`
	Score  float32                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}
