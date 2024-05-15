package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"quizzy_game/internal/dataTypes"
)

func GetQuestionsWeb(w http.ResponseWriter, r *http.Request) {

	questions := GetQuestions()

	fmt.Printf("got /questions request\n")
	for _, question := range questions {
		io.WriteString(w, question.String())
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

func GetQuestions() []dataTypes.Question {

	type QuestionResponse struct {
		ResponseCode int                  `json:"response_code"`
		Questions    []dataTypes.Question `json:"results"`
	}

	questionRequestUrl := "https://opentdb.com/api.php?amount=10"
	responseBody := getSessionRequest(questionRequestUrl)

	var qr QuestionResponse
	var err = json.Unmarshal(responseBody, &qr)
	if err != nil {
		log.Fatal(err)
	}
	return qr.Questions
}
