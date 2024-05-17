package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"quizzy_game/internal/dataTypes"
	"strconv"
	"time"
)

func GetQuestionsWeb(w http.ResponseWriter, r *http.Request) {

	for x := 10; x < 33; x += 7 {
		timer := time.NewTimer(7 * time.Second)
		questions := GetQuestions(x, dataTypes.Easy, dataTypes.MultipleChoice)
		// https://opentdb.com/api.php?amount=10&category=9&difficulty=easy&type=multiple

		fmt.Printf("got /questions request\n")
		for _, question := range questions {
			io.WriteString(w, question.String())
		}
		<-timer.C
	}

}

func GetCategoriesWeb(w http.ResponseWriter, r *http.Request) {

	categories := GetCategories()

	fmt.Printf("got /categories request with %d categories\n", len(categories))
	for _, cat := range categories {
		io.WriteString(w, cat.String())
	}
}

func GetCategories() []dataTypes.Category {

	type CategoryResponse struct {
		Categories []dataTypes.Category `json:"trivia_categories"`
	}

	categoryRequestUrl := "https://opentdb.com/api_category.php"
	responseBody := getRequest(categoryRequestUrl)

	var cr CategoryResponse
	var err = json.Unmarshal(responseBody, &cr)
	if err != nil {
		log.Fatal(err)
	}
	return cr.Categories
}

func GetQuestions(categoryId int, difficulty dataTypes.Difficulty, quizType dataTypes.QuestionType) []dataTypes.Question {

	type QuestionResponse struct {
		ResponseCode int                  `json:"response_code"`
		Questions    []dataTypes.Question `json:"results"`
	}

	questionRequestUrl := "https://opentdb.com/api.php?amount=10"
	questionRequestUrl += "&category=" + strconv.Itoa(categoryId)
	questionRequestUrl += "&difficulty=" + string(difficulty)
	questionRequestUrl += "&type=" + string(quizType)
	fmt.Println("Requesting questions on: ", questionRequestUrl)

	var qr QuestionResponse
	retries := 0
	for {
		if retries >= 3 {
			fmt.Println("GetQuestions: Max retries exceeded!")
			return []dataTypes.Question{}
		}
		responseBody := getSessionRequest(questionRequestUrl)

		var err = json.Unmarshal(responseBody, &qr)

		if err != nil {
			log.Fatal(err)
			break
		} else if qr.ResponseCode == 5 {
			printResponseDescription(qr.ResponseCode)
			fmt.Println("Retry again in 5 Seconds....")
			time.Sleep(5 * time.Second)
			break
		} else if qr.ResponseCode != 0 {
			printResponseDescription(qr.ResponseCode)
			break
		} else {
			return qr.Questions
		}

	}
	return qr.Questions

}

func printResponseDescription(responseCode int) {
	msg := fmt.Sprintf("ResponseCode %d: ", responseCode)
	switch responseCode {
	case 0:
		msg += "**Success** Returned results successfully."
	case 1:
		msg += "**No Results** Could not return results. The API doesn't have enough questions for your query. (Ex. Asking for 50 Questions in a Category that only has 20.)"
	case 2:
		msg += "**Invalid Parameter** Contains an invalid parameter. Arguements passed in aren't valid. (Ex. Amount = Five)"
	case 3:
		msg += "**Token Not Found** Session Token does not exist."
	case 4:
		msg += "**Response Empty** Request has returned no questions for the specified query.Cause: Cannot fullfil request query(amount/selection). Can also be caused by Exhaust options of current Token -> then Token reset is necessary."
	case 5:
		msg += "**Rate Limit** Too many requests have occurred. Each IP can only access the API once every 5 seconds."
	default:
		msg = "ERROR: UNKNOW RESPONSE CODE: " + strconv.Itoa(responseCode)
	}
	fmt.Println(msg)
}
