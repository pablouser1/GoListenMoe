package constants

import (
	"github.com/pablouser1/GoListenMoe/models"
)

var JPOP_GENRE = models.Genre{
	Id:       "jpop",
	Stream:   "https://listen.moe/stream",
	Fallback: "https://listen.moe/fallback",
	Socket:   "wss://listen.moe/gateway_v2",
}

var KPOP_GENRE = models.Genre{
	Id:       "kpop",
	Stream:   "https://listen.moe/kpop/stream",
	Fallback: "https://listen.moe/kpop/fallback",
	Socket:   "wss://listen.moe/kpop/gateway_v2",
}
