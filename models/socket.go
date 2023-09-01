package models

import "encoding/json"

type SocketRes struct {
	Op int64 `json:"op"`
	D  json.RawMessage
}

type SendData struct {
	Op int64 `json:"op"`
}

type HeartbeatData struct {
	Message   string `json:"message"`
	Heartbeat int64  `json:"heartbeat"`
}

type PlayingData struct {
	Song       Song        `json:"song"`
	Requester  interface{} `json:"requester"`
	Event      interface{} `json:"event"`
	StartTime  string      `json:"startTime"`
	LastPlayed []Song      `json:"lastPlayed"`
	Listeners  int64       `json:"listeners"`
}
