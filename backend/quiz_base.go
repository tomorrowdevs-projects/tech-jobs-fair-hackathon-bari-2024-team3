package main

import (
	"fmt"
	"quizzy_game/api"
	"time"
)

func Main_Loop() {

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
