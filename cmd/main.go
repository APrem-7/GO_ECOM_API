package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

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
