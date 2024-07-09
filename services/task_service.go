package services

import (
	"TimeTracker/models"
	"TimeTracker/repository"
	"math"
	"strconv"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewTaskService(
	taskRepository *repository.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (ts *TaskService) GetLaborCostsByUserId(userId int) ([]*models.Task, *models.ResponseError) {
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

func (ts *TaskService) StartTask(startTask *models.StartTask) *models.ResponseError {
	return ts.taskRepository.StartTask(startTask)
}

func (ts *TaskService) EndTask(endTask *models.EndTask) *models.ResponseError {
	return ts.taskRepository.EndTask(endTask)
}
