package task

import (
	"context"

	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent"
	todoappv1 "github.com/szpp-dev-team/hands-on-todo-app/proto-gen/go/todoapp/v1"
	"golang.org/x/exp/slog"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) GetTask(ctx context.Context, req *todoappv1.GetTaskRequest) (*todoappv1.GetTaskResponse, error) {
	panic("implement here")
}
