package websocket

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type ClientManager struct {
	Clients map[string]*Client
	Mutex   sync.Mutex
	Db      *sql.DB
}

func (cm *ClientManager) AddClient(client *Client, id string) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	client.Online = true
	client.Id = id
	cm.Clients[id] = client
	go client.ListenForMessages(cm)

	rows, err := cm.Db.Query("SELECT content FROM messages WHERE user_id = $1 AND seen = 0 ORDER BY created_at ASC", id)
	if err != nil {
		log.Println("Failed to scan unread messages:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			log.Println("Failed to scan unread messages:", err)
			continue
		}
		client.SendMessage(cm, content)
	}

	_, err = cm.Db.Exec("UPDATE messages SET seen = 1 WHERE user_id = $1 AND seen = 0", id)
	if err != nil {
		log.Println("Failed to scan unread messages:", err)
	}
}

func (cm *ClientManager) RemoveClient(client *Client) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	if val, ok := cm.Clients[client.Id]; ok {
		val.Online = false
		fmt.Println("I am disconnecting it")
	}
}

func (cm *ClientManager) Broadcast(message string) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	for _, client := range cm.Clients {
		var seen bool
		if client.Online {
			fmt.Println("Still online user id -> ", client.Id)
			client.Send <- []byte(message)
			seen = true
		} else {
			fmt.Println("Offline user id -> ", client.Id)
		}

		_, err := cm.Db.Exec("INSERT INTO messages (user_id, content, seen, created_at) VALUES ($1, $2, $3, $4)", client.Id, message, seen, time.Now())
		if err != nil {
			log.Printf("Failed to save message for user %v: %v", client.Id, err)
		}
	}
}
