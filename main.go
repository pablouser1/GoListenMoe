package main

import (
	"flag"
	"fmt"

	"github.com/pablouser1/GoListenMoe/helpers"
	"github.com/pablouser1/GoListenMoe/socket"
	"github.com/pablouser1/GoListenMoe/ui"
)

func main() {
	fallback := flag.Bool("f", false, "Use fallback MP3")
	player := flag.String("p", "native", "Player to use")
	flag.Parse()

	genreStr := "jpop"

	if len(flag.Args()) > 0 {
		genreStr = flag.Args()[0]
	}

	genre, err := helpers.GetGenreById(genreStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	playing, err := socket.Start(genre.Socket)

	if err != nil {
		fmt.Println(err)
		return
	}

	playerStop, err := helpers.StartPlayer(*player, genre, *fallback)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Blocking function, shows the UI
	ui.Cli(playing)

	// Cleanup
	socket.Stop()
	playerStop()
}
