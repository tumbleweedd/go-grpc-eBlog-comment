package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/internal/model"
)

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (commentRepo *CommentRepository) GetCommentsByPostId(postId int) ([]model.Comment, error) {
	var comments []model.Comment

	query := fmt.Sprintf(`select * from %s c where c.post_id=$1`, commentTable)

	err := commentRepo.db.Select(&comments, query, postId)

	return comments, err
}

func (commentRepo *CommentRepository) GetCommentById(commentId, postId int) (model.Comment, error) {
	var comment model.Comment

	query := fmt.Sprintf(`select * from %s c where c.comment_id=$1 and c.post_id=$2`, commentTable)

	err := commentRepo.db.Get(&comment, query, commentId, postId)

	return comment, err
}

func (commentRepo *CommentRepository) AddComment(commentBody string, postId, userId int) error {
	query := fmt.Sprintf(`insert into %s (body, date_creation, post_id, user_id) 
										values($1, current_timestamp, $2, $3)`, commentTable)
	_, err := commentRepo.db.Exec(query, commentBody, postId, userId)

	return err
}

func (commentRepo *CommentRepository) DeleteComment(commentId, postId int) error {
	query := fmt.Sprintf(`delete from %s c where c.comment_id=$1 and c.post_id=$2`, commentTable)

	_, err := commentRepo.db.Exec(query, commentId, postId)

	return err
}
