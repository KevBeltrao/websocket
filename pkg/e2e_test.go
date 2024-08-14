package websocket_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	gorillaWebsocket "github.com/gorilla/websocket"
	"github.com/kevbeltrao/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := gorillaWebsocket.Upgrader{}
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connectionection", http.StatusBadRequest)
		return
	}
	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			break
		}

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			break
		}
	}
}

func TestE2EWebSocket(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(websocketHandler))
	defer server.Close()

	u := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/"}
	conn1, _, err := gorillaWebsocket.DefaultDialer.Dial(u.String(), nil)
	require.NoError(t, err)
	defer conn1.Close()

	conn2, _, err := gorillaWebsocket.DefaultDialer.Dial(u.String(), nil)
	require.NoError(t, err)
	defer conn2.Close()

	room := websocket.NewRoom("e2eRoom")
	go room.Run()

	clientConn1 := websocket.NewConnection(conn1)
	clientConn2 := websocket.NewConnection(conn2)

	room.Register <- clientConn1
	room.Register <- clientConn2

	message := []byte("E2E test message")
	room.Broadcast <- message

	time.Sleep(100 * time.Millisecond)

	_, receivedMessage1, err := conn1.ReadMessage()
	require.NoError(t, err)
	assert.Equal(t, message, receivedMessage1)

	_, receivedMessage2, err := conn2.ReadMessage()
	require.NoError(t, err)
	assert.Equal(t, message, receivedMessage2)
}
