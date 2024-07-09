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

//func (us UserService) AddTask(user *models.Users) (*models.Users, *models.ResponseError) {
//	responseErr := validateUser(user)
//	if responseErr != nil {
//		return nil, responseErr
//	}
//	return us.userRepository.AddTask(user)
//}

//func (ts *TaskService) UpdateTask(task *models.Task) *models.ResponseError {
//	responseErr := validateTaskId(task.ID) //дубли?
//	if responseErr != nil {
//		return responseErr
//	}
//	responseErr = validateTask(task)
//	if responseErr != nil {
//		return responseErr
//	}
//	return ts.taskRepository.UpdateTask(task)
//}

//func (ts *TaskService) DeleteTask(userId int) *models.ResponseError {
//	responseErr := validateTaskId(userId) // наличие
//	if responseErr != nil {
//		return responseErr
//	}
//	return ts.userRepository.DeleteUser(userId)
//}
//
//func (us *UserService) GetTask(taskId int) (*models.Task, *models.ResponseError) {
//	task, responseErr := us.userRepository.GetTask(taskId)
//	if responseErr != nil {
//		return nil, responseErr
//	}
//	return task, nil
//}
//
//func (us *UserService) GetAllTasks() ([]*models.Task, *models.ResponseError) {
//	return us.taskRepositroy.GetAllTasks, nil
//}
//func countLabor(tasks []*models.Task) ([]*models.Task, *models.ResponseError) {
//	laborCosts := 0
//	for i := 0; i < 0; i++ {
//		task := &tasks[i]
//		timeTask := task.StartTask.Add(-task.EndTask)
//	}
//}

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
