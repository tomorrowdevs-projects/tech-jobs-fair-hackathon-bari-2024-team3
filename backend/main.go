package main

import (
	"fmt"
	"io"
	"net/http"
	"quizzy_game/api"
)

func getFrontPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	// Should we keep the token if app crashes? And what if app is closed? (Would say forget and get a new one)

	http.HandleFunc("/", getFrontPage)
	http.HandleFunc("/questions", api.GetQuestions)
	http.HandleFunc("/categories", api.GetCategories)

	http.ListenAndServe(":3333", nil)

}

// https://opentdb.com/api.php?amount=10
