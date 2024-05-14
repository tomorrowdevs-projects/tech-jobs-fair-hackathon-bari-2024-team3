package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"quizzy_game/internal/dataTypes"
)

type QuestionResponse struct {
	ResponseCode int                  `json:"response_code"`
	Questions    []dataTypes.Question `json:"results"`
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {

	questionRequestUrl := "https://opentdb.com/api.php?amount=10"
	responseBody := getSessionRequest(questionRequestUrl)

	var qr QuestionResponse
	var err = json.Unmarshal(responseBody, &qr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got /questions request\n")
	for i := 0; i < len(qr.Questions); i++ {
		io.WriteString(w, qr.Questions[i].String())
	}
}

type CategoryResponse struct {
	Categories []dataTypes.Category `json:"trivia_categories"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	questionRequestUrl := "https://opentdb.com/api_category.php"
	responseBody := getRequest(questionRequestUrl)

	var cr CategoryResponse
	var err = json.Unmarshal(responseBody, &cr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got /categories request with %d categories\n", len(cr.Categories))
	for i := 0; i < len(cr.Categories); i++ {
		io.WriteString(w, cr.Categories[i].String())
	}
}
