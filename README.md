# Go E-Commerce API

A fast, robust, and scalable E-Commerce REST API built with Go.

## 🚀 Tech Stack

- **Language:** [Go](https://go.dev/) (v1.25.5)
- **Router:** [go-chi](https://github.com/go-chi/chi) - A lightweight, idiomatic, and composable router for building Go HTTP services.
- **Database:** PostgreSQL
- **Database Driver:** [pgx/v5](https://github.com/jackc/pgx)
- **Database Toolkit:** [sqlc](https://sqlc.dev/) - Compiles SQL to type-safe Go code.

## 📂 Project Structure

- `cmd/`: Application entry points (`api.go`).
- `internal/`: Private application and library code.
  - `adapters/postgres/`: Database-specific implementations, migrations, and `sqlc` generated code.
  - `orders/`: Business logic, domain models, and HTTP handlers for the orders domain.
  - `products/`: Business logic, domain models, and HTTP handlers for the products domain.
  - `json/`: Reusable utility functions for JSON encoding/request handling.

## 🛠 Prerequisites

Before you start, make sure you have the following installed:
- Go (1.25.5 or later)
- Docker and Docker Compose
- `sqlc` CLI (optional, if you plan to modify SQL queries)

## 💻 Getting Started

### 1. Start the Database

The project uses Docker Compose to easily spin up a PostgreSQL instance.

```bash
docker-compose up -d
```

This will start a PostgreSQL database running on `localhost:5432` with:
- **User:** `postgres`
- **Password:** `postgres`
- **Database:** `ecom`

### 2. Configure Environment

Ensure you have a `.env` file in the root directory configured with your database connection details and other necessary environment variables.

Example `.env` (adjust according to your specific setup):
```env
# Example environment variables
PORT=:8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/ecom?sslmode=disable
```

### 3. Run Database Migrations

*(Make sure your database schema is initialized. You may need a tool like `goose` or `migrate` depending on your setup, or manually execute the SQL files in `internal/adapters/postgres/migrations/`)*

### 4. Run the Application

Start the HTTP server:

```bash
go run cmd/api.go
```

## 🗄 Database Workflow (sqlc)

This project uses `sqlc` to generate type-safe Go code from SQL. 

1. Write your SQL queries in `internal/adapters/postgres/sqlc/queries.sql`.
2. Define schema changes in `internal/adapters/postgres/migrations/`.
3. Run `sqlc generate` in the root directory.
4. The generated Go code will be available in the `internal/adapters/postgres/sqlc/` package (as `repo`).

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page.
