## Architecture flow:
HTTP → Handler → Service → Repository → Database
## ARCHITECTURE DIAGRAM (MERMAID)
```mermaid
graph TD
    A[Client / Frontend] -->|HTTP Requests| B[Chi Router]
    B --> C[Middleware]
    C --> D[Handlers]
    D --> E[Product Handler]
    D --> F[Order Handler]
    E --> G[Product Service]
    F --> H[Order Service]
    G --> I[Product Logic]
    H --> J[Order Logic]
    I --> K[Repository]
    J --> K
    K --> L[SQLC Queries]
    L --> M[(PostgreSQL)]
    subgraph Infrastructure
        N[Docker Compose]
        O[Env Config]
    end
    N --> M
    O --> B
```
## GETTING STARTED
### Prerequisites:
- Go 1.25+
- Docker & Docker Compose
### Start database:
docker-compose up -d  
Postgres runs on localhost:5432,
User: postgres,
Password: postgres,
DB: ecom.
### Set environment:
Create .env file:
PORT=:8080,
DATABASE_URL=postgres://postgres:postgres@localhost:5432/ecom?sslmode=disable;
Run migrations:
apply SQL files from:
internal/adapters/postgres/migrations/;
Start server:
go run cmd/api.go
defaults to API endpoints:
GET /health - Response: "all good"
GET /products - Returns list of products with ID, Name, PriceInCents, Description, Quantity, timestamps (Prices stored in cents; e.g., 8999 = $89.99)
POST /orders - Request includes customer_id and items list; Responses include status codes like 201 (Created), 400 (Bad request), etc.
Working with SQLC:
Queries live in internal/adapters/postgres/sqlc/queries.sql;
after editing queries, run sqlc generate; do not edit generated files directly.
things to contribute are welcome via issues and pull requests.
