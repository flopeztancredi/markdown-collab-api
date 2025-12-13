# Variables
APP_NAME := markdown-collab
MAIN_PATH := ./cmd/api

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

## build: Build the application binary
build:
	@echo "Building..."
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

## run: Run the application locally
run:
	@echo "Running..."
	go run $(MAIN_PATH)

## test: Run tests
test:
	@echo "Testing..."
	go test -v -race -cover ./...

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

## tidy: Tidy dependencies
tidy:
	go mod tidy

## docker-build: Build Docker image
docker-build:
	docker compose build

## docker-up: Start Docker containers
docker-up:
	docker compose up -d

## docker-down: Stop Docker containers
docker-down:
	docker compose down

## docker-logs: Show Docker logs
docker-logs:
	docker compose logs -f

## docker-restart: Restart Docker containers
docker-restart: docker-down docker-up
