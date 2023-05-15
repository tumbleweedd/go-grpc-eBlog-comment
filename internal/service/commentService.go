package service

import (
	"context"
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/client"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/internal/model"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/internal/repository"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/pkg/pb"
	"net/http"
)

type CommentService struct {
	commentRepo repository.Comment
	userSvc     client.UserServiceClient
}

func (commentService *CommentService) GetCommentsByPostId(ctx context.Context, request *pb2.GetCommentsByPostIdRequest) (*pb2.GetCommentsByPostIdResponse, error) {
	allCommentsByPostId, err := commentService.commentRepo.GetCommentsByPostId(int(request.GetPostId()))
	if err != nil {
		return &pb2.GetCommentsByPostIdResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
	}
	allUsersData, err := commentService.userSvc.GetUserList()
	if err != nil {
		return &pb2.GetCommentsByPostIdResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
	}

	commentMap := make(map[string]*pb2.CommentBody)

	for _, user := range allUsersData.GetData() {
		setOfCommentsForThisUser, err := commentService.getSetOfCommentsForUser(allCommentsByPostId, user)
		if err != nil {
			return &pb2.GetCommentsByPostIdResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		commentMap[user.GetUsername()] = &pb2.CommentBody{
			Body: setOfCommentsForThisUser,
		}
	}

	return &pb2.GetCommentsByPostIdResponse{Status: http.StatusOK, Comments: commentMap}, nil
}

func (commentService *CommentService) getSetOfCommentsForUser(allCommentsByPostId []model.Comment, user *pb2.UserData) ([]string, error) {
	var setOfCommentsForThisUser []string
	userId, err := commentService.userSvc.GetUserIdByUsername(user.GetLastname())
	if err != nil {
		return nil, err
	}

	for _, comment := range allCommentsByPostId {
		if comment.UserId == int(userId.GetUserId()) {
			setOfCommentsForThisUser = append(setOfCommentsForThisUser, comment.Body)
		}
	}

	return setOfCommentsForThisUser, nil
}

func (commentService *CommentService) GetCommentById(ctx context.Context, request *pb2.GetCommentByIdRequest) (*pb2.GetCommentByIdResponse, error) {
	postId := request.GetPostId()
	commentId := request.GetCommentId()

	comment, err := commentService.commentRepo.GetCommentById(int(commentId), int(postId))
	if err != nil {
		return &pb2.GetCommentByIdResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
	}

	user, err := commentService.userSvc.GetUserById(comment.UserId)
	if err != nil {
		return &pb2.GetCommentByIdResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
	}

	return &pb2.GetCommentByIdResponse{
		Status:   http.StatusOK,
		Username: user.Data.GetUsername(),
		Body:     comment.Body,
	}, nil
}

func (commentService *CommentService) AddComment(ctx context.Context, request *pb2.AddCommentRequest) (*pb2.AddCommentResponse, error) {
	commentBody := request.GetBody()
	postId := request.GetPostId()
	userId := request.GetUserId()

	loggedUserProfile, err := commentService.userSvc.GetUserById(int(userId))

	if err != nil {
		return &pb2.AddCommentResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	err = commentService.commentRepo.AddComment(commentBody, int(postId), int(userId))
	if err != nil {
		return &pb2.AddCommentResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println(loggedUserProfile.Data.GetUsername(), ":", commentBody)

	return &pb2.AddCommentResponse{
		Status:   http.StatusOK,
		Username: loggedUserProfile.Data.GetUsername(),
		Body:     commentBody,
	}, nil
}

func (commentService *CommentService) DeleteComment(ctx context.Context, request *pb2.DeleteCommentRequest) (*pb2.DeleteCommentResponse, error) {
	postId := request.GetPostId()
	commentId := request.GetCommentId()

	err := commentService.commentRepo.DeleteComment(int(commentId), int(postId))
	if err != nil {
		return &pb2.DeleteCommentResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
	}

	return &pb2.DeleteCommentResponse{
		Status: http.StatusOK,
	}, nil
}

func NewCommentService(commentRepo repository.Comment, userSvc client.UserServiceClient) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		userSvc:     userSvc,
	}
}
