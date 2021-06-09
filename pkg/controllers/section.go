package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"log"
)

// GetAll return all tasks found in db or error
func GetSectionsByCourseID(courseID int) (sections []models.Section, err error) {
	sections, err = db.GetSectionsByCourseID(courseID)
	if err != nil {
		return
	}
	if len(sections) == 0 {
		log.Println("Empty data")
	}
	return
}

