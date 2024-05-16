package dataTypes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TODO: Need to have a participant/user type defined
type Quiz struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	QuizStatus   QuizStatus      `json:"quizStatus"`
	Category     Category        `json:"category"`
	Difficulty   Difficulty      `json:"difficulty"`
	Type         QuestionType    `json:"type"`
	Questions    []QuestionTuple `json:"questions"`
	Participants map[string]int  `json:"participants"`
}

func (q Quiz) String() string {
	qJSON, err := json.MarshalIndent(q, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error converting Quiz to JSON: %v", err)
	}
	return string(qJSON)
}

func (q Quiz) RemainingQuestions() int {
	unaskedCounter := 0
	for _, question := range q.Questions {
		if !question.IsNotAsked() {
			unaskedCounter++
		}
	}

	return unaskedCounter
}

func (q Quiz) ParticipantsAsString() string {
	var names []string
	for name := range q.Participants {
		names = append(names, name)
	}
	return strings.Join(names, ", ")
}

func (q Quiz) ScoreBoard() string {
	var scoresList []string
	for name, points := range q.Participants {
		scoresList = append(scoresList, fmt.Sprintf("%s: %d", name, points))
	}
	return strings.Join(scoresList, ", ")

}
