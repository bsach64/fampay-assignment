package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/bsach64/fampay-assignment/internal/ytapi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Could not load environment variables: %v\n", err)
	}

	// API_KEYS has a list of all valid api keys that can be used
	apiKeysStr := os.Getenv("API_KEYS")
	apiKeys := strings.Split(apiKeysStr, ",")

	ytClient := ytapi.NewClient(apiKeys)
	start2024 := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	videoData, err := ytClient.PublishedVideosByDate("cricket", start2024)

	log.Println(videoData.Items)
}
