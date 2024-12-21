package models

import "time"

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
	ID             int
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []string
	Favorited      bool
	FavoritesCount int
	Author         Profile
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
