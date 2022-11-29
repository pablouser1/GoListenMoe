package player

import (
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

var streamer beep.StreamSeekCloser

func Start(url string, fallback bool) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error sending HTTP request")
	}

	var format beep.Format

	if fallback {
		// Enable fallback mode (using mp3)
		streamer, format, err = mp3.Decode(resp.Body)
	} else {
		streamer, format, err = vorbis.Decode(resp.Body)
	}

	if err != nil {
		log.Fatal("Error decoding")
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}

func Stop() {
	streamer.Close()
}
