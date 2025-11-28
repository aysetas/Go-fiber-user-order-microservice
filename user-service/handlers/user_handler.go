package handlers

import (
	"github.com/gofiber/fiber/v2"
	"main.go/user-service/services"
)


func GetUsers(c *fiber.Ctx) error{

	users := services.GetAllUsers()

	return c.JSON(users)

}