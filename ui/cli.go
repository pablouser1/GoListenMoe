package ui

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/pablouser1/GoListenMoe/models"
)

func showAlbum(albums []models.Album) string {
	album := "None"
	if len(albums) > 0 {
		album = albums[0].Name
	}
	return "Album: " + album
}

func WriteToScreen(now models.Song, last models.Song, listeners int64, start string) {
	fmt.Print("\033[H\033[2J") // Clear screen
	parseStart, err := time.Parse(time.RFC3339, start)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	diff := math.Round(t.Sub(parseStart).Minutes())

	// Currently playing
	fmt.Printf("Now Playing (started %g mins ago)\n", diff)
	fmt.Println("Title: " + now.Title)
	fmt.Println("Artist: " + now.Artists[0].Name)
	fmt.Println(showAlbum(now.Albums))

	fmt.Print("\n")

	// Latest song
	fmt.Println("Latest song:")
	fmt.Println("Title: " + last.Title)
	fmt.Println("Artist: " + last.Artists[0].Name)
	fmt.Println(showAlbum(last.Albums))

	fmt.Print("\n")

	// Misc
	fmt.Println("Listeners: " + strconv.FormatInt(listeners, 10))
}

func Cli(playing chan models.PlayingData) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			now := <-playing
			WriteToScreen(now.Song, now.LastPlayed[0], now.Listeners, now.StartTime)
		}
	}()

	<-interrupt
}
