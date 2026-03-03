package grpc

import (
	taskpb "github.com/Vostok1611/project-protos/proto/tasks"
	userpb "github.com/Vostok1611/project-protos/proto/users"
	"github.com/Vostok1611/tasks-service/internal/task"
)

type Handler struct {
	svc        task.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.TaskService)
