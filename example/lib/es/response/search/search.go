package search

type Body[T any] struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Shards   BodyShards  `json:"_shards"`
	Hits     BodyHits[T] `json:"hits"`
}

type BodyShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type BodyHits[T any] struct {
	Total    BodyHitsTotal     `json:"total"`
	MaxScore interface{}       `json:"max_score"`
	Hits     []BodyHitsHits[T] `json:"hits"`
}

type BodyHitsTotal struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type BodyHitsHits[T any] struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	Id     string      `json:"_id"`
	Score  interface{} `json:"_score"`
	Source T           `json:"_source"`
	Sort   []int64     `json:"sort"`
}
