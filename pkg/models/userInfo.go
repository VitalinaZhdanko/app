package models

type UserInfo struct {
	ID       int    `json:"id"`
	Login string `json:"login"`
	FIO string `json:"fio"`
	Course string `json:"course"`
	Score int `json:"score"`
}
