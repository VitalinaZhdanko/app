package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
)

func GetCourseByLanguageID(languageID int) (course *models.Course, err error) {
	course, err = db.GetCourseByLanguage(languageID)
	if err != nil {
		return
	}
	return
}
func GetCourseIDByUserID(userID int) (courseID int, err error) {
	courseID, err = db.GetCourseIDByUserID(userID)
	if err != nil {
		return
	}
	return
}

