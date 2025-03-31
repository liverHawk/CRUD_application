package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	title       string
	content     string
	unpublished bool
	AuthorID    uint
	Author      User `gorm:"foreignKey:AuthorID"` // article belongs to user
}

func CreateArticle(db *gorm.DB, a *Article) int {
	result := db.Create(a)

	if result.Error != nil {
		return 500
	} else {
		return 200
	}
}

func GetArticle(db *gorm.DB, a *Article, id string) int {
	result := db.First(&a, id)

	if result.Error != nil {
		return 500
	} else {
		return 200
	}
}

func UpdateArticle(db *gorm.DB, a *Article, id string) int {
	result := db.Model(&Article{}).Where("id = ?", id).Updates(a)

	if result.Error != nil {
		return 500
	} else {
		return 200
	}
}

func DeleteArticle(db *gorm.DB, id string) int {
	result := db.Delete(&Article{}, id)

	if result.Error != nil {
		return 500
	} else {
		return 200
	}
}

func ExistArticle(db *gorm.DB, id string) bool {
	var article Article
	result := db.First(&article, id)

	if result.Error != nil {
		return false
	} else {
		return true
	}
}
