package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/pablouser1/GoListenMoe/player"
	"github.com/pablouser1/GoListenMoe/socket"
)

type Stream struct {
	standard string
	fallback string
	socket   string
}

type FinalStream struct {
	stream string
	socket string
}

func loop() {
	// After user exists close everything
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Exiting...")
	player.Stop()
	socket.Stop()
}

func streamPicker(genre string, fallback bool) FinalStream {
	jpop := Stream{
		standard: "https://listen.moe/stream",
		fallback: "https://listen.moe/fallback",
		socket:   "wss://listen.moe/gateway_v2",
	}

	kpop := Stream{
		standard: "https://listen.moe/kpop/stream",
		fallback: "https://listen.moe/kpop/fallback",
		socket:   "wss://listen.moe/kpop/gateway_v2",
	}

	pickedStream := kpop
	if genre == "jpop" || genre == "j" {
		pickedStream = jpop
	}

	stream := pickedStream.standard
	if fallback {
		stream = pickedStream.fallback
	}

	return FinalStream{
		stream: stream,
		socket: pickedStream.socket,
	}

}

func main() {
	fallback := flag.Bool("f", false, "Use fallback MP3")
	flag.Parse()

	genre := "jpop"

	if len(flag.Args()) > 0 {
		genre = flag.Args()[0]
	}

	streaming := streamPicker(genre, *fallback)

	socket.Start(streaming.socket)
	player.Start(streaming.stream, *fallback)
	loop()
}
