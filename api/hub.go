package api

import (
	"log"
	"sync"
)

type Hub struct {
	clients        []*Client
	addedClients   chan *Client
	removedClients chan *Client
	mutex          sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:        []*Client{},
		addedClients:   make(chan *Client),
		removedClients: make(chan *Client),
		mutex:          sync.Mutex{},
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.addedClients:
			h.handleAddClient(c)
		case c := <-h.removedClients:
			h.handleRemoveClient(c)
		}
	}
}

func (h *Hub) AddClient(c *Client) {
	h.addedClients <- c
}

func (h *Hub) RemoveClient(c *Client) {
	h.removedClients <- c
}

func (h *Hub) handleAddClient(c *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.clients = append(h.clients, c)
}

func (h *Hub) handleRemoveClient(c *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	c.Close()

	for i, x := range h.clients {
		if x.id == c.id {
			h.clients[i] = nil
			h.clients = append(h.clients[:i], h.clients[i+1:]...)
			break
		}
	}
}

func (h *Hub) handleMessageReceived(c *Client, data []byte) {
	log.Printf("message received: %s\n", string(data))
}
