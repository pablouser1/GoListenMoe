package helpers

import (
	"fmt"
	"os"

	"github.com/pablouser1/GoListenMoe/constants"
	"github.com/pablouser1/GoListenMoe/models"
)

func GetGenreById(genreStr string) models.Genre {
	var genre models.Genre

	switch genreStr {
	case "jpop", "j":
		genre = constants.JPOP_GENRE
	case "kpop", "k":
		genre = constants.KPOP_GENRE
	default:
		fmt.Println("Genre not available!")
		os.Exit(1)
	}

	return genre
}
