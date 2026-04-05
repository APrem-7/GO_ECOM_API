# A REST API for E-commerce

Built with Go. No fluff, no framework magic — just clean layers, typed SQL, and a Postgres database that Docker spins up for you in one command.

## WHAT IT DOES
Two core things:

- **Browse products** — `GET /products` returns every product in the store with name, description, price, and stock quantity.
- **Place orders** — `POST /orders` creates an order for a customer. It validates that every item exists, checks stock availability, and wraps everything in a database transaction so you never end up with a half-created order.

There’s also a `GET /health` endpoint that just says "all good" — because sometimes that’s all you need to hear.

## Why did the Go developer quit their job?
Because they didn’t get a **CHAN**-ce to do anything interesting.

## STACK
- **Go 1.25.5** — Language
- **chi** — Lightweight HTTP router
- **PostgreSQL 16** — Database
- **pgx/v5** — Postgres driver
- **sqlc** — Type-safe SQL → Go generator
- **Docker Compose** — Local database setup

## PROJECT LAYOUT
```
cmd/
  ├─ api.go      → server setup, routes, middleware
  └─ main.go     → entry point, env + DB config
total/ 
internal/
  ├─ products/   → handler + service for products
  └─ orders/     → handler + service for orders (transaction + stock check)
adapters/
postgres/
sqlc/   → queries + generated code
json/       → JSON helpers
env/        → environment config
```

## Architecture flow:
HTTP → Handler → Service → Repository → Database

## ARCHITECTURE DIAGRAM (MERMAID)
graph TD;
A[Client / Frontend] -->|HTTP Requests| B[Chi Router];
B --> C[Middleware];
C --> D[Handlers];
D --> E[Product Handler];
D --> F[Order Handler];
E --> G[Product Service];
F --> H[Order Service];
G --> I[Product Logic];
H --> J[Order Logic];
I --> K[Repository];
J --> K;
k --> L[SQLC Queries];
l --> M[(PostgreSQL)];
sublgraph Infrastructure;
n [Docker Compose]; o [Env Config]; end;
n --> M; o --> B;
 
## GETTING STARTED
### Prerequisites:
- Go 1.25+
- Docker & Docker Compose
### Start database:
docker-compose up -d  
pPostgres runs on localhost:5432,
uUser: postgres,
pPassword: postgres,
pDB: ecom.
### Set environment:
tCreate .env file:
pPORT=:8080,
pDATABASE_URL=postgres://postgres:postgres@localhost:5432/ecom?sslmode=disable;
'tRun migrations:
apply SQL files from:
internal/adapters/postgres/migrations/;
'tStart server:
go run cmd/api.go'
defaults to API endpoints:
bGET /health - Response: "all good"
bGET /products - Returns list of products with ID, Name, PriceInCents, Description, Quantity, timestamps (Prices stored in cents; e.g., 8999 = $89.99)
bPOST /orders - Request includes customer_id and items list; Responses include status codes like 201 (Created), 400 (Bad request), etc.
'tWorking with SQLC:
dQueries live in internal/adapters/postgres/sqlc/queries.sql;
after editing queries, run sqlc generate; do not edit generated files directly.
'things to contribute are welcome via issues and pull requests.
