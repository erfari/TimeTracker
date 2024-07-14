package tasks

import (
	"TimeTracker/internal/entity"
	"math"
	"net/http"
	"strconv"
)

type TaskService struct {
	taskRepository *TaskRepository
}

func NewTaskService(
	taskRepository *TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (ts *TaskService) GetLaborCostsByUserId(userId int) ([]*entity.Task, *entity.ResponseError) {
	tasks, responseErr := ts.taskRepository.GetLabors(userId)
	if responseErr != nil {
		return nil, responseErr
	}
	for i, v := range tasks {
		timeDuration := v.EndTask.Sub(v.StartTask)
		hours := timeDuration.Hours()
		hours, mf := math.Modf(hours)
		minutes := mf * 60
		tasks[i].LaborCosts = "hours: " + strconv.Itoa(int(hours)) + " minutes: " + strconv.Itoa(int(minutes))
	}
	return tasks, nil
}

func (ts *TaskService) AddTask(task *entity.Task) (*entity.Task, *entity.ResponseError) {
	responseErr := validateTask(task)
	if responseErr != nil {
		return nil, responseErr
	}
	task, err := ts.taskRepository.AddTask(task)
	if err != nil {
		return nil, err
	}
	return task, err
}

func (ts *TaskService) DeleteTask(taskID int) *entity.ResponseError {
	err := BeginTransaction(ts.taskRepository)
	if err != nil {
		return &entity.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}
	//Добавить роллбек
	responseErr := ts.taskRepository.DeleteTask(taskID)
	if responseErr != nil {
		return responseErr
	}
	CommitTransaction(ts.taskRepository)
	return nil
}

func (ts *TaskService) GetTask(taskID int) (*entity.Task, *entity.ResponseError) {
	task, responseErr := ts.taskRepository.GetTask(taskID)
	if responseErr != nil {
		return nil, responseErr
	}
	return task, nil
}

func validateTask(task *entity.Task) *entity.ResponseError {
	//if task.PassportNumber == "" || len([]rune(user.PassportNumber)) > 6 {
	//	return &entity.ResponseError{
	//		Message: "Invalid Passport Passport Number",
	//		Status:  http.StatusBadRequest,
	//	}
	//}
	return nil
}

func (ts *TaskService) StartTask(startTask *entity.StartTask) *entity.ResponseError {
	return ts.taskRepository.StartTask(startTask)
}

func (ts *TaskService) EndTask(endTask *entity.EndTask) *entity.ResponseError {
	return ts.taskRepository.EndTask(endTask)
}

func (ts *TaskService) UpdateTask(task *entity.Task) *entity.ResponseError {
	responseErr := validateTask(task)
	if responseErr != nil {
		return responseErr
	}
	return ts.taskRepository.UpdateTask(task)
}
