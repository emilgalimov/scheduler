package models

import (
	"time"
)

type User struct {
	ID uint64
}

type Task struct {
	ID          uint64
	Name        string
	Description string
	Stages      []*TaskStage
}

type TaskStage struct {
	ID               uint64
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
