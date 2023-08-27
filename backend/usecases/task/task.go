package task

import (
	"time"

	"github.com/samber/lo"
	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent"
	enttag "github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent/tag"
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
	task, err := i.entClient.Task.Query().WithTags().Where(enttask.ID(int(req.Id))).Only(ctx)
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
	tagIDs := []int{}

	// タグの ID 取得・作成
	for _, t := range req.Tags {
		tag, err := i.entClient.Tag.Query().Where(enttag.NameEQ(t)).Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else {
			// タグが存在していたら ID を追加して continue
			tagIDs = append(tagIDs, tag.ID)
			continue
		}
		// タグが存在しないので作成して ID を追加
		tag, err = i.entClient.Tag.Create().SetName(t).Save(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		tagIDs = append(tagIDs, tag.ID)
	}

	task, err := i.entClient.Task.Create().
		SetName(req.Name).
		SetNillableDescription(req.Description).
		SetNillableDeadline(toTime(req.Deadline)).
		SetNillableCompletedAt(toTime(req.CompletedAt)).
		SetCreatedAt(time.Now()).
		AddTagIDs(tagIDs...).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateTaskResponse{
		Task: toPbTask(task),
	}, nil
}

func (i *Interactor) GetTaskList(ctx context.Context, req *pb.GetTaskListRequest) (*pb.GetTaskListResponse, error) {
	tasks, err := i.entClient.Task.Query().WithTags().All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetTaskListResponse{
		Tasks: lo.Map(tasks, func(task *ent.Task, _ int) *pb.Task {
			return toPbTask(task)
		}),
	}, nil
}

func (i *Interactor) SearchByTag(ctx context.Context, req *pb.SearchByTagRequest) (*pb.SearchByTagResponse, error) {
	tags, err := i.entClient.Tag.Query().WithTasks().All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	tasks := make([]*pb.Task, 0)
	for _, tag := range tags {
		for _, task := range tag.Edges.Tasks {
			tasks = append(tasks, toPbTask(task))
		}
	}
	return &pb.SearchByTagResponse{
		Tasks: tasks,
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
		Tags: lo.Map(t.Edges.Tags, func(tag *ent.Tag, _ int) string {
			return tag.Name
		}),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: toTimestamppb(t.UpdatedAt),
	}
}
