package grpc

import (
	"log"
	"net"

	taskpb "github.com/Vostok1611/project-protos/proto/tasks"
	userpb "github.com/Vostok1611/project-protos/proto/users"
	"github.com/Vostok1611/tasks-service/internal/task"
	"google.golang.org/grpc"
)

// RunGRPC запускает gRPC-сервер для задач
func RunGRPC(svc task.TaskService, us userpb.UserServiceClient) error {
	// 1. Слушаем порт 50052 (Tasks-сервис)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}
	// 2. Создаём gRPC-сервер
	s := grpc.NewServer()
	// 3. Создаём обработчик с нашим сервисом и клиентом пользователей
	handler := NewHandler(svc, us)
	// 4. Регистрируем обработчик в сервере
	taskpb.RegisterTaskServiceServer(s, handler)

	log.Println("Tasks gRPC server listening on:50052")
	// 5. Запускаем сервер (блокируется до остановки)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
