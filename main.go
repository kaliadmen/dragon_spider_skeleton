package main

import (
	"${APP_NAME}/data"
	"${APP_NAME}/handlers"
	"${APP_NAME}/middleware"
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//application holds data all the needed for Dragon Spider App
type application struct {
	App        *dragonSpider.DragonSpider
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
	wg         sync.WaitGroup
}

func main() {
	ds := initApplication()

	go ds.listenForShutdown()

	err := ds.App.ListenAndServe()
	ds.App.ErrorLog.Println(err)
}

func (a *application) shutdown() {
	//put clean up task here

	//block until wait group is empty
	a.wg.Wait()

}

func (a *application) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s := <-quit

	a.App.InfoLog.Println(s.String())
	a.shutdown()

	os.Exit(0)
}
