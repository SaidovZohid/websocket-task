package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id     string
	Conn   *websocket.Conn
	Send   chan []byte
	Online bool
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) WriteMessage(messageType int, data []byte) error {
	return c.Conn.WriteMessage(messageType, data)
}

func (c *Client) ListenForMessages(cm *ClientManager) {
	defer c.Close()
	for message := range c.Send {
		err := c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			cm.RemoveClient(c)
			break
		}
	}
}

func (c *Client) SendMessage(cm *ClientManager, msg string) {
	err := c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Printf("Error sending message: %v", err)
		cm.RemoveClient(c)
	}
}
