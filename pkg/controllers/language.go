package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"log"
)

// GetAll return all tasks found in db or error
func GetAllLanguage() (tasks []models.Language, err error) {
	tasks, err = db.GetAllLanguage()
	if err != nil {
		return
	}
	if len(tasks) == 0 {
		log.Println("Empthy data")
	}
	return
}