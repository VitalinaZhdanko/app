package models


type QuestionAndAnswer struct {
	TestID int `json:"test_id"`
	QuestionID       int    `json:"question_id"`
	QuestionDescription string `json:"question_description"`
	Cost int `json:"cost"`
	Answer1 string `json:"answer1"`
	IsTrue1 bool `json:"is_true1"`
	Answer2 string `json:"answer2"`
	IsTrue2 bool `json:"is_true2"`
	Answer3 string `json:"answer3"`
	IsTrue3 bool `json:"is_true3"`
}

