package genres

import "github.com/pablouser1/GoListenMoe/models"

var KPOP = models.Genre{
	Id:       "kpop",
	Stream:   "https://listen.moe/kpop/stream",
	Fallback: "https://listen.moe/kpop/fallback",
	Socket:   "wss://listen.moe/kpop/gateway_v2",
}
