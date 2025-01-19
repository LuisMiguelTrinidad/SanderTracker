package router

import (
	"github.com/LuisMiguelTrinidad/Sandertracker/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", controllers.GetBooks)
	app.Get("/api/v1/books/:id", controllers.GetBook)
	app.Post("/api/v1/books", controllers.NewBook)
	app.Delete("/api/v1/books/:id", controllers.DeleteBook)
	app.Put("/api/v1/books/:id", controllers.UpdateBook)
}
