package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/LuisMiguelTrinidad/Sandertracker/logging"
)

func LogRequest(c *fiber.Ctx) error {
	method := c.Method()
	url := c.OriginalURL()

	if len(url) > 1 && url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	err := c.Next()
	if err != nil {
		var status int
		switch e := err.(type) {
		case *fiber.Error:
			status = e.Code
		default:
			status = fiber.StatusInternalServerError
		}
		logging.RequestInfoLog(url, method, status)
		return err
	}

	status := c.Response().StatusCode()
	logging.RequestInfoLog(url, method, status)
	return nil
}
