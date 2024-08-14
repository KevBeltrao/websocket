package websocket

import (
	"github.com/kevbeltrao/websocket/pkg/utils"
	"github.com/kevbeltrao/websocket/pkg/websocket"
)

type Room = websocket.Room

var NewRoom = websocket.NewRoom

type Connection = websocket.Connection

var NewConnection = websocket.NewConnection

var SetLogger = utils.SetLogger
