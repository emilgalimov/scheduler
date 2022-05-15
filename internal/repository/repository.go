package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/models"
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
	const sql = `INSERT INTO users DEFAULT VALUES RETURNING id`

	err = r.pool.QueryRow(ctx, sql).Scan(&ID)
	return
}

func (r *TasksRepository) CreateTask(ctx context.Context, task models.Task) (ID uint64, err error) {
	const sql = `INSERT INTO tasks(name, description) VALUES ($1, $2) RETURNING id`

	err = r.pool.QueryRow(ctx, sql, task.Name, task.Description).Scan(&ID)
	return
}

func (r *TasksRepository) GetTask(ctx context.Context, ID uint64) (task models.Task, err error) {
	const sql = `SELECT id, name, description FROM tasks WHERE id = $1`

	err = r.pool.QueryRow(ctx, sql, ID).Scan(
		&task.ID,
		&task.Name,
		&task.Description,
	)
	return
}

func (r *TasksRepository) CreateTaskStage(ctx context.Context, taskStage models.TaskStage, taskID uint64) (ID uint64, err error) {
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

func (TasksRepository) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) CreateUserTask(ctx context.Context, userID uint64, taskID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) GetUserTask(ctx context.Context, userID uint64, taskID uint64) ([]*models.UserTask, error) {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) GetAllUserTasks(ctx context.Context, userID uint64) ([]*models.UserTask, error) {
	//TODO implement me
	panic("implement me")
}

func (TasksRepository) DeleteUserTask(ctx context.Context, userID uint64, taskID uint64) error {
	//TODO implement me
	panic("implement me")
}
