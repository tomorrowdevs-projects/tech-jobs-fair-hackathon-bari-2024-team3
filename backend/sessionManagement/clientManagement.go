package sessionManagement

import (
	"fmt"
	"net"
	"strings"
)

var allConnectedClients = make(map[string]net.Conn)

func SocketListenerLoop() {

	//Connect via terminal on: $ telnet 127.0.0.1 8080
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Listening for %s on: %s\n", listener.Addr().Network(), listener.Addr().String())

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	connectionAddr := conn.RemoteAddr().String()
	allConnectedClients[connectionAddr] = conn
	fmt.Println("Accepted connection from", connectionAddr)

	for {
		conn.Write([]byte("How can we be of any help today?\n"))
		readBuffer := make([]byte, 10000)
		conn.Read(readBuffer)
		msg := string(readBuffer)

		if strings.Contains(msg, "bye") {
			println("Closing down connection")
			break
		} else if strings.Contains(msg, "greetAll") {
			fmt.Println("Greetings from ", connectionAddr)
			broadcastToAllClients("Hello everyone!\n")

		} else {
			println("Reading:", msg)
		}
	}

	defer conn.Close()

}

func broadcastToAllClients(msg string) {
	broadcastToClients(msg, allConnectedClients)
}

func broadcastToClients(msg string, clients map[string]net.Conn) {
	for _, client := range clients {
		client.Write([]byte(msg))
	}
}
