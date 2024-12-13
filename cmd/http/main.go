package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	addEndpoint      = "add"
	subtractEndpoint = "subtract"
)

func proxyRequest(c *fiber.Ctx, operation string) error {
	query := c.Request().URI().QueryString()
	resp, err := http.Get("http://localhost:3003/calculator/" + operation + "?" + string(query))
	if err != nil {
		return c.Status(500).SendString("Error connecting to Service 3")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(500).SendString("Error reading response")
	}

	return c.Send(body)
}

func main() {
	app := fiber.New()

	app.Get("/calculator/"+addEndpoint, func(c *fiber.Ctx) error {
		return proxyRequest(c, addEndpoint)
	})

	app.Get("/calculator/subtract", func(c *fiber.Ctx) error {
		return proxyRequest(c, subtractEndpoint)
	})

	log.Fatal(app.Listen(":3000"))
}
