package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

func (t tserver) CreateTaskStage(ctx context.Context, taskStageRequest *pb.CreateTaskStageRequest) (*pb.TaskStageID, error) {

	if _, err := t.repo.GetTask(ctx, taskStageRequest.TaskID); err != nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %v not found", taskStageRequest.TaskID)
	}

	taskStage := taskStageRequest.TaskStage

	mFromStart, _ := time.ParseDuration(strconv.FormatUint(taskStage.MinutesFromStart, 10) + "m")

	durationM, _ := time.ParseDuration(strconv.FormatUint(taskStage.DurationMinutes, 10) + "m")

	ID, err := t.repo.CreateTaskStage(ctx, models.TaskStage{
		Name:             taskStage.Name,
		Description:      taskStage.Description,
		MinutesFromStart: mFromStart,
		DurationMinutes:  durationM,
	}, taskStageRequest.TaskID)

	return &pb.TaskStageID{ID: ID}, err
}
