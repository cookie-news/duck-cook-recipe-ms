package controller

import (
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/usecase"
)

type Controller struct {
	recipeRepository     repository.RecipeRepository
	likeRecipeRepository repository.LikeRecipeRepository
	recipeStorage        repository.RecipeStorageRepository
	commentRecipeUseCase usecase.CommentRecipeUseCase
	userUseCase          usecase.UserUseCase
	storageUseCase       usecase.StorageUseCase
}

func NewController(
	recipeRepository repository.RecipeRepository,
	likeRecipeRepository repository.LikeRecipeRepository,
	recipeStorage repository.RecipeStorageRepository,
	commentRecipeUseCase usecase.CommentRecipeUseCase,
	userUseCase usecase.UserUseCase,
	storageUseCase usecase.StorageUseCase,
) Controller {
	return Controller{
		recipeRepository,
		likeRecipeRepository,
		recipeStorage,
		commentRecipeUseCase,
		userUseCase,
		storageUseCase,
	}
}
