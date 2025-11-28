package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/user-service/handlers"
)


func main(){


	app := fiber.New()

	app.Get("/users", handlers.GetUsers)

	app.Listen(":3001")

}