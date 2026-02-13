HealthCheck API Service
A lightweight, production-ready Health Check API service built with Go (Golang). This service provides endpoints to monitor the health status of your applications and services.

📋 Table of Contents
Overview

Features

Architecture

Tech Stack

Prerequisites

Installation

Configuration

API Endpoints

Usage Examples

Project Structure

Development

Docker Support

Testing

Monitoring

Contributing

License

🎯 Overview
HealthCheck API is a simple yet powerful service that provides health monitoring endpoints for your applications. It follows Go best practices and clean architecture principles, making it easy to extend and maintain. The service returns comprehensive health information including system uptime, version, timestamp, and service dependencies status.

✨ Features
Health Check Endpoints: Basic and detailed health status information

Structured Logging: Built-in logger with different log levels

Graceful Shutdown: Proper handling of shutdown signals

Configuration Management: Environment-based configuration

Docker Support: Containerization ready

Structured Responses: JSON-formatted API responses

Request Logging: Middleware for logging all incoming requests

Modular Architecture: Clean separation of concerns

Production Ready: Includes timeout handling and error management

🏗 Architecture
The project follows a clean architecture pattern with clear separation of concerns:

text
┌─────────────────────────────────────┐
│           cmd/api (Entrypoint)      │
├─────────────────────────────────────┤
│           internal/api (Handlers)   │
├─────────────────────────────────────┤
│           internal/domain (Models)  │
├─────────────────────────────────────┤
│           internal/config           │
├─────────────────────────────────────┤
│           pkg/logger                │
└─────────────────────────────────────┘
Architecture Layers:
Domain Layer (internal/domain): Core business logic and models

Application Layer (internal/api): HTTP handlers and routing

Infrastructure Layer (pkg/logger): External concerns and utilities

Config Layer (internal/config): Configuration management

🛠 Tech Stack
Language: Go 1.21+

HTTP Server: Native net/http

Logging: Custom structured logger

Configuration: Environment variables

Container: Docker

📋 Prerequisites
Before you begin, ensure you have the following installed:

Go 1.21 or higher

Docker (optional, for containerization)

Make (optional, for using Makefile commands)

Git

🚀 Installation
1. Clone the Repository
bash
git clone https://github.com/yourusername/healthcheck-api.git
cd healthcheck-api
2. Install Dependencies
bash
go mod download
go mod tidy
3. Build the Application
bash
# Using Go directly
go build -o bin/healthcheck-api cmd/api/main.go

# Using Make
make build
4. Run the Application
bash
# Direct run
go run cmd/api/main.go

# Using Make
make run

# Using the binary
./bin/healthcheck-api
⚙️ Configuration
The application can be configured using environment variables:

Variable	Description	Default	Required
SERVER_PORT	Port the server listens on	8080	No
APP_VERSION	Application version	1.0.0	No
START_TIME	Custom start time	Current time	No
Configuration Example
Create a .env file:

env
SERVER_PORT=8080
APP_VERSION=1.0.0
Or set environment variables directly:

bash
export SERVER_PORT=9090
export APP_VERSION=2.0.0
go run cmd/api/main.go
📡 API Endpoints
1. Basic Health Check
Returns basic health status of the service.

Endpoint: GET /health

Response:

json
{
    "status": "healthy",
    "timestamp": "2024-01-15T10:30:00Z",
    "version": "1.0.0",
    "uptime": "2m30s"
}
2. Detailed Health Check
Returns detailed health information including dependency status.

Endpoint: GET /health/detailed

Response:

json
{
    "status": "healthy",
    "timestamp": "2024-01-15T10:30:00Z",
    "version": "1.0.0",
    "uptime": "2m30s",
    "services": {
        "database": "connected",
        "cache": "connected"
    }
}
3. Welcome Message
Endpoint: GET /

Response:

text
Welcome to HealthCheck API. Try /health or /health/detailed
💻 Usage Examples
Using cURL
bash
# Basic health check
curl http://localhost:8080/health

# Detailed health check
curl http://localhost:8080/health/detailed

# Welcome message
curl http://localhost:8080/
Using Python
python
import requests

response = requests.get('http://localhost:8080/health')
print(response.json())
Using JavaScript
javascript
fetch('http://localhost:8080/health')
    .then(response => response.json())
    .then(data => console.log(data));
📁 Project Structure
text
healthcheck-api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entrypoint
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   └── health_handler.go    # HTTP handlers
│   │   └── router/
│   │       └── router.go            # Route configuration
│   ├── config/
│   │   └── config.go                 # Configuration management
│   └── domain/
│       └── health.go                  # Domain models
├── pkg/
│   └── logger/
│       └── logger.go                  # Structured logging
├── Dockerfile                          # Docker configuration
├── Makefile                            # Build automation
├── go.mod                              # Go module file
├── go.sum                              # Go module checksum
└── README.md                           # Documentation
🛠 Development
Available Make Commands
bash
make build    # Build the application
make run      # Run the application
make test     # Run tests
make clean    # Clean build artifacts
make tidy     # Tidy dependencies
make docker-build # Build Docker image
make docker-run   # Run Docker container
Code Formatting
bash
# Format all Go files
go fmt ./...

# Run linter (if golint is installed)
golint ./...
Adding New Endpoints
Create a new handler in internal/api/handlers/

Register the route in internal/api/router/router.go

Add business logic in domain layer if needed

🐳 Docker Support
Build Docker Image
bash
# Using Docker directly
docker build -t healthcheck-api:1.0.0 .

# Using Make
make docker-build
Run Docker Container
bash
# Using Docker directly
docker run -p 8080:8080 healthcheck-api:1.0.0

# Using Make
make docker-run
Docker Compose (Optional)
Create a docker-compose.yml:

yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - SERVER_PORT=8080
      - APP_VERSION=1.0.0
🧪 Testing
Run Tests
bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Using Make
make test
Example Test (Add to health_handler_test.go)
go
package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthCheck(t *testing.T) {
    handler := NewHealthHandler("1.0.0")
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    
    handler.HealthCheck(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("Expected status OK, got %v", w.Code)
    }
}
📊 Monitoring
The health check endpoints can be integrated with monitoring tools like:

Prometheus: Can scrape the /health endpoint

Kubernetes: Use as liveness and readiness probes

Consul: Service health checking

Custom monitoring scripts

Kubernetes Probe Example
yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10

readinessProbe:
  httpGet:
    path: /health/detailed
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
🤝 Contributing
Contributions are welcome! Please follow these steps:

Fork the repository

Create a feature branch (git checkout -b feature/amazing-feature)

Commit your changes (git commit -m 'Add amazing feature')

Push to the branch (git push origin feature/amazing-feature)

Open a Pull Request

Contribution Guidelines
Write clear commit messages

Add tests for new features

Update documentation

Follow Go best practices

Ensure all tests pass

📝 License
This project is licensed under the MIT License - see the LICENSE file for details.

👥 Authors
Hamed Safarzadeh - Initial work

🙏 Acknowledgments
Go community for best practices

Clean Architecture principles

Contributors and users

📧 Contact
For questions or support, please open an issue or contact:

Email: h.safarzadeh2014@gmail.com

GitHub: @hamed-hsb

🚀 Future Enhancements
Add Prometheus metrics

Implement more detailed health checks (database, cache, external services)

Add authentication/authorization

Create OpenAPI/Swagger documentation

Add more configuration options

Implement caching

Add rate limiting

Create Helm chart for Kubernetes deployment

Made with ❤️ using Go