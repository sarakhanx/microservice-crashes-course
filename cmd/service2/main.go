package main

import (
    "github.com/gofiber/fiber/v2"
    "microservices/internal/core/services"
    "microservices/internal/handlers"
    "log"
)

func main() {
    app := fiber.New()
    
    worldService := services.NewWorldService()
    worldHandler := handlers.NewWorldHandler(worldService)
    
    app.Get("/world", worldHandler.HandleWorld)
    
    log.Fatal(app.Listen(":3002"))
}