package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: "host=localhost user=postgres password=postgres dbname=ecom sslmode=disabled"},
	}
	//Database
	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		panic(err)

	}
	defer conn.Close(ctx)

	api := application{
		config: cfg,
	}
	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	h := api.mount()
	api.run(h)
	// api.run(api.mount())

	if err := api.run(api.mount()); err != nil {
		log.Printf("the server has failed to start %s", err)
		os.Exit(1)
	}

}
