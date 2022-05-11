package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (t tserver) CreateTask(ctx context.Context, task *pb.Task) (*pb.TaskID, error) {
	ID, err := t.repo.CreateTask(ctx, models.Task{
		Name:        task.Name,
		Description: task.Description,
	})
	return &pb.TaskID{ID: ID}, err
}
