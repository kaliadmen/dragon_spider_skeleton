package main

import (
	"net/http"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(mux chi.Router) {
		//add api routes

	})

	return r
}
