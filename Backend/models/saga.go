package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Saga struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Status string             `json:"status" bson:"status"`
	Books  []Book             `json:"books" bson:"books"`
}
