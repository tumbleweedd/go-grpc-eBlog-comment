package model

import "time"

type Comment struct {
	Id           int       `json:"comment_id" db:"comment_id"`
	Body         string    `json:"body" db:"body"`
	DateCreation time.Time `json:"date_creation" db:"date_creation"`
	PostId       int       `json:"post_id" db:"post_id"`
	UserId       int       `json:"user_id" db:"user_id"`
}

type CommentDTO struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}
