package grpc

import (
	"context"

	userpb "github.com/guarref/project-protos/proto/user"
	"github.com/guarref/users-service/internal/user"
)

type Handler struct {
	svc *user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	newUser := user.User{Email: req.Email}

	createdUser, err := h.svc.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{User: &userpb.User{Id: createdUser.ID, Email: createdUser.Email}}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {

	findUser, err := h.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.User{Id: findUser.ID, Email: findUser.Email}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return &userpb.UpdateUserResponse{}, nil
}
