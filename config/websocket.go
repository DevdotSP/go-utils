package config

import (
	"log"
	"sync"
	"github.com/gofiber/websocket/v2"
)

var (
	clients = make(map[string]*websocket.Conn)
	mu      sync.Mutex
)

func WebSocketHandler(c *websocket.Conn) {
	userID := c.Params("id")
	if userID == "" {
		log.Println("Missing user ID")
		return
	}

	mu.Lock()
	clients[userID] = c
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, userID)
		mu.Unlock()
		c.Close()
	}()

	for {
		if _, _, err := c.ReadMessage(); err != nil {
			log.Printf("User %s disconnected", userID)
			break
		}
	}
}

