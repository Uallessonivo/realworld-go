package models

import "time"

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
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	Author    Profile
}
