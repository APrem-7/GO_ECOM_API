package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	//methods that are going to run and moiunt your api and you can add graceful shutdown also

	config config
	//later on add logger and db driver

}

// mount mounts the routes to the router
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) //important for rate limiting
	r.Use(middleware.RealIP)    //important for rate linmitng and analytucs and tracing
	r.Use(middleware.Logger)    //logger
	r.Use(middleware.Recoverer) //recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	return r
}

// run
func (app *application) run(h http.Handler) error {
	return http.ListenAndServe(":8080", h)
}

type config struct {
	addr string //address to listen on
	db   dbConfig
}

type dbConfig struct {
	dsn string //data source name its like the username password dbname and stuff
}
