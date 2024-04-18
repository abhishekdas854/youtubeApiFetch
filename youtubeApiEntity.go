package main

import "time"

// type YoutubeAPiResponse struct {
// 	NextPageToken string    `json:"nextPageToken"`
// 	PrevPageToken string    `json:"prevPageToken"`
// 	Items         []Snippet `json:"items"`
// }

// type Snippet struct {
// 	PublishedAt string    `json:"publishedAt"`
// 	Description string    `json:"description"`
// 	Title       string    `json:"title"`
// 	Thumbnails  Thumbnail `json:"thumbnails"`
// }

// type Thumbnail struct {
// 	ThumbnailUrl string `json:"thumbnailUrl"`
// }

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

// Thumbnails struct represents the "thumbnails" object
type Thumbnails struct {
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

// ID struct represents the "id" object
type ID struct {
	Kind       string `json:"kind"`
	VideoID    string `json:"videoId"`
	ChannelID  string `json:"channelId"`
	PlaylistID string `json:"playlistId"`
}

// Snippet struct represents the "snippet" object
type Snippet struct {
	PublishedAt          time.Time             `json:"publishedAt"`
	ChannelID            string                `json:"channelId"`
	Title                string                `json:"title"`
	Description          string                `json:"description"`
	Thumbnails           map[string]Thumbnails `json:"thumbnails"`
	ChannelTitle         string                `json:"channelTitle"`
	LiveBroadcastContent string                `json:"liveBroadcastContent"`
}

// Item struct represents the "items" array
type Item struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

// SearchResult struct represents the root object
type YoutubeAPiResponse struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	PrevPageToken string   `json:"prevPageToken"`
	RegionCode    string   `json:"regionCode"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Item   `json:"items"`
}
