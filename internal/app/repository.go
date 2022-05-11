package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
)

type Repository interface {
	CreateUser(context.Context) (uint64, error)
	CreateTask(context.Context, models.Task) (uint64, error)
}
