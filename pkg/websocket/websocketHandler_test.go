package websocket_test

import (
	"net/http"

	gorillaWebsocket "github.com/gorilla/websocket"
)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := gorillaWebsocket.Upgrader{}
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer connection.Close()

	for {
		_, _, err := connection.ReadMessage()
		if err != nil {
			break
		}
	}
}
