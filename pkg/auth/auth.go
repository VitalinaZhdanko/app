package auth

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"app/diplom/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration


func RegisterUser(user *models.User) (err error) {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	user.Password = string(hash)
	_, err = db.AddUser(user)
	return
}

// LoginUser handles user login
func LoginUser(user *models.User) (res models.Login, err error) {
	userFromDB, err := db.GetUserByUsername(user.Username)
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword(
		[]byte(userFromDB.Password),
		[]byte(user.Password),
	); err != nil {
		return
	}
	userToken, err := token.Create(userFromDB)
	if err != nil {
		return
	}

	res.Authorization = userToken

	res.Id = userFromDB.ID
	return
}
