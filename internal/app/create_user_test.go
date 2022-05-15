package app

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"testing"
)

func TestTserver_CreateUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateUserMock.Return(1, nil)
	ctx := context.Background()

	svc := NewServer(mockRepo)

	User, err := svc.CreateUser(ctx, &pb.CreateUserRequest{})

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), User.ID)
}
