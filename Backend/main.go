package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/LuisMiguelTrinidad/Sandertracker/controllers"
	"github.com/LuisMiguelTrinidad/Sandertracker/database"
	"github.com/LuisMiguelTrinidad/Sandertracker/middleware"
	"github.com/LuisMiguelTrinidad/Sandertracker/router"
	"github.com/LuisMiguelTrinidad/Sandertracker/utils/logging"
)

func setupServer() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New())
	app.Use(middleware.LogRequest)
	router.SetupRoutes(app)
	controllers.InitControllers()
	return app
}

func startServer(app *fiber.App) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logging.SystemInfoLog(fmt.Sprintf("Starting server on port %s", port))
	if err := app.Listen(":" + port); err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Failed to start server: %v", err))
	}
}

func gracefulShutdown(app *fiber.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	logging.SystemInfoLog("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Error during shutdown: %v", err))
	}
}

func main() {
	defer logging.CloseLogFile()
	defer database.CloseMongoDB()

	logging.InitLogger()
	database.InitMongoDB()

	app := setupServer()
	go startServer(app)
	gracefulShutdown(app)
}
