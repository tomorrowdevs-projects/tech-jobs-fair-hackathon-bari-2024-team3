package dataTypes

type QuizStatus string

const (
	StatusInitialized     = "initialized"
	StatusStart           = "start"
	StatusQuizTime        = "quizTime"
	StatusQuizTimeEnded   = "quizTimeEnded"
	StatusEvaluation      = "evaluation"
	StatusEvaluationEnded = "evaluationEnded"
	StatusEnded           = "ended"
)
