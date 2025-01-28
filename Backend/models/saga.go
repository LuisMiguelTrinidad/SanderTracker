package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Saga struct {
	Id     primitive.ObjectID `json:"id"`
	Name   string             `json:"name"`
	Status string             `json:"status"`
	Books  []Book             `json:"books"`
}
