package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/bsach64/fampay-assignment/internal/database"
)

func backgroundQuery(s state, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := fetchAndWrite(s)
			if err != nil {
				log.Printf("could not query yt api: %v\n", err)
			}
		}
	}
}

func fetchAndWrite(s state) error {
	start2024 := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	Data, err := s.ytClient.PublishedVideosByDate("cricket", start2024)
	if err != nil {
		return err
	}

	for _, item := range Data.Items {
		tb, err := json.Marshal(item.Snippet.Thumbnails)
		if err != nil {
			return err
		}

		addVideoParams := database.AddVideoParams{
			VideoID:     item.ID.VideoID,
			Title:       item.Snippet.Title,
			Description: sql.NullString{String: item.Snippet.Description, Valid: true},
			PublishedAt: item.Snippet.PublishedAt,
			ChannelID:   item.Snippet.ChannelID,
			Thumbnails:  tb,
		}
		err = s.db.AddVideo(context.Background(), addVideoParams)
		if err != nil {
			return err
		}
	}
	return nil
}
