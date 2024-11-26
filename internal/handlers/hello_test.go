package handlers

import (
    "github.com/gofiber/fiber/v2"
    "io"
    "net/http/httptest"
    "testing"
)

// Mock HelloService
type mockHelloService struct{}

func (m *mockHelloService) SayHello() string {
    return "hello"
}

func TestHelloHandler_HandleHello(t *testing.T) {
    // Arrange
    app := fiber.New()
    mockService := &mockHelloService{}
    handler := NewHelloHandler(mockService)
    app.Post("/hello", handler.HandleHello)

    // Act
    req := httptest.NewRequest("POST", "/hello", nil)
    resp, err := app.Test(req)

    // Assert
    if err != nil {
        t.Fatalf("Failed to test request: %v", err)
    }
    if resp.StatusCode != 200 {
        t.Errorf("Expected status code 200, got %v", resp.StatusCode)
    }

    body, _ := io.ReadAll(resp.Body)
    if string(body) != "hello" {
        t.Errorf("Expected body 'hello', got '%v'", string(body))
    }
}