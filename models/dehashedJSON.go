package models

type EmailSearchJSON struct {
	Balance int `json:"balance"`
	Entries []struct {
		ID             string `json:"id"`
		Email          string `json:"email"`
		IPAddress      string `json:"ip_address"`
		Username       string `json:"username"`
		Password       string `json:"password"`
		HashedPassword string `json:"hashed_password"`
		Name           string `json:"name"`
		Vin            string `json:"vin"`
		Address        string `json:"address"`
		Phone          string `json:"phone"`
		DatabaseName   string `json:"database_name"`
	} `json:"entries"`
	Success bool   `json:"success"`
	Took    string `json:"took"`
	Total   int    `json:"total"`
}
