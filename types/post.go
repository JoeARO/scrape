package types

type Posts []Post

type Post struct {
	Kind string   `json:"kind"`
	Data PostData `json:"data"`
}

type PostData struct {
	RestrictedContent bool    `json:"over_18"`
	Score             int     `json:"score"`
	Pinned            bool    `json:"pinned"`
	Permalink         string  `json:"permalink"`
	Author            string  `json:"author"`
	Thumbnail         string  `json:"thumbnail"`
	UpVotes           int     `json:"ups"`
	DownVotes         int     `json:"downs"`
	Url               string  `json:"url"`
	Title             string  `json:"title"`
	CreatedAt         float64 `json:"created_utc"`
	Domain            string  `json:"domain"`
	Media             Media   `json:"media"`
	Preview           Preview `jsons:"preview"`
}

type Media struct {
	Content RedditVideo `json:"reddit_video"`
}

type RedditVideo struct {
	Url    string `json:"fallback_url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	IsGif  bool   `json:"is_gif"`
}

type Preview struct {
	Images []Image `json:"images"`
}

type Image struct {
	Source SourceImage `json:"source"`
}

type SourceImage struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
