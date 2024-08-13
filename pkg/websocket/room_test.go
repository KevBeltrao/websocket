package websocket_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	gorillaWebsocket "github.com/gorilla/websocket"
	"github.com/kevbeltrao/websocket/pkg/mock"
	"github.com/kevbeltrao/websocket/pkg/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRoom(t *testing.T) {
	roomID := "testRoom"
	room := websocket.NewRoom(roomID)

	assert.Equal(t, roomID, room.Id)
	assert.NotNil(t, room.Clients)
	assert.NotNil(t, room.Broadcast)
	assert.NotNil(t, room.Register)
	assert.NotNil(t, room.Unregister)
}

func TestRoomAddAndRemoveClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(websocketHandler))
	defer server.Close()

	websocketUrl := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/"}
	connection, _, err := gorillaWebsocket.DefaultDialer.Dial(websocketUrl.String(), nil)
	require.NoError(t, err)
	defer connection.Close()

	room := websocket.NewRoom("testRoom")

	go room.Run()

	clientConnection := websocket.NewConnection(connection)

	room.Register <- clientConnection
	time.Sleep(100 * time.Millisecond)

	require.Contains(t, room.Clients, clientConnection)

	room.Unregister <- clientConnection
	time.Sleep(100 * time.Millisecond)

	assert.NotContains(t, room.Clients, clientConnection)
}

func TestRoomBroadcastMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnection1 := mock.NewMockConnectionInterface(ctrl)
	mockConnection2 := mock.NewMockConnectionInterface(ctrl)

	room := websocket.NewRoom("testRoom")
	go room.Run()

	message := []byte("test message")
	mockConnection1.EXPECT().SendMessage(message).Return(nil)
	mockConnection2.EXPECT().SendMessage(message).Return(nil)

	room.Register <- mockConnection1
	room.Register <- mockConnection2

	room.Broadcast <- message

	time.Sleep(100 * time.Millisecond)

	ctrl.Finish()
}

func TestRoomRun(t *testing.T) {
	room := websocket.NewRoom("testRoom")

	go room.Run()

	done := make(chan bool)
	go func() {
		server := httptest.NewServer(http.HandlerFunc(websocketHandler))
		defer server.Close()

		websocketUrl := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/"}
		conn, _, err := gorillaWebsocket.DefaultDialer.Dial(websocketUrl.String(), nil)
		require.NoError(t, err)
		defer conn.Close()

		clientConnection := websocket.NewConnection(conn)
		room.Register <- clientConnection
		time.Sleep(100 * time.Millisecond)
		room.Unregister <- clientConnection
		time.Sleep(100 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		assert.True(t, true)
	case <-time.After(1 * time.Second):
		t.Error("Test timed out")
	}
}
