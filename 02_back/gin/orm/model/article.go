package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	title string
	content string
	Author User `gorm:"embedded"`
}