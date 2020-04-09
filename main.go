package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/websocket"
)

const TextType = 1

func main() {
	app := fiber.New()
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(cors.New())
	clients := NewHub()

	app.Get("/hello/:name?", func(c *fiber.Ctx) {
		name := c.Params("name")
		if name != "" {
			c.Status(fiber.StatusOK).Send("Hello " + name + " !!")
		} else {
			c.Status(fiber.StatusOK).Send("Hello World !!")
		}
	})

	app.Get("/ws", websocket.New(HandleWebsocket(clients)))

	app.Listen(4500)
}
