package handlers

import (
	"microservices/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type HelloHandler struct {
	service ports.HelloService
}

func NewHelloHandler(service ports.HelloService) *HelloHandler {
	return &HelloHandler{service: service}
}

func (h *HelloHandler) HandleHello(c *fiber.Ctx) error {
	return c.SendString(h.service.SayHello())
}
