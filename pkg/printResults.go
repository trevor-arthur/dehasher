package pkg

import (
	"fmt"

	"github.com/trevor-arthur/dehasher/models"
)

func PrintResults(results models.SearchJSON) {
	for i := 0; i < len(results.Entries); i++ {
		fmt.Println("----------------------------------")
		if len(results.Entries[i].Name) > 0 {
			name := results.Entries[i].Name
			fmt.Println("Name:", name)
		}
		if len(results.Entries[i].Email) > 0 {
			email := results.Entries[i].Email
			fmt.Println("Email:", email)
		}
		if len(results.Entries[i].Username) > 0 {
			username := results.Entries[i].Username
			fmt.Println("Username:", username)
		}
		if len(results.Entries[i].Password) > 0 {
			password := results.Entries[i].Password
			fmt.Println("Password:", password)
		}
		if len(results.Entries[i].HashedPassword) > 0 {
			hashedPassword := results.Entries[i].HashedPassword
			fmt.Println("Hashed Password:", hashedPassword)
		}
	}
	fmt.Println("----------------------------------")
}
