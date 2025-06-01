package ws

import (
	"github.com/Defenestrationq/forum-api/internal/entity"
	"github.com/Defenestrationq/forum-api/internal/usecase"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn  *websocket.Conn
	send  chan entity.Message
	disID int
}

type Hub struct {
	clients    map[*Client]bool
	chat       chan entity.Message
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
	msgUseCase usecase.MsgUseCase
	disUseCase usecase.DisUseCase
}

func NewHub(msg usecase.MsgUseCase, dis usecase.DisUseCase) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		chat:       make(chan entity.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mu:         sync.Mutex{},
		msgUseCase: msg,
		disUseCase: dis,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
			}
			h.mu.Unlock()

		case message := <-h.chat:
			h.mu.Lock()
			for client := range h.clients {
				if client.disID == message.DiscussionID {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
			h.mu.Unlock()
		}
	}
}
