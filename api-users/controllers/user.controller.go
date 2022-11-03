package controllers

import (
	"api-users/models"
	"api-users/responses"
	"api-users/services"
	"api-users/validations"
	"context"
	"github.com/labstack/echo/v4"
)

var (
	ctx  context.Context = context.TODO()
	user models.Users
)

func CreateUser(c echo.Context) error {
	userErr := c.Bind(&user)
	if userErr != nil {
		return c.JSON(responses.BadRequest(userErr, "Error: Body não encontrado"))
	}

	message, validErr := validations.ValidCreateUser(ctx, user)
	if validErr != nil {
		return c.JSON(responses.BadRequest(validErr, message))
	}

	result, err := services.CreateUser(ctx, user)
	if err != nil {
		return c.JSON(responses.InternalError(err, "Erro ao inserir usuario"))
	}

	return c.JSON(responses.Created(result))
}

func GetUser(c echo.Context) error {
	userId := c.Param("userId")

	result, err := services.GetUserById(ctx, userId, user)
	if err != nil {
		return c.JSON(responses.InternalError(err, "Erro interno ao procurar usuario"))
	}

	return c.JSON(responses.Ok(result))
}

func UpdateUser(c echo.Context) error {
	userId := c.Param("userId")

	if err := c.Bind(&user); err != nil {
		return c.JSON(responses.BadRequest(err, "Erro, Body não encontrado"))
	}

	message, validErr := validations.ValidCreateUser(ctx, user)
	if validErr != nil {
		return c.JSON(responses.BadRequest(validErr, message))
	}

	result, err := services.UpdateUserById(ctx, userId, user)
	if err != nil {
		return c.JSON(responses.InternalError(err, "Erro interno ao atualizar usuario"))
	}

	return c.JSON(responses.Ok(result))
}

func DeleteUser(c echo.Context) error {
	userId := c.Param("userId")

	result, err := services.DeleteUserById(ctx, userId, user)
	if err != nil {
		return c.JSON(responses.InternalError(err, "Erro interno ao tentar deletar usuario"))
	}

	return c.JSON(responses.Ok(result))
}

func GetAllUsers(c echo.Context) error {
	result, err := services.GetAllUsers(ctx, user)
	if err != nil {
		return c.JSON(responses.InternalError(err, "Erro interno ao tentar buscar usuarios"))
	}

	return c.JSON(responses.Ok(result))
}
