package models

type Lesson struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	SectionID int `json:"sectionID"`
	Youtube string `json:"youtube"`
}
