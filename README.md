# Go Listen.moe
Cross platform Listen.moe client made using Go

## Installation
```bash
go get && go build .
```

## How to use
```bash
./GoListenMoe jpop --fallback # (or kpop)
```

## TODO
* Improve UI
* Play Vorbis correctly (right now using fallback stream)
* Hot-swapping between JPop and Kpop
* Tidy code

## External libraries used
* Beep (https://github.com/faiface/beep)
* Gorilla websocket (https://github.com/gorilla/websocket)
