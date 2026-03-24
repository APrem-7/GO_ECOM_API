package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/APrem-7/GO_ECOM_API/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost port=5432 user=postgres password=postgres dbname=ecom sslmode=disable")},
	}
	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)

	}
	defer conn.Close(ctx)

	logger.Info("database connected", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
	}

	h := api.mount()
	api.run(h)
	// api.run(api.mount())

	if err := api.run(api.mount()); err != nil {
		log.Printf("the server has failed to start %s", err)
		os.Exit(1)
	}

}
