package domain

import (
	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	hub      *WebsocketHub
	id       string
	socket   *websocket.Conn
	outbound chan []byte
}

func NewWebSocketClient(hub *WebsocketHub, socket *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (w *WebsocketClient) Write() {
	for {
		select {
		case message, ok := <-w.outbound:
			// Socket is close
			if !ok {
				w.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
