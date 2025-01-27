package pgdb

import "time"

type DateFields struct {
	CreateDate time.Time  `gorm:"column:create_date" json:"create_date"`
	UpdateDate time.Time  `gorm:"column:update_date" json:"update_date"`
	DeleteDate *time.Time `gorm:"column:delete_date" json:"delete_date"`
}

type JobState string

var (
	JobStateNew     JobState = "new"
	JobStatePending JobState = "pending"
	JobStateStart   JobState = "start"
	JobStateStop    JobState = "stop"
	JobStateDone    JobState = "done"
)

type Jobs struct {
	ID    string   `gorm:"column:id" json:"id"`
	Name  string   `gorm:"column:name" json:"name"`
	State JobState `gorm:"column:state" json:"state"`

	DateFields
}

func (m *Jobs) TableName() string {
	return "jobs"
}

type TaskStatus string

var (
	TaskStatusNew  TaskStatus = "new"
	TaskStatusDone TaskStatus = "done"
)

type Tasks struct {
	ID          string     `gorm:"column:id" json:"id"`
	JobID       string     `gorm:"column:job_id" json:"job_id"`
	ExecuteTime int64      `gorm:"column:execute_time" json:"execute_time"`
	Status      TaskStatus `gorm:"column:status" json:"status"`

	DateFields
}

func (m *TaskStatus) TableName() string {
	return "tasks"
}

type ExecutorStatus string

var (
	ExecutorStatusIdle    ExecutorStatus = "idle"
	ExecutorStatusRunning ExecutorStatus = "running"
)

type Executors struct {
	ID     string         `gorm:"column:id"`
	Name   string         `gorm:"column:name"`
	Status ExecutorStatus `gorm:"column:status"`

	DateFields
}
