package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/LuisMiguelTrinidad/Sandertracker/config"
	"github.com/LuisMiguelTrinidad/Sandertracker/logging"
	"github.com/LuisMiguelTrinidad/Sandertracker/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	defer logging.CloseLogFile()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New())
	defer config.CloseMongoDB()

	router.SetupRoutes(app)

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
		}
		logging.LogInfo(fmt.Sprintf("Starting server on port %s", port))
		if err := app.Listen(":" + port); err != nil {
			logging.LogFatal(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	<-sigChan
	logging.LogInfo("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		logging.LogFatal(fmt.Sprintf("Error during shutdown: %v", err))
	}
}
