package mpv

import (
	"os"
	"os/exec"

	"github.com/pablouser1/GoListenMoe/constants/enums"
	"github.com/pablouser1/GoListenMoe/models"
)

const modifier float64 = 1

var cmd *exec.Cmd

func Init(genre models.Genre, fallback bool, action chan uint8) {
	go func() {
		for {
			a := <-action

			switch a {
			case enums.ACTION_PLAY:
				// Play
				play(genre, fallback)
			case enums.ACTION_TOGGLE:
				// Toggle
				toggle()
			case enums.ACTION_VOLUME_UP:
				// Volume up
				volume(modifier)
			case enums.ACTION_VOLUME_DOWN:
				// Volume down
				volume(-modifier)
			case enums.ACTION_VOLUME_MUTE:
				// Mute
				mute()
			case enums.ACTION_STOP:
				// Stop
				stop()
			}
		}
	}()
}

func play(genre models.Genre, fallback bool) error {
	var stream string

	if fallback {
		stream = genre.Fallback
	} else {
		stream = genre.Stream
	}

	cmd = exec.Command("mpv", "--vid=no", stream)
	err := cmd.Start()

	if err != nil {
		return err
	}

	return nil
}

func toggle() {
	// TODO: Implement toggle
}

func volume(delta float64) {
	// TODO: Implement volume
}

func mute() {
	// TODO: Implement mute
}

func stop() {
	// Send SIGINT
	cmd.Process.Signal(os.Interrupt)
}
