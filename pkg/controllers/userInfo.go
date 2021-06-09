package controllers

import (
	"app/diplom/pkg/db"
	"app/diplom/pkg/models"
	"fmt"
)

func GetUserInfoByUser(user string) (userInfo *models.UserInfo, err error) {
	fmt.Println("1")
	userInfo, err = db.GetUserInfoByUserID(user)
	if err != nil {
		return
	}

	fmt.Println(userInfo)
	return
}