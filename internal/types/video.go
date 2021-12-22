package types

import "time"

type Video struct {
	ID          string `json:"id,omitempty"`
	IDHash      string `json:"id_hash,omitempty"`
	Views       uint32 `json:"views,omitempty"`
	Likes       uint32 `json:"likes,omitempty"`
	Dislikes    uint32 `json:"dislikes,omitempty"`
	Comments    uint32 `json:"comments,omitempty"`
	PublishedAt int64  `json:"published_at,omitempty"`
	Subscribers uint32 `json:"subscribers,omitempty"`
}

func (v Video) LikesPerView() float64 {
	if v.Views == 0 {
		return 0
	}
	return float64(v.Likes) / float64(v.Views) / 100
}

func (v Video) DaysSincePublish() int32 {
	timestamp := time.Unix(0, v.PublishedAt*int64(time.Millisecond))
	now := time.Now()
	return int32(now.Sub(timestamp).Hours() / 24)
}

func (v Video) CommentsPerView() float64 {
	if v.Views == 0 {
		return 0
	}
	return float64(v.Comments) / float64(v.Views) / 100
}
