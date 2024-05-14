package dataTypes

import (
	"encoding/json"
	"fmt"
)

type Quiz struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Category     string       `json:"category"`
	Difficulty   string       `json:"difficulty"`
	Type         QuestionType `json:"type"`
	Questions    []Question   `json:"questions"`
	Participants []string     `json:"participants"`
}

func (q Quiz) String() string {
	qJSON, err := json.MarshalIndent(q, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error converting Quiz to JSON: %v", err)
	}
	return string(qJSON)
}
