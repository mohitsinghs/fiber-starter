package main

import (
	"encoding/json"
	"github.com/gofiber/websocket"
	"github.com/google/uuid"
	"sync"
	"sync/atomic"
)

// connection hub
type Hub struct {
	count   int64
	clients map[*Client]bool
	mt      sync.RWMutex
}

type Active struct {
	Count int64 `json:"active"`
}

func NewHub() *Hub {
	return &Hub{
		count:   0,
		clients: make(map[*Client]bool),
		mt:      sync.RWMutex{},
	}
}

// add client
func (h *Hub) add(conn *websocket.Conn) *Client {
	h.mt.Lock()
	defer h.mt.Unlock()
	id := uuid.New()
	client := &Client{
		id:   id,
		conn: conn,
		hub:  h,
	}
	atomic.AddInt64(&h.count, 1)
	h.clients[client] = true
	return client
}

// remove client
func (h *Hub) remove(client *Client) {
	h.mt.Lock()
	defer h.mt.Unlock()
	delete(h.clients, client)
	atomic.AddInt64(&h.count, -1)
}

// send msg to all client
func (h *Hub) sendAll(msg []byte) {
	h.mt.RLock()
	defer h.mt.RUnlock()
	for c, _ := range h.clients {
		c.send(msg)
	}
}

func (h *Hub) broadcastActive() {
	h.mt.RLock()
	active, _ := json.Marshal(Active{
		Count: h.count,
	})
	h.mt.RUnlock()
	h.sendAll(active)
}
