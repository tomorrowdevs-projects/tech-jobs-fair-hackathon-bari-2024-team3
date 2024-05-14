package dataTypes

type QuizStatus string

const (
	Initializing = "initializing"
	Holding      = "holding"
	QuizTime     = "quizTime"
	Evaluation   = "evaluation"
	Ended        = "ended"
)
