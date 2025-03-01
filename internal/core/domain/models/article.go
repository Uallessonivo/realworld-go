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
	Articles []Article
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
