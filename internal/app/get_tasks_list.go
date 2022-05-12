package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (t *tserver) GetTasksList(ctx context.Context, rq *pb.GetTasksListRequest) (*pb.GetTasksListResponse, error) {
	tasks, err := t.repo.GetAllTasks(ctx)

	tasksToResponse := &pb.GetTasksListResponse{
		Tasks: nil,
	}

	tasksToResponse.Tasks = makeTasksResponseFromTasks(tasks)

	return tasksToResponse, err
}

func makeTasksResponseFromTasks(tasks []*models.Task) []*pb.Task {
	var toReturn []*pb.Task
	for _, t := range tasks {
		var taskStagesToReturn []*pb.TaskStage
		for _, ts := range t.Stages {
			taskStagesToReturn = append(taskStagesToReturn, &pb.TaskStage{
				ID:               ts.ID,
				Name:             ts.Name,
				Description:      ts.Description,
				MinutesFromStart: uint64(ts.MinutesFromStart.Minutes()),
				DurationMinutes:  uint64(ts.DurationMinutes.Minutes()),
			})
		}
		toReturn = append(toReturn, &pb.Task{
			ID:          t.ID,
			Name:        t.Name,
			Description: t.Description,
			TaskStages:  taskStagesToReturn,
		})
	}

	return toReturn
}
