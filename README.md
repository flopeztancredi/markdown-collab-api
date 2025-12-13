# Markdown Collab API

A collaborative markdown editing tool built with Go and Gin framework, following Onion Architecture principles.

## Getting Started

### Prerequisites

- Go 1.23
- Docker & Docker Compose
- [golang-migrate](https://github.com/golang-migrate/migrate)

### Running with Docker

```bash
# Start the application
make docker-up

# Run database migrations
make migrate-up

# View logs
make docker-logs

# Stop the application
make docker-down
```

### Running Locally

```bash
go mod tidy
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

| Variable     | Description              | Default                                                              |
|--------------|--------------------------|----------------------------------------------------------------------|
| APP_NAME     | Application name         | markdown-collab-api                                                  |
| APP_PORT     | Server port              | 8080                                                                 |
| GIN_MODE     | Gin mode (debug/release) | release                                                              |
| DATABASE_URL | PostgreSQL connection    | postgres://postgres:postgres@localhost:5432/markdown?sslmode=disable |

## Database

### Migrations

```bash
# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down

# Create new migration
make migrate-create name=migration_name
```

## Testing

```bash
make test
```
