package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"log"
)

func GetTasksByLessonID(lessonID int) (tasks []models.Task, err error) {
	tasks, err = db.GetTasksByLessonID(lessonID)
	if err != nil {
		return
	}
	if len(tasks) == 0 {
		log.Println("Empty data")
	}
	return
}

func GetAllTask() (tasks []models.Task, err error) {
	tasks, err = db.GetAllTask()
	if err != nil {
		return
	}
	if len(tasks) == 0 {
		log.Println("Empthy data")
	}
	return
}