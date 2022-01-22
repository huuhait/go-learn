package identity

import (
	"github.com/gofiber/fiber/v2"

	"github.com/huuhait/go-learn/config"
	"github.com/huuhait/go-learn/models"
)

func Login(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	payload := new(Payload)

	if err := c.BodyParser(payload); err != nil {
		return c.JSON("LỖI")
	}

	var user *models.User
	config.Database.First(&user)

	if !user.CheckPasswordHash(payload.Password) {
		return c.JSON("LỖI")
	}

	session, _ := config.Store.Get(c)

	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		return c.JSON("LỖI lưu cookie")
	}

	return c.Status(201).JSON(201)
}
