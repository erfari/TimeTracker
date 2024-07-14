package entity

import "time"

type Task struct {
	ID             int       `json:"id"`
	TaskName       string    `json:"task_name"`
	AssigneeName   string    `json:"assignee_name"`
	AssigneeUserId int       `json:"assignee_user_id"`
	StartTask      time.Time `json:"start_task"`
	EndTask        time.Time `json:"end_task"`
	CreatedAt      time.Time `json:"created_at"`
	LaborCosts     string    `json:"labor_costs"`
}

// одинаковые структуры для читаемости и расширяемости, во избежание расширении модели если в одной структуре добавится уникальное поле
type StartTask struct {
	TaskName string `json:"task_name"`
	UserName string `json:"user_name"`
}

type EndTask struct {
	TaskName string `json:"task_name"`
	UserName string `json:"user_name"`
}
