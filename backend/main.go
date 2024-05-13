package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {

	var session_token, err = http.Get("https://opentdb.com/api_token.php?command=request")

	http.HandleFunc("/", getRoot)

	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3333", nil)

}

// https://opentdb.com/api.php?amount=10
