package sessionManagement

import (
	"fmt"
	"log"
	"net"
	"net/http"
	quizmanagement "quizzy_game/quizManagement"
	usermanagement "quizzy_game/userManagement"

	"github.com/gorilla/websocket"
)

var allConnectedClients = make(map[string]net.Conn)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleRequest(request string) {

	// TODO Add some handler logic deciding if the request needs to go to userManaging or quizManaging
	// if isUserReq {
	usermanagement.HandleUser(request)
	// }else{
	quizmanagement.HandleQuiz(request)
	// }

}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))
		handleRequest(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
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
