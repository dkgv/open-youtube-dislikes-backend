package types

import "time"

type Video struct {
	IDHash      string `json:"idHash,omitempty"`
	Views       int64  `json:"views,omitempty"`
	Likes       int64  `json:"likes,omitempty"`
	Dislikes    int64  `json:"dislikes,omitempty"`
	Comments    *int64 `json:"comments,omitempty"`
	PublishedAt int64  `json:"publishedAt,omitempty"`
	Subscribers int64  `json:"subscribers,omitempty"`
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
	return float64(*v.Comments) / float64(v.Views) / 100
}
