package mpv

import (
	"fmt"
	"os/exec"

	"github.com/pablouser1/GoListenMoe/models"
)

var cmd *exec.Cmd

func Start(genre models.Genre, fallback bool) (func(), error) {
	var stream string

	if fallback {
		stream = genre.Fallback
	} else {
		stream = genre.Stream
	}

	cmd = exec.Command("mpv", "--vid=no", stream)
	err := cmd.Start()

	if err != nil {
		return nil, err
	}

	return Stop, nil
}

func Stop() {
	// SIGINT gets propagated, no need to manually close it
	fmt.Println("\nClosing mpv")
}
