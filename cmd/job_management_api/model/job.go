package model

import "job_executors/thirdparty/pgdb"

type JobCreateRequest struct {
	Name string `json:"name"`
}

type JobUpdateRequest struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	State pgdb.JobState `json:"state"`
}
