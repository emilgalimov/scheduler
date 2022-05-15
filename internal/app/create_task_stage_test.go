package app

import (
	"context"
	"errors"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"testing"
)

func TestTserver_CreateTaskStage(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateTaskStageMock.Return(1, nil)
	mockRepo.GetTaskMock.Return(models.Task{}, nil)

	svc := NewServer(mockRepo)

	TaskStageID, err := svc.CreateTaskStage(ctx, &pb.CreateTaskStageRequest{
		TaskStage: &pb.TaskStage{Name: "testName", Description: "testDescription", MinutesFromStart: 1, DurationMinutes: 1},
		TaskID:    1,
	})

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), TaskStageID.ID)

}

func TestTserver_ReturnsErrorWhenTaskNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetTaskMock.Return(models.Task{}, errors.New("test"))

	svc := NewServer(mockRepo)

	TaskStageID, err := svc.CreateTaskStage(ctx, &pb.CreateTaskStageRequest{
		TaskStage: &pb.TaskStage{Name: "testName", Description: "testDescription", MinutesFromStart: 1, DurationMinutes: 1},
		TaskID:    1,
	})

	assert.Nil(t, TaskStageID)
	assert.Errorf(t, err, "task with ID %v not found", 1)
}
