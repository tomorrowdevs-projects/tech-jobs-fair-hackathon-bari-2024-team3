package dataTypes

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ParticipantsTuple struct {
	Ref   *User
	Score int
}

// TODO: Need to have a participant/user type defined
type Quiz struct {
	Id            string                       `json:"id"`
	Name          string                       `json:"name"`
	QuizStatus    QuizStatus                   `json:"quizStatus"`
	Category      Category                     `json:"category"`
	Difficulty    Difficulty                   `json:"difficulty"`
	Type          QuestionType                 `json:"type"`
	Questions     map[string]QuestionTriple    `json:"questions"`
	Participants  map[string]ParticipantsTuple `json:"participants"`
	StatusChannel *chan QuizStatus             `json:"-"`
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
		if question.IsNotAsked() {
			unaskedCounter++
		}
	}

	return unaskedCounter
}

func (q Quiz) ParticipantsAsString() string {
	var names []string
	for _, pt := range q.Participants {
		names = append(names, pt.Ref.Name)
	}
	return strings.Join(names, ", ")
}

func (q Quiz) ScoreBoard() string {
	var scoresList []string
	for _, pt := range q.Participants {
		scoresList = append(scoresList, fmt.Sprintf("%s: %d", pt.Ref.Name, pt.Score))
	}
	return strings.Join(scoresList, ", ")

}
