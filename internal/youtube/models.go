package youtube

import "time"

type BaseResponse struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

type VideosListResponse struct {
	BaseResponse
	Items []VideoItem `json:"items"`
}

type VideoItem struct {
	Kind           string       `json:"kind"`
	Etag           string       `json:"etag"`
	Id             string       `json:"id"`
	Snippet        VideoSnippet `json:"snippet"`
	ContentDetails struct {
		Duration        string `json:"duration"`
		Dimension       string `json:"dimension"`
		Definition      string `json:"definition"`
		Caption         string `json:"caption"`
		LicensedContent bool   `json:"licensedContent"`
		ContentRating   struct {
		} `json:"contentRating"`
		Projection string `json:"projection"`
	} `json:"contentDetails"`
	Statistics struct {
		ViewCount     string `json:"viewCount"`
		LikeCount     string `json:"likeCount"`
		FavoriteCount string `json:"favoriteCount"`
		CommentCount  string `json:"commentCount"`
	} `json:"statistics"`
}

type VideoSnippet struct {
	PublishedAt time.Time `json:"publishedAt"`
	ChannelId   string    `json:"channelId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnails  struct {
		Default struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"default"`
		Medium struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"medium"`
		High struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"high"`
		Standard struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"standard"`
		Maxres struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"maxres"`
	} `json:"thumbnails"`
	ChannelTitle         string   `json:"channelTitle"`
	Tags                 []string `json:"tags"`
	CategoryId           string   `json:"categoryId"`
	LiveBroadcastContent string   `json:"liveBroadcastContent"`
	Localized            struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"localized"`
}

type ChannelsListResponse struct {
	BaseResponse
	Items []ChannelItem `json:"items"`
}

type ChannelItem struct {
	Kind       string `json:"kind"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Statistics struct {
		ViewCount             string `json:"viewCount"`
		SubscriberCount       string `json:"subscriberCount"`
		HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
		VideoCount            string `json:"videoCount"`
	} `json:"statistics"`
}

type CommentThreadResponse struct {
	BaseResponse
	Comments []Comment `json:"items"`
}

type Comment struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	Id      string `json:"id"`
	Snippet struct {
		VideoId         string `json:"videoId"`
		TopLevelComment struct {
			Kind    string `json:"kind"`
			Etag    string `json:"etag"`
			Id      string `json:"id"`
			Snippet struct {
				VideoId               string `json:"videoId"`
				TextDisplay           string `json:"textDisplay"`
				TextOriginal          string `json:"textOriginal"`
				AuthorDisplayName     string `json:"authorDisplayName"`
				AuthorProfileImageUrl string `json:"authorProfileImageUrl"`
				AuthorChannelUrl      string `json:"authorChannelUrl"`
				AuthorChannelId       struct {
					Value string `json:"value"`
				} `json:"authorChannelId"`
				CanRate      bool      `json:"canRate"`
				ViewerRating string    `json:"viewerRating"`
				LikeCount    int       `json:"likeCount"`
				PublishedAt  time.Time `json:"publishedAt"`
				UpdatedAt    time.Time `json:"updatedAt"`
			} `json:"snippet"`
		} `json:"topLevelComment"`
		CanReply        bool `json:"canReply"`
		TotalReplyCount int  `json:"totalReplyCount"`
		IsPublic        bool `json:"isPublic"`
	} `json:"snippet"`
}
