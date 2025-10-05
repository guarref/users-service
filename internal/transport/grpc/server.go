package grpc

import (
	"fmt"
	"net"

	userpb "github.com/guarref/project-protos/proto/user"
	"github.com/guarref/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.UserService) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("gRPC fail to listen on port 50051: %v", err)
	}

	grpcSrv := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))
	if err := grpcSrv.Serve(listener); err != nil {
		return fmt.Errorf("Error of running server: %v", err)
	}

	// 1. net.Listen на ":50051"
	// 2. grpc.NewServer()
	// 3. userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))
	// 4. grpcSrv.Serve(listener)
	return nil
}
