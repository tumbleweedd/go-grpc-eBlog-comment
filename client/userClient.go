package client

import (
	"context"
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitUserServiceClient(url string) UserServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := UserServiceClient{
		Client: pb.NewUserServiceClient(cc),
	}

	return c
}

func (c *UserServiceClient) GetUserList() (*pb.GetUserListResponse, error) {
	req := &pb.GetUserListRequest{}

	return c.Client.GetUserList(context.Background(), req)
}

func (c *UserServiceClient) GetUserById(userId int) (*pb.GetLoggedUserProfileResponse, error) {
	req := &pb.GetLoggedUserProfileRequest{
		UserId: int64(userId),
	}

	return c.Client.GetLoggedUserProfile(context.Background(), req)
}

func (c *UserServiceClient) GetUserIdByUsername(username string) (*pb.GetUserIdByUsernameResponse, error) {
	req := &pb.GetUserIdByUsernameRequest{
		Username: username,
	}

	return c.Client.GetUserIdByUsername(context.Background(), req)
}
