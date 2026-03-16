package main

import (
	"log"
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

	h := api.mount()
	api.run(h)
	// api.run(api.mount())

	if err := api.run(api.mount()); err != nil {
		log.Printf("the server has failed to start %s", err)
		os.Exit(1)
	}

}
