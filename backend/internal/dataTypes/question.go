package dataTypes

import (
	"encoding/json"
	"fmt"
)

type Question struct {
	Type          QuestionType `json:"type"`
	Difficulty    string       `json:"difficulty"`
	Category      string       `json:"category"`
	Question      string       `json:"question"`
	CorrectAnswer string       `json:"correct_answer"`
	WrongAnswer   []string     `json:"incorrect_answers"`
}

func (q Question) String() string {
	// Convert the Question struct to a JSON string for better visibility
	qJSON, err := json.MarshalIndent(q, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error converting Question to JSON: %v", err)
	}
	return string(qJSON)
}
