package app

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestTserver_GetUserTasksByTime(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	ctx := context.Background()
	userID := uint64(1)

	startTime, _ := time.Parse(time.RFC3339, "2022-05-10T10:00:00Z")
	mockRepo.GetAllUserTasksMock.Expect(ctx, userID).Return([]*models.UserTask{
		{
			Task: &models.Task{
				ID:          1,
				Name:        "name",
				Description: "description",
				Stages: []*models.TaskStage{
					{
						ID:               1,
						Name:             "first",
						Description:      "first desc",
						MinutesFromStart: 0 * time.Minute,
						DurationMinutes:  2 * time.Minute,
					},
					{
						ID:               2,
						Name:             "second",
						Description:      "second desc",
						MinutesFromStart: 2 * time.Minute,
						DurationMinutes:  4 * time.Minute,
					},
					{
						ID:               3,
						Name:             "third",
						Description:      "third desc",
						MinutesFromStart: 4 * time.Minute,
						DurationMinutes:  6 * time.Minute,
					},
				},
			},
			User: &models.User{
				ID: 1,
			},
			StartTime: startTime,
		},
	}, nil)

	svc := NewServer(mockRepo)

	timeFrom, _ := time.Parse(time.RFC3339, "2022-05-10T10:02:00Z")
	timeTo, _ := time.Parse(time.RFC3339, "2022-05-10T10:03:00Z")

	tasks, err := svc.GetUserTasksByTime(ctx, &pb.GetUserTasksByTimeRequest{
		UserId:   userID,
		TimeFrom: timestamppb.New(timeFrom),
		TimeTo:   timestamppb.New(timeTo),
	})

	assert.Nil(t, err)
	assert.Equal(t, tasks, &pb.TasksList{
		Tasks: []*pb.Task{
			{
				ID:          1,
				Name:        "name",
				Description: "description",
				TaskStages: []*pb.TaskStage{
					{
						ID:               2,
						Name:             "second",
						Description:      "second desc",
						MinutesFromStart: 2,
						DurationMinutes:  4,
					},
				},
			},
		},
	})
}

func TestTserver_GetUserTasksByTimeReturnsEmpty(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	mockRepo := NewRepositoryMock(mc)

	ctx := context.Background()
	userID := uint64(1)

	mockRepo.GetAllUserTasksMock.Expect(ctx, userID).Return(nil, nil)

	svc := NewServer(mockRepo)

	timeFrom, _ := time.Parse(time.RFC3339, "2022-05-10T10:02:00Z")
	timeTo, _ := time.Parse(time.RFC3339, "2022-05-10T10:03:00Z")

	tasks, err := svc.GetUserTasksByTime(ctx, &pb.GetUserTasksByTimeRequest{
		UserId:   userID,
		TimeFrom: timestamppb.New(timeFrom),
		TimeTo:   timestamppb.New(timeTo),
	})

	assert.Nil(t, err)
	assert.Equal(t, tasks, &pb.TasksList{
		Tasks: nil,
	})
}
