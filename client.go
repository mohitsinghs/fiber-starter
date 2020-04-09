package main

import (
	"github.com/gofiber/websocket"
	"github.com/google/uuid"
	"sync"
)

// socket connection
type Client struct {
	id   uuid.UUID
	conn *websocket.Conn
	hub  *Hub
	mt   sync.Mutex
}

// send msg to one client
func (c *Client) send(msg []byte) {
	c.mt.Lock()
	defer c.mt.Unlock()
	err := c.conn.WriteMessage(TextType, msg)
	if err != nil {
		c.hub.remove(c)
	}
}
