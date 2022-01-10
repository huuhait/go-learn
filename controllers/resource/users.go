package resource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuhait/go-learn/config"
	"github.com/huuhait/go-learn/controllers/entities"
	"github.com/huuhait/go-learn/models"
)

func userToEntity(user *models.User) entities.User {
	return entities.User{
		ID:        user.ID,
		UID:       user.UID,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}
}

func GetUser(c *fiber.Ctx) error {
	var user *models.User

	if result := config.Database.First(&user, 1); result.Error != nil {
		return c.JSON("LỖI RỒI")
	}

	return c.JSON(userToEntity(user))
}

func CreateUser(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	payload := new(Payload)

	if err := c.BodyParser(payload); err != nil {
		c.Status(422).JSON("LỖI")
	}

	user := &models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	config.Database.Create(&user)

	return c.JSON(userToEntity(user))
}
