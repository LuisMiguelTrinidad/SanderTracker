package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/LuisMiguelTrinidad/Sandertracker/config"
	"github.com/LuisMiguelTrinidad/Sandertracker/logging"
	"github.com/LuisMiguelTrinidad/Sandertracker/middleware"
	"github.com/LuisMiguelTrinidad/Sandertracker/router"
)

func main() {
	defer logging.CloseLogFile()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New())
	app.Use(middleware.LogRequest)
	defer config.CloseMongoDB()

	router.SetupRoutes(app)

	// Mover el middleware de logging al final

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
		}
		logging.SystemInfoLog(fmt.Sprintf("Starting server on port %s", port))
		if err := app.Listen(":" + port); err != nil {
			logging.SystemFatalLog(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	<-sigChan
	logging.SystemInfoLog("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Error during shutdown: %v", err))
	}
}
