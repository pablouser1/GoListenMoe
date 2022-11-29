package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pablouser1/GoListenMoe/player"
	"github.com/pablouser1/GoListenMoe/socket"
	"github.com/urfave/cli/v2"
)

const JPOP_STEAM string = "https://listen.moe/stream"
const KPOP_STEAM string = "https://listen.moe/kpop/stream"

const JPOP_FALLBACK_STEAM string = "https://listen.moe/fallback"
const KPOP_FALLBACK_STEAM string = "https://listen.moe/kpop/fallback"

const JPOP_SOCKET string = "wss://listen.moe/gateway_v2"
const KPOP_SOCKET string = "wss://listen.moe/kpop/gateway_v2"

func loop() {
	// After user exists close everything
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Exiting...")
	player.Stop()
	socket.Stop()
}

func main() {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "fallback",
			Aliases: []string{"f"},
			Usage:   "Use MP3 stream",
			Value:   false,
		},
	}

	app := &cli.App{
		Name: "GoListenMoe",
		Commands: []*cli.Command{
			{
				Name:    "jpop",
				Aliases: []string{"j"},
				Usage:   "Listen to J-Pop",
				Action: func(cCtx *cli.Context) error {
					var STREAM_URL string
					if cCtx.Bool("fallback") {
						STREAM_URL = JPOP_FALLBACK_STEAM
					} else {
						STREAM_URL = JPOP_STEAM
					}
					socket.Start(JPOP_SOCKET)
					player.Start(STREAM_URL, cCtx.Bool("fallback"))
					loop()
					return nil
				},
				Flags: flags,
			},
			{
				Name:    "kpop",
				Aliases: []string{"k"},
				Usage:   "Listen to K-Pop",
				Action: func(cCtx *cli.Context) error {
					var STREAM_URL string
					if cCtx.Bool("fallback") {
						STREAM_URL = KPOP_FALLBACK_STEAM
					} else {
						STREAM_URL = KPOP_STEAM
					}
					socket.Start(KPOP_SOCKET)
					player.Start(STREAM_URL, cCtx.Bool("fallback"))
					loop()
					return nil
				},
				Flags: flags,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
