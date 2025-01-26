package apis

import (
	"encoding/json"
	"job_executors/mocks"
	"job_executors/thirdparty/pgdb"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/jobs")

	jobRepo := mocks.NewMockIJobRepo(t)

	mockReturn := []pgdb.Jobs{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}

	jobRepo.EXPECT().
		ListJobs(req.Context()).
		Return(
			mockReturn,
			nil,
		)

	taskRepo := mocks.NewMockITaskRepo(t)

	jobAPI := NewJobAPI(
		jobRepo,
		taskRepo,
	)

	assert.NoError(t, jobAPI.listJobs(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Message string      `json:"message"`
		Data    []pgdb.Jobs `json:"data"`
	}

	assert.NoError(
		t,
		json.Unmarshal(
			rec.Body.Bytes(),
			&resp,
		),
	)
	assert.Equal(t, resp.Data, mockReturn)

}
