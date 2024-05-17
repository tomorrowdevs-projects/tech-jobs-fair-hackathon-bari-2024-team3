package sessionManagement

import (
	"fmt"
	"log"
	"net/http"
	"quizzy_game/internal/dataTypes"
	quizmanagement "quizzy_game/quizManagement"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, user *dataTypes.User) {
	var mutex sync.Mutex // Create a mutex to synchronize writes to the WebSocket connection

	go func() {
		for {
			data, more := <-user.MsgChannel
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
			close(user.MsgChannel)
			return
		}
		msg := string(p)

		input := strings.Fields(msg)
		if len(input) > 1 && input[0] == "setUsername" {
			newUsername := input[1]
			user.Name = newUsername
			user.MsgChannel <- "Your Username has been updated to: " + newUsername

		}

		// Adding data to the channel should also be synchronized
		mutex.Lock()
		go quizmanagement.HandleQuizUpdate(msg, user)
		err = conn.WriteMessage(messageType, p)
		mutex.Unlock()

		if err != nil {
			log.Println(err)
			close(user.MsgChannel)
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

	responseChan := make(chan string, 1)
	newUserId := uuid.NewString()
	newUser := dataTypes.User{
		Id:         newUserId,
		Name:       "user" + newUserId[0:4],
		MsgChannel: responseChan,
	}
	reader(ws, &newUser)
}
