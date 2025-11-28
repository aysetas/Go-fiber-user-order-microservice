package clients

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FetchUsers() ([]User, error){
	resp, err := http.Get("http://localhost:3001/users")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
