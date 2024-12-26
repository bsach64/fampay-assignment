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

type state struct {
	ytClient ytapi.Client
	db       *database.Queries
}

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

	stat := state{
		ytClient: ytapi.NewClient(apiKeys),
		db:       database.New(db),
	}

	go backgroundQuery(stat, 10*time.Second)
	select {}
}
