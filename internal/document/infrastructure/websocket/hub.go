package websocket

import (
	"sync"

	"github.com/google/uuid"
)

type Hub struct {
	rooms      map[uuid.UUID]*Room
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

type Room struct {
	documentID uuid.UUID
	clients    map[*Client]bool
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		rooms:      make(map[uuid.UUID]*Room),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.addClient(client)
		case client := <-h.unregister:
			h.removeClient(client)
		}
	}
}

func (h *Hub) addClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	room, exists := h.rooms[client.documentID]
	if !exists {
		room = &Room{
			documentID: client.documentID,
			clients:    make(map[*Client]bool),
			broadcast:  make(chan []byte, 256),
		}
		h.rooms[client.documentID] = room
		go h.runRoom(room)
	}
	room.clients[client] = true
	client.room = room
}

func (h *Hub) removeClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if client.room != nil {
		delete(client.room.clients, client)
		if len(client.room.clients) == 0 {
			close(client.room.broadcast)
			delete(h.rooms, client.documentID)
		}
	}
	close(client.send)
}

func (h *Hub) runRoom(room *Room) {
	for message := range room.broadcast {
		for client := range room.clients {
			select {
			case client.send <- message:
			default:
				h.unregister <- client
			}
		}
	}
}

func (h *Hub) Register() chan<- *Client {
	return h.register
}

func (h *Hub) Unregister() chan<- *Client {
	return h.unregister
}
