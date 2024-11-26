# Go Microservices Study Guide

This guide will help you understand the concepts and implementation details of this microservices project.

## 1. Project Overview

### Basic Concepts
- **Microservices**: Independent services that work together
- **API Gateway**: Central entry point that routes requests to appropriate services
- **Clean Architecture**: Separation of concerns and dependency management

### What You'll Learn
- Building microservices with Go
- Using the Fiber web framework
- Implementing clean architecture
- Writing tests for Go services
- Service communication via HTTP

## 2. Step-by-Step Tutorial

### Step 1: Understanding the Project Structure
```bash
.
├── cmd/                # Entry points for our applications
├── internal/          
    ├── core/          # Business logic
    └── handlers/      # HTTP handlers
```

### Step 2: Core Components

1. **Ports (Interfaces)**
```go
// internal/core/ports/service.go
type HelloService interface {
    SayHello() string
}

type WorldService interface {
    SayWorld() string
}
```
These interfaces define the contract that our services must implement.

2. **Services (Implementation)**
```go
// internal/core/services/hello.go
type helloService struct{}

func (s *helloService) SayHello() string {
    return "hello"
}
```

### Step 3: HTTP Handlers
The handlers connect our services to HTTP endpoints:
```go
// Example handler implementation
func (h *HelloHandler) HandleHello(c *fiber.Ctx) error {
    return c.SendString(h.service.SayHello())
}
```

### Step 4: Running the Services

1. Start each service separately:
```bash
# Terminal 1
go run cmd/service1/main.go

# Terminal 2
go run cmd/service2/main.go

# Terminal 3
go run cmd/gateway/main.go
```

2. Test the endpoints:
```bash
# Test Hello Service
curl -X POST http://localhost:3000/hello

# Test World Service
curl http://localhost:3000/world
```

## 3. Key Learning Points

### Clean Architecture
1. **Dependency Rule**
   - Inner layers don't know about outer layers
   - Business logic (core) is independent of delivery mechanism

2. **Interface Segregation**
   - Services are defined by interfaces
   - Easy to mock for testing
   - Loose coupling between components

### Microservices Principles
1. **Service Independence**
   - Each service runs independently
   - Services can be deployed separately
   - Failure isolation

2. **API Gateway Pattern**
   - Single entry point for clients
   - Request routing
   - Can implement cross-cutting concerns

## 4. Practical Exercises

1. **Add New Feature**
   - Create a new service that returns "Goodbye"
   - Implement the interface, service, and handler
   - Add tests for the new service

2. **Enhance Gateway**
   - Add error handling
   - Implement request logging
   - Add basic authentication

3. **Service Communication**
   - Modify services to communicate with each other
   - Implement a composite endpoint that combines responses

## 5. Testing Strategy

1. **Unit Tests**
```go
func TestHelloService_SayHello(t *testing.T) {
    service := NewHelloService()
    result := service.SayHello()
    if result != "hello" {
        t.Errorf("Expected 'hello', got %v", result)
    }
}
```

2. **Integration Tests**
```go
func TestHelloHandler_HandleHello(t *testing.T) {
    app := fiber.New()
    // Setup handler and test HTTP endpoint
}
```

## 6. Common Challenges and Solutions

1. **Service Discovery**
   - Currently using hardcoded URLs
   - Could implement service registry
   - Consider using Consul or etcd

2. **Error Handling**
   - Implement circuit breakers
   - Add timeout handling
   - Proper error propagation

3. **Monitoring**
   - Add metrics collection
   - Implement health checks
   - Logging strategy

## 7. Next Steps

1. **Advanced Features**
   - Add database integration
   - Implement caching
   - Add authentication/authorization
   - Use message queues for async communication

2. **Production Considerations**
   - Containerization with Docker
   - Kubernetes deployment
   - Monitoring and logging
   - CI/CD pipeline

## 8. Resources

- [Go Documentation](https://golang.org/doc/)
- [Fiber Framework](https://gofiber.io/)
- [Clean Architecture Blog Post](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Microservices Pattern](https://microservices.io/patterns/microservices.html) 