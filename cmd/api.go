package main

import (
	"log"
	"net/http"
	"time"

	"github.com/APrem-7/GO_ECOM_API/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	//methods that are going to run and moiunt your api and you can add graceful shutdown also

	config config
	//later on add logger
	db *pgx.Conn
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

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good\n"))
	})
	productHandler := products.NewHandler(nil)
	r.Get("/products", productHandler.ListProducts)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()

}

type config struct {
	addr string //address to listen on
	db   dbConfig
}

type dbConfig struct {
	dsn string //data source name its like the username password dbname and stuff
}
