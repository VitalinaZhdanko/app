package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
)

func GetTestByLessonID(lessonID int) (test *models.Test, err error) {
	test, err = db.GetTestByLessonID(lessonID)
	if err != nil {
		return
	}

	return
}