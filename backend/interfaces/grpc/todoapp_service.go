package grpc

import (
	"context"

	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent"
	pb "github.com/szpp-dev-team/hands-on-todo-app/proto-gen/go/todoapp/v1"
	"github.com/szpp-dev-team/hands-on-todo-app/usecases/task"
)

type todoappServiceServer struct {
	taskInteractor *task.Interactor
}

func NewTodoappServiceServer(entClient *ent.Client) pb.TodoappServiceServer {
	return &todoappServiceServer{
		taskInteractor: task.NewInteractor(entClient),
	}
}

func (s *todoappServiceServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	return s.taskInteractor.GetTask(ctx, req)
}

// CreateTask implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) CreateTask(context.Context, *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	panic("unimplemented")
}

// GetTaskList implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) GetTaskList(context.Context, *pb.GetTaskListRequest) (*pb.GetTaskListResponse, error) {
	panic("unimplemented")
}

// SearchByTag implements todoappv1.TodoappServiceServer.
func (*todoappServiceServer) SearchByTag(context.Context, *pb.SearchByTagRequest) (*pb.SearchByTagResponse, error) {
	panic("unimplemented")
}
