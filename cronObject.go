package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CronObj struct {
	nextPageToken string
	store         Storage
}

const (
	url            = "https://www.googleapis.com/youtube/v3/search"
	key            = "AIzaSyD6OUXH2GgxLwPI3qvyb2L2fox30WaETIs"
	part           = "snippet"
	typeValue      = "video"
	query          = "football"
	publishedAfter = "2024-03-01T00:00:00Z"
	order          = "date"
)

func NewCronObject(store Storage) *CronObj {

	return &CronObj{
		store: store,
	}
}

func (s *CronObj) FetchYoutubeData() {
	for range time.Tick(10 * time.Second) {
		requestUrl := s.generateRequestUrl()

		resp, err := http.Get(requestUrl)

		if err != nil {
			log.Fatal("Error while requesting youtube api data: ", err)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatal("Required status code from youtube is not 200", resp.StatusCode)
			return
		}

		var youtubeApiResponse YoutubeAPiResponse

		err = json.NewDecoder(resp.Body).Decode(&youtubeApiResponse)

		if err != nil {
			log.Fatal("Error in unmarshalling youtube api response: ", err)
			return
		}

		s.createDbEntry(youtubeApiResponse)

	}
}

func (s *CronObj) generateRequestUrl() string {

	requestUrl := url + "?" + "key=" + key + "&part=" + part + "&type=" + typeValue + "&query=" + query + "&publishedAfter=" + publishedAfter + "&order=" + order

	if s.nextPageToken != "" {
		requestUrl += "&pageToken=" + s.nextPageToken
	}

	log.Println("Request url: ", requestUrl)

	return requestUrl
}

func (s *CronObj) createDbEntry(youtubeApiResponse YoutubeAPiResponse) {
	var videoInfos []VideoInfo

	// log.Println("youtube Api resp: ", youtubeApiResponse)

	for i := 0; i < len(youtubeApiResponse.Items); i++ {
		var videoInfo VideoInfo
		videoInfo.Description = youtubeApiResponse.Items[i].Snippet.Description
		videoInfo.ThumbnailUrl = youtubeApiResponse.Items[i].Snippet.Thumbnails["default"].URL
		videoInfo.Title = youtubeApiResponse.Items[i].Snippet.Title
		videoInfo.DateTime = youtubeApiResponse.Items[i].Snippet.PublishedAt

		videoInfos = append(videoInfos, videoInfo)

	}

	err := s.store.AddInDb(videoInfos)

	if err != nil {
		log.Fatal("Error While storing in db: ", err)
	}
}
