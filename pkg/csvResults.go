package pkg

import (
	"fmt"
	"time"

	"github.com/trevor-arthur/dehasher/models"
	tas "github.com/trevor-arthur/taslib/pkg/csv"
)

// CSVresults writes dehashed search results to a csv file.
func CSVresults(results models.SearchJSON) {

	// Name of file to output results to
	now := time.Now().Local()
	dateFormat := fmt.Sprintf("%d%02d%02d-%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	outputFile := "dehashed-results-" + dateFormat + ".csv"

	// Create headers
	headers := []string{"email", "password", "hashed password", "username", "name"}

	rowCount := len(results.Entries)

	var rowData [][]string

	for i := 0; i < len(results.Entries); i++ {

		// Get values for row
		email := results.Entries[i].Email
		password := results.Entries[i].Password
		hashedPassword := results.Entries[i].HashedPassword
		username := results.Entries[i].Username
		name := results.Entries[i].Name

		// Create row
		row := []string{email, password, hashedPassword, username, name}
		rowData = append(rowData, row)
	}

	tas.WriteToCSV(outputFile, headers, rowCount, rowData)
}
