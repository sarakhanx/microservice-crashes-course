# Learning Hexagonal Microservice Api Gateway

A microservice-based API Server application built with Go and Fiber framework that provides basic arithmetic operations through REST API endpoints.

## Architecture

The project follows a Hexagonal architecture pattern with the following components:

- HTTP Gateway Service (Port 3000)
- Calculator Service (Port 3003)

## Features

- Addition operation (`/calculator/add`)
- Subtraction operation (`/calculator/subtract`)
- RESTful API endpoints
- Microservice architecture
- Error handling

## API Endpoints

- GET {localhost:3000}/calculator/add?a={number}&b={number}
- GET {localhost:3000}/calculator/subtract?a={number}&b={number}

## Project Structure

```d
.
├── cmd/
│ ├── http/ # HTTP Gateway Service
│ │ └── main.go
│ └── gateway/ # Gateway Service
│ ├── main.go
│ └── main_test.go
├── internal/
│ ├── adapters/
│ ├── core/
│ │ ├── ports/
│ │ └── services/
│ ├── domain/
│ └── handlers/
└── go.mod
```

## Getting Started

### Prerequisites

- Go 1.x or higher
- Git

### Installation

1. Clone the repository

   ```bash
   git clone [repository-url]
   ```

2. Install dependencies

   ```bash
   go mod download
   ```

### Running the Services

1. Start the Calculator Service (Port 3003)

   ```bash
   go run cmd/http/main.go
   ```

## Testing

Run the tests using:

   ```bash
   go test ./...
   ```
