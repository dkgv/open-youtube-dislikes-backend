package types

type VideoDetails struct {
	Views       int64 `json:"views"`
	Likes       int64 `json:"likes"`
	Dislikes    int64 `json:"dislikes"`
	Comments    int64 `json:"comments"`
	PublishedAt int64 `json:"published_at"`
}

func (v VideoDetails) LikePerView() float64 {
	if v.Views == 0 {
		return 0
	}
	return float64(v.Likes) / float64(v.Views) / 100
}
