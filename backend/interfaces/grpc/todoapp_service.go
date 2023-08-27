package grpc

import (
	"context"

	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent"
	todoappv1 "github.com/szpp-dev-team/hands-on-todo-app/proto-gen/go/todoapp/v1"
	"github.com/szpp-dev-team/hands-on-todo-app/usecases/task"
	"golang.org/x/exp/slog"
)

type todoappServiceServer struct {
	taskInteractor *task.Interactor
}

func NewTodoappServiceServer(entClient *ent.Client, logger *slog.Logger) todoappv1.TodoappServiceServer {
	return &todoappServiceServer{
		taskInteractor: task.NewInteractor(entClient, logger),
	}
}

func (i *todoappServiceServer) CreateTask(ctx context.Context, req *todoappv1.CreateTaskRequest) (*todoappv1.CreateTaskResponse, error) {
	return i.taskInteractor.CreateTask(ctx, req)
}

func (i *todoappServiceServer) GetTask(ctx context.Context, req *todoappv1.GetTaskRequest) (*todoappv1.GetTaskResponse, error) {
	return i.taskInteractor.GetTask(ctx, req)
}

func (i *todoappServiceServer) GetTaskList(ctx context.Context, req *todoappv1.GetTaskListRequest) (*todoappv1.GetTaskListResponse, error) {
	return i.taskInteractor.GetTaskList(ctx, req)
}

func (i *todoappServiceServer) SearchByTag(ctx context.Context, req *todoappv1.SearchByTagRequest) (*todoappv1.SearchByTagResponse, error) {
	return i.taskInteractor.SearchByTag(ctx, req)
}
