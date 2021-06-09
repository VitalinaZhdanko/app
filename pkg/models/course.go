package models

type Course struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	LanguageID int `json:"language_id"`
}

