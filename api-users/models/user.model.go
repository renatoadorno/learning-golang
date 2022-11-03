package models

import (
	"api-users/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	Name     string `json:"name,omitempty"`
	Email    string `bson:"email" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var (
	userColl *mongo.Collection = database.UserCollection("users")
)

func Insert(ctx context.Context, user Users) (interface{}, error) {
	newUser := Users{
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

func UniqueEmail(ctx context.Context, user Users) error {
	newUser := Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	// objEmail, _ := primitive.ObjectIDFromHex(newUser.Email)
	err := userColl.FindOne(ctx, bson.M{"email": newUser.Email}).Decode(&user)
	if err != nil {
		return err
	}

	return nil
}

func GetById(ctx context.Context, user Users, id string) (Users, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	err := userColl.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateById(ctx context.Context, user Users, id string) (interface{}, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	userUpdate := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}
	result, err := userColl.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": userUpdate})
	if err != nil {
		return result, err
	}

	var updated Users
	if result.MatchedCount == 1 {
		err := userColl.FindOne(ctx, bson.M{"_id": objId}).Decode(&updated)
		if err != nil {
			return updated, err
		}
	}

	return updated, nil
}

func DeleteById(ctx context.Context, user Users, id string) (*mongo.DeleteResult, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := userColl.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetAll(ctx context.Context, user Users) ([]Users, error) {
	var users []Users

	results, err := userColl.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var single Users
		if err = results.Decode(&single); err != nil {
			return users, err
		}

		users = append(users, single)
	}

	return users, nil
}
