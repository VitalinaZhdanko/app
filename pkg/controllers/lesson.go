package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"log"
)

func GetLessonsBySectionID(sectionID int) (lessons []models.Lesson, err error) {
	lessons, err = db.GetLessonsBySectionID(sectionID)
	if err != nil {
		return
	}
	if len(lessons) == 0 {
		log.Println("Empty data")
	}
	return
}

func GetLessonByID(lessonID int) (lesson *models.Lesson, err error) {
	lesson, err = db.GetLessonByID(lessonID)
	if err != nil {
		return
	}
	return
}