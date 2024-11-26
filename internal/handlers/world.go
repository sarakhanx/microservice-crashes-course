package handlers

import (
    "github.com/gofiber/fiber/v2"
    "microservices/internal/core/ports"
)

type WorldHandler struct {
    service ports.WorldService
}

func NewWorldHandler(service ports.WorldService) *WorldHandler {
    return &WorldHandler{service: service}
}

func (h *WorldHandler) HandleWorld(c *fiber.Ctx) error {
    return c.SendString(h.service.SayWorld())
}