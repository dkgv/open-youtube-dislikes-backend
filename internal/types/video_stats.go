package types

type VideoDetails struct {
	Views       int64 `json:"views"`
	Likes       int64 `json:"likes"`
	Dislikes    int64 `json:"dislikes"`
	Comments    int64 `json:"comments"`
	PublishedAt int64 `json:"published_at"`
}
