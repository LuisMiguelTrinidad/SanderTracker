package models

import (
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Saga struct {
	gorm.Model
	Id     primitive.ObjectID `json:"id"`
	Name   string             `json:"name"`
	Status string             `json:"status"`
	Books  []Book             `json:"books"`
}
