package main

import (
	"log"

	"github.com/LuisMiguelTrinidad/Sandertracker/config"
	"github.com/LuisMiguelTrinidad/Sandertracker/controllers"
	"github.com/LuisMiguelTrinidad/Sandertracker/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	if err := config.InitMongoDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer config.CloseMongoDB()

	// Inicializar controladores después de la conexión a MongoDB
	if err := controllers.InitializeControllers(); err != nil {
		log.Fatalf("Failed to initialize controllers: %v", err)
	}

	router.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
