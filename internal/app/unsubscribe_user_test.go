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

func TestTserver_UnsubscribeUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	ctx := context.Background()

	userID := uint64(1)
	taskID := uint64(1)
	mockRepo.GetTaskMock.Expect(ctx, taskID).Return(models.Task{}, nil)
	mockRepo.GetUserTaskMock.Expect(ctx, userID, taskID).Return([]models.UserTask{}, nil)
	mockRepo.DeleteUserTaskMock.Expect(ctx, userID, taskID).Return(nil)

	svc := NewServer(mockRepo)
	response, err := svc.UnsubscribeUser(ctx, &pb.UnsubscribeUserRequest{UserID: userID, TaskID: taskID})

	assert.Nil(t, err)
	assert.Equal(t, true, response.IsUnsubscribed)
}

func TestTserver_UnsubscribeUserReturnRepoError(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	ctx := context.Background()

	userID := uint64(1)
	taskID := uint64(1)
	mockRepo.GetTaskMock.Expect(ctx, taskID).Return(models.Task{}, nil)
	mockRepo.GetUserTaskMock.Expect(ctx, userID, taskID).Return([]models.UserTask{}, nil)
	mockRepo.DeleteUserTaskMock.Expect(ctx, userID, taskID).Return(errors.New("repo error"))

	svc := NewServer(mockRepo)
	response, err := svc.UnsubscribeUser(ctx, &pb.UnsubscribeUserRequest{UserID: userID, TaskID: taskID})

	assert.Nil(t, response)
	assert.Errorf(t, err, "repo error")
}

func TestTserver_UnsubscribeUserReturnsErrorWhenTaskNotExists(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	ctx := context.Background()

	userID := uint64(1)
	taskID := uint64(1)
	mockRepo.GetTaskMock.Expect(ctx, taskID).Return(models.Task{}, errors.New("not found"))

	svc := NewServer(mockRepo)
	response, err := svc.UnsubscribeUser(ctx, &pb.UnsubscribeUserRequest{UserID: userID, TaskID: taskID})

	assert.Nil(t, response)
	assert.Errorf(t, err, "task with ID %v not found", 1)
}

func TestTserver_UnsubscribeUserReturnErrorWhenUserNotSubscribed(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	ctx := context.Background()

	userID := uint64(1)
	taskID := uint64(1)
	mockRepo.GetTaskMock.Expect(ctx, taskID).Return(models.Task{}, nil)
	mockRepo.GetUserTaskMock.Expect(ctx, userID, taskID).Return(nil, errors.New("not found"))

	svc := NewServer(mockRepo)
	response, err := svc.UnsubscribeUser(ctx, &pb.UnsubscribeUserRequest{UserID: userID, TaskID: taskID})

	assert.Nil(t, response)
	assert.Errorf(t, err, "user not subscribed")
}
