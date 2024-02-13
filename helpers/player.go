package helpers

import (
	"fmt"

	"github.com/pablouser1/GoListenMoe/models"
	"github.com/pablouser1/GoListenMoe/players/mpv"
	"github.com/pablouser1/GoListenMoe/players/native"
)

func InitPlayer(p string, genre models.Genre, fallback bool, action chan uint8) error {
	var err error = nil

	switch p {
	case "native":
		native.Init(genre, fallback, action)
	case "mpv":
		mpv.Init(genre, fallback, action)
	default:
		err = fmt.Errorf("invalid player")
	}

	return err
}
