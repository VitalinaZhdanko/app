package models

type Test struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	LessonID int `json:"lesson_id"`
}
