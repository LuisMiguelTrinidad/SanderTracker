package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	Id         primitive.ObjectID `json:"id"`
	Token      string             `json:"token"`
	Expiration time.Time          `json:"expiration"`
	UserId     primitive.ObjectID `json:"userId"`
}
