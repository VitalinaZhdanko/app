package controllers

import (
	"app/diplom/pkg/db"
)

func WriteOnCourse(userID int, courseID int) (message string, err error) {
	message, err = db.WriteOnCourse(userID, courseID)
	if err != nil {
		return
	}
	return
}

func ChangeCourse(userID int, courseID int) (message string, err error) {
	message, err = db.ChangeCourse(userID, courseID)
	if err != nil {
		return
	}
	return
}