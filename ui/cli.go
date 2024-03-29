package ui

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pablouser1/GoListenMoe/constants/enums"
	"github.com/pablouser1/GoListenMoe/models"
)

func showAlbum(albums []models.Album) string {
	album := "None"
	if len(albums) > 0 {
		album = albums[0].Name
	}
	return "Album: " + album
}

func writeToScreen(now models.Song, last models.Song, listeners int64, start string) {
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

func Cli(playing chan models.PlayingData, action chan uint8) {
	go func() {
		for {
			now := <-playing
			writeToScreen(now.Song, now.LastPlayed[0], now.Listeners, now.StartTime)
		}
	}()

	// Start playing
	action <- enums.ACTION_PLAY

	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		aTmp, _ := reader.ReadString('\n')
		a := strings.TrimRight(aTmp, "\n")

		switch a {
		case "p":
			action <- enums.ACTION_TOGGLE
		case "+":
			action <- enums.ACTION_VOLUME_UP
		case "-":
			action <- enums.ACTION_VOLUME_DOWN
		case "m":
			action <- enums.ACTION_VOLUME_MUTE
		case "q":
			action <- enums.ACTION_STOP
			running = false
		}

		// Remove that line
		fmt.Printf("\033[1A\033[K")
	}
}
