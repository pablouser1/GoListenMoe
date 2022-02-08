# Go Listen.moe
Listen.moe client made using Go

## Installation
This project depends on mplayer for playing audio

```bash
go get && go build .
```

## How to use
```bash
./GoListenMoe jpop # (or kpop)
```

## TODO
* Build an actual UI using gocui

## Known issues
* Some requests made through the webproxy trigger an error (Invalid OP code, maybe related to the JAP/KOR characters) and disconnects the socket
