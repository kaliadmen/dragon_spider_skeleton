package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(mux chi.Router) {
		//add api routes

	})

	return r
}
