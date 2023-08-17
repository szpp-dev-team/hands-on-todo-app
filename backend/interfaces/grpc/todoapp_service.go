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

// CreateTask implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) CreateTask(context.Context, *todoappv1.CreateTaskRequest) (*todoappv1.CreateTaskResponse, error) {
	panic("unimplemented")
}

// GetTask implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) GetTask(context.Context, *todoappv1.GetTaskRequest) (*todoappv1.GetTaskResponse, error) {
	panic("unimplemented")
}

// GetTaskList implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) GetTaskList(context.Context, *todoappv1.GetTaskListRequest) (*todoappv1.GetTaskListResponse, error) {
	panic("unimplemented")
}

// SearchByTag implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) SearchByTag(context.Context, *todoappv1.SearchByTagRequest) (*todoappv1.SearchByTagResponse, error) {
	panic("unimplemented")
}
