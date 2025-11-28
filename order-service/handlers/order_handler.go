package handlers

import (
	"order-service/services"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := services.GetOrdersWithUser()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(orders)
}
