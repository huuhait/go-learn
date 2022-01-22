package public

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huuhait/go-learn/config"
)

func GetCount(c *fiber.Ctx) error {
	var count int64

	count_str, err := config.RedisClient.Get("count")
	if err != nil {
		count = 1
		config.RedisClient.Set("count", count)

		return c.JSON(count)
	}

	if i, err := strconv.ParseInt(count_str, 10, 64); err != nil {
		return c.JSON("Count không phải là một số nguyên")
	} else {
		count = i
	}

	count++

	config.RedisClient.Set("count", count)

	return c.JSON(count)
}
