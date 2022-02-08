package socket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pablouser1/GoListenMoe/model"
	"github.com/pablouser1/GoListenMoe/viewer"
	"github.com/sacOO7/gowebsocket"
)

var socket gowebsocket.Socket

func sendHeartBeat() {
	data := model.SendData{
		Op: 9,
	}
	data_bytes, _ := json.Marshal(data)

	socket.SendBinary(data_bytes)
}

func setHeartbeat(repeat int64) {
	sendHeartBeat()
	ticker := time.NewTicker(time.Duration(repeat) * time.Millisecond)
	go func() {
		<-ticker.C
		sendHeartBeat()
	}()
}

func handleMessage(msg_str string, _ gowebsocket.Socket) {
	var msg model.SocketRes
	msg_bytes := []byte(msg_str)
	json.Unmarshal(msg_bytes, &msg)
	switch msg.Op {
	case 0:
		var data model.HeartbeatData
		json.Unmarshal(msg.D, &data)
		setHeartbeat(data.Heartbeat)
	case 1:
		var data model.PlayingData
		json.Unmarshal(msg.D, &data)
		album := "None"
		if len(data.Song.Albums) > 0 {
			album = data.Song.Albums[0].Name
		}
		viewer.WriteToScreen(data.Song.Title, data.Song.Artists[0].Name, album)
	}
}

func Start(url string) {
	socket = gowebsocket.New(url)
	socket.OnConnected = func(_ gowebsocket.Socket) {
		fmt.Println("Connected to server")
	}
	socket.OnConnectError = func(err error, _ gowebsocket.Socket) {
		fmt.Println("Recieved connect error ", err)
	}
	socket.OnTextMessage = handleMessage
	socket.OnDisconnected = func(err error, _ gowebsocket.Socket) {
		fmt.Printf("Disconnected from socket")
		if err != nil {
			fmt.Println(err)
		}
	}
	socket.Connect()
}

func Stop() {
	socket.Close()
}
