package main

import (
	"log"

	"github.com/guarref/users-service/internal/database"
	"github.com/guarref/users-service/internal/user"
)

func main() {
	database.InitDB()
	repo := user.NewUserRepository(database.DB)
	svc := user.NewUserService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}