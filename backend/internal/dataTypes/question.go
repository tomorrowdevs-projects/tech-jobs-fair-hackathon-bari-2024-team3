package dataTypes

import (
	"encoding/json"
	"fmt"
)

type QuestionTuple struct {
	Ref           *Question
	IsAskedStatus bool
}

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

func (q Question) ToTuple() QuestionTuple {
	return QuestionTuple{&q, false}
}

func (qt QuestionTuple) IsNotAsked() bool {
	return !qt.IsAskedStatus
}
