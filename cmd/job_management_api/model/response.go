package model

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type JobDeleteResponse struct {
	ID string `json:"id"`
}
