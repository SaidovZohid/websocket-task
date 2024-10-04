package websocket

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(cm *ClientManager, db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer ws.Close()

	client := &Client{Conn: ws, Send: make(chan []byte)}
	cm.AddClient(client, id)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			cm.RemoveClient(client)
			break
		}
		log.Printf("Received message: %s", msg)
	}
}
