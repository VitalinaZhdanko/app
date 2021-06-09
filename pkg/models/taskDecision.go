package models

import (
	"encoding/json"
	"io"
)

type TaskDecision struct {
	UserID int `json:"userID"`
	TaskDecision string `json:"taskDecision"`
}

// PopulateFromRequest add fields from bytes to struct
func (u *TaskDecision) PopulateFromRequest(requestBody io.Reader) (err error) {
	decoder := json.NewDecoder(requestBody)

	err = decoder.Decode(u)
	return
}
