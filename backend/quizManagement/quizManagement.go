package quizmanagement

import (
	"fmt"
	"quizzy_game/api"
	"quizzy_game/internal/dataTypes"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var quizzes = make(map[string]*dataTypes.Quiz)
var categories = make(map[int]dataTypes.Category)

func init() {

	var dbCategories []dataTypes.Category = api.GetCategories()
	for _, category := range dbCategories {
		categories[category.Id] = category
	}

}

func HandleQuiz(quizUpdate string) {

	/*
		commands:
		createQuiz FirstQuiz Christian
		joinQuiz <quizId> Jack
		joinQuiz 05126db1-dbb2-49a3-b187-9bbe46bc48a4 Jack
		print
	*/

	fmt.Println("Quiz update: ", quizUpdate)
	input := strings.Fields(quizUpdate)

	switch input[0] {
	case "createQuiz":
		createQuiz(input[1], categories[0], dataTypes.Easy, dataTypes.MultipleChoice, input[2])
	case "joinQuiz":
		joinQuiz(input[1], input[2])
	case "leaveQuiz":
		leaveQuiz(input[1], input[2])
	case "startQuiz":
		startQuiz(input[1])
	case "print":
		fmt.Println("Printing Quizzes:")
		for _, quiz := range quizzes {
			println(quiz.String())
		}
	}

}

func createQuiz(name string, category dataTypes.Category, difficulty dataTypes.Difficulty, quizType dataTypes.QuestionType, userId string) {

	participants := make(map[string]int)
	participants[userId] = 0
	questions := api.GetQuestions(category.Id, difficulty, quizType)
	questionTuples := []dataTypes.QuestionTuple{}
	for _, question := range questions {
		questionTuples = append(questionTuples, question.ToTuple())
	}
	newQuiz := dataTypes.Quiz{
		Id:           uuid.NewString(),
		Name:         name,
		QuizStatus:   dataTypes.StatusInitialized,
		Category:     category,
		Difficulty:   dataTypes.Easy,
		Type:         dataTypes.MultipleChoice,
		Questions:    questionTuples,
		Participants: participants,
	}

	quizzes[newQuiz.Id] = &newQuiz
	fmt.Printf("Sucessfully created Quiz %s with ID: %s\n", newQuiz.Name, newQuiz.Id)
}

// Use a reference to user instead of a userIdString
func joinQuiz(quizID string, userId string) {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Quiz not found.")
		return
	}
	if _, ok := quiz.Participants[userId]; !ok {
		// User does not exist, so insert the user
		quiz.Participants[userId] = 0
		fmt.Printf("Added user %s to Quiz: %s\n", userId, quizID)
	}
}

// Use a reference to user instead of a userIdString
func leaveQuiz(quizID string, userId string) {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Quiz not found.")
		return
	}
	if _, ok := quiz.Participants[userId]; ok {
		delete(quiz.Participants, userId)
		fmt.Printf("Deleted user %s from Quiz %s: ", userId, quizID)
	}
}

func startQuiz(quizID string) {
	quiz, ok := quizzes[quizID]
	switch {
	case !ok:
		fmt.Println("Quiz not found.")
		return
	case quiz.QuizStatus != dataTypes.StatusInitialized:
		fmt.Println("Quiz could not start. Expected status Initialized but got " + string(quiz.QuizStatus))
		return
	case len(quiz.Participants) < 2:
		fmt.Println("Quiz could not start. Expected PARTICIPANTS to contain more than 2. got ", len(quiz.Participants))
		return
	default:
		quiz.QuizStatus = dataTypes.StatusStart
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusStart)
	}
}

func startQuestion(quizID string) {
	quiz, ok := quizzes[quizID]
	switch {
	case !ok:
		fmt.Println("Quiz not found.")
		return
	case quiz.QuizStatus != dataTypes.StatusStart || quiz.QuizStatus != dataTypes.StatusEvaluation:
		fmt.Println("Quiz could not start. Expected status Start or Evaluation but got " + string(quiz.QuizStatus))
		return
	case quiz.RemainingQuestions() < 1:
		fmt.Println("Quiz could not start. No more Questions to ask.")
		return
	default:
		quiz.QuizStatus = dataTypes.StatusQuizTime
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusQuizTime)
	}

}

func endQuestion(quizID string) {
	quiz, ok := quizzes[quizID]
	switch {
	case !ok:
		fmt.Println("Quiz not found.")
		return
	case quiz.QuizStatus != dataTypes.StatusQuizTime:
		fmt.Println("Quiz could not start. Expected status QuizTime but got " + string(quiz.QuizStatus))
		return
	default:
		quiz.QuizStatus = dataTypes.StatusEvaluation
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusEvaluation)
	}

}

func questionLoop(quizID string) {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Question Loop: Quiz not found.")
		return
	}
	if quiz.RemainingQuestions() < 1 {
		fmt.Println("Question Loop: No questions remaining.")
		return
	}

	var wg sync.WaitGroup
	broadcastChannel := make(chan string)
	receiverChannel := make(chan string)
	statusChannel := make(chan dataTypes.QuizStatus)

	startTimerRoutine(&wg, statusChannel)

	for _, question := range quiz.Questions {
		if question.IsAskedStatus {
			break
		}
		startQuestion(quizID)

		for status := range statusChannel {
			if status == dataTypes.StatusQuizTimeEnded {
				fmt.Println("Question Time Ended !")
				endQuestion(quizID)
			} else if status == dataTypes.StatusEvaluationEnded {
				fmt.Println("Evaluation Time Ended!")
				break
			}
		}
	}

	wg.Wait() // Wait for all goroutines to finish
	close(broadcastChannel)
	close(receiverChannel)
	close(statusChannel)

}

func endQuiz(quizID string) {
	quiz, ok := quizzes[quizID]
	switch {
	case !ok:
		fmt.Println("Quiz not found.")
		return
	default:
		quiz.QuizStatus = dataTypes.StatusEnded
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusEnded)
	}

}

func resetQuiz(quizID string) {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Quiz not found.")
		return
	}
	fmt.Println("Resetting Quiz with ID: ", quiz.Id)

	for _, question := range quiz.Questions {
		question.IsAskedStatus = false
	}
	fmt.Println("Questions reset for Quiz with ID: ", quiz.Id)

	for userId := range quiz.Participants {
		quiz.Participants[userId] = 0
	}
	fmt.Println("Scores for users reset for Quiz with ID: ", quiz.Id)

	quiz.QuizStatus = dataTypes.StatusInitialized
	fmt.Println("Quiz Status updated to: ", dataTypes.StatusInitialized)

}

func startTimerRoutine(wg *sync.WaitGroup, statusChannel chan dataTypes.QuizStatus) {
	answerTimeout := 10 * time.Second
	evaluationTimeout := 5 * time.Second

	wg.Add(1)
	go func() {
		for status := range statusChannel {

			if status == dataTypes.StatusQuizTime {
				answerTimer := time.NewTimer(answerTimeout)
				fmt.Println("Answer timer started!")
				<-answerTimer.C
				statusChannel <- dataTypes.StatusQuizTimeEnded
				fmt.Println("Answer timer ENDED! EvalStatus broadcasted")

			} else if status == dataTypes.StatusEvaluation {
				evalTimer := time.NewTimer(evaluationTimeout)
				fmt.Println("Evaluation timer started!")
				<-evalTimer.C
				statusChannel <- dataTypes.StatusEvaluationEnded
				fmt.Println("Evaluation timer ENDED! QuizTime broadcasted")

			} else if status == dataTypes.StatusEnded {
				fmt.Println("Shutting down the timer GoRoutine!")
				defer wg.Done()
				return
			}
		}

	}()

}
