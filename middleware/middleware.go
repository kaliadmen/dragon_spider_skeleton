package middleware

import (
	"${APP_NAME}/data"
	dragonSpider "github.com/kaliadmen/dragon_spider"
)

type Middleware struct {
	App    *dragonSpider.DragonSpider
	Models data.Models
}
