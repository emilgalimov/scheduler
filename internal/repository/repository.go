package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
	"time"
)

type TasksRepository struct {
	pool *pgxpool.Pool
}

func NewTasksRepository(pool *pgxpool.Pool) *TasksRepository {
	return &TasksRepository{
		pool: pool,
	}
}

func (r *TasksRepository) CreateUser(ctx context.Context) (ID uint64, err error) {
	//language=PostgreSQL
	const sql = `INSERT INTO users DEFAULT VALUES RETURNING id`

	err = r.pool.QueryRow(ctx, sql).Scan(&ID)
	return
}

func (r *TasksRepository) CreateTask(ctx context.Context, task models.Task) (ID uint64, err error) {
	//language=PostgreSQL
	const sql = `INSERT INTO tasks(name, description) VALUES ($1, $2) RETURNING id`

	err = r.pool.QueryRow(ctx, sql, task.Name, task.Description).Scan(&ID)
	return
}

func (r *TasksRepository) GetTask(ctx context.Context, ID uint64) (task models.Task, err error) {
	//language=PostgreSQL
	const sqlT = `SELECT id, name, description FROM tasks WHERE id = $1`

	err = r.pool.QueryRow(ctx, sqlT, ID).Scan(
		&task.ID,
		&task.Name,
		&task.Description,
	)

	//language=PostgreSQL
	const sqlTS = `SELECT id, name, description, minutes_from_start, duration_minutes FROM task_stages WHERE task_id = $1`

	rows, err := r.pool.Query(ctx, sqlTS, ID)
	for {
		if !rows.Next() {
			break
		}
		values, _ := rows.Values()
		task.Stages = append(task.Stages, models.TaskStage{
			ID:               uint64(values[0].(int64)),
			Name:             values[1].(string),
			Description:      values[2].(string),
			MinutesFromStart: time.Duration(values[3].(int64)) * time.Millisecond,
			DurationMinutes:  time.Duration(values[4].(int64)) * time.Millisecond,
		})
	}

	return task, err
}

func (r *TasksRepository) CreateTaskStage(ctx context.Context, taskStage models.TaskStage, taskID uint64) (ID uint64, err error) {
	//language=SQL
	const sql = `INSERT INTO 
		task_stages(name, description, minutes_from_start, duration_minutes, task_id) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id
		`

	err = r.pool.QueryRow(ctx, sql,
		taskStage.Name,
		taskStage.Description,
		taskStage.MinutesFromStart.Milliseconds(),
		taskStage.DurationMinutes.Milliseconds(),
		taskID,
	).Scan(&ID)
	return
}

func (TasksRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) CreateUserTask(ctx context.Context, userID uint64, taskID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) GetUserTask(ctx context.Context, userID uint64, taskID uint64) ([]models.UserTask, error) {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) GetAllUserTasks(ctx context.Context, userID uint64) ([]models.UserTask, error) {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) DeleteUserTask(ctx context.Context, userID uint64, taskID uint64) error {
	//TODO implement me
	panic("implement me")
}
