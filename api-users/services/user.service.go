package services

import (
	"api-users/database"
	"api-users/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userColl *mongo.Collection = database.UserCollection("users")
)

func Insert(ctx context.Context, user models.Users) (interface{}, error) {
	newUser := models.Users{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := userColl.InsertOne(ctx, newUser)
	if err != nil {
		return result, err
	}

	return result, nil
}

func Get(ctx context.Context, userId string, user models.Users) (interface{}, error) {
	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userColl.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	return user, err
}
