package main

import (
	"log"
	"microservices/internal/core/services"
	"microservices/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	helloService := services.NewHelloService()
	helloHandler := handlers.NewHelloHandler(helloService)

	app.Post("/hello", helloHandler.HandleHello)

	log.Fatal(app.Listen(":3001"))
}
