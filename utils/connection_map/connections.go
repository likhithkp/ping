package connectionmap

import (
	"log"
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

func RemoveStaleConnections() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Checking stale connections...")
		log.Println("connections map", Connections)

		if len(Connections) == 0 {
			continue
		}

		for userId, user := range Connections {
			elapsed := time.Since(user.LastHeartBeat)

			if elapsed >= 45*time.Second {
				log.Printf("User %s last heartbeat activeness has crossed 45s, removing user from connection map\n", userId)
				Remove(userId)
				user.Ws.Close()
			}
		}
	}
}

func UpdateHeartbeat(userId string) {
	Connections[userId].LastHeartBeat = time.Now().UTC()
}
