package helper

import (
	"firstApp/conf"
	"firstApp/models"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSal(pwd []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)

}

func ComparePasswords(user *models.User, plainPwd []byte) bool {
	findUser := conf.DB.Where("username = ?", user.Username).First(&user)
	if findUser.Error != nil {
		return false
	}
	byteHash := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil

}
