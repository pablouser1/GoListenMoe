package main

import (
	"flag"
	"fmt"

	"github.com/pablouser1/GoListenMoe/helpers"
	"github.com/pablouser1/GoListenMoe/player/socket"
	"github.com/pablouser1/GoListenMoe/player/stream"
	"github.com/pablouser1/GoListenMoe/ui"
)

func main() {
	fallback := flag.Bool("f", false, "Use fallback MP3")
	flag.Parse()

	genreStr := "jpop"

	if len(flag.Args()) > 0 {
		genreStr = flag.Args()[0]
	}

	genre := helpers.GetGenreById(genreStr)

	playing, err := socket.Start(genre.Socket)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = stream.Start(genre, *fallback)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Blocking function, shows the UI
	ui.Cli(playing)

	// Cleanup
	stream.Stop()
	socket.Stop()
}
