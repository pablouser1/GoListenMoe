package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pablouser1/GoListenMoe/player"
	"github.com/pablouser1/GoListenMoe/socket"
)

const JPOP_STEAM string = "https://listen.moe/fallback"
const KPOP_STEAM string = "https://listen.moe/kpop/fallback"

const JPOP_SOCKET string = "wss://listen.moe/gateway_v2"
const KPOP_SOCKET string = "wss://listen.moe/kpop/gateway_v2"

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
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

	socket.Start(SOCKET_URL)
	player.Start(STREAM_URL)
	// After user exists close everything
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Exiting...")
	player.Stop()
	socket.Stop()
}
