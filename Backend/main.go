package main

import (
	"log"

	"github.com/LuisMiguelTrinidad/Sandertracker/config"
	"github.com/LuisMiguelTrinidad/Sandertracker/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	err := config.InitMongoDB()
	if err != nil {
		panic(err)
	}
	defer config.CloseMongoDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	// Your application logic here
}
