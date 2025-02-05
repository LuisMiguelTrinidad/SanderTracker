package middleware

import (
	"fmt"

	"github.com/LuisMiguelTrinidad/Sandertracker/logging"

	"github.com/gofiber/fiber/v2"
)

func LogRequest(c *fiber.Ctx) error {
	url := c.OriginalURL()
	// Trim trailing slash if it exists and it's not the root path
	if len(url) > 1 && url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	statusCode := c.Response().StatusCode()
	// Color code the status code
	coloredStatus := fmt.Sprint(statusCode)
	switch {
	case statusCode >= 200 && statusCode < 300:
		coloredStatus = "\033[32m" + coloredStatus + "\033[0m" // Green for 2xx
	case statusCode >= 300 && statusCode < 400:
		coloredStatus = "\033[36m" + coloredStatus + "\033[0m" // Cyan for 3xx
	case statusCode >= 400 && statusCode < 500:
		coloredStatus = "\033[33m" + coloredStatus + "\033[0m" // Yellow for 4xx
	case statusCode >= 500:
		coloredStatus = "\033[31m" + coloredStatus + "\033[0m" // Red for 5xx
	}

	// Get request method and color code it
	method := c.Method()
	coloredMethod := method
	switch method {
	case "GET":
		coloredMethod = "\033[32m" + method + "\033[0m" // Green
	case "POST":
		coloredMethod = "\033[34m" + method + "\033[0m" // Blue
	case "PUT":
		coloredMethod = "\033[33m" + method + "\033[0m" // Yellow
	case "DELETE":
		coloredMethod = "\033[31m" + method + "\033[0m" // Red
	}

	logging.LogInfo(fmt.Sprintf("%v %v %v", coloredMethod, coloredStatus, url))
	return c.Next()
}
