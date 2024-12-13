package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// Mock HTTP server for service1
func setupMockService1() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
}

// Mock HTTP server for service2
func setupMockService2() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("world"))
	}))
}

func setupTestApp(service1URL, service2URL string) *fiber.App {
	app := fiber.New()

	app.Post("/hello", func(c *fiber.Ctx) error {
		resp, err := http.Post(service1URL+"/hello", "application/json", nil)
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
		resp, err := http.Get(service2URL + "/world")
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

	return app
}

func TestGatewayIntegration(t *testing.T) {
	// Setup mock services
	service1 := setupMockService1()
	service2 := setupMockService2()
	defer service1.Close()
	defer service2.Close()

	// Setup test app with mock service URLs
	app := setupTestApp(service1.URL, service2.URL)

	// Test hello endpoint
	t.Run("Test POST /hello", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/hello", nil)
		resp, err := app.Test(req)

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
	})

	// Test world endpoint
	t.Run("Test GET /world", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/world", nil)
		resp, err := app.Test(req)

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
	})
}
