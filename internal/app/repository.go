package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
)

type Repository interface {
	CreateUser(context.Context) (uint64, error)
	CreateTask(context.Context, models.Task) (uint64, error)
	GetTask(ctx context.Context, ID uint64) (*models.Task, error)
	CreateTaskStage(ctx context.Context, taskStage models.TaskStage, taskID uint64) (uint64, error)
}
