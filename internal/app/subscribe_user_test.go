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

func TestTserver_SubscribeUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	mockRepo.CreateUserTaskMock.Return(nil)
	mockRepo.GetTaskMock.Return(&models.Task{}, nil)
	mockRepo.GetUserTaskMock.Return(nil, nil)

	ctx := context.Background()
	svc := NewServer(mockRepo)

	response, err := svc.SubscribeUser(ctx, &pb.SubscribeUserRequest{UserID: 1, TaskID: 1})

	assert.Nil(t, err)
	assert.Equal(t, true, response.IsSubscribed)
}

func TestTserver_SubscribeUserReturnsRepoError(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateUserTaskMock.Return(errors.New("repo error"))
	mockRepo.GetTaskMock.Return(&models.Task{}, nil)
	mockRepo.GetUserTaskMock.Return(nil, nil)

	ctx := context.Background()
	svc := NewServer(mockRepo)

	response, err := svc.SubscribeUser(ctx, &pb.SubscribeUserRequest{UserID: 1, TaskID: 1})

	assert.Nil(t, response)
	assert.Errorf(t, err, "repo error")
}

func TestTserver_SubscribeUserReturnsErrorWhenTaskNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetTaskMock.Return(nil, errors.New("test"))

	svc := NewServer(mockRepo)

	response, err := svc.SubscribeUser(ctx, &pb.SubscribeUserRequest{UserID: 1, TaskID: 1})

	assert.Nil(t, response)
	assert.Errorf(t, err, "task with ID %v not found", 1)
}

func TestTserver_SubscribeUserReturnsErrorWhenUserAlreadySubscribed(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetTaskMock.Return(&models.Task{}, nil)
	mockRepo.GetUserTaskMock.Return([]*models.UserTask{}, nil)

	svc := NewServer(mockRepo)

	response, err := svc.SubscribeUser(ctx, &pb.SubscribeUserRequest{UserID: 1, TaskID: 1})

	assert.Nil(t, response)
	assert.Errorf(t, err, "user already subscribed")
}
