package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bsach64/fampay-assignment/internal/database"
)

func writeErrorMsg(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	errMsg := struct {
		ErrorMsg string `json:"error"`
	}{
		ErrorMsg: msg,
	}
	json.NewEncoder(w).Encode(errMsg)
}

func (s *state) handleGetVideos(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			writeErrorMsg(w, err.Error(), 400)
			return
		}
		if page <= 0 {
			writeErrorMsg(w, "page should be greater than 0", 400)
			return
		}
	}

	rows, err := s.db.GetVideos(context.Background(), database.GetVideosParams{Offset: int32((page - 1) * 5), Limit: 5})
	if err != nil {
		writeErrorMsg(w, err.Error(), 500)
		return
	}

	var PVideos PaginatedVideos
	PVideos.Videos = make([]Video, 0)
	PVideos.NextPage = int32(page + 1)

	for _, r := range rows {
		var tb Tb
		err := json.Unmarshal(r.Thumbnails, &tb)
		if err != nil {
			writeErrorMsg(w, err.Error(), 500)
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
