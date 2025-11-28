package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/order-service/handlers"
)

func main(){

	app := fiber.New()

	app.Get("/orders", handlers.GetOrders)

	app.Listen(":3002")
}