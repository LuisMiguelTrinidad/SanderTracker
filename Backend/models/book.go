package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type Book struct {
	Id          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Cover       string             `json:"cover"`
	Pages       int                `json:"pages"`
	Isbn        int                `json:"isbn"`
	PublishDate time.Time          `json:"publishDate"`
}
