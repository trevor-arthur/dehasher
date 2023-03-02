package pkg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/trevor-arthur/dehasher/models"
)

func Search(query string, apiEmail string, apiKey string) models.SearchJSON {

	// Formated API request
	apiReq := "https://api.dehashed.com/search?query=" + query
	req, err := http.NewRequest("GET", apiReq, nil)
	if err != nil {
		log.Fatalln("[!] Error with Search GET request:", err)
	}

	// Add Basic Auth and Headers
	req.SetBasicAuth(apiEmail, apiKey)
	req.Header.Set("Accept", "application/json")

	// Send request and capture response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("[!] Error with Search GET request response:", err)
	}
	defer resp.Body.Close()

	// Convert HTTP response into readable format
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[!] Error with reading Search response:", err)
	}

	var resObj models.SearchJSON
	json.Unmarshal(resBody, &resObj)

	return resObj
}
