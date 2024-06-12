package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // This is the one which gives us all the IDs, If ORM comes into the picture then we don't have to do this
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}