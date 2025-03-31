package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
}

func CreateUser(db *gorm.DB, u *User) int {
	result := db.Create(u)

	if result.Error != nil {
		return 500
	} else {
		return 200
	}
}

func GetUser(db *gorm.DB, u *User, id string) int {
	result := db.First(&u, id)

	if result.Error != nil {
		return 500
	} else {
		return 200
	}
}
