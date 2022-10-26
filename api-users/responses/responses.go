package responses

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}

func BadRequest(err error) (int, UserResponse) {
	status := http.StatusBadRequest
	res := UserResponse{
		Status:  http.StatusBadRequest,
		Message: "error",
		Data:    &echo.Map{"data": err.Error},
	}

	return status, res
}

func InternalError(err error) (int, UserResponse) {
	status := http.StatusInternalServerError
	res := UserResponse{
		Status:  http.StatusInternalServerError,
		Message: "error",
		Data:    &echo.Map{"data": err.Error},
	}

	return status, res
}

func Created(result interface{}) (int, UserResponse) {
	status := http.StatusCreated
	res := UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &echo.Map{"data": result},
	}

	return status, res
}

func Ok(result interface{}) (int, UserResponse) {
	status := http.StatusOK
	res := UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": result},
	}

	return status, res
}
