# Markdown Collab

A collaborative markdown editing tool built with Go and Gin framework, following Onion Architecture principles.

## Getting Started

### Prerequisites

- Go 1.23
- Docker & Docker Compose

### Running with Docker

```bash
# Start the application
make docker-up

# View logs
make docker-logs

# Stop the application
make docker-down
```

### Running Locally

```bash
# Install dependencies
go mod tidy

# Run the application
make run
```

## API Endpoints

| Method | Endpoint  | Description         |
|--------|-----------|---------------------|
| GET    | `/health` | Health check status |

## Configuration

Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
```

| Variable   | Description           | Default   |
|------------|-----------------------|-----------|
| APP_NAME   | Application name      | gin-api   |
| APP_PORT   | Server port           | 8080      |
| GIN_MODE   | Gin mode (debug/release) | release |

## Testing

```bash
make test
```
