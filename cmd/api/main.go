package main

import (
	"github.com/huuhait/go-learn/config"
	"github.com/huuhait/go-learn/models"
	"github.com/huuhait/go-learn/routes"
)

func main() {
	config.InitConfig()

	config.Database.AutoMigrate(models.User{}, models.Music{})

	routes.InitRouter()
}
