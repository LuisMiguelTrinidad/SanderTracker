package router

import (
	"github.com/LuisMiguelTrinidad/Sandertracker/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/books", controllers.GetBooks)
	v1.Get("/books/:id", controllers.GetBook)
	v1.Post("/books", controllers.CreateBook)
	v1.Delete("/books/:id", controllers.DeleteBook)
	v1.Put("/books/:id", controllers.UpdateBook)
}
