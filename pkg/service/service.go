package service

import (
	"context"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/client"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/pb"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/pkg/repository"
)

type Comment interface {
	GetCommentsByPostId(ctx context.Context, request *pb.GetCommentsByPostIdRequest) (*pb.GetCommentsByPostIdResponse, error)
	GetCommentById(ctx context.Context, request *pb.GetCommentByIdRequest) (*pb.GetCommentByIdResponse, error)
	AddComment(ctx context.Context, request *pb.AddCommentRequest) (*pb.AddCommentResponse, error)
	DeleteComment(ctx context.Context, request *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error)
}

type Service struct {
	Comment
	pb.UnsafeCommentServiceServer
}

func NewService(r *repository.Repository, userSvc client.UserServiceClient) *Service {
	return &Service{
		Comment: NewCommentService(r.Comment, userSvc),
	}
}
