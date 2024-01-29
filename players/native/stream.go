package native

import (
	"io"
	"net/http"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/vorbis"
	"github.com/pablouser1/GoListenMoe/models"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

var streamer beep.StreamSeekCloser

func Start(genre models.Genre, fallback bool) (func(), error) {
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
		return nil, err
	}

	// Decode http body
	var format beep.Format
	streamer, format, err = decoder(res.Body)

	if err != nil {
		return nil, err
	}

	// Start audio stream
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	return Stop, nil
}

func Stop() {
	streamer.Close()
}
