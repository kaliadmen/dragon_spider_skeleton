package main

import (
	dragonSpider "github.com/kaliadmen/dragon_spider/v2"
	"log"
	"${APP_NAME}/data"
	"${APP_NAME}/handlers"
	"${APP_NAME}/middleware"
	"os"
)

func initApplication() *application {
	//root directory of application
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//initialize DragonSpider
	ds := &dragonSpider.DragonSpider{}
	err = ds.New(path)
	if err != nil {
		log.Fatal(err)
	}

	ds.AppName = "myapp"

	appMiddleware := &middleware.Middleware{
		App: ds,
	}

	appHandlers := &handlers.Handlers{
		App: ds,
	}

	app := &application{
		App:        ds,
		Handlers:   appHandlers,
		Middleware: appMiddleware,
	}

	//add default dragon spider routes with application added routes
	app.App.Routes = app.routes()

	//add default dragon spider models
	app.Models = data.New(app.App.Db.Pool)
	//add default dragon spider handlers
	appHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app

}
