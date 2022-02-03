package handlers

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	pbUsers "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users/api"
)

type Users struct {
	pbUsers.UnimplementedUsersServer
}

func (u *Users) CreateUser(ctx context.Context, request *pbUsers.CreateUserRequest) (*pbUsers.CreateUserResponse, error) {
	if strings.TrimSpace(request.GetEmail()) == "" {
		return nil, errors.New("email should not be empty")
	}

	if strings.TrimSpace(request.GetPassword()) == "" {
		return nil, errors.New("password should not be empty")
	}

	return &pbUsers.CreateUserResponse{Id: 1}, nil
}

func (u *Users) ListUsers(ctx context.Context, request *pbUsers.ListUsersRequest) (*pbUsers.ListUsersResponse, error) {
	return &pbUsers.ListUsersResponse{
		Users: []*pbUsers.User{
			{
				Id:        1,
				Firstname: "Alex",
				Lastname:  "Dummy",
				Email:     "alex@dummy.com",
			},
		},
	}, nil
}

func (u *Users) GetUser(ctx context.Context, request *pbUsers.GetUserRequest) (*pbUsers.User, error) {
	return &pbUsers.User{
		Id:        1,
		Firstname: "Alex",
		Lastname:  "Dummy",
		Email:     "alex@dummy.com",
	}, nil
}
