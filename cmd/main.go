package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/trevor-arthur/dehasher/pkg"
)

func main() {

	// Load .env file
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// Get API key from .env
	apiKey := os.Getenv("DEHASHED_API_KEY")
	apiEmail := os.Getenv("DEHASHED_API_EMAIL")

	// Get search query from arg (i.e. ./dehasher "email:@trevorarthur.com")
	query := os.Args[1]

	// Get email search results
	results := pkg.Search(query, apiEmail, apiKey)

	// Print email search results
	pkg.PrintResults(results)

}
