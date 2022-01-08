package youtube

type VideosListResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []struct {
		Kind           string `json:"kind"`
		Etag           string `json:"etag"`
		Id             string `json:"id"`
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
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}
