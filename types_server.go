package main

import "time"

type Video struct {
	VideoID      string    `json:"videoId"`
	ChannelID    string    `json:"channelId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Thumbnails   Tb        `json:"thumbnails"`
	ChannelTitle string    `json:"channelTitle"`
	PublishedAt  time.Time `json:"publishedAt"`
}

type Tb struct {
	Default struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"default"`
	Medium struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"medium"`
	High struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"high"`
}

type PaginatedVideos struct {
	NextPage int32   `json:"next_page"`
	Videos   []Video `json:"videos"`
}
