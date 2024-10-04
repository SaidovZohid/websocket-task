package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SaidovZohid/websocket-task/database"
	"github.com/SaidovZohid/websocket-task/websocket"
	wb "github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestWebSocketConnectionWithUserID(t *testing.T) {
	conn, err := database.Make()
	if err != nil {
		log.Fatal("Database Make:", err)
	}

	clientManager := &websocket.ClientManager{
		Clients: make(map[string]*websocket.Client),
		Db:      conn,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleConnections(clientManager, conn, w, r)
	}))
	defer server.Close()

	url := "ws" + server.URL[len("http"):] + "?user_id=1"
	client1, _, err := wb.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err, "Client 1 should connect successfully")
	defer client1.Close()

	time.Sleep(100 * time.Millisecond)

	clientManager.Broadcast("Hello, Client 1!")

	_, msg, err := client1.ReadMessage()
	assert.NoError(t, err, "Client 1 should receive the message")
	assert.Equal(t, "Hello, Client 1!", string(msg), "Client 1 should receive the correct message")

	client1.Close()

	url = "ws" + server.URL[len("http"):] + "?user_id=2"
	client2, _, err := wb.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err, "Client 2 should connect successfully")
	defer client2.Close()

	clientManager.Broadcast("Hello, Client 2!")

	_, msg, err = client2.ReadMessage()
	assert.NoError(t, err, "Client 2 should receive the message")
	assert.Equal(t, "Hello, Client 2!", string(msg), "Client 2 should receive the correct message")
}
