package models

type Answer struct {
	ID int `json:"id"`
	Description string `json:"description"`
	IsTrue bool `json:"is_true"`
	QuestionID int `json:"question_id"`
}

