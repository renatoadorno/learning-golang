package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"password" validate:"required"`
}
