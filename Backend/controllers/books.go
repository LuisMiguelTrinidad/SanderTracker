package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/LuisMiguelTrinidad/Sandertracker/config"
	"github.com/LuisMiguelTrinidad/Sandertracker/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBooks(c *fiber.Ctx) error {
	books := config.Db.Collection("Books")
	cursor, err := books.Find(context.Background(), bson.D{})
	fmt.Println(cursor)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var results []models.Book
	if err := cursor.All(c.Context(), &results); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(results)
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update book")
}
