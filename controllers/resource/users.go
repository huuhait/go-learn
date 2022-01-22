package resource

import (
	"log"

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

	session, _ := config.Store.Get(c)
	email := session.Get("email")
	log.Println(email)

	if email == nil {
		return c.JSON("CHƯA ĐĂNG NHẬP!")
	}

	if result := config.Database.First(&user, "email = ?", email); result.Error != nil {
		session.Destroy()
		return c.JSON("KO TÌM THẤY USER")
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
