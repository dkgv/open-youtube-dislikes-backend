package types

import "time"

type Video struct {
	IDHash      string  `json:"idHash,omitempty"`
	Views       int64   `json:"views,omitempty"`
	Likes       int64   `json:"likes,omitempty"`
	Dislikes    int64   `json:"dislikes,omitempty"`
	Comments    int64   `json:"comments,omitempty"`
	PublishedAt int64   `json:"publishedAt,omitempty"`
	Subscribers int64   `json:"subscribers,omitempty"`
	DurationSec int32   `json:"durationSec,omitempty"`
	Positive    float64 `json:"positive,omitempty"`
	Negative    float64 `json:"negative,omitempty"`
	Neutral     float64 `json:"neutral,omitempty"`
	Compound    float64 `json:"compound,omitempty"`
}

func (v Video) ViewsPerLike() float64 {
	if v.Likes == 0 {
		return 0
	}
	return float64(v.Views) / float64(v.Likes) / 100
}

func (v Video) DaysSincePublish() int32 {
	timestamp := time.Unix(0, v.PublishedAt*int64(time.Millisecond))
	now := time.Now()
	return int32(now.Sub(timestamp).Hours() / 24)
}

func (v Video) ViewsPerComment() float64 {
	if v.Comments == 0 {
		return 0
	}
	return float64(v.Views) / float64(v.Comments) / 100
}

func (v Video) LikesPerComment() float64 {
	if v.Comments == 0 {
		return 0
	}
	return float64(v.Likes) / float64(v.Comments) / 100
}

func (v Video) DaysPerLike() float64 {
	if v.Likes == 0 {
		return 0
	}
	return float64(v.DaysSincePublish()) / float64(v.Likes)
}

func (v Video) DaysPerComment() float64 {
	if v.Comments == 0 {
		return 0
	}
	return float64(v.DaysSincePublish()) / float64(v.Comments)
}

func (v Video) SubscribersPerLike() float64 {
	if v.Likes == 0 {
		return 0
	}
	return float64(v.Subscribers) / float64(v.Likes) / 100
}
