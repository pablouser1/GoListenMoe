package socket

import (
	"log"

	"github.com/pablouser1/GoListenMoe/models"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var done = false

func Start(socket string) (chan models.PlayingData, error) {
	playing := make(chan models.PlayingData)
	c, _, err := websocket.DefaultDialer.Dial(socket, nil)
	if err != nil {
		return nil, err
	}

	conn = c

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Panic("Couldn't read WebSocket message", err)
			}
			handleMessage(msg, playing)
		}
	}()

	return playing, nil
}

func Stop() {
	ticker.Stop()
	done = true
}
