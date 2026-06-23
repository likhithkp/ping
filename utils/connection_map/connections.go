package connectionmap

import (
	"log"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
)

var mu *sync.RWMutex

type User struct {
	Ws            *websocket.Conn
	LastHeartBeat time.Time
}

/*
Use redis to maintain active connection map once the app is scaled horizontally
*/
var Connections = make(map[string]*User)

func Add(userId string, c *websocket.Conn) {
	mu.Lock()
	Connections[userId] = &User{
		Ws:            c,
		LastHeartBeat: time.Now().UTC(),
	}
	mu.Unlock()
}

func Remove(userId string) {
	mu.Lock()
	delete(Connections, userId)
	mu.Unlock()
}

func RemoveStaleConnections() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Checking stale connections...")

		if len(Connections) == 0 {
			continue
		}

		for userId, user := range Connections {
			elapsed := time.Since(user.LastHeartBeat)

			if elapsed >= 60*time.Second {
				log.Printf("Removed offline user %s from ws connection map\n", userId)
				Remove(userId)
				user.Ws.Close()
			}
		}

	}
}

func UpdateHeartbeat(userId string) {
	mu.Lock()
	Connections[userId].LastHeartBeat = time.Now().UTC()
	mu.Unlock()
}
