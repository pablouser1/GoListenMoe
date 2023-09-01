package socket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pablouser1/GoListenMoe/models"
	"github.com/pablouser1/GoListenMoe/viewer"
)

var conn *websocket.Conn
var done = false
var ticker *time.Ticker

func sendHeartBeat() {
	data := models.SendData{
		Op: 9,
	}
	conn.WriteJSON(data)
}

func setHeartbeat(repeat int64) {
	sendHeartBeat()
	ticker = time.NewTicker(time.Duration(repeat) * time.Millisecond)
	go func() {
		for !done {
			<-ticker.C
			sendHeartBeat()
		}
	}()
}

func handleMessage(in []byte) {
	var msg models.SocketRes
	json.Unmarshal(in, &msg)
	switch msg.Op {
	case 0:
		var data models.HeartbeatData
		json.Unmarshal(msg.D, &data)
		setHeartbeat(data.Heartbeat)
	case 1:
		var data models.PlayingData
		json.Unmarshal(msg.D, &data)
		viewer.WriteToScreen(data.Song, data.LastPlayed[0], data.Listeners, data.StartTime)
	}
}

func Start(url string) {
	conn_l, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Couldn't connect to websocket", err)
	}
	conn = conn_l

	go func() {
		for {
			if done {
				conn.Close()
				break
			}
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Panic("Couldn't read WebSocket message", err)
			}
			handleMessage(msg)
		}
	}()
}

func Stop() {
	ticker.Stop()
	done = true
}
