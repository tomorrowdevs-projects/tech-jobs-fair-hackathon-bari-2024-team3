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

func HandleQuizUpdate(quizUpdate string, responseChannel chan string) {

	fmt.Println("Quiz update: ", quizUpdate)
	input := strings.Fields(quizUpdate)
	if len(input) < 1 {
		responseChannel <- "Not enough input parameters! Try: \n\tcreateQuiz \n\tjoinQuiz \n\tleaveQuiz \n\tstartQui \n\tresetQuiz \n\tprint"
		return
	}

	switch input[0] {
	case "createQuiz":
		if len(input) < 3 {
			responseChannel <- "Not enough input parameters! Try: createQuiz <quizName> <username>"
			return
		}
		//TODO Fix hardcoded Values: Categories, difficulty and type
		responseChannel <- createQuiz(input[1], categories[0], dataTypes.Easy, dataTypes.MultipleChoice, input[2])
	case "joinQuiz":
		if len(input) < 3 {
			responseChannel <- "Not enough input parameters! Try: joinQuiz <quizID> <username>"
			return
		}
		responseChannel <- joinQuiz(input[1], input[2])
	case "leaveQuiz":
		if len(input) < 3 {
			responseChannel <- "Not enough input parameters! Try: leaveQuiz <quizID> <username>"
			return
		}
		responseChannel <- leaveQuiz(input[1], input[2])
	case "startQuiz":
		if len(input) < 2 {
			responseChannel <- "Not enough input parameters! Try: startQuiz <quizID>"
			return
		}
		responseChannel <- startQuiz(input[1])
	case "resetQuiz":
		if len(input) < 2 {
			responseChannel <- "Not enough input parameters! Try: resetQuiz <quizID>"
			return
		}
		responseChannel <- resetQuiz(input[1])
	case "print":
		quizPrintString := "Printing Quizzes: \n"
		for _, quiz := range quizzes {
			quizPrintString += quiz.String()
		}
		fmt.Println(quizPrintString)
		responseChannel <- quizPrintString

	}

}

func createQuiz(name string, category dataTypes.Category, difficulty dataTypes.Difficulty, quizType dataTypes.QuestionType, userId string) string {

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
	return fmt.Sprintf("QuizID: %s, QuizStatus: %s, participants: %s ", newQuiz.Id, newQuiz.QuizStatus, newQuiz.ParticipantsAsString())
}

// Use a reference to user instead of a userIdString
func joinQuiz(quizID string, userId string) string {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Quiz not found.")
		return "Error joining Quiz: " + quizID
	}
	if _, ok := quiz.Participants[userId]; !ok {
		// User does not exist, so insert the user
		quiz.Participants[userId] = 0
		fmt.Printf("Added user %s to Quiz: %s\n", userId, quizID)
	}
	return fmt.Sprintf("QuizID: %s, QuizStatus: %s, participants: %s\n", quiz.Id, quiz.QuizStatus, quiz.ParticipantsAsString())

}

// Use a reference to user instead of a userIdString
func leaveQuiz(quizID string, userId string) string {
	quiz, ok := quizzes[quizID]
	if !ok {
		return "Error leaving Quiz. Quiz not found. ID: " + quizID
	}
	if _, ok := quiz.Participants[userId]; ok {
		delete(quiz.Participants, userId)
		fmt.Printf("Deleted user %s from Quiz %s: ", userId, quizID)
	}
	return fmt.Sprintf("User: %s left Quiz QuizID: %s\n", userId, quiz.Id)

}

func startQuiz(quizID string) string {
	quiz, ok := quizzes[quizID]
	switch {
	case !ok:
		return "Error Starting. Quiz not found. ID: " + quizID
	case quiz.QuizStatus != dataTypes.StatusInitialized:
		return fmt.Sprintf("Quiz could not start. Expected status Initialized but got " + string(quiz.QuizStatus))
	case len(quiz.Participants) < 2:
		return fmt.Sprintf("Quiz could not start. Expected PARTICIPANTS to contain more than 2. got %d.", len(quiz.Participants))
	default:
		quiz.QuizStatus = dataTypes.StatusStart
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusStart)

		var wg sync.WaitGroup
		broadcastChannel := make(chan string)
		receiverChannel := make(chan string)
		statusChannel := make(chan dataTypes.QuizStatus)

		wg.Add(1)
		go timerRoutine(&wg, statusChannel)

		wg.Add(1)
		go questionLoopRoutine(quizID, &wg, statusChannel, broadcastChannel)

		wg.Wait() // Wait for all goroutines to finish

		quiz.QuizStatus = dataTypes.StatusEnded
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusEnded)
		broadcastChannel <- fmt.Sprintf("Quiz %s finished. Scoreboard is: %s", quiz.Name, quiz.ScoreBoard())

		close(broadcastChannel)
		close(receiverChannel)
		close(statusChannel)
	}
	return fmt.Sprintf("Sucessfully started quiz with ID: %s\n", quiz.Id)
}

func resetQuiz(quizID string) string {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Quiz not found.")
		return "Error reset quiz with ID: " + quizID
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
	return fmt.Sprintf("Quiz Status updated to: %s", dataTypes.StatusInitialized)

}

func questionLoopRoutine(quizID string, wg *sync.WaitGroup, statusChannel chan dataTypes.QuizStatus, broadcastChannel chan string) {
	quiz, ok := quizzes[quizID]
	if !ok {
		fmt.Println("Question Loop: Quiz not found.")
		return
	}
	fmt.Println(quiz.String())
	if quiz.RemainingQuestions() < 1 {
		fmt.Println("Question Loop: No questions remaining.")
		return
	}

	for _, question := range quiz.Questions {
		if question.IsAskedStatus {
			break
		}
		quiz.QuizStatus = dataTypes.StatusQuizTime
		fmt.Println("Quiz Status updated to: ", dataTypes.StatusQuizTime)
		broadcastChannel <- question.Ref.Question
		statusChannel <- dataTypes.StatusQuizTime
		question.IsAskedStatus = true

		for status := range statusChannel {
			if status == dataTypes.StatusQuizTimeEnded {
				fmt.Println("Quiz Time Ended!")
				quiz.QuizStatus = dataTypes.StatusEvaluation
				fmt.Println("Quiz Status updated to: ", dataTypes.StatusEvaluation)
				break
				// } else if status == dataTypes.StatusEvaluationEnded {
				// 	fmt.Println("Evaluation Time Ended!")
				// 	break
			}
		}
	}

	defer wg.Done()
}

func timerRoutine(wg *sync.WaitGroup, statusChannel chan dataTypes.QuizStatus) {
	answerTimeout := 10 * time.Second
	evaluationTimeout := 5 * time.Second

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

}
