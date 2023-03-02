package pkg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/trevor-arthur/dehasher/models"
)

func EmailSearch(email string, apiKey string) models.EmailSearchJSON {

	// Formated API request
	apiReq := "https://api.dehashed.com/search?query=email:" + email
	req, err := http.NewRequest("GET", apiReq, nil)
	if err != nil {
		log.Fatalln("[!] Error with EmailSearch GET request:", err)
	}

	// Add Basic Auth and Headers
	req.SetBasicAuth("sammy@carefulsecurity.com", apiKey)
	req.Header.Set("Accept", "application/json")

	// Send request and capture response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("[!] Error with EmailSearch GET request response:", err)
	}
	defer resp.Body.Close()

	// Convert HTTP response into readable format
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[!] Error with reading EmailSearch response:", err)
	}

	var resObj models.EmailSearchJSON
	json.Unmarshal(resBody, &resObj)

	return resObj
}
