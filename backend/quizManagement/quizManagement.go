package quizmanagement

import (
	"fmt"
	"quizzy_game/internal/dataTypes"
)

var quizzes = make(map[string]dataTypes.Quiz)

func HandleQuiz(quizUpdate string) {

	/* TODO:
	/ Take ind the quiz commands/request from the WebSocket and decide what to do and which of the functions to call.
	/
	*/

}

func createQuiz(category dataTypes.Category) {
	action := "CREATE"
	fmt.Printf("You are trying to %s Quiz %s, but the functionality has not yet been implemented!", action)
}

func joinQuiz(id string) {
	action := "JOIN"
	fmt.Printf("You are trying to %s Quiz %s, but the functionality has not yet been implemented!", action, id)
}

func startQuiz(id string) {
	action := "START"
	fmt.Printf("You are trying to %s Quiz %s, but the functionality has not yet been implemented!", action, id)

}

func startQuestion(id string) {
	action := "START"
	fmt.Printf("You are trying to %s QUESTION %s, but the functionality has not yet been implemented!", action, id)

}

func endQuestion(id string) {
	action := "END"
	fmt.Printf("You are trying to %s QUESTION %s, but the functionality has not yet been implemented!", action, id)

}

func questionLoop(id string) {
	action := "RUN Question Loop"
	fmt.Printf("You are trying to %s for %s, but the functionality has not yet been implemented!", action, id)

}

func endQuiz(id string) {
	action := "END"
	fmt.Printf("You are trying to %s Quiz %s, but the functionality has not yet been implemented!", action, id)

}
