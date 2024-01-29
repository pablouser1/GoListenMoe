package helpers

import (
	"fmt"

	"github.com/pablouser1/GoListenMoe/models"
	"github.com/pablouser1/GoListenMoe/players/mpv"
	"github.com/pablouser1/GoListenMoe/players/native"
)

func StartPlayer(p string, genre models.Genre, fallback bool) (func(), error) {
	var stop func()
	var err error

	switch p {
	case "native":
		stop, err = native.Start(genre, fallback)
	case "mpv":
		stop, err = mpv.Start(genre, fallback)
	default:
		stop = nil
		err = fmt.Errorf("invalid player")
	}

	return stop, err
}
