package models

import "gorm.io/gorm"

type ArticleUser struct {
	gorm.Model
	User      User
	UserId    uint
	Articles  []Article
	Favorited []Favorite
}
