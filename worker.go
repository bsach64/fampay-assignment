package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/bsach64/fampay-assignment/internal/database"
)

func (s *state) backgroundQuery(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.fetchAndWrite()
			if err != nil {
				log.Printf("could not query yt api: %v\n", err)
			}
		}
	}
}

func (s *state) fetchAndWrite() error {
	Data, err := s.ytClient.PublishedVideosByDate("cricket")
	if err != nil {
		return err
	}

	for _, item := range Data.Items {
		tb, err := json.Marshal(item.Snippet.Thumbnails)
		if err != nil {
			return err
		}

		addVideoParams := database.AddVideoParams{
			VideoID:      item.ID.VideoID,
			Title:        item.Snippet.Title,
			Description:  sql.NullString{String: item.Snippet.Description, Valid: true},
			PublishedAt:  item.Snippet.PublishedAt,
			ChannelID:    item.Snippet.ChannelID,
			Thumbnails:   tb,
			ChannelTitle: item.Snippet.ChannelTitle,
		}
		err = s.db.AddVideo(context.Background(), addVideoParams)
		if err != nil {
			return err
		}
	}
	return nil
}
