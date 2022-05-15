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
	if err != nil {
		return models.Task{}, err
	}

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

func (r *TasksRepository) GetAllTasks(ctx context.Context) (tasks []models.Task, err error) {
	//language=PostgreSQL
	const sql = `SELECT 
    t.id, 
    t.name, 
    t.description, 
    ts.id, 
    ts.name, 
    ts.description, 
    ts.minutes_from_start,
    ts.duration_minutes
    FROM tasks t LEFT JOIN task_stages ts on t.id = ts.task_id ORDER BY t.ID, ts.minutes_from_start`

	row, err := r.pool.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	tasksMap := make(map[uint64]*models.Task)

	for row.Next() {
		values, _ := row.Values()

		taskId := uint64(values[0].(int64))

		if _, ok := tasksMap[taskId]; !ok {
			tasksMap[taskId] = &models.Task{
				ID:          taskId,
				Name:        values[1].(string),
				Description: values[2].(string),
			}
		}

		if values[3] == nil {
			continue
		}

		tasksMap[taskId].Stages = append(tasksMap[taskId].Stages,
			models.TaskStage{
				ID:               uint64(values[3].(int64)),
				Name:             values[4].(string),
				Description:      values[5].(string),
				MinutesFromStart: time.Duration(values[6].(int64)) * time.Millisecond,
				DurationMinutes:  time.Duration(values[7].(int64)) * time.Millisecond,
			})
	}

	for _, v := range tasksMap {
		tasks = append(tasks, *v)
	}

	return tasks, err
}

func (r *TasksRepository) CreateUserTask(ctx context.Context, userID uint64, taskID uint64) (err error) {
	//language=SQL
	const sql = `INSERT INTO 
		user_tasks(user_id, task_id, start_time) 
		VALUES ($1, $2, now())
		`

	r.pool.QueryRow(ctx, sql, userID, taskID)

	return
}

func (r *TasksRepository) GetUserTask(ctx context.Context, userID uint64, taskID uint64) (userTask models.UserTask, err error) {
	//language=SQL
	const sql = `SELECT start_time FROM user_tasks WHERE user_id = $1 AND task_id = $2`

	err = r.pool.QueryRow(ctx, sql, userID, taskID).Scan(&userTask.StartTime)

	if err != nil {
		return models.UserTask{}, err
	}

	task, err := r.GetTask(ctx, taskID)

	if err != nil {
		return models.UserTask{}, err
	}
	userTask.Task = task

	userTask.User = models.User{ID: userID}

	return userTask, err
}

func (r *TasksRepository) GetAllUserTasks(ctx context.Context, userID uint64) (userTasks []models.UserTask, err error) {
	//language=SQL
	const sql = "SELECT task_id, start_time FROM user_tasks WHERE user_id = $1"

	row, _ := r.pool.Query(ctx, sql, userID)

	for row.Next() {
		values, _ := row.Values()

		task, errTask := r.GetTask(ctx, uint64(values[0].(int64)))
		if errTask != nil {
			return nil, errTask
		}

		userTasks = append(userTasks, models.UserTask{
			User:      models.User{ID: userID},
			Task:      task,
			StartTime: values[1].(time.Time),
		})

	}

	return userTasks, nil
}

func (r *TasksRepository) DeleteUserTask(ctx context.Context, userID uint64, taskID uint64) (err error) {
	//language=SQL
	const sql = `DELETE FROM user_tasks WHERE user_id = $1 AND task_id = $2`

	r.pool.QueryRow(ctx, sql, userID, taskID)

	return
}
