package helpers

import (
	"fmt"

	"github.com/pablouser1/GoListenMoe/constants/genres"
	"github.com/pablouser1/GoListenMoe/models"
)

func GetGenreById(genreStr string) (models.Genre, error) {
	var genre models.Genre

	switch genreStr {
	case "jpop", "j":
		genre = genres.JPOP
	case "kpop", "k":
		genre = genres.KPOP
	default:
		return models.Genre{}, fmt.Errorf("genre not available")
	}

	return genre, nil
}
