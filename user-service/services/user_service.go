package services

import "main.go/user-service/models"


func GetAllUsers() []models.User{
	return []models.User{
		{ID: 1, Name: "Ayşe"},
		{ID: 2, Name: "Mehmet"},
	}
}
