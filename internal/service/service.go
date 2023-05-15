package service

import (
	"context"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/client"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/internal/repository"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/pkg/pb"
)

type Comment interface {
	GetCommentsByPostId(ctx context.Context, request *pb2.GetCommentsByPostIdRequest) (*pb2.GetCommentsByPostIdResponse, error)
	GetCommentById(ctx context.Context, request *pb2.GetCommentByIdRequest) (*pb2.GetCommentByIdResponse, error)
	AddComment(ctx context.Context, request *pb2.AddCommentRequest) (*pb2.AddCommentResponse, error)
	DeleteComment(ctx context.Context, request *pb2.DeleteCommentRequest) (*pb2.DeleteCommentResponse, error)
}

type Service struct {
	Comment
	pb2.UnsafeCommentServiceServer
}

func NewService(r *repository.Repository, userSvc client.UserServiceClient) *Service {
	return &Service{
		Comment: NewCommentService(r.Comment, userSvc),
	}
}
