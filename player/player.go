package player

import (
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var streamer beep.StreamSeekCloser

func Start(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error sending HTTP request")
	}

	l_streamer, format, err := mp3.Decode(resp.Body)
	streamer = l_streamer
	if err != nil {
		log.Fatal("Error decoding")
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}

func Stop() {
	streamer.Close()
}
