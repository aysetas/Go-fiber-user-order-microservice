package handlers

import (
	"github.com/gofiber/fiber/v2"
	"main.go/order-service/services"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := services.GetOrdersWithUser()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(orders)
}
