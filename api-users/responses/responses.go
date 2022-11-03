package responses

import (
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BadRequest(_err error, msg string) (int, ErrorResponse) {
	status := http.StatusBadRequest

	return status, ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

func InternalError(_err error, msg string) (int, ErrorResponse) {
	status := http.StatusInternalServerError

	return status, ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

func Ok(result interface{}) (int, SuccessResponse) {
	status := http.StatusOK

	return status, SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	}
}

func Created(result interface{}) (int, SuccessResponse) {
	status := http.StatusCreated

	return status, SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
}
