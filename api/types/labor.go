package types

type Labor struct {
	UserID          int    `json:"user_id"`
	CountTasks      int    `json:"count_tasks"`
	TimeResultLabor string `json:"time_result"`
}
