package main

import (
	"duck-cook-recipe/api"
	"duck-cook-recipe/controller"
	"duck-cook-recipe/pkg/mongo"
	"duck-cook-recipe/repository/recipe_repository"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server api.Server
}

func NewAppConfig() AppConfig {
	_ = godotenv.Load()

	mongoDb := mongo.Connect()

	recipeRepository := recipe_repository.New(mongoDb)

	controller := controller.NewController(recipeRepository, nil)
	server := api.NewServer(controller)

	return AppConfig{
		Server: *server,
	}
}
