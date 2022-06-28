package handlers

import (
	"${APP_NAME}/data"
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"net/http"
)

//Handlers is the type for handlers, and gives access to DragonSpider and models
type Handlers struct {
	App    *dragonSpider.DragonSpider
	Models data.Models
}

//Home is the handler to render a home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering page:", err)
	}
}
