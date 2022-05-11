package models

import "time"

type User struct {
	ID int
}

type Task struct {
	ID          int
	Name        string
	Description string
	stages      []TaskStage
}

type TaskStage struct {
	Name             string
	Description      string
	MinutesFromStart time.Duration
	DurationMinutes  time.Duration
}

type UserTask struct {
	user      User
	task      Task
	startTime time.Time
}
