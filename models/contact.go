package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contact model
type Contact struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserIDOne string             `json:"useridone,omitempty" bson:"useridone,omitempty"`
	UserIDTwo string             `json:"useridtwo,omitempty" bson:"useridtwo,omitempty"`
	Timestamp int64              `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
