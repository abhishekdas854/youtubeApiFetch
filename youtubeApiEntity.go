package main

type YoutubeAPiResponse struct {
	NextPageToken string `json:"nextPageToken"`
	PrevPageToken string `json:""prevPageToken`
}

type Item struct {
}

type Snippet struct {
	PublishedAt string `json:"publishedAt"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
