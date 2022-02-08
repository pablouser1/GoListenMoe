package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/sacOO7/gowebsocket"
)

const JPOP_STEAM string = "https://listen.moe/stream"
const KPOP_STEAM string = "https://listen.moe/kpop/stream"

const JPOP_SOCKET string = "wss://listen.moe/gateway_v2"
const KPOP_SOCKET string = "wss://listen.moe/kpop/gateway_v2"

type SocketRes struct {
	Op int64 `json:"op"`
	D  json.RawMessage
}

type SendData struct {
	Op int64 `json:"op"`
}

type HeartbeatData struct {
	Message   string `json:"message"`
	Heartbeat int64  `json:"heartbeat"`
}

type PlayingData struct {
	Song       Song        `json:"song"`
	Requester  interface{} `json:"requester"`
	Event      interface{} `json:"event"`
	StartTime  string      `json:"startTime"`
	LastPlayed []Song      `json:"lastPlayed"`
	Listeners  int64       `json:"listeners"`
}

type Song struct {
	ID       int64         `json:"id"`
	Title    string        `json:"title"`
	Sources  []interface{} `json:"sources"`
	Artists  []Album       `json:"artists"`
	Albums   []Album       `json:"albums"`
	Duration int64         `json:"duration"`
}

type Album struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	NameRomaji *string `json:"nameRomaji"`
	Image      *string `json:"image"`
}

func sendHeartBeat(socket gowebsocket.Socket) {
	data := SendData{
		Op: 9,
	}
	data_str, _ := json.Marshal(data)

	socket.SendText(string(data_str))
}

func setHeartbeat(socket gowebsocket.Socket, repeat int64) {
	sendHeartBeat(socket)
	ticker := time.NewTicker(time.Duration(repeat) * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				sendHeartBeat(socket)
			}
		}
	}()
}

func handleMessage(msg_str string, socket gowebsocket.Socket) {
	var msg SocketRes
	msg_bytes := []byte(msg_str)
	json.Unmarshal(msg_bytes, &msg)
	switch msg.Op {
	case 0:
		var data HeartbeatData
		json.Unmarshal(msg.D, &data)
		setHeartbeat(socket, data.Heartbeat)
	case 1:
		var data PlayingData
		json.Unmarshal(msg.D, &data)
		fmt.Print("\033[H\033[2J")
		fmt.Println("Now Playing:")
		fmt.Println("Title: " + data.Song.Title)
		fmt.Println("Artist: " + data.Song.Artists[0].Name)
		fmt.Println("Album: " + data.Song.Albums[0].Name)
	}
}

func main() {
	mode := "jpop"
	var STREAM_URL string
	var SOCKET_URL string

	// Setup args
	if len(os.Args) == 2 {
		mode = os.Args[1]
	}

	// MODE
	switch mode {
	case "kpop":
		STREAM_URL = KPOP_STEAM
		SOCKET_URL = KPOP_SOCKET
	case "jpop":
		STREAM_URL = JPOP_STEAM
		SOCKET_URL = JPOP_SOCKET
	default:
		fmt.Println("Invalid mode")
		os.Exit(1)
	}
	socket := gowebsocket.New(SOCKET_URL)
	socket.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to server")
	}
	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Recieved connect error ", err)
	}
	socket.OnTextMessage = handleMessage
	socket.Connect()

	// Player
	fmt.Println("Starting player")
	player := exec.Command("mplayer", STREAM_URL)
	player.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			fmt.Println("Exiting...")
			socket.Close()
			return
		}
	}
}
