package api

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type Client struct {
	hub      *Hub
	id       uuid.UUID
	socket   *websocket.Conn
	outgress chan []byte
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		id:       uuid.NewV4(),
		socket:   socket,
		outgress: make(chan []byte),
	}
}

func (c *Client) Close() {
	if c.socket == nil {
		return
	}

	_ = c.socket.WriteMessage(websocket.CloseMessage, []byte{})
	_ = c.socket.Close()
	c.socket = nil
	close(c.outgress)
}

func (c *Client) Send(message interface{}) {
	data, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	c.outgress <- data
}

func (c *Client) readQueue() {
	defer func() {
		c.hub.removedClients <- c
	}()

	for {
		_, data, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.hub.handleMessageReceived(c, data)
	}
}

func (c *Client) writeQueue() {
	defer func() {
		c.hub.removedClients <- c
	}()

	for {
		select {
		case data, ok := <-c.outgress:
			if !ok {
				return
			}
			if err := c.socket.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		}
	}
}

func (c *Client) Run() {
	go c.readQueue()
	go c.writeQueue()
}
