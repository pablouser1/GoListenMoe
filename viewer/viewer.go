package viewer

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/pablouser1/GoListenMoe/model"
)

func showAlbum(albums []model.Album) string {
	album := "None"
	if len(albums) > 0 {
		album = albums[0].Name
	}
	return "Album: " + album
}

func WriteToScreen(now model.Song, last model.Song, listeners int64, start string) {
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
