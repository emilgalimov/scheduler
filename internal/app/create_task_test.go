package app

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"testing"
)

func TestTserver_CreateTask(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateTaskMock.Return(1, nil)
	ctx := context.Background()
	svc := NewServer(mockRepo)

	TaskID, err := svc.CreateTask(ctx, &pb.Task{
		Name:        "test name",
		Description: "test description",
	})

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), TaskID.ID)
}
