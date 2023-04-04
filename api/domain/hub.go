package domain

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// Client manager
// With CheckOrigin for allowing clients
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Mutex for race conditions
// Using go channels to registry the websocket clients
type WebsocketHub struct {
	clients    []*WebsocketClient
	register   chan *WebsocketClient
	unregister chan *WebsocketClient
	mutex      *sync.Mutex
}

func NewHub() *WebsocketHub {
	return &WebsocketHub{
		clients:    make([]*WebsocketClient, 0),
		register:   make(chan *WebsocketClient),
		unregister: make(chan *WebsocketClient),
		mutex:      &sync.Mutex{},
	}
}

func (hub *WebsocketHub) RegistryClient(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	client := NewWebSocketClient(hub, socket)
	hub.register <- client

	go client.Write()
}

func (hub *WebsocketHub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnected(client)
		case client := <-hub.unregister:
			hub.onDisconnected(client)
		}
	}
}

func (hub *WebsocketHub) onConnected(client *WebsocketClient) {
	log.Println("Client connected from: ", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	client.id = client.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, client)
}

func (hub *WebsocketHub) onDisconnected(client *WebsocketClient) {
	log.Println("Client disconnected from: ", client.socket.RemoteAddr())
	client.socket.Close()

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	i := -1
	for j, c := range hub.clients {
		if c.id == client.id {
			i = j
		}
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}

func (hub *WebsocketHub) Broadcast(message any, ignore *WebsocketClient) {
	// Serialize
	data, _ := json.Marshal(message)
	for _, client := range hub.clients {
		if client != ignore {
			fmt.Println("Entre")
			fmt.Println(client)
			client.outbound <- data
		}
	}
}
