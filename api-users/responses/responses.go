package responses

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    interface{} `json:"data"`
}

func BadRequest(err error, msg string) (int, UserResponse) {
	status := http.StatusBadRequest

	return status, UserResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
		Data:    &echo.Map{"error": err},
	}
}

func InternalError(err error, msg string) (int, UserResponse) {
	status := http.StatusInternalServerError

	return status, UserResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Data:    &echo.Map{"error": err},
	}
}

func Ok(result interface{}) (int, UserResponse) {
	status := http.StatusOK

	return status, UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	}
}

func Created(result interface{}) (int, UserResponse) {
	status := http.StatusCreated

	return status, UserResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
}
