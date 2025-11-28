package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"order-service/config"
	"order-service/handlers"
)

func main() {

	config.LoadEnv()

	port := config.GetEnv("APP_PORT")

	if port == "" {
		port = "3002"
	}
	fmt.Println("Order Service çalışıyor... Port:", port)

	app := fiber.New()

	app.Get("/orders", handlers.GetOrders)

}
