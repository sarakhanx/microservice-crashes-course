package handlers

import (
    "github.com/gofiber/fiber/v2"
    "io"
    "net/http/httptest"
    "testing"
)

// Mock WorldService
type mockWorldService struct{}

func (m *mockWorldService) SayWorld() string {
    return "world"
}

func TestWorldHandler_HandleWorld(t *testing.T) {
    // Arrange
    app := fiber.New()
    mockService := &mockWorldService{}
    handler := NewWorldHandler(mockService)
    app.Get("/world", handler.HandleWorld)

    // Act
    req := httptest.NewRequest("GET", "/world", nil)
    resp, err := app.Test(req)

    // Assert
    if err != nil {
        t.Fatalf("Failed to test request: %v", err)
    }
    if resp.StatusCode != 200 {
        t.Errorf("Expected status code 200, got %v", resp.StatusCode)
    }

    body, _ := io.ReadAll(resp.Body)
    if string(body) != "world" {
        t.Errorf("Expected body 'world', got '%v'", string(body))
    }
}