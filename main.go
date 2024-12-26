package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Could not load environment variables: %v\n", err)
	}

	// API_KEYS has a list of all valid api keys that can be used
	apiKeysStr := os.Getenv("API_KEYS")
	_ = strings.Split(apiKeysStr, ",")

}
