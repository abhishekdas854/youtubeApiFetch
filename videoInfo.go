package main

import "time"

type VideoInfo struct {
	Id           int64
	Title        string
	Description  string
	ThumbnailUrl string
	DateTime     time.Time
}
