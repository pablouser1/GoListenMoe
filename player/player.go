package player

import (
	"os/exec"
	"syscall"
)

var player *exec.Cmd

func Start(url string) {
	player = exec.Command("mplayer", url)
	player.Start()
}

func Stop() {
	player.Process.Signal(syscall.SIGTERM)
}
