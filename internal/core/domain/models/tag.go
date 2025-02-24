package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Tag      string
	Articles []Article
}
