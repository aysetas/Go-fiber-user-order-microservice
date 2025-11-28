package main

import (
	"fmt"
	"user-service/config"
)

func main() {

	config.LoadEnv()

	port := config.GetEnv("APP_PORT")

	if port == "" {
		port = "3001"
	}
	fmt.Println("User Service çalışıyor... Port:", port)

}
