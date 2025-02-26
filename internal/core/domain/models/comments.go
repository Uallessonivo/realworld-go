package models

import (
	"gorm.io/gorm"
)

type CommentRequest struct {
	Body string
}

type CommentResponse struct {
	Comment Comment
}

type CommentsResponse struct {
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Article   Article
	ArticleId uint
	Author    ArticleUser
	AuthorId  uint
	Body      string
}
