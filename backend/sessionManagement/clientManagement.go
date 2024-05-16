package sessionManagement

import (
	"fmt"
	"log"
	"net/http"
	quizmanagement "quizzy_game/quizManagement"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	responseChan := make(chan string, 1)
	var mutex sync.Mutex // Create a mutex to synchronize writes to the WebSocket connection

	go func() {
		for {
			data, more := <-responseChan
			if more {
				mutex.Lock()
				err := conn.WriteMessage(websocket.TextMessage, []byte(data))
				mutex.Unlock()
				if err != nil {
					fmt.Println("Error writing to WebSocket:", err)
					break
				}
			} else {
				break
			}
		}
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			close(responseChan)
			return
		}

		// Adding data to the channel should also be synchronized
		mutex.Lock()
		quizmanagement.HandleQuizUpdate(string(p), responseChan)
		err = conn.WriteMessage(messageType, p)
		mutex.Unlock()

		if err != nil {
			log.Println(err)
			close(responseChan)
			return
		}
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}
