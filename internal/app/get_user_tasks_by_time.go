package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"time"
)

func (t *tserver) GetUserTasksByTime(ctx context.Context, rq *pb.GetUserTasksByTimeRequest) (*pb.TasksList, error) {

	userTasks, _ := t.repo.GetAllUserTasks(ctx, rq.UserId)

	var tasks []models.Task

	for _, ut := range userTasks {
		ut.Task.Stages = filterTaskStagesByTime(ut.Task.Stages, ut.StartTime, rq.TimeFrom.AsTime(), rq.TimeTo.AsTime())
		tasks = append(tasks, ut.Task)
	}

	returnTasks := makeTasksResponseFromTasks(tasks)
	return &pb.TasksList{Tasks: returnTasks}, nil
}

func filterTaskStagesByTime(
	stages []models.TaskStage,
	startTime time.Time,
	fromTime time.Time,
	toTime time.Time,
) []models.TaskStage {

	var filteredStages []models.TaskStage
	for _, ts := range stages {
		tsStartTime := startTime.Add(ts.MinutesFromStart)
		if (tsStartTime.Equal(fromTime) || tsStartTime.After(fromTime)) && tsStartTime.Before(toTime) {
			filteredStages = append(filteredStages, ts)
		}
	}

	return filteredStages
}
