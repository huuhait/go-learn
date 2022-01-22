package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/huuhait/go-learn/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Store = session.New()
var Database *gorm.DB
var RedisClient *services.RedisClient

func InitConfig() {
	dsn := "host=103.159.51.229 user=root password=changeme dbname=fake port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database = db
	RedisClient = services.NewRedisClient("localhost:6379")
}
