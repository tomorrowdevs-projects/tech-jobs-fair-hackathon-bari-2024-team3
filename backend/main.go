package main

import (
	"fmt"
	"io"
	"net/http"
	"quizzy_game/api"
	"quizzy_game/sessionManagement"
)

func getFrontPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	http.HandleFunc("/", getFrontPage)
	http.HandleFunc("/questions", api.GetQuestionsWeb)
	http.HandleFunc("/categories", api.GetCategoriesWeb)
	http.HandleFunc("/ws", sessionManagement.WsEndpoint)

	http.ListenAndServe(":3333", nil)

}
