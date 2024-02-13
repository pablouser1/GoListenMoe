package native

import (
	"io"
	"net/http"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
	"github.com/pablouser1/GoListenMoe/constants/enums"
	"github.com/pablouser1/GoListenMoe/models"
)

const modifier float64 = 1

var streamer beep.StreamSeekCloser
var ctrl *beep.Ctrl
var vol *effects.Volume

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
	var decoder func(rc io.ReadCloser) (s beep.StreamSeekCloser, format beep.Format, err error)

	// Setup appropriate stream and decoder
	if fallback {
		// Use fallback url and mp3 decoder
		stream = genre.Fallback
		decoder = mp3.Decode
	} else {
		// Use standard url and vorbis decoder
		stream = genre.Stream
		decoder = vorbis.Decode
	}

	// Get HTTP Stream
	res, err := http.Get(stream)
	if err != nil {
		return err
	}

	// Decode http body
	var format beep.Format
	streamer, format, err = decoder(res.Body)

	if err != nil {
		return err
	}

	// Start audio stream
	ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	vol = &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(vol)

	return nil
}

func toggle() {
	speaker.Lock()
	ctrl.Paused = !ctrl.Paused
	speaker.Unlock()
}

func volume(delta float64) {
	speaker.Lock()
	vol.Volume += delta
	speaker.Unlock()
}

func mute() {
	speaker.Lock()
	vol.Silent = !vol.Silent
	speaker.Unlock()
}

func stop() {
	streamer.Close()
}
