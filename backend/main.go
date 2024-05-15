package main

import (
	"fmt"
	"io"
	"net/http"
	"quizzy_game/api"
	"quizzy_game/sessionManagement"
	"time"
)

func getFrontPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	// Should we keep the token if app crashes? And what if app is closed? (Would say forget and get a new one)

	http.HandleFunc("/", getFrontPage)
	http.HandleFunc("/questions", api.GetQuestionsWeb)
	http.HandleFunc("/categories", api.GetCategoriesWeb)
	http.HandleFunc("/ws", sessionManagement.WsEndpoint)
	main_loop()
	http.ListenAndServe(":3333", nil)

}

func main_loop() {

	questions := api.GetQuestions()

	for i, question := range questions {
		fmt.Println("Welcome to the quiz, you have 30 seconds to answer the questions.")

		// answer possibilities
		answers := append(question.WrongAnswer, question.CorrectAnswer)

		timer := time.NewTimer(30 * time.Second) // Set the timer to 30 seconds

		// scores := map [string]int

		fmt.Println("Question", i+1, ":", question.Question)

		for i, answer := range answers {
			println("Option ", i+1, ":", answer)
		}

		var userAnswer string

		fmt.Print("Answer: ")
		fmt.Scanln(&userAnswer)
		<-timer.C // Wait until the timer fires

		if userAnswer == question.CorrectAnswer {

			fmt.Println("Time's up! You answer is correct")
		} else {
			fmt.Println("Time's up! You answer is wrong")
		}
	}
}

// https://opentdb.com/api.php?amount=10
