package models

import "time"

type UserTask struct {
	ID          int       `json:"id"`
	TaskName    int       `json:"task_name"`
	TaskStartAt time.Time `db:"taskStartAt"`
	TaskEndAt   time.Time `db:"taskEndAt"`
}
