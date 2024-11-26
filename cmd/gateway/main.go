package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/hello", func(c *fiber.Ctx) error {
		resp, err := http.Post("http://localhost:3001/hello", "application/json", nil)
		if err != nil {
			return c.Status(500).SendString("Error connecting to Service 1")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(500).SendString("Error reading response")
		}

		return c.Send(body)
	})

	app.Get("/world", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://localhost:3002/world")
		if err != nil {
			return c.Status(500).SendString("Error connecting to Service 2")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(500).SendString("Error reading response")
		}

		return c.Send(body)
	})

	app.Get("/calculator", func(c *fiber.Ctx) error {
		query := c.Request().URI().QueryString()
		resp, err := http.Get("http://localhost:3003/calculator?" + string(query))
		if err != nil {
			return c.Status(500).SendString("Error connecting to Service 3")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(500).SendString("Error reading response")
		}

		return c.Send(body)
	})

	log.Fatal(app.Listen(":3000"))
}
