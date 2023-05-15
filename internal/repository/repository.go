package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/internal/model"
)

type Comment interface {
	GetCommentsByPostId(postId int) ([]model.Comment, error)
	GetCommentById(commentId, postId int) (model.Comment, error)
	AddComment(commentBody string, postId, userId int) error
	DeleteComment(commentId, postId int) error
}

type Repository struct {
	Comment Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Comment: NewCommentRepository(db),
	}
}
