package main

import (
	"encoding/json"
	"github.com/gofiber/websocket"
)

type Event struct {
	Type string            `json:"event"`
	Data map[string]string `json:"data"`
}

func HandleWebsocket(clients *Hub) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		client := clients.add(c)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				clients.remove(client)
				clients.broadcastActive()
				break
			} else {
				event := Event{}
				json.Unmarshal(msg, &event)
				switch event.Type {
				case "hello":
					client.send([]byte("Hello World !!"))
					break
				}
			}
		}
	}
}
