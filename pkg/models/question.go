package models

type Question struct {
	ID       int    `json:"id"`
	Description string `json:"description"`
	Cost int `json:"cost"`
	TestID int `json:"test_id"`
}
