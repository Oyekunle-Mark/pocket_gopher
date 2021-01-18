package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var msg *message

		err := c.socket.ReadJSON(&msg)

		if err != nil {
			return
		}

		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)

		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			return
		}
	}
}
