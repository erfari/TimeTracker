package tasks

import (
	"TimeTracker/internal/entity"
	"database/sql"
	"net/http"
	"time"
)

type TaskRepository struct {
	DbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTaskRepository(dbHandler *sql.DB) *TaskRepository {
	return &TaskRepository{
		DbHandler: dbHandler,
	}
}

func (tr TaskRepository) StartTask(task *entity.StartTask) *entity.ResponseError {
	query := `
		UPDATE tasks
  		SET assignee_name=us.name, assignee_user_id=us.id, start_task=current_timestamp, end_task=current_timestamp
  		FROM (
  			SELECT * FROM users WHERE name=$1
  			) AS us
  		WHERE task_name=$2;`
	rows, err := tr.DbHandler.Exec(
		query, task.UserName, task.TaskName)
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &entity.ResponseError{
			Message: "Task not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (tr TaskRepository) EndTask(task *entity.EndTask) *entity.ResponseError {
	query := `
		UPDATE tasks
  		SET assignee_name=us.name, assignee_user_id=us.id, end_task=current_timestamp
  		FROM (
  			SELECT * FROM users WHERE name=$1
  			) AS us
  		WHERE task_name=$2;`
	rows, err := tr.DbHandler.Exec(
		query, task.UserName, task.TaskName)
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &entity.ResponseError{
			Message: "Task not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (tr TaskRepository) GetLabors(userId int) ([]*entity.Task, *entity.ResponseError) {
	query := `
			SELECT task_name, assignee_name, start_task, end_task
			FROM tasks AS t
			JOIN users AS u ON u.id = t.assignee_user_id
			WHERE u.id = $1
			ORDER BY end_task - start_task DESC;
`
	rows, err := tr.DbHandler.Query(query, userId)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	tasks := make([]*entity.Task, 0)
	var taskName, assigneeName string
	var startTask, endTask time.Time
	for rows.Next() {
		err := rows.Scan(&taskName, &assigneeName, &startTask, &endTask)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		task := &entity.Task{
			TaskName:     taskName,
			AssigneeName: assigneeName,
			StartTask:    startTask,
			EndTask:      endTask,
		}
		tasks = append(tasks, task)
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return tasks, nil
}

func (tr TaskRepository) AddTask(task *entity.Task) (*entity.Task, *entity.ResponseError) {
	query := `
    			INSERT INTO tasks(task_name)
    			VALUES($1)
    			RETURNING id`
	rows, err := tr.DbHandler.Query(query, task.TaskName)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var taskId int
	for rows.Next() {
		err := rows.Scan(&taskId)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &entity.Task{
		ID:             taskId,
		TaskName:       task.TaskName,
		AssigneeName:   task.AssigneeName,
		AssigneeUserId: task.AssigneeUserId,
		StartTask:      task.StartTask,
		EndTask:        task.EndTask,
		CreatedAt:      task.CreatedAt,
		LaborCosts:     task.LaborCosts,
	}, nil
}

// fv
func (tr TaskRepository) UpdateTask(task *entity.Task) *entity.ResponseError {
	query := `
    	UPDATE tasks SET task_name=$1, assignee_name=$2, patronymic=$3, address=$4 WHERE id=$5
`
	rows, err := tr.DbHandler.Exec(
		query, task.TaskName)
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &entity.ResponseError{
			Message: "Task not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

// ss
func (tr TaskRepository) DeleteTask(taskID int) *entity.ResponseError {
	query := `
		DELETE FROM tasks
		WHERE id=$1
		RETURNING id`
	rows, err := tr.transaction.Query(
		query, taskID)
	defer rows.Close()
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	for rows.Next() {
		err := rows.Scan(&taskID)
		if err != nil {
			return &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (tr TaskRepository) GetTask(taskID int) (*entity.Task, *entity.ResponseError) {
	query := `
			SELECT t.id, t.task_name, t.assignee_name, t.assignee_user_id, t.start_task, t.end_task, t.created_at
			FROM tasks AS t
			WHERE t.id = $1;`
	rows, err := tr.DbHandler.Query(query, taskID)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var taskName, assigneeName string
	var startTask, endTask, createdAt time.Time
	var id, assigneeUserId int
	for rows.Next() {
		err := rows.Scan(&id, &taskName, &assigneeName, &assigneeUserId, &startTask, &endTask, &createdAt)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &entity.Task{
		ID:             id,
		TaskName:       taskName,
		AssigneeName:   assigneeName,
		AssigneeUserId: assigneeUserId,
		StartTask:      startTask,
		EndTask:        endTask,
		CreatedAt:      createdAt,
	}, nil
}
