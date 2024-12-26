package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bsach64/fampay-assignment/internal/database"
)

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
	Videos []Video `json:"videos"`
}

func (s *state) handleGetVideos(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.GetVideos(context.Background(), database.GetVideosParams{ID: 0, Limit: 5})
	if err != nil {
		w.WriteHeader(500)
		errMsg := struct {
			ErrorMsg string `json:"error"`
		}{
			ErrorMsg: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	var PVideos PaginatedVideos
	PVideos.Videos = make([]Video, 0)

	for _, r := range rows {
		var tb Tb
		err := json.Unmarshal(r.Thumbnails, &tb)
		if err != nil {
			w.WriteHeader(500)
			errMsg := struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: fmt.Sprintf("%v", err),
			}
			json.NewEncoder(w).Encode(errMsg)
			return
		}
		v := Video{
			ChannelID:    r.ChannelID,
			VideoID:      r.VideoID,
			ChannelTitle: r.ChannelTitle,
			Title:        r.Title,
			Description:  r.Description.String,
			PublishedAt:  r.PublishedAt,
			Thumbnails:   tb,
		}
		PVideos.Videos = append(PVideos.Videos, v)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PVideos)
}

func (s *state) server() {
	http.HandleFunc("GET /videos", s.handleGetVideos)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Print(err)
	}
}
