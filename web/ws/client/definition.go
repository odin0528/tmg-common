package client

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type IWsClient interface {
	Send([]byte)
	SendThenClose([]byte, time.Duration)
	Close() error
	IsClose() bool
	GetID() string
}

type WsClient struct {
	id          string
	socket      *websocket.Conn
	sendChannal chan []byte
	isClose     int64
	closeChan   chan bool
	mutex       sync.Mutex
}
