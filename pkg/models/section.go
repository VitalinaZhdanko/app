package models

type Section struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CourseID int `json:"courseID"`
}

