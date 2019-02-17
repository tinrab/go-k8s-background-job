package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Transport struct {
	port   uint16
	router *mux.Router
	hub    *Hub
}

func NewTransport(port uint16) *Transport {
	return &Transport{
		port:   port,
		router: mux.NewRouter(),
		hub:    NewHub(),
	}
}

func (t *Transport) Run() error {
	go t.hub.Run()

	t.router.HandleFunc("/ws", t.handleWebsocket)

	return http.ListenAndServe(fmt.Sprintf(":%d", t.port), t.router)
}

func (t *Transport) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := NewClient(t.hub, socket)
	t.hub.AddClient(client)

	client.Run()
}
