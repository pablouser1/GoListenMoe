package socket

import (
	"time"

	"github.com/pablouser1/GoListenMoe/models"
)

var ticker *time.Ticker

func sendHeartbeat() {
	data := models.SendData{
		Op: 9,
	}
	conn.WriteJSON(data)
}

func setHeartbeat(repeat int64) {
	sendHeartbeat()
	ticker = time.NewTicker(time.Duration(repeat) * time.Millisecond)
	go func() {
		for !done {
			<-ticker.C
			sendHeartbeat()
		}
	}()
}
