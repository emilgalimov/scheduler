package app

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"testing"
	"time"
)

func TestTserver_GetTasksList(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)

	mockRepo.GetAllTasksMock.Return([]*models.Task{
		{
			ID:          1,
			Name:        "name",
			Description: "description",
			Stages: []*models.TaskStage{
				{
					ID:               1,
					Name:             "nam",
					Description:      "desc",
					MinutesFromStart: 1 * time.Minute,
					DurationMinutes:  2 * time.Minute,
				},
			},
		},
	}, nil)
	ctx := context.Background()
	svc := NewServer(mockRepo)

	tasks, err := svc.GetTasksList(ctx)

	assert.Nil(t, err)
	assert.Equal(t, tasks, &pb.TasksList{
		Tasks: []*pb.Task{
			{
				ID:          1,
				Name:        "name",
				Description: "description",
				TaskStages: []*pb.TaskStage{
					{
						ID:               1,
						Name:             "nam",
						Description:      "desc",
						MinutesFromStart: 1,
						DurationMinutes:  2,
					},
				},
			},
		},
	})
}
