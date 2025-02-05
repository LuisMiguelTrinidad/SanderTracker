package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/LuisMiguelTrinidad/Sandertracker/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	BookRoutes(api)
}

func BookRoutes(group fiber.Router) {
	books := group.Group("/books")

	books.Get("/", controllers.GetBooks)
	books.Get("/:id", controllers.GetBook)
	books.Post("/", controllers.CreateBook)
	books.Put("/:id", controllers.UpdateBook)
	books.Delete("/:id", controllers.DeleteBook)
}
