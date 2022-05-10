package models

import "time"

type User struct {
	ID int
}

type Task struct {
	ID          int
	name        string
	description string
	stages      []TaskStage
}

type TaskStage struct {
	name             string
	description      string
	minutesFromStart time.Duration
	durationMinutes  time.Duration
}

type UserTask struct {
	user      User
	task      Task
	startTime time.Time
}
