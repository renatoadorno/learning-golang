package services

import (
	"api-users/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, user models.Users) (interface{}, error) {
	result, err := models.Insert(ctx, user)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetUserById(ctx context.Context, userId string, user models.Users) (models.Users, error) {
	result, err := models.GetById(ctx, user, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func UpdateUserById(ctx context.Context, userId string, user models.Users) (interface{}, error) {
	result, err := models.UpdateById(ctx, user, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func DeleteUserById(ctx context.Context, userId string, user models.Users) (*mongo.DeleteResult, error) {
	result, err := models.DeleteById(ctx, user, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetAllUsers(ctx context.Context, user models.Users) ([]models.Users, error) {
	result, err := models.GetAll(ctx, user)
	if err != nil {
		return result, err
	}

	return result, nil
}
