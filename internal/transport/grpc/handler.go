package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/Vostok1611/project-protos/proto/tasks"
	userpb "github.com/Vostok1611/project-protos/proto/users"
	"github.com/Vostok1611/tasks-service/internal/task"
)

type Handler struct {
	svc        task.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.TaskService, uc userpb.UserServiceClient) *Handler {
	return &Handler{
		svc:        svc,
		userClient: uc,
	}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	//Проверка пользователя
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %s not found: %w", req.UserId, err)
	}
	//Внутренняя логика:
	t, err := h.svc.CreateTask(req.Title, req.IsDone, req.UserId)
	if err != nil {
		return nil, err
	}
	// Ответ
	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		},
	}, nil
}
func (h *Handler) GetTaskById(ctx context.Context, req *taskpb.GetTaskByIdRequest) (*taskpb.GetTaskByIdResponse, error) {
	// 1. Получаем задачу из БД по ID
	t, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &taskpb.GetTaskByIdResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		},
	}, nil
}

func (h *Handler) GetAllTasks(ctx context.Context, req *taskpb.GetAllTasksRequest) (*taskpb.GetAllTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}

	pbTasks := make([]*taskpb.Task, 0, len(tasks))
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		})
	}
	return &taskpb.GetAllTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) GetTasksByUser(ctx context.Context, req *taskpb.GetTasksByUserRequest) (*taskpb.GetTasksByUserResponse, error) {
	if req.UserId == "" {
		return nil, fmt.Errorf("user_id cannot be empty")
	}

	tasks, err := h.svc.GetTasksByUserID(req.UserId)
	if err != nil {
		return nil, err
	}

	pbTasks := make([]*taskpb.Task, 0, len(tasks))
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     t.ID,
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: t.UserID,
		})
	}

	return &taskpb.GetTasksByUserResponse{Tasks: pbTasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	// 1. Распаковываем optional-поля
	title := ""
	if req.Title != nil {
		title = *req.Title
	}

	isDone := ""
	if req.IsDone != nil {
		isDone = *req.IsDone
	}

	// 2. Вызываем сервис для обновления
	// userId не передаём, так как его нет в запросе
	updatedTask, err := h.svc.UpdateTask(req.Id, title, isDone, "")
	if err != nil {
		return nil, err
	}

	// 3. Преобразуем ответ
	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     updatedTask.ID,
			Title:  updatedTask.Task,
			IsDone: updatedTask.IsDone,
			UserId: updatedTask.UserID,
		},
	}, nil

}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	// 1. Вызываем сервис для удаления задачи по ID
	err := h.svc.DeleteTask(req.Id)
	if err != nil {
		return nil, err
	}
	// 2. Возвращаем подтверждение
	return &taskpb.DeleteTaskResponse{
		Success: true,
	}, nil
}
