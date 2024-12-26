package main

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bsach64/fampay-assignment/internal/database"
	"github.com/bsach64/fampay-assignment/internal/ytapi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("could not load environment variables: %v\n", err)
	}

	// API_KEYS has a list of all valid api keys that can be used
	apiKeysStr := os.Getenv("API_KEYS")
	apiKeys := strings.Split(apiKeysStr, ",")

	dbURL := os.Getenv("DATABASE")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("could not establish connection to db: %v\n", err)
	}

	_ = database.New(db)

	ytClient := ytapi.NewClient(apiKeys)
	start2024 := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	videoData, err := ytClient.PublishedVideosByDate("cricket", start2024)

	log.Println(videoData.Items)
}
