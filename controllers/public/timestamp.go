package public

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetTimestamp(c *fiber.Ctx) error {
	return c.JSON(time.Now().Unix())
}
