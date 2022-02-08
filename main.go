package main

import (
	"fmt"
	"os"

	"github.com/pablouser1/GoListenMoe/player"
	"github.com/pablouser1/GoListenMoe/socket"
	"github.com/pablouser1/GoListenMoe/viewer"
)

const JPOP_STEAM string = "https://listen.moe/stream"
const KPOP_STEAM string = "https://listen.moe/kpop/stream"

const JPOP_SOCKET string = "wss://listen.moe/gateway_v2"
const KPOP_SOCKET string = "wss://listen.moe/kpop/gateway_v2"

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

	player.Start(STREAM_URL)
	socket.Start(SOCKET_URL)
	viewer.Init()
	// This will block the thread
	viewer.Poll()
	// After user exists close everything
	socket.Stop()
	player.Stop()
}
