package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func ChangePassword(user *models.User) (err error) {

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hash)

	err = db.ChangePassword(user)

	return
}
