package models

import "time"

type CreateArticle struct {
	title       string
	description string
	body        string
	tagList     []string
}

type UpdateArticle struct {
	title       string
	description string
	body        string
}

type Article struct {
	slug           string
	title          string
	description    string
	body           string
	tagList        []string
	favorited      bool
	favoritesCount int
	author         Profile
	comments       []Comment
	createdAt      time.Time
	updatedAt      time.Time
}
