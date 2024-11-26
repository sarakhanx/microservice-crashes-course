package handlers

import (
	"microservices/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type calculatorHandler struct {
	service ports.CalculatorService
}

func NewAddHandler(service ports.CalculatorService) *calculatorHandler {
	return &calculatorHandler{service: service}
}

func (h *calculatorHandler) Add(c *fiber.Ctx) error {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameter a"})
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameter b"})
	}
	return c.JSON(fiber.Map{"result": h.service.Add(a, b)})
}

// func (h *calculatorHandler) Subtract(c *fiber.Ctx) error {
// 	return c.JSON(fiber.Map{"result": h.service.Subtract(a, b)})
// }
