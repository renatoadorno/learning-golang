package routes

import (
	"api-users/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetUser)
	e.PUT("/user/:userId", controllers.UpdateUser)
	e.DELETE("/user/:userId", controllers.DeleteUser)
	e.GET("/users", controllers.GetAllUsers)
}
