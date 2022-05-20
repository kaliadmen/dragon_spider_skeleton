package handlers

import (
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"myApp/data"
	"net/http"
)

type Handlers struct {
	App    *dragonSpider.DragonSpider
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering page:", err)
	}
}
