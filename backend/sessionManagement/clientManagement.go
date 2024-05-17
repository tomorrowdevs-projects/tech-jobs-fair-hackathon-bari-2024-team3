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

var users = make(map[string]*dataTypes.User)

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

	user.MsgChannel <- fmt.Sprintf("UserID: %s", user.Id)

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

	parts := strings.Split(r.URL.Path, "/")
	var id string
	if len(parts) > 2 {
		id = strings.TrimSpace(parts[2])
	}

	var user *dataTypes.User

	if id != "" {
		fmt.Println("User Connected with UserID: ", id)
		user = users[id]
		if user != nil {
			user.MsgChannel = make(chan string, 1)
		}
	}

	if user == nil {
		fmt.Println("User Connected without UserID or  with invalid UserID. Creating new User!")
		responseChan := make(chan string, 1)
		newUserId := uuid.NewString()
		newUser := dataTypes.User{
			Id:         newUserId,
			Name:       "user" + newUserId[0:4],
			MsgChannel: responseChan,
		}
		users[newUser.Id] = &newUser
		user = &newUser
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection

	reader(ws, user)
}
