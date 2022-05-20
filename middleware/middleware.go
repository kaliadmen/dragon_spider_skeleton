package middleware

import (
	dragonSpider "github.com/kaliadmen/dragon_spider"
	"myApp/data"
)

type Middleware struct {
	App    *dragonSpider.DragonSpider
	Models data.Models
}
