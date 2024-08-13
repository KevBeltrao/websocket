package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionInterface interface {
	SendMessage(message []byte) error
	Close() error
}

type Connection struct {
	Connection *websocket.Conn
	Mutex      sync.Mutex
}

func NewConnection(connection *websocket.Conn) *Connection {
	return &Connection{
		Connection: connection,
	}
}

func (connection *Connection) SendMessage(message []byte) error {
	connection.Mutex.Lock()
	defer connection.Mutex.Unlock()
	return connection.Connection.WriteMessage(websocket.TextMessage, message)
}

func (connection *Connection) Close() error {
	connection.Mutex.Lock()
	defer connection.Mutex.Unlock()
	return connection.Connection.Close()
}
