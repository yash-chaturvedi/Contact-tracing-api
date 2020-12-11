package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	BirthDate   string             `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	PhoneNumber string             `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt   int64              `json:"createdat,omitempty" bson:"createdat,omitempty"`
}
