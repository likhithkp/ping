package connectionmap

import (
	"time"

	"github.com/gofiber/contrib/websocket"
)

type User struct {
	Ws            *websocket.Conn
	LastHeartBeat time.Time
}

/*
Use redis to maintain active connection map once the app is scaled horizontally
*/
var Connections = make(map[string]*User)

func Add(userId string, c *websocket.Conn) {
	Connections[userId] = &User{
		Ws:            c,
		LastHeartBeat: time.Now().UTC(),
	}
}

func Remove(userId string) {
	delete(Connections, userId)
}
