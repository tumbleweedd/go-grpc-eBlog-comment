package client

import (
	"context"
	"fmt"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb2.UserServiceClient
}

func InitUserServiceClient(url string) UserServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := UserServiceClient{
		Client: pb2.NewUserServiceClient(cc),
	}

	return c
}

func (c *UserServiceClient) GetUserList() (*pb2.GetUserListResponse, error) {
	req := &pb2.GetUserListRequest{}

	return c.Client.GetUserList(context.Background(), req)
}

func (c *UserServiceClient) GetUserById(userId int) (*pb2.GetLoggedUserProfileResponse, error) {
	req := &pb2.GetLoggedUserProfileRequest{
		UserId: int64(userId),
	}

	return c.Client.GetLoggedUserProfile(context.Background(), req)
}

func (c *UserServiceClient) GetUserIdByUsername(username string) (*pb2.GetUserIdByUsernameResponse, error) {
	req := &pb2.GetUserIdByUsernameRequest{
		Username: username,
	}

	return c.Client.GetUserIdByUsername(context.Background(), req)
}
