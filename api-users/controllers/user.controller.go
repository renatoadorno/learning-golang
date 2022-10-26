package controllers

import (
	"api-users/models"
	"api-users/responses"
	"api-users/services"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	ctx  context.Context = context.TODO()
	user models.Users
)

func CreateUser(c echo.Context) error {
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(responses.BadRequest(err, "error body"))
	}

	validate := validator.New()
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(responses.BadRequest(err, "error validate"))
	}

	result, err := services.Insert(ctx, user)
	if err != nil {
		return c.JSON(responses.InternalError(err, "Internal server error"))
	}

	return c.JSON(responses.Created(result))
}

func GetUser(c echo.Context) error {
	userId := c.Param("userId")

	result, err := services.Get(ctx, userId, user)

	if err != nil {
		return c.JSON(responses.InternalError(err, "Internal server Error"))
	}

	return c.JSON(responses.Ok(result))
}
