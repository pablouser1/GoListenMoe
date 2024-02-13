# Go Listen.moe
Cross platform Listen.moe client made using Go

## Installation
```bash
go get && go build .
```

## How to use
```bash
./GoListenMoe [-f] [-p mpv,native] jpop # (or kpop)
```

* -f: Forces fallback mode, currently required for native to work properly
* -p: Player to use, options are: native, mpv

## TODO
* Improve UI
* Play Vorbis correctly in native (right now using fallback stream)
* Hot-swapping between JPop and Kpop
* Add compability for Windows when using mpv

## External libraries used
* Beep (https://github.com/faiface/beep)
* Gorilla websocket (https://github.com/gorilla/websocket)
