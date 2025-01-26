package model

import "job_executors/thirdparty/pgdb"

type Tasks struct {
	ExecuteTime int64 `json:"execute_time"`
}

type TaskCreateRequest struct {
	JobID     string  `json:"job_id"`
	TasksList []Tasks `json:"tasks"`
}

type TaskUpdateRequest struct {
	ID     string          `json:"id"`
	Status pgdb.TaskStatus `json:"status"`
}
