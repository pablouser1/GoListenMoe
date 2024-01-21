package genres

import "github.com/pablouser1/GoListenMoe/models"

var JPOP = models.Genre{
	Id:       "jpop",
	Stream:   "https://listen.moe/stream",
	Fallback: "https://listen.moe/fallback",
	Socket:   "wss://listen.moe/gateway_v2",
}
