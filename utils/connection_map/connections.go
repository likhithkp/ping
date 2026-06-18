package connectionmap

import (
	"github.com/gofiber/contrib/websocket"
)

/*
Use redis to maintain active connection map once the app is called horizontally
*/
var Connections = make(map[string]*websocket.Conn)

func Add(userId string, c *websocket.Conn) {
	Connections[userId] = c
}

func Remove(userId string) {
	delete(Connections, userId)
}
