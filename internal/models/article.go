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
	createdAt      time.Time
	updatedAt      time.Time
	favorited      bool
	favoritesCount int
	author         Profile
}

type Articles struct {
	articles      []Article
	articlesCount int
}
