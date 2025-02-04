package main

import (
	"github.com/LuisMiguelTrinidad/Sandertracker/config"
	"github.com/LuisMiguelTrinidad/Sandertracker/router"
	"github.com/LuisMiguelTrinidad/Sandertracker/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	defer utils.Logger.Sync()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New())
	defer config.CloseMongoDB()

	router.SetupRoutes(app)

	utils.Logger.Error("Starting server on port 3000...")
	if err := app.Listen(":3000"); err != nil {
		utils.Logger.Fatalf("Failed to start server: %v", err)
	}
}
