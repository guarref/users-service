package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/guarref/users-service/internal/database"
	"github.com/guarref/users-service/internal/transport/grpc"
	"github.com/guarref/users-service/internal/user"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("failed to automigrate users table: %v", err)
	}
	repo := user.NewUserRepository(database.DB)
	svc := user.NewUserService(repo)

	if err := grpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
}
