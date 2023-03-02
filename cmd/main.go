package main

import (
	"fmt"
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

	// Get search query from arg (i.e. ./dehasher "@trevorarthur.com")
	query := os.Args[1]

	// Get email search results
	emailResults := pkg.EmailSearch(query, apiKey)

	for i := 0; i < len(emailResults.Entries); i++ {
		fmt.Println("----------------------------------")
		if len(emailResults.Entries[i].Name) > 0 {
			name := emailResults.Entries[i].Name
			fmt.Println("Name:", name)
		}
		if len(emailResults.Entries[i].Email) > 0 {
			email := emailResults.Entries[i].Email
			fmt.Println("Email:", email)
		}
		if len(emailResults.Entries[i].Username) > 0 {
			username := emailResults.Entries[i].Username
			fmt.Println("Username:", username)
		}
		if len(emailResults.Entries[i].Password) > 0 {
			password := emailResults.Entries[i].Password
			fmt.Println("Password:", password)
		}
		if len(emailResults.Entries[i].HashedPassword) > 0 {
			hashedPassword := emailResults.Entries[i].HashedPassword
			fmt.Println("Hashed Password:", hashedPassword)
		}
	}
}
