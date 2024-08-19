package client

import (
	"errors"
	"mgmt/common/web/response"
	"sync/atomic"
	"time"
	"xxx/common/configs"
	"xxx/common/web/response"
	"xxx/common/web/ws"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func NewWsClient(socket *websocket.Conn) *WsClient {
	return &WsClient{
		id:          uuid.New().String(),
		socket:      socket,
		sendChannal: make(chan []byte),
	}
}

func (client *WsClient) InitReader(callback func(client *WsClient, message []byte, err error)) {
	if client.socket == nil {
		return
	}
	go func() {
		for {
			_, msg, err := client.socket.ReadMessage()

			if client.keepAlive(msg) {
				continue
			}

			callback(client, msg, err)
			if err != nil {
				return
			}
		}
	}()
}

func (client *WsClient) InitSender() {
	go func() {
		for {
			if client.IsClose() {
				return
			}
			msg, ok := <-client.sendChannal
			if ok {
				if decodeType := configs.GetInt(configs.SECTION_SYSTEM, configs.SYSTEM_WEBSOCKET_DECODE_MODE, configs.ENABLE_WS_BASE64); decodeType == configs.ENABLE_WS_MSG_PACK {
					client.socket.WriteMessage(websocket.BinaryMessage, msg)
				} else {
					client.socket.WriteMessage(websocket.TextMessage, msg)
				}
			}
		}
	}()
}

func (client *WsClient) Send(message []byte) {
	go func() {
		client.mutex.Lock()
		defer client.mutex.Unlock()

		if client.IsClose() {
			return
		}

		defer func() {
			if err := recover(); err != nil {
				client.SetClose()
			}
		}()

		client.sendChannal <- message
	}()
}

func (client *WsClient) SendThenClose(message []byte, delayTime time.Duration) {
	client.Send(message)

	go func() {
		time.Sleep(delayTime)
		client.Close()
	}()
}

func (client *WsClient) Close() error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if client.IsClose() {
		return errors.New(response.MSG_WS_IS_CLOSED)
	}

	if client.socket == nil {
		return errors.New("websocket is nil")
	}

	client.SetClose()

	client.socket.Close()

	close(client.sendChannal)

	return nil
}

func (client *WsClient) IsClose() bool {
	return atomic.CompareAndSwapInt64(&client.isClose, 1, 1)
}

func (client *WsClient) SetClose() {
	atomic.StoreInt64(&client.isClose, 1)
}

func (client *WsClient) GetID() string {
	return client.id
}

func (client *WsClient) keepAlive(msg []byte) bool {

	event, err := ws.ParseEvent(msg)

	if err != nil {
		return false
	}

	if event.Event != ws.EVENT_WS_KEEP_ALIVE {
		return false
	}

	client.Send(ws.GetEventResponse(ws.EVENT_WS_KEEP_ALIVE, response.CODE_SUCCESS, response.MSG_SUCCESS, nil))
	return true
}
