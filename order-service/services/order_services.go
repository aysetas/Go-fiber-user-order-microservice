package services

import (
	"order-service/clients"
	"order-service/models"
)

var orders = []models.Order{
	{ID: 1, UserID: 1, Item: "Laptop"},
	{ID: 2, UserID: 2, Item: "Kulaklık"},
}

func GetOrdersWithUser() ([]models.Order, error) {
	users, err := clients.FetchUsers()
	if err != nil {
		return nil, err
	}

	for i, order := range orders {
		for _, u := range users {
			if u.ID == order.UserID {
				orders[i].User = u
			}
		}
	}

	return orders, nil
}
