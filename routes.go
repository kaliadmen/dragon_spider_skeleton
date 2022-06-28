package main

import (
	"github.com/go-chi/chi/v5"
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"net/http"
)

func (a *application) routes() *chi.Mux {
	//middleware must come before any routes

	//add routes here
	a.get("/", a.Handlers.Home)

	//static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	a.App.Routes.Mount("/dragon_spider", dragonSpider.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
