package main

import (
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"myApp/data"
	"myApp/handlers"
	"myApp/middleware"
)

//application holds data all the needed for Dragon Spider App
type application struct {
	App        *dragonSpider.DragonSpider
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	ds := initApplication()

	ds.App.ListenAndServe()
}