package task

import (
	"time"

	"github.com/samber/lo"
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
		Task: toPbTask(task),
	}, nil
}

func (i *Interactor) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task, err := i.entClient.Task.Create().
		SetName(req.Name).
		SetNillableDescription(req.Description).
		SetNillableDeadline(toTime(req.Deadline)).
		SetNillableCompletedAt(toTime(req.CompletedAt)).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateTaskResponse{
		Task: toPbTask(task),
	}, nil
}

func (i *Interactor) GetTaskList(ctx context.Context, req *pb.GetTaskListRequest) (*pb.GetTaskListResponse, error) {
	tasks, err := i.entClient.Task.Query().All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetTaskListResponse{
		Tasks: lo.Map(tasks, func(task *ent.Task, _ int) *pb.Task {
			return toPbTask(task)
		}),
	}, nil
}

// helper
func toTimestamppb(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func toTime(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

func toPbTask(t *ent.Task) *pb.Task {
	return &pb.Task{
		Id:          int32(t.ID),
		Name:        t.Name,
		Description: t.Description,
		Deadline:    toTimestamppb(t.Deadline),
		CompletedAt: toTimestamppb(t.CompletedAt),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   toTimestamppb(t.UpdatedAt),
	}
}
