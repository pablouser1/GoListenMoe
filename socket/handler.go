package socket

import (
	"encoding/json"

	"github.com/pablouser1/GoListenMoe/models"
)

func handleMessage(in []byte, playing chan models.PlayingData) {
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
		playing <- data
	}
}
