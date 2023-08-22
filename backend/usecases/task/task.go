package task

import (
	"time"

	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent"
	enttask "github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent/task"
	pb "github.com/szpp-dev-team/hands-on-todo-app/proto-gen/go/todoapp/v1"
	"golang.org/x/exp/slog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task, err := i.entClient.Task.Query().Where(enttask.ID(int(req.Id))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the task is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetTaskResponse{
		Task: &pb.Task{
			Id:          int32(task.ID),
			Name:        task.Name,
			Description: task.Description,
			Deadline:    toTimestamppb(task.Deadline),
			CompletedAt: toTimestamppb(task.CompletedAt),
			CreatedAt:   timestamppb.New(task.CreatedAt),
			UpdatedAt:   toTimestamppb(task.UpdatedAt),
		},
	}, nil
}

// helper
func toTimestamppb(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
