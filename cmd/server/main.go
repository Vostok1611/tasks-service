package main

import (
	"log"

	"github.com/Vostok1611/tasks-service/internal/database"
	"github.com/Vostok1611/tasks-service/internal/task"
	transportgrpc "github.com/Vostok1611/tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Репозиторий и сервис задач
	repo := task.NewTaskRepository(db)
	svc := task.NewTaskService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
