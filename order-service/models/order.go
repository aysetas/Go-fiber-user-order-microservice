package models

import "main.go/order-service/clients"

type Order struct {
	ID     int          `json:"id"`
	UserID int          `json:"user_id"`
	Item   string       `json:"item"`
	User   clients.User `json:"user"`
}
