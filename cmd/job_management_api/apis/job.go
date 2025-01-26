package apis

import (
	"context"
	"job_executors/cmd/job_management_api/model"
	"job_executors/thirdparty/pgdb"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IJobRepo interface {
	CreateJob(ctx context.Context, req model.JobCreateRequest) (*pgdb.Jobs, error)
	GetJob(ctx context.Context, id string) (*pgdb.Jobs, error)
	ListJobs(ctx context.Context) ([]pgdb.Jobs, error)
	ListJobsByState(ctx context.Context, jobState pgdb.JobState) ([]pgdb.Jobs, error)
	UpdateJob(ctx context.Context, req model.JobUpdateRequest) (*pgdb.Jobs, error)
	DeleteJob(ctx context.Context, id string) error
}

type ITaskRepo interface {
	CreateTask(ctx context.Context, req model.TaskCreateRequest) ([]pgdb.Tasks, error)
	DeleteTask(ctx context.Context, jobID string) error
	UpdateTask(ctx context.Context, req model.JobUpdateRequest) (*pgdb.Tasks, error)
}

type JobAPI struct {
	jobRepo  IJobRepo
	taskRepo ITaskRepo
}

func NewJobAPI(jobRepo IJobRepo, taskRepo ITaskRepo) *JobAPI {
	return &JobAPI{jobRepo, taskRepo}
}

func (a *JobAPI) Setup(g echo.Group) {
	g.GET("/jobs", a.listJobs)
	g.POST("/jobs/create", a.createJob)
	g.POST("/jobs/update", a.updateJob)
	g.POST("/tasks/update", a.updateTask)
	g.DELETE("/jobs/:id", a.deleteJob)
}

func (a *JobAPI) listJobs(c echo.Context) error {

	ctx := c.Request().Context()
	jobs, err := a.jobRepo.ListJobs(ctx)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		model.BaseResponse{
			Message: "success",
			Data:    jobs,
		},
	)
}

func (a *JobAPI) createJob(c echo.Context) error {

	var req model.JobCreateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	ctx := c.Request().Context()
	job, err := a.jobRepo.CreateJob(ctx, req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		model.BaseResponse{
			Message: "success",
			Data:    job,
		},
	)
}

func (a *JobAPI) updateJob(c echo.Context) error {

	var req model.JobUpdateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	ctx := c.Request().Context()

	updatedJob, err := a.jobRepo.UpdateJob(ctx, req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		model.BaseResponse{
			Message: "success",
			Data:    updatedJob,
		},
	)
}

func (a *JobAPI) updateTask(c echo.Context) error {

	var req model.TaskCreateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	ctx := c.Request().Context()

	job, err := a.jobRepo.GetJob(ctx, req.JobID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	if job.State == pgdb.JobStateStart {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: "unable to update tasks on running job",
			},
		)
	}

	// delete all task in request job id
	err = a.taskRepo.DeleteTask(ctx, req.JobID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	// create new task id
	tasks, err := a.taskRepo.CreateTask(ctx, req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		model.BaseResponse{
			Message: "success",
			Data:    tasks,
		},
	)
}

func (a *JobAPI) deleteJob(c echo.Context) error {

	jobID := c.Param("id")

	ctx := c.Request().Context()
	job, err := a.jobRepo.GetJob(ctx, jobID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	if job.State == pgdb.JobStateStart {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: "unable to delete running job",
			},
		)
	}

	err = a.taskRepo.DeleteTask(
		ctx,
		jobID,
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	err = a.jobRepo.DeleteJob(
		ctx,
		jobID,
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.BaseResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		model.BaseResponse{
			Message: "success",
			Data: model.JobDeleteResponse{
				ID: jobID,
			},
		})
}
