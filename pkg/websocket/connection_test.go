package websocket_test

import (
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	gorillaWebsocket "github.com/gorilla/websocket"
	"github.com/kevbeltrao/websocket/pkg/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConnection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(websocketHandler))
	defer server.Close()

	websocketUrl := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/"}
	gorillaConnection, _, err := gorillaWebsocket.DefaultDialer.Dial(websocketUrl.String(), nil)
	require.NoError(t, err)
	defer gorillaConnection.Close()

	connection := websocket.NewConnection(gorillaConnection)
	assert.NotNil(t, connection)
	assert.Equal(t, gorillaConnection, connection.Connection)
}

func TestSendMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(websocketHandler))
	defer server.Close()

	websocketUrl := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/"}
	gorillaConnection, _, err := gorillaWebsocket.DefaultDialer.Dial(websocketUrl.String(), nil)
	require.NoError(t, err)
	defer gorillaConnection.Close()

	connection := websocket.NewConnection(gorillaConnection)
	err = connection.SendMessage([]byte("test message"))
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(websocketHandler))
	defer server.Close()

	websocketUrl := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/"}
	gorillaConnection, _, err := gorillaWebsocket.DefaultDialer.Dial(websocketUrl.String(), nil)
	require.NoError(t, err)

	connection := websocket.NewConnection(gorillaConnection)
	err = connection.Close()
	assert.NoError(t, err)

	err = gorillaConnection.WriteMessage(gorillaWebsocket.TextMessage, []byte("test"))
	assert.Error(t, err)
}
