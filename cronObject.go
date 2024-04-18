package main

import "database/sql"

type CronObj struct {
	db            *sql.DB
	nextPageToken string
}

var url = "https://www.googleapis.com/youtube/v3/search"

func (s *CronObj) fetchYoutubeData() {

}
