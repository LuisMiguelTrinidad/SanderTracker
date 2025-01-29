package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type Book struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Author      string             `json:"author" bson:"author"`
	Cover       string             `json:"cover" bson:"cover"`
	Pages       int                `json:"pages" bson:"pages"`
	Isbn        string             `json:"isbn" bson:"isbn"`
	PublishDate time.Time          `json:"publishDate" bson:"publishDate"`
}
