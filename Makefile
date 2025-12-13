# Variables
APP_NAME := markdown-collab-api
MAIN_PATH := ./cmd/api
DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/markdown?sslmode=disable

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

## migrate-up: Run database migrations
migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

## migrate-down: Rollback database migrations
migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

## migrate-create: Create a new migration (usage: make migrate-create name=migration_name)
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

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
