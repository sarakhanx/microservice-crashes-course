package main

import (
	"log"
	"microservices/internal/core/services"
	"microservices/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	calculatorService := services.NewCalculatorService()
	calculatorHandler := handlers.NewAddHandler(calculatorService)

	app.Get("/calculator", calculatorHandler.Add)

	log.Fatal(app.Listen(":3003"))
}
