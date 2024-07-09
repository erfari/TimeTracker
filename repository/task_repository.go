package repository

import (
	"TimeTracker/models"
	"database/sql"
	"net/http"
	"time"
)

type TaskRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTaskRepository(dbHandler *sql.DB) *TaskRepository {
	return &TaskRepository{
		dbHandler: dbHandler,
	}
}

func (tr TaskRepository) StartTask(task *models.StartTask) *models.ResponseError {
	query := `
		UPDATE tasks
  		SET assignee_name=us.name, assignee_user_id=us.id, start_task=current_timestamp, end_task=current_timestamp
  		FROM (
  			SELECT * FROM users WHERE name=$1
  			) AS us
  		WHERE task_name=$2;`
	rows, err := tr.dbHandler.Exec(
		query, task.UserName, task.TaskName)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "Task not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (tr TaskRepository) EndTask(task *models.EndTask) *models.ResponseError {
	query := `
		UPDATE tasks
  		SET assignee_name=us.name, assignee_user_id=us.id, end_task=current_timestamp
  		FROM (
  			SELECT * FROM users WHERE name=$1
  			) AS us
  		WHERE task_name=$2;`
	rows, err := tr.dbHandler.Exec(
		query, task.UserName, task.TaskName)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "Task not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (tr TaskRepository) GetLabors(userId int) ([]*models.Task, *models.ResponseError) {
	query := `
			SELECT task_name, assignee_name, start_task, end_task
			FROM tasks AS t
			JOIN users AS u ON u.id = t.assignee_user_id
			WHERE u.id = $1
			ORDER BY end_task - start_task DESC;
`
	rows, err := tr.dbHandler.Query(query, userId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	tasks := make([]*models.Task, 0)
	var taskName, assigneeName string
	var startTask, endTask time.Time
	for rows.Next() {
		err := rows.Scan(&taskName, &assigneeName, &startTask, &endTask)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		user := &models.Task{
			TaskName:     taskName,
			AssigneeName: assigneeName,
			StartTask:    startTask,
			EndTask:      endTask,
		}
		tasks = append(tasks, user)
	}
	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return tasks, nil
}
