# Go Microservices Learning Guide

Welcome to the Go Microservices Learning Project! This guide is designed to help you understand the concepts and implementation details of building microservices using Go and the Fiber framework.

## Project Structure

Understanding the project structure is crucial for navigating and comprehending the codebase. Here's a breakdown:
```
.
├── cmd/ # Entry points for our applications
│ ├── gateway/ # API Gateway service
│ ├── service1/ # Hello service
│ ├── service2/ # World service
│ └── calculator-service/ # Calculator service
└── internal/ # Internal application code
├── core/ # Business logic
│ ├── ports/ # Interfaces
│ └── services/ # Service implementations
└── handlers/ # HTTP handlers
```


## Key Concepts

### Microservices
- **Definition**: Independent services that work together to form a complete application.
- **Benefits**: Scalability, flexibility, and ease of deployment.

### API Gateway
- **Role**: Acts as a central entry point that routes requests to appropriate services.
- **Benefits**: Simplifies client interactions and can handle cross-cutting concerns like authentication.

### Clean Architecture
- **Principles**: Separation of concerns and dependency management.
- **Goal**: Keep business logic independent from delivery mechanisms.

## Learning Objectives

1. **Building Microservices with Go**
   - Understand how to structure a Go project for microservices.
   - Learn to use the Fiber web framework for handling HTTP requests.

2. **Implementing Clean Architecture**
   - Explore how to separate business logic from other concerns.
   - Use interfaces to define service contracts.

3. **Service Communication via HTTP**
   - Learn how services communicate with each other using HTTP requests.

## Step-by-Step Tutorial

### Step 1: Understanding the Project Structure
- Familiarize yourself with the directory layout and the purpose of each folder.

### Step 2: Core Components
- **Ports (Interfaces)**: Define the contract that services must implement.
  ```go:internal/core/ports/service.go
  startLine: 3
  endLine: 14
  ```

- **Services (Implementation)**: Implement the business logic.
  ```go:internal/core/services/hello.go
  startLine: 3
  endLine: 11
  ```

### Step 3: HTTP Handlers
- Connect services to HTTP endpoints.
  ```go:internal/handlers/hello.go
  startLine: 17
  endLine: 19
  ```

### Step 4: Running the Services
- Start each service separately in different terminal windows.
  ```bash
  # Terminal 1
  go run cmd/service1/main.go

  # Terminal 2
  go run cmd/service2/main.go

  # Terminal 3
  go run cmd/gateway/main.go
  ```

## Practical Exercises

1. **Add a New Feature**
   - Create a new service that returns "Goodbye".
   - Implement the interface, service, and handler.
   - Add tests for the new service.

2. **Enhance the Gateway**
   - Add error handling and request logging.
   - Implement basic authentication.

3. **Service Communication**
   - Modify services to communicate with each other.
   - Implement a composite endpoint that combines responses.

## Testing Strategy

1. **Unit Tests**
   - Test individual components in isolation.
   ```go:internal/core/services/hello_test.go
   startLine: 5
   endLine: 16
   ```

2. **Integration Tests**
   - Test how components work together.
   ```go:internal/handlers/hello_test.go
   startLine: 17
   endLine: 39
   ```

## Next Steps
1. **Advanced Features**
   - Add database integration and caching.
   - Implement authentication/authorization.

2. **Production Considerations**
   - Containerization with Docker.
   - Deployment with Kubernetes.
   - Set up monitoring and logging.

##### Happy learning!# microservice-crashes-course
