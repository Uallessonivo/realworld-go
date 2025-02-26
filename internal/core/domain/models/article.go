package models

import (
	"gorm.io/gorm"
)

type ArticleRequest struct {
	Title       string
	Description string
	Body        string
	TagList     []string
}

type ArticleResponse struct {
	Article Article
}

type Article struct {
	gorm.Model
	Slug        string
	Title       string
	Description string
	Body        string
	Author      ArticleUser
	AuthorId    uint
	Tags        []Tag
	Comments    []Comment
}

type ArticleUser struct {
	gorm.Model
	User      User
	UserId    uint
	Articles  []Article
	Favorited []Favorite
}
