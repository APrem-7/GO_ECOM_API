# GO_ECOM_API

A REST API for e-commerce, built with Go. No fluff, no framework magic — just clean layers, typed SQL, and a Postgres database that Docker spins up for you in one command.

---

## What it does

Two core things:

- **Browse products** — `GET /products` returns every product in the store with name, description, price, and stock quantity.
- **Place orders** — `POST /orders` creates an order for a customer. It validates that every item in the order actually exists, checks that stock is available, and wraps the whole thing in a database transaction so you never end up with a half-created order.

There's also a `GET /health` endpoint that just says `"all good"` — because sometimes that's all you need to hear.

> Why did the Go developer quit their job?  
> Because they didn't get a CHAN-ce to do anything interesting. 🥁

---

## Stack

| Thing | What it is |
|---|---|
| [Go 1.25.5](https://go.dev/) | The language |
| [chi](https://github.com/go-chi/chi) | Lightweight HTTP router — no magic, just middleware and routes |
| [PostgreSQL 16](https://www.postgresql.org/) | The database |
| [pgx/v5](https://github.com/jackc/pgx) | Postgres driver — fast and idiomatic |
| [sqlc](https://sqlc.dev/) | Generates type-safe Go from raw SQL — write SQL, get Go structs |
| Docker Compose | Spins up the database locally without any setup pain |

---

## Project Layout

```
cmd/
  api.go      ← server setup, routes, middleware
  main.go     ← entry point, env + DB config

internal/
  products/   ← handler + service for listing products
  orders/     ← handler + service for creating orders (with tx + stock check)
  adapters/
    postgres/
      sqlc/   ← all DB queries and generated code live here
  json/       ← tiny helpers for reading/writing JSON
  env/        ← reads env vars with a fallback default
```

The architecture is intentionally simple: HTTP handler → service → database. No over-engineering.

### Architecture Diagram

```mermaid
graph TD
    Client([Client])

    subgraph cmd ["cmd/"]
        Router["chi Router\napi.go"]
    end

    subgraph products ["internal/products/"]
        PH["Handler"]
        PS["Service"]
    end

    subgraph orders ["internal/orders/"]
        OH["Handler"]
        OS["Service"]
    end

    subgraph adapters ["internal/adapters/postgres/sqlc/"]
        DB["Queries\n(sqlc-generated)"]
    end

    Postgres[(PostgreSQL 16)]

    Client -->|HTTP request| Router
    Router -->|GET /products| PH
    Router -->|POST /orders| OH
    PH --> PS
    OH --> OS
    PS --> DB
    OS --> DB
    DB --> Postgres
```

### Order Request Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant H as orders.Handler
    participant S as orders.Service
    participant Q as sqlc Queries
    participant DB as PostgreSQL

    C->>H: POST /orders { customer_id, items }
    H->>H: Decode & validate request body
    H->>S: CreateOrder(ctx, params)

    S->>DB: BEGIN transaction

    loop For each order item
        S->>Q: GetProduct(ctx, product_id)
        Q-->>S: Product (or 404 Not Found)
        S->>S: Check stock availability
    end

    S->>Q: CreateOrder(ctx, ...)
    Q-->>S: Order record

    loop For each order item
        S->>Q: CreateOrderItem(ctx, ...)
        S->>Q: DeductStock(ctx, product_id, qty)
    end

    S->>DB: COMMIT transaction
    DB-->>S: OK
    S-->>H: Order created
    H-->>C: 201 Created { order }
```

---

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose

### 1. Start the database

```bash
docker-compose up -d
```

This starts a Postgres 16 container on `localhost:5432` with:

- **User:** `postgres`  
- **Password:** `postgres`  
- **Database:** `ecom`

### 2. Set up your environment

Create a `.env` file in the project root:

```env
PORT=:8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/ecom?sslmode=disable
```

### 3. Run migrations

Apply the SQL schema from `internal/adapters/postgres/migrations/` to your database. You can use a tool like [`goose`](https://github.com/pressly/goose) or just run the migration files directly against the database.

### 4. Start the server

```bash
go run cmd/api.go
```

The server starts on the port defined in your `PORT` env var (default `:8080`).

---

## API

### `GET /health`
Returns `200 OK` with `"all good"`. Use it for health checks or just reassurance.

---

### `GET /products`
Returns a JSON array of all products.

```json
[
  {
    "ID": 1,
    "Name": "Mechanical Keyboard",
    "PriceInCenters": 8999,
    "Description": { "String": "Clicky and satisfying", "Valid": true },
    "Quantity": 42,
    "CreatedAt": "...",
    "UpdatedAt": "..."
  }
]
```

> Note: Prices are stored in **cents** (e.g. `8999` = $89.99).

---

### `POST /orders`
Creates a new order. Validates stock availability and runs everything in a single transaction.

**Request body:**
```json
{
  "customer_id": 7,
  "items": [
    { "product_id": 1, "quantity": 2 }
  ]
}
```

**Responses:**
- `201 Created` — order created, returns the order object
- `400 Bad Request` — missing customer ID or items
- `404 Not Found` — a product ID doesn't exist
- `507 Insufficient Storage` — not enough stock for one of the items
- `500 Internal Server Error` — something went wrong on our end

---

## Working with sqlc

SQL queries live in `internal/adapters/postgres/sqlc/queries.sql`. When you add or change a query, regenerate the Go code:

```bash
sqlc generate
```

The generated types and functions land in the same `sqlc/` directory. Don't edit the generated files directly — they'll get overwritten.

---

## Contributing

Issues and pull requests are welcome. If something's broken or you've got an idea, open an issue and we can talk through it.

