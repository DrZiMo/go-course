package main

import (
	"net/http"
	"time"
)

func (app *application) Serve() error {
	server := http.Server{
		Addr:        "8080",
		ReadTimeout: 2 * time.Second,
		Handler:     app.routes(),
	}

	return server.ListenAndServe()
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)

	return mux
}
