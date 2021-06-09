package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"log"
)

func GetQuestionAndAnswerByTestID(testID int) (questionAndAnswer []models.QuestionAndAnswer, err error) {
	questionAndAnswer, err = db.GetQuestionAndAnswerByTestID(testID)
	if err != nil {
		return
	}
	if len(questionAndAnswer) == 0 {
		log.Println("Empty data")
	}
	return
}
