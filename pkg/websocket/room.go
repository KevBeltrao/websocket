package websocket

import (
	"sync"

	"github.com/kevbeltrao/websocket/pkg/utils"
)

type Room struct {
	Id         string
	Clients    map[ConnectionInterface]bool
	Broadcast  chan []byte
	Register   chan ConnectionInterface
	Unregister chan ConnectionInterface
	Mutex      sync.Mutex
}

func NewRoom(id string) *Room {
	return &Room{
		Id:         id,
		Clients:    make(map[ConnectionInterface]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan ConnectionInterface),
		Unregister: make(chan ConnectionInterface),
	}
}

func (room *Room) addClient(connection ConnectionInterface) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	room.Clients[connection] = true
	utils.LogInfo("Client added to room %s", room.Id)
}

func (room *Room) removeClient(connection ConnectionInterface) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	if _, ok := room.Clients[connection]; ok {
		delete(room.Clients, connection)
		connection.Close()
		utils.LogInfo("Client removed from room %s", room.Id)
	}
}

func (room *Room) broadcastMessage(message []byte) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for connection := range room.Clients {
		go func(connection ConnectionInterface) {
			err := connection.SendMessage(message)
			if err != nil {
				utils.LogError("Error sending message to client: %v", err)
			}
		}(connection)
	}
}

func (room *Room) Run() {
	for {
		select {
		case connection := <-room.Register:
			room.addClient(connection)
		case connection := <-room.Unregister:
			room.removeClient(connection)
		case message := <-room.Broadcast:
			room.broadcastMessage(message)
		}
	}
}
