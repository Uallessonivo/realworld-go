package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	Favorite      Article
	FavoriteId    uint
	FavoritedBy   ArticleUser
	FavoritedById uint
}
