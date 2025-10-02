package main

func main() {
  database.InitDB()
  repo := user.NewRepository(database.DB)
  svc  := user.NewService(repo)

  if err := transportgrpc.RunGRPC(svc); err != nil {
    log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
  }
}