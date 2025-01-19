package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	gorm.Model
	Id         primitive.ObjectID `json:"id"`
	Token      string             `json:"token"`
	Expiration time.Time          `json:"expiration"`
	UserId     primitive.ObjectID `json:"userId"`
}
