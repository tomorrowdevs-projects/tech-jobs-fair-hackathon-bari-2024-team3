package sessionManagement

import (
	"fmt"
	"net"
	"strings"
)

var peers = make(map[string]net.Conn)

func SocketListenerLoop() {

	//Connect via terminal on: $ telnet 127.0.0.1 8080
	listener, err := net.Listen("tcp", ":8080")
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

		addString := conn.RemoteAddr().String()
		peers[conn.RemoteAddr().String()] = conn
		fmt.Printf("Handling connection: %s\n", addString)

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	fmt.Println("Accepted connection from", conn.RemoteAddr())
	conn.Write([]byte("Welcome to the game!\n"))

	for {
		conn.Write([]byte("How can we be of any help today?\n"))
		readBuffer := make([]byte, 10000)
		conn.Read(readBuffer)
		msg := string(readBuffer)
		if strings.Contains(msg, "bye") {
			println("Closing down connection")
			break
		} else if strings.Contains(msg, "greetAll") {
			fmt.Println("Greetings from ", conn.RemoteAddr())
			greetAll()

		} else {
			println("Reading:", msg)
		}
	}

	defer conn.Close()

}

func greetAll() {
	for _, peer := range peers {
		peer.Write([]byte("Hello everyone!\n"))
	}
}
