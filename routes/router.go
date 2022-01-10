package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/huuhait/go-learn/controllers/public"
	"github.com/huuhait/go-learn/controllers/resource"
)

func InitRouter() {
	app := fiber.New()

	app.Get("/public/timestamp", public.GetTimestamp)
	app.Get("/resource/users/me", resource.GetUser)
	app.Post("/resource/users/create", resource.CreateUser)

	log.Fatal(app.Listen(":3000"))
}
