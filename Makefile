.PHONY: help build run test clean lint

help:
	@echo "jxcf-api - Go REST API for article analysis"
	@echo ""
	@echo "Available commands:"
	@echo "  make build      - Build the binary"
	@echo "  make run        - Run the server"
	@echo "  make test       - Run tests"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make lint       - Run linter"
	@echo "  make fmt        - Format code"

build:
	@echo "Building jxcf-api..."
	go build -o bin/jxcf-api cmd/server/main.go

run:
	@echo "Running jxcf-api..."
	go run cmd/server/main.go

test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated in coverage.html"

clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

lint:
	@echo "Running linter..."
	golangci-lint run

fmt:
	@echo "Formatting code..."
	go fmt ./...
	deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
