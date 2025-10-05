package grpc

import (
	"context"
	"errors"

	userpb "github.com/guarref/project-protos/proto/user"
	"github.com/guarref/users-service/internal/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc *user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(_ context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	newUser := user.User{Email: req.Email}

	createdUser, err := h.svc.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{User: &userpb.User{Id: createdUser.ID, Email: createdUser.Email}}, nil
}

func (h *Handler) GetUser(_ context.Context, req *userpb.User) (*userpb.User, error) {

	findUser, err := h.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.User{Id: findUser.ID, Email: findUser.Email}, nil
}

func (h *Handler) UpdateUser(_ context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {

	if req.Email == "" {
		return nil, errors.New("email can't be empty")
	}
	id := req.Id
	newUser := user.User{Email: req.Email}

	updatedUser, err := h.svc.UpdateUserByID(id, newUser)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{User: &userpb.User{Id: updatedUser.ID, Email: updatedUser.Email}}, nil
}

func (h *Handler) DeleteUser(_ context.Context, req *userpb.User) (*emptypb.Empty, error) {

	id := req.Id

	err := h.svc.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) ListUsers(_ context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {

	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	arrUsers := make([]*userpb.User, 0, len(users))

	for _, val := range users {
		oneUser := userpb.User{Id: val.ID, Email: val.Email}
		arrUsers = append(arrUsers, &oneUser)
	}

	return &userpb.ListUsersResponse{Users: arrUsers}, nil
}
