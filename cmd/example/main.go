package main

import (
	"net/http"

	"github.com/kevbeltrao/websocket"
)

func main() {
	room := websocket.NewRoom("exampleRoom")
	go room.Run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.NewUpgrader()
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			return
		}
		clientConnection := websocket.NewConnection(connection)
		room.Register <- clientConnection
	})

	http.ListenAndServe(":8000", nil)
}
