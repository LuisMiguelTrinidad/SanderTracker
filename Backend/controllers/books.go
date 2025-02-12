package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/LuisMiguelTrinidad/Sandertracker/database"
	"github.com/LuisMiguelTrinidad/Sandertracker/models"
)

var bookCollection *mongo.Collection

func InitBooks() {
	bookCollection = database.Db.Collection("books")
}

func GetBooks(c *fiber.Ctx) error {
	cursor, err := bookCollection.Find(c.Context(), bson.D{})

	if err != nil {
		return c.Status(500).SendString("Error fetching books")
	}

	var results []models.Book
	if err := cursor.All(c.Context(), &results); err != nil {
		return c.Status(500).SendString("Error fetching books")
	}

	return c.JSON(results)
}

func GetBook(c *fiber.Ctx) error {
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}
	cursor := bookCollection.FindOne(c.Context(), bson.D{{Key: "_id", Value: objectId}})
	if cursor.Err() != nil {
		return c.Status(404).SendString("Book not found")
	}
	result := models.Book{}
	if err := cursor.Decode(&result); err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.JSON(result)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	result, err := bookCollection.InsertOne(c.Context(), book)
	if err != nil {
		return c.Status(500).SendString("Failed to create book")
	}

	return c.Status(201).JSON(result)
}

func DeleteBook(c *fiber.Ctx) error {
	objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	result, err := bookCollection.DeleteOne(c.Context(), bson.D{{Key: "_id", Value: objectId}})

	if err != nil {
		return c.Status(500).SendString("Error deleting book")
	}

	if result.DeletedCount == 0 {
		return c.Status(404).SendString("Book not found")
	}

	return c.SendString("Deleted book " + objectId.Hex())
}

func UpdateBook(c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	var updateData models.Book
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	if updateData == (models.Book{}) {
		return c.Status(400).SendString("No fields provided for update")
	}

	update := bson.M{"$set": updateData}

	result, err := bookCollection.UpdateOne(c.Context(), bson.D{{Key: "_id", Value: objectID}}, update)
	if err != nil {
		return c.Status(500).SendString("Failed to update book")
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Book not found")
	}
	return c.Status(200).SendString("Book updated")
}
