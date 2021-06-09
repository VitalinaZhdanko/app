package models

type Task struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Score int `json:"score"`
	LessonID int `json:"lessonID"`
}

