package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// or <api>
type application struct {
	config config 
	
}


type config struct {
	addr string 
}



func (app *application) mount() http.Handler {
	// mux := http.NewServeMux()



	// mux.HandleFunc("GET /v1/health", app.healthCheckHandler)


	// using the "Chi" package, we can group endpoints and middlewares
	// it's possible to use middlewares with standard libraries, however it's much easier to chain middlewares with Chi.

	// posts 

	
	// users 

	// auth 

	r := chi.NewRouter()


	// A good base middleware stack 
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP) // deprecated due to risk of ip spoofing 
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// r.Get("/health", app.healthCheckHandler)


	r.Route("/v1", func (r chi.Router){
		r.Get("/health", app.healthCheckHandler)
	})


	return r
}

func (app* application) run(mux http.Handler) error {


	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	log.Printf("Server has started at %s", app.config.addr)



	return srv.ListenAndServe()
}