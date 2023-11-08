package main

import (
	"duck-cook-recipe/api"
	"duck-cook-recipe/controller"
	"duck-cook-recipe/pkg/mongo"
	"duck-cook-recipe/pkg/supabase"
	"duck-cook-recipe/repository/comment_repository"
	"duck-cook-recipe/repository/like_repository"
	"duck-cook-recipe/repository/recipe_repository"
	"duck-cook-recipe/repository/supabase_repository"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server api.Server
}

func NewAppConfig() AppConfig {
	_ = godotenv.Load()

	mongoDb := mongo.Connect()
	supabase := supabase.ConnectStorage()

	storageRecipe := supabase_repository.New(supabase)

	recipeRepository := recipe_repository.New(mongoDb)
	commentRepository := comment_repository.New(mongoDb)
	likeRepository := like_repository.New(mongoDb)

	controller := controller.NewController(recipeRepository, commentRepository, likeRepository, storageRecipe)
	server := api.NewServer(controller)

	return AppConfig{
		Server: *server,
	}
}
