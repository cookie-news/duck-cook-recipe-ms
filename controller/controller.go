package controller

import (
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/usecase"
)

type Controller struct {
	recipeRepository     repository.RecipeRepository
	likeRecipeRepository repository.LikeRecipeRepository
	recipeStorage        repository.RecipeStorage
	commentRecipeUseCase usecase.CommentRecipeUseCase
	userUseCase          usecase.UserUseCase
}

func NewController(
	recipeRepository repository.RecipeRepository,
	likeRecipeRepository repository.LikeRecipeRepository,
	recipeStorage repository.RecipeStorage,
	commentRecipeUseCase usecase.CommentRecipeUseCase,
	userUseCase usecase.UserUseCase,
) Controller {
	return Controller{
		recipeRepository,
		likeRecipeRepository,
		recipeStorage,
		commentRecipeUseCase,
		userUseCase,
	}
}
