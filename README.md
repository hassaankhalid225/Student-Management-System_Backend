# Student Management System - Backend (Go)

This is the backend for the Student Management System built with Go, Gin, GORM, and PostgreSQL.

## 🚀 Getting Started

### 1. Prerequisites
- [Go](https://golang.org/doc/install) (v1.21+)
- [Docker Desktop](https://www.docker.com/products/docker-desktop) (for PostgreSQL and Redis)

### 2. Setup Infrastructure
The easiest way to run the database and cache is via Docker:
```bash
docker-compose up -d
```
This will start:
- **PostgreSQL**: `localhost:5432` (User: postgres, Pass: postgres)
- **Redis**: `localhost:6379`

### 3. Environment Configuration
Copy the `.env.example` file to `.env`:
```bash
cp .env.example .env
```
Ensure the values in `.env` match your local setup.

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Run the Server
```bash
go run cmd/server/main.go
```
The API will be available at `http://localhost:8080/api/v1`.
You can check the health at `http://localhost:8080/health`.

## 📁 Project Structure
- `cmd/server`: Application entry point.
- `internal/application`: Use cases and DTOs.
- `internal/domain`: Business entities and repository interfaces.
- `internal/infrastructure`: External tools (Database, JWT, Redis).
- `internal/interfaces/http`: API Handlers and Middlewares.

## 🛠️ Common Commands
- **Build**: `go build -o server cmd/server/main.go`
- **Test**: `go test ./...`
- **Lint**: `golangci-lint run` (if installed)
