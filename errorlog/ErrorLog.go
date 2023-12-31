package errorlog

import (
	"errors"

	"github.com/antigloss/go/logger"
	"github.com/gofiber/fiber/v2"
)

func LogError(err string) error {
	logger.Error(err)
	return errors.New(err)
}

func Log_n_send(err string, c *fiber.Ctx, code int, message string) error {
	logger.Error(err)
	return c.Status(code).JSON(fiber.Map{"error": message, "status": false})
}
